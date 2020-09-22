package containerapi

import (
	"context"
	"fmt"
	"io"
	"time"
)

// Configuration passed to the ContainerProvider.Create method.
type ContainerConfig struct {
	name  string
	image string
	env   map[string]string
	command []string
}

// Set the container name, the image to use, and the environment to pass to the container.
func NewContainerConfig(name, image string, env map[string]string, opts ...ConfigOption) *ContainerConfig {
	if env == nil {
		env = map[string]string{}
	}

	cfg := &ContainerConfig{
		name:  name,
		image: image,
		env:   env,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// Sets the environment variable given by the key to val for the container.
func (cfg *ContainerConfig) SetEnv(key, val string) {
	cfg.env[key] = val
}

type ConfigOption func(config *ContainerConfig)

func WithCommand(cmd ...string) ConfigOption {
	return func(config *ContainerConfig) {
		config.command = cmd
	}
}

// A ContainerProvider offers basic control over a container lifecycle
//
// For a standard container lifecycle, use the following pattern:
// 1. call Create()
// 2. defer Remove()
// 3. call Wait()
// 4. call Start()
// 5. defer Stop()
// 6. call Logs(), copy until EOF
// 7. check the Wait() channels for the actual exit code
// 8. call CopyFrom() if required
type ContainerProvider interface {
	// Create a new container with the given ContainerConfig, returns the ID of the container for later use.
	//
	// Note: Currently all containers are started with
	// * Host networking ("--net=host")
	Create(ctx context.Context, cfg *ContainerConfig) (string, error)
	// Remove a container that is not running.
	Remove(ctx context.Context, containerId string) error
	// Start a container that was created previously.
	Start(ctx context.Context, containerID string) error
	// Stop a container. Calling this on a stopped container will return nil.
	Stop(ctx context.Context, containerID string, timeout *time.Duration) error
	// Wait can be used to receive the exit code once the container stops.
	Wait(ctx context.Context, containerID string) (<-chan int64, <-chan error)
	// Logs returns reads for stdout and stderr of a running container
	Logs(ctx context.Context, containerID string) (io.ReadCloser, io.ReadCloser, error)
	// CopyFrom copies files/folders from the container to the local filesystem
	CopyFrom(ctx context.Context, container, sourcePath, destPath string) error
	// Close should be called once this provider is no longer needed.
	Close() error
}

var providers = map[string]func(context.Context) (ContainerProvider, error){
	"podman": NewPodmanProvider,
	"docker": NewDockerProvider,
}

// Returns the list of available providers
func Providers() []string {
	keys := make([]string, 0, len(providers))
	for k := range providers {
		keys = append(keys, k)
	}
	return keys
}

// Returns a new, initialized ContainerProvider of the given provider type.
func NewProvider(ctx context.Context, provider string) (ContainerProvider, error) {
	providerFun, ok := providers[provider]
	if !ok {
		return nil, fmt.Errorf("unknown container provider '%s'", provider)
	}
	return providerFun(ctx)
}
