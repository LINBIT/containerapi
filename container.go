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
}

// Set the container name, the image to use, and the environment to pass to the container.
func NewContainerConfig(name, image string, env map[string]string) *ContainerConfig {
	if env == nil {
		env = map[string]string{}
	}

	return &ContainerConfig{
		name:  name,
		image: image,
		env:   env,
	}
}

// Sets the environment variable given by the key to val for the container.
func (cfg *ContainerConfig) SetEnv(key, val string) {
	cfg.env[key] = val
}

// A ContainerProvider offers basic control over a container lifecycle
type ContainerProvider interface {
	// Create a new container with the given ContainerConfig, returns the ID of the container for later use.
	//
	// Note: Currently all containers are started with
	// * Remove on exit ("--rm")
	// * Host networking ("--net=host")
	Create(ctx context.Context, cfg *ContainerConfig) (string, error)
	// Start a container that was created previously.
	Start(ctx context.Context, containerID string) error
	// Stop a container that was started previously.
	Stop(ctx context.Context, containerID string, timeout *time.Duration) error
	// Wait can be used to receive the exit code once the container stops.
	Wait(ctx context.Context, containerID string) (<-chan int64, <-chan error)
	// Logs returns reads for stdout and stderr of a running container
	Logs(ctx context.Context, containerID string) (io.ReadCloser, io.ReadCloser, error)
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
