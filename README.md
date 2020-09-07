# containerapi

Go bindings to manage containers in a runtime agnostic way.

Current plan is to support:

* Docker
* podman

## Why?

While podman claims to provide an API that is compatible with Docker, in our experience this is not always the case.

Edge cases include:
* Diverging container states (podman does not recongize the "removed" state)
* podman fails to remove containers that are marked as "autoremove" via Docker compatible API.

## Development

### Regenerate API bindings

```
go generate ./...
```

### Run tests

```
go test ./...
```
