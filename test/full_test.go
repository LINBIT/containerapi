package test

import (
	"bytes"
	"context"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/LINBIT/containerapi"
)

func containerName(name string) string {
	return strings.ReplaceAll(name, "/", "-")
}

func TestRun(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/hello-world", nil)
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

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
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		stdout, stderr, err := provider.Logs(ctx, id)
		assert.NoError(t, err)

		assertIOEquals(t, []byte(expectedLogStdout), stdout, []byte(expectedLogStderr), stderr)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

func TestContainerLifecycle(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("false"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(1), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

func TestCopyFrom(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		tmp, err := os.MkdirTemp("", "*")
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := os.RemoveAll(tmp)
			assert.NoError(t, err)
		})

		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("sleep", "1"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		err = provider.CopyFrom(ctx, id, "/bin/busybox", tmp)
		assert.NoError(t, err)

		assert.FileExists(t, tmp+"/busybox")
	})
}

func TestLifecycle(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("echo", "foobar"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		// Assert that the logs are accessible even after the container exits.
		time.Sleep(1 * time.Second)

		stdout, stderr, err := provider.Logs(ctx, id)
		assert.NoError(t, err)

		assertIOEquals(t, []byte("foobar\n"), stdout, []byte{}, stderr)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}

		// Assert we can still access logs after container completed
		stdout, stderr, err = provider.Logs(ctx, id)
		assert.NoError(t, err)

		assertIOEquals(t, []byte("foobar\n"), stdout, []byte{}, stderr)

		// Assert that calling stop on an already stopped container works without error
		timeout := 1 * time.Second
		err = provider.Stop(ctx, id, &timeout)
		assert.NoError(t, err)

		err = provider.Remove(ctx, id)
		assert.NoError(t, err)

		// Asserts that the container is unknown to the container runtime after Remove()
		time.Sleep(1 * time.Second)
		err = provider.Start(ctx, id)
		assert.Error(t, err)
	})
}

func TestRunWithCancel(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil, containerapi.WithCommand("tail", "-f", "/dev/null"))
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			timeout := 1 * time.Second
			err := provider.Stop(ctx, id, &timeout)
			assert.NoError(t, err)
			err = provider.Remove(ctx, id)
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

const volumeTestScript = `
cat /readonly/source > /readwrite/dest
echo "newcontent" > /readonly/source || true
touch /readonly/invalid || true
`

var testSourceContent = []byte("foobar")

func TestRunWithVolumes(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		tempDir, err := os.MkdirTemp("", containerName(t.Name())+"-*")
		if !assert.NoError(t, err) {
			t.Fatal("failed to create tempdir")
		}
		t.Cleanup(func() {
			err := os.RemoveAll(tempDir)
			assert.NoError(t, err)
		})

		roDir := tempDir + "/readonly"
		rwDir := tempDir + "/readwrite"
		source := roDir + "/source"
		dest := rwDir + "/dest"

		err = os.Mkdir(roDir, 0o755)
		assert.NoError(t, err)

		err = os.Mkdir(rwDir, 0o755)
		assert.NoError(t, err)

		err = os.WriteFile(source, testSourceContent, 0o644)
		assert.NoError(t, err)

		roBind := containerapi.Mount{
			HostPath:      roDir,
			ContainerPath: "/readonly",
			ReadOnly:      true,
		}

		rwBind := containerapi.Mount{
			HostPath:      rwDir,
			ContainerPath: "/readwrite",
		}

		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil,
			containerapi.WithCommand("sh", "-ec", volumeTestScript),
			containerapi.WithMounts(roBind, rwBind),
		)

		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}

		assert.FileExists(t, dest)
		assert.NoFileExists(t, roDir+"/invalid")

		actualSrc, err := os.ReadFile(source)
		assert.NoError(t, err)
		assert.Equal(t, testSourceContent, actualSrc, "readonly data changed")

		actualDest, err := os.ReadFile(dest)
		assert.NoError(t, err)
		assert.Equal(t, testSourceContent, actualDest)
	})
}

func TestRunWithDnsSettings(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testipv4 := net.ParseIP("1.1.1.1")
		testipv6 := net.ParseIP("2606:4700:4700::64")

		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/busybox", nil,
			containerapi.WithCommand("cat", "/etc/resolv.conf"),
			containerapi.WithDNSServers(testipv4, testipv6),
			containerapi.WithDNSSearchDomains("test", "containerapi.test"),
		)
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		stdout, stderr, err := provider.Logs(ctx, id)
		assert.NoError(t, err)

		const expectedOut = "search test containerapi.test\nnameserver 1.1.1.1\nnameserver 2606:4700:4700::64\n"
		assertIOEquals(t, []byte(expectedOut), stdout, []byte(""), stderr)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

func TestRunWithExtraHosts(t *testing.T) {
	runOnAllProviders(t, 30*time.Second, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/alpine", nil,
			containerapi.WithCommand("getent", "hosts", "extra.example.com"),
			containerapi.WithPullConfig(containerapi.PullIfNotExists),
			containerapi.WithExtraHosts(containerapi.ExtraHost{HostName: "extra.example.com", IP: "1.0.0.1"}),
		)
		id, err := provider.Create(ctx, testConfig)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})

		returnCodeChan, errChan := provider.Wait(ctx, id)

		err = provider.Start(ctx, id)
		assert.NoError(t, err)

		stdout, stderr, err := provider.Logs(ctx, id)
		assert.NoError(t, err)

		const expectedOut = "1.0.0.1           extra.example.com  extra.example.com\n"
		assertIOEquals(t, []byte(expectedOut), stdout, []byte(""), stderr)

		select {
		case r := <-returnCodeChan:
			assert.Equal(t, int64(0), r)
		case err := <-errChan:
			assert.FailNow(t, err.Error())
		}
	})
}

func TestRunWithImagePull(t *testing.T) {
	runOnAllProviders(t, 1*time.Minute, func(ctx context.Context, provider containerapi.ContainerProvider, t *testing.T) {
		_ = exec.CommandContext(ctx, provider.Command(), "image", "rm", "docker.io/alpine").Run()

		failingConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/alpine", nil)
		id, err := provider.Create(ctx, failingConfig)
		if !assert.Error(t, err) {
			t.Cleanup(func() {
				err := provider.Remove(ctx, id)
				assert.NoError(t, err)
			})
		}

		configWithPull := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/alpine", nil, containerapi.WithPullConfig(containerapi.PullIfNotExists))
		id, err = provider.Create(ctx, configWithPull)
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := provider.Remove(ctx, id)
			assert.NoError(t, err)
		})
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

func assertIOEquals(t *testing.T, expectedStdout []byte, stdout io.ReadCloser, expectedStderr []byte, stderr io.ReadCloser) bool {
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

	outEqual := assert.Equal(t, expectedStdout, stdoutBuff.Bytes())
	errEqual := assert.Equal(t, expectedStderr, stderrBuff.Bytes())
	return outEqual && errEqual
}
