# containerapi

Go bindings to manage containers in a runtime agnostic way.

Currently supports:

* Docker
* Podman >= 4.0.0.

## Why?

While Podman and Docker API is mostly compatible, we noticed instances where the exact behaviour was slightly different.
Having a small set of primitives to cover the most basic container use cases, we can check for divergent behaviour
in CI and ensure we offer a consistent API for both.

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
