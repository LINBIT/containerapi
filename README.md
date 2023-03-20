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

We use [`virter`](https://github.com/linbit/virter) for provisioning development machines.

### Run tests

```
go test ./...
```

Run podman tests in a VM:

```
virter vm run --name podman-test --id 101 --provision virter/provision-podman.toml alma-8
virter vm exec --provision virter/exec-test.toml podman-test
```

Run docker tests in a VM:

```
virter vm run --name docker-test --id 102 --provision virter/provision-docker.toml alma-8
virter vm exec --provision virter/exec-test.toml docker-test
```
