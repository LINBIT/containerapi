package containerapi

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumePullStream(t *testing.T) {
	tests := []struct {
		name           string
		stream         string
		wantDownloaded bool
		wantErr        string
	}{
		{
			name: "image already up to date",
			stream: `{"status":"Pulling from library/alpine","id":"latest"}
{"status":"Digest: sha256:abc"}
{"status":"Status: Image is up to date for alpine:latest"}`,
			wantDownloaded: false,
		},
		{
			name: "newly downloaded image",
			stream: `{"status":"Pulling from library/alpine","id":"latest"}
{"status":"Pulling fs layer","progressDetail":{},"id":"abc"}
{"status":"Downloading","progressDetail":{"current":1,"total":2},"id":"abc"}
{"status":"Pull complete","progressDetail":{},"id":"abc"}
{"status":"Digest: sha256:def"}
{"status":"Status: Downloaded newer image for alpine:latest"}`,
			wantDownloaded: true,
		},
		{
			name: "all layers already exist still counts as no-op",
			stream: `{"status":"Pulling from library/alpine","id":"latest"}
{"status":"Already exists","progressDetail":{},"id":"abc"}
{"status":"Status: Image is up to date for alpine:latest"}`,
			wantDownloaded: false,
		},
		{
			name:    "engine error in stream",
			stream:  `{"errorDetail":{"message":"manifest unknown"},"error":"manifest unknown"}`,
			wantErr: "manifest unknown",
		},
		{
			name:           "empty stream",
			stream:         ``,
			wantDownloaded: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			downloaded, err := consumePullStream(strings.NewReader(tc.stream))
			if tc.wantErr != "" {
				assert.EqualError(t, err, tc.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantDownloaded, downloaded)
		})
	}
}
