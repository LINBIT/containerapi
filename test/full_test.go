package test

import (
	"bytes"
	"context"
	"github.com/LINBIT/containerapi"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"sync"
	"testing"
	"time"
)

func containerName(name string) string {
	return strings.ReplaceAll(name, "/", "-")
}

func TestRun(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/hello-world", nil)
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		returnCodeChan, errChan := provider.Wait(ctx, id)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

const (
	logTestScript = `
echo stdout1
echo stderr1 >&2
echo stdout2
echo stderr2 >&2
sleep 1 # ensures we don't miss the container lifecycle
`
	expectedLogStdout = "stdout1\nstdout2\n"
	expectedLogStderr = "stderr1\nstderr2\n"
)

func TestRunWithLogs(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("sh", "-ec", logTestScript))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		stdout, stderr, err := provider.Logs(ctx, id)
		assert.NoError(t, err)

		returnCodeChan, errChan := provider.Wait(ctx, id)

		stdoutBuff := bytes.Buffer{}
		stderrBuff := bytes.Buffer{}

		vg := sync.WaitGroup{}
		vg.Add(2)
		go func() {
			defer vg.Done()
			_, err := io.Copy(&stdoutBuff, stdout)
			assert.NoError(t, err)
		}()
		go func() {
			defer vg.Done()
			_, err := io.Copy(&stderrBuff, stderr)
			assert.NoError(t, err)
		}()

		vg.Wait()

		assert.Equal(t, expectedLogStdout, stdoutBuff.String())
		assert.Equal(t, expectedLogStderr, stderrBuff.String())

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

func TestRunWithCancel(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("sleep", "10"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			timeout := 1 * time.Second
			err := provider.Stop(ctx, id, &timeout)
			assert.NoError(t, err)
		})

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		waitCtx, cancel := context.WithCancel(ctx)
		returnCodeChan, errChan := provider.Wait(waitCtx, id)

		cancel()

		select {
		case r := <-returnCodeChan:
			assert.FailNow(t, "got return code for cancelled Wait(): %d", r)
		case err := <-errChan:
			assert.Error(t, err)
		}
	})
}

// Tries to run the test function on any provider, skipping those that are not available
// If no provider is available, the whole test will fail
func runOnAllProviders(t *testing.T, timeout time.Duration, lambda func(context.Context, containerapi.ContainerProvider, *testing.T)) {
	anySuccess := false
	for _, provider := range containerapi.Providers() {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)

		success := t.Run(provider, func(t *testing.T) {
			provider, err := containerapi.NewProvider(ctx, provider)
			if err != nil {
				t.Skipf("failed to initialize provider: %v", err)
			}

			lambda(ctx, provider, t)
		})

		if success {
			anySuccess = true
		}
	}

	if !anySuccess {
		t.Fatal("expected (at least) one test to succeed")
	}
}
