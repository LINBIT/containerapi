package containerapi

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/LINBIT/containerapi/internal/podmanapi"
	"github.com/LINBIT/containerapi/internal/podmanapi/containers"
	"github.com/LINBIT/containerapi/internal/podmanapi/images"
	"github.com/LINBIT/containerapi/internal/podmanapi/system"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

// swagger command from https://github.com/go-swagger/go-swagger/releases/
// swagger file from https://storage.googleapis.com/libpod-master-releases/swagger-latest-master.yaml
// with slight modifications, see ./swagger-2020-09-02.yaml.patch
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f ./podman_swagger.yaml --with-flatten expand --tags=containers --tags=system --tags=images --skip-models -t ./internal -c podmanapi --keep-spec-order

type PodmanProvider struct {
	client *podmanapi.ProvidesaContainerCompatibleInterface
}

func (d PodmanProvider) Close() error {
	return nil
}

func (d PodmanProvider) Create(ctx context.Context, cfg *ContainerConfig) (string, error) {
	existsParams := images.NewLibpodImageExistsParamsWithContext(ctx)
	existsParams.Name = cfg.image
	_, err := d.client.Images.LibpodImageExists(existsParams)
	if cfg.pullConfig != nil && cfg.pullConfig(cfg.image, err == nil) {
		pullParams := images.NewLibpodImagesPullParamsWithContext(ctx)
		pullParams.Reference = &cfg.image
		log.WithField("image", cfg.image).Info("pulling...")

		read, write, err := os.Pipe()
		if err != nil {
			return "", err
		}
		defer write.Close()

		go func() {
			defer read.Close()
			scan := bufio.NewScanner(read)
			for scan.Scan() {
				line := scan.Text()
				type item struct {
					Stream string `json:"stream,omitempty"`
				}
				data := item{}
				err := json.Unmarshal([]byte(line), &data)
				if err != nil {
					log.WithError(err).WithField("line", line).Errorf("failed to read pull log")
				}

				if data.Stream != "" {
					log.Info(data.Stream)
				}
			}
		}()

		// Some explanation is required here: We force the result to be interpreted as raw bytes in the success case.
		// This is needed, as:
		// * podman does not set an appropriate Content-Type header for the response
		// * go-swagger defaults to assume JSON output
		// * podman will send newline delimited json
		// * go-swagger will stop after reading the first line of json
		// This means that we return to early from the ImagePull function, so the image isn't actually pulled in time
		// to create the container. Forcing the response to be read by the runtime.ByteStreamConsumer means we always
		// wait for the full response, which waits for the image pull to complete.
		_, err = d.client.Images.LibpodImagesPull(pullParams, write, func(operation *runtime.ClientOperation) {
			originalReader := operation.Reader
			operation.Reader = runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
				if response.Code() == 200 {
					consumer = runtime.ByteStreamConsumer(runtime.ClosesStream)
				}
				return originalReader.ReadResponse(response, consumer)
			})
		})
		if err != nil {
			return "", err
		}
		log.WithField("image", cfg.image).Infof("pull complete")
	}

	params := containers.NewLibpodCreateContainerParamsWithContext(ctx)
	timeout := uint64(0)

	mounts := make([]*containers.LibpodCreateContainerParamsBodyMountsItems0, len(cfg.mounts))
	for i, b := range cfg.mounts {
		// "Z" enables SELinux relabels so the content is accessible for the container
		opts := []string{"Z"}
		if b.ReadOnly {
			opts = append(opts, "ro")
		}

		mounts[i] = &containers.LibpodCreateContainerParamsBodyMountsItems0{
			Source:      b.HostPath,
			Destination: b.ContainerPath,
			Options:     opts,
			Type:        "bind",
		}
	}

	servers := make([]string, len(cfg.dnsServers))
	for i, server := range cfg.dnsServers {
		servers[i] = server.String()
	}

	params.WithCreate(containers.LibpodCreateContainerBody{
		Image: cfg.image,
		Netns: &containers.LibpodCreateContainerParamsBodyNetns{
			Nsmode: "host",
		},
		DNSServers: servers,
		DNSSearch:  cfg.dnsSearchDomains,
		Command:    cfg.command,
		Env:        cfg.env,
		Name:       cfg.name,
		// In case we pass a 0 value to the /container/<name>/stop API, this timeout will be used.
		// Setting this to 0 means: send SIGKILL immediately
		StopTimeout: &timeout,
		Mounts:      mounts,
	})
	resp, err := d.client.Containers.LibpodCreateContainer(params)
	if err != nil {
		return "", err
	}
	return resp.GetPayload().ID, nil
}

func (d PodmanProvider) Remove(ctx context.Context, containerID string) error {
	params := containers.NewLibpodRemoveContainerParamsWithContext(ctx)
	params.WithName(containerID)
	_, err := d.client.Containers.LibpodRemoveContainer(params)
	return err
}

