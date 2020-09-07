package test

import (
	"context"
	"github.com/LINBIT/containerapi"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func containerName(name string) string {
	return strings.ReplaceAll(name, "/", "-")
}

func TestRun(t *testing.T) {
	anySuccess := false
	for _, provider := range containerapi.Providers() {
		runSuccess := false
		t.Run(provider, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			t.Cleanup(cancel)

			provider, err := containerapi.NewProvider(ctx, provider)
			if err != nil {
				t.Skipf("failed to initialize provider: %v", err)
			}

			testConfig := containerapi.NewContainerConfig(containerName(t.Name()), "docker.io/hello-world", nil)
			id, err := provider.Create(ctx, testConfig)
			assert.NoError(t, err)

			err = provider.Start(ctx, id)
			assert.NoError(t, err)

			returnCodeChan, errChan := provider.Wait(ctx, id)

			select {
			case r := <- returnCodeChan:
				assert.Equal(t, int64(0), r)
				runSuccess = true
			case err := <- errChan:
				assert.FailNow(t, err.Error())
			}
		})

		if runSuccess {
			anySuccess = true
		}
	}

	if !anySuccess {
		t.Fatal("expected (at least) one test to succeed")
	}
}
