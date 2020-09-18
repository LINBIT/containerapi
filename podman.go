package containerapi

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/LINBIT/containerapi/internal/podmanapi"
	"github.com/LINBIT/containerapi/internal/podmanapi/containers"
	"github.com/LINBIT/containerapi/internal/podmanapi/system"

	"github.com/docker/docker/pkg/stdcopy"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

// swagger command from https://github.com/go-swagger/go-swagger/releases/
// swagger file from https://storage.googleapis.com/libpod-master-releases/swagger-latest-master.yaml
// with slight modifications, see ./swagger-2020-09-02.yaml.patch
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f ./podman_swagger.yaml --with-flatten expand --tags=containers --tags=system --skip-models -t ./internal -c podmanapi --keep-spec-order

type PodmanProvider struct {
	client *podmanapi.ProvidesaContainerCompatibleInterface
}

func (d PodmanProvider) Close() error {
	return nil
}

func (d PodmanProvider) Create(ctx context.Context, cfg *ContainerConfig) (string, error) {
	params := containers.NewLibpodCreateContainerParamsWithContext(ctx)
	params.WithCreate(containers.LibpodCreateContainerBody{
		Image: cfg.image,
		Netns: &containers.LibpodCreateContainerParamsBodyNetns{
			Nsmode: "host",
		},
		Command: cfg.command,
		Env:    cfg.env,
		Name:   cfg.name,
		Remove: true,
	})
	resp, err := d.client.Containers.LibpodCreateContainer(params)
	if err != nil {
		return "", err
	}
	return resp.GetPayload().ID, nil
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