func (d PodmanProvider) Start(ctx context.Context, containerID string) error {
	params := containers.NewLibpodStartContainerParamsWithContext(ctx)
	params.WithName(containerID)
	_, err := d.client.Containers.LibpodStartContainer(params)
	return err
}

func (d PodmanProvider) Stop(ctx context.Context, containerID string, timeout *time.Duration) error {
	params := containers.NewLibpodStopContainerParamsWithContext(ctx)
	params.WithName(containerID)
	seconds := int64(timeout.Seconds())
	params.WithT(&seconds)
	_, err := d.client.Containers.LibpodStopContainer(params)
	if _, ok := err.(*containers.LibpodStopContainerNotModified); ok {
		return nil
	}
	return err
}

func (d PodmanProvider) Wait(ctx context.Context, containerID string) (<-chan int64, <-chan error) {
	params := containers.NewLibpodWaitContainerParamsWithContext(ctx)
	params.WithName(containerID)

	errChan := make(chan error)
	statusChan := make(chan int64)

	go func() {
		defer close(errChan)
		defer close(statusChan)

		resp, err := d.client.Containers.LibpodWaitContainer(params)
		if err != nil {
			errChan <- err
			return
		}
		v, err := strconv.Atoi(strings.TrimSpace(resp.GetPayload()))
		if err != nil {
			errChan <- err
			return
		}
		statusChan <- int64(v)
	}()

	return statusChan, errChan
}

// Inserts a '\n' after every write operation
type newlineInserter struct {
	inner io.WriteCloser
}

func (i newlineInserter) Write(p []byte) (int, error) {
	n, err := i.inner.Write(p)
	if err != nil {
		return n, err
	}
	_, err = i.inner.Write([]byte{'\n'})
	return n, err
}

func (i newlineInserter) Close() error {
	return i.inner.Close()
}

// Copies the output of podman's Logs() call into a stdout and stderr pipe
func (d PodmanProvider) Logs(ctx context.Context, container string) (io.ReadCloser, io.ReadCloser, error) {
	params := containers.NewLibpodLogsFromContainerParamsWithContext(ctx)
	yes := true
	params.WithName(container)
	params.WithFollow(&yes)
	params.WithStdout(&yes)
	params.WithStderr(&yes)

	readOut, writeOut, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}
	readErr, writeErr, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	readHttp, writeHttp, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create http pipe: %w", err)
	}

	// Podman does not send any newline characters with each line. We have to insert them ourselves.
	// This wrapper will add a '\n' on every .Write() call. The StdCopy() call will only ever call .Write()
	// on a full line, so this inserts newlines at the correct location.
	// See also: https://github.com/containers/podman/issues/6539
	wrappedWriteOut := &newlineInserter{writeOut}
	wrappedWriteErr := &newlineInserter{writeErr}

	go func() {
		defer wrappedWriteOut.Close()
		defer wrappedWriteErr.Close()
		defer readHttp.Close()
		_, err := stdcopy.StdCopy(wrappedWriteOut, wrappedWriteErr, readHttp)
		if err != nil {
			log.WithField("err", err).Warn("failed to copy logs content to pipes")
		}
	}()

	go func() {
		defer writeHttp.Close()
		_, err = d.client.Containers.LibpodLogsFromContainer(params, writeHttp)
		if err != nil {
			log.WithField("err", err).Warn("failed to copy logs content to pipes")
		}
	}()

	return readOut, readErr, err
}

// CopyFrom copies files/directories from a container to the host file system
// NOTE: We can't use what `podman cp` uses here, because it is not yet
// implemented in the API. See https://github.com/containers/podman/issues/6050.
// Thus, we just cheap out and call `podman cp` from the code here.
func (d PodmanProvider) CopyFrom(ctx context.Context, container, source, dest string) error {
	cmd := exec.CommandContext(ctx, "podman", "cp", container+":"+source, dest)

	_, err := cmd.Output()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			log.Warn("Execution of \"podman cp\" failed. Details:")
			fmt.Fprint(log.StandardLogger().Out, string(e.Stderr))
		}
		return fmt.Errorf("failed to copy files from container: %w", err)
	}

	return nil
}

func NewPodmanProvider(ctx context.Context) (ContainerProvider, error) {
	socketDir, ok := os.LookupEnv("XDG_RUNTIME_DIR")
	if !ok {
		socketDir = "/run"
	}
	socket := socketDir + "/podman/podman.sock"

	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
				parts := strings.SplitN(addr, ":", 2)
				dialer := net.Dialer{}
				return dialer.DialContext(ctx, "unix", parts[0])
			},
			DisableCompression: true,
		},
	}

	transport := httptransport.NewWithClient(socket, "v1.0.0/", []string{"http"}, httpClient)
	client := podmanapi.New(transport, strfmt.Default)

	_, err := client.System.SystemVersion(system.NewSystemVersionParamsWithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("connection check failed: %w", err)
	}

	return &PodmanProvider{
		client: client,
	}, nil
}
