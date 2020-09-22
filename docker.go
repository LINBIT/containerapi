package containerapi

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	log "github.com/sirupsen/logrus"
)

type DockerProvider struct {
	client *client.Client
}

func (d DockerProvider) Create(ctx context.Context, cfg *ContainerConfig) (string, error) {
	dockerEnv := make([]string, 0, len(cfg.env))
	for k, v := range cfg.env {
		dockerEnv = append(dockerEnv, fmt.Sprintf("%s=%s", k, v))
	}

	config := &container.Config{
		Image: cfg.image,
		Env:   dockerEnv,
		Cmd:   cfg.command,
	}

	hostConfig := &container.HostConfig{
		NetworkMode: "host",
	}

	resp, err := d.client.ContainerCreate(ctx, config, hostConfig, nil, cfg.name)
	if err != nil {
		return "", fmt.Errorf("failed to create docker container: %w", err)
	}
	return resp.ID, nil
}

func (d DockerProvider) Remove(ctx context.Context, containerID string) error {
	return d.client.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
}

func (d DockerProvider) Start(ctx context.Context, containerID string) error {
	return d.client.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
}

func (d DockerProvider) Stop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return d.client.ContainerStop(ctx, containerID, timeout)
}

func (d DockerProvider) Wait(ctx context.Context, containerID string) (<-chan int64, <-chan error) {
	msgChan := make(chan int64)
	errChan := make(chan error)
	message, err := d.client.ContainerWait(ctx, containerID, container.WaitConditionNextExit)

	go func() {
		defer close(msgChan)
		defer close(errChan)

		select {
		case <- ctx.Done():
			errChan <- ctx.Err()
		case e := <- err:
			errChan <- e
		case msg := <-message:
			if msg.Error != nil {
				errChan <- fmt.Errorf("error waiting on container end: %s", msg.Error.Message)
				return
			}
			msgChan <- msg.StatusCode
		}
	}()

	return msgChan, errChan
}

func (d DockerProvider) Logs(ctx context.Context, containerID string) (io.ReadCloser, io.ReadCloser, error) {
	options := types.ContainerLogsOptions{
		Follow:     true,
		ShowStdout: true,
		ShowStderr: true,
	}
	combined, err := d.client.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get logs from docker: %w", err)
	}

	readOut, writeOut, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}
	readErr, writeErr, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	go func() {
		defer writeOut.Close()
		defer writeErr.Close()
		defer combined.Close()

		_, err := stdcopy.StdCopy(writeOut, writeErr, combined)
		if err != nil {
			log.WithField("err", err).Warn("failed to copy logs content to pipes")
		}
	}()

	return readOut, readErr, nil
}

func (d DockerProvider) CopyFrom(ctx context.Context, container, source, dest string) error {
	readTar, _, err := d.client.CopyFromContainer(ctx, container, source)
	if err != nil {
		return fmt.Errorf("failed to copy localpkgs from container: %w", err)
	}

	tar := exec.CommandContext(ctx, "tar", "-x", "-C", dest, "-f", "-")

	tar.Stdin = readTar

	tar.Stdout = os.Stdout
	tar.Stderr = os.Stderr

	err = tar.Run()
	if err != nil {
		return fmt.Errorf("failed to extract tar archive: %w", err)
	}
	return nil
}

func (d DockerProvider) Close() error {
	return d.client.Close()
}

func NewDockerProvider(ctx context.Context) (ContainerProvider, error) {
	apiclient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("could not connect to Docker: %w", err)
	}

	_, err = apiclient.ServerVersion(ctx)
	if err != nil {
		return nil, fmt.Errorf("connection check failed: %w", err)
	}

	return &DockerProvider{
		client: apiclient,
	}, nil
}
