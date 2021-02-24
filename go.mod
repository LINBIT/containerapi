module github.com/LINBIT/containerapi

go 1.14

require (
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/containerd/containerd v1.4.1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20201020191947-73dc6a680cdd+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.14
	github.com/go-openapi/validate v0.20.2
	github.com/go-swagger/go-swagger v0.26.1
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/moby/term v0.0.0-20200915141129-7f0af18e79f2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2-0.20190823105129-775207bd45b6 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	gotest.tools/v3 v3.0.3 // indirect
)

// We need a new go-swagger version, but this dependency breaks definition expansion, so we must replace it with an
// older version.
replace github.com/go-openapi/loads v0.20.2 => github.com/go-openapi/loads v0.19.0
