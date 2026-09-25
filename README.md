# mini-container-orchestrator

A small container orchestrator written in Go, built as a learning project. Status: early scaffold — the runtime abstraction and its Docker implementation are in place; scheduling, registry, and API layers are not yet implemented.

## Project layout

```
cmd/
  orchestrator/   entry point for the orchestrator service
  nodeagent/      entry point for the node agent
  cli/            entry point for the CLI

internal/
  scheduler/      placement/scheduling logic
  runtime/        container runtime abstraction (Runtime interface + Docker implementation)
  registry/       node/task registry

pkg/
  api/            shared API types
  api/cli/        CLI-facing API helpers
  api/nodeagent/  node agent-facing API helpers
```

`cmd/` holds the buildable binaries; `internal/` holds implementation details private to this module; `pkg/` holds code intended to be reusable/importable.

### Runtime layer

`internal/runtime` defines a `Runtime` interface (create/start/stop/remove/status/list containers, plus `Close`) so the orchestrator, node agent, and CLI depend on an abstraction rather than a concrete container engine. `DockerRuntime` is the current implementation, backed by the [Docker Engine SDK](https://github.com/moby/moby) (`github.com/moby/moby/client`). Construct one with `runtime.NewDockerRuntime()`, which connects to the local Docker daemon via `client.FromEnv`.

## Requirements

- Go 1.24.3+
- A running Docker daemon (for `DockerRuntime`)

## Build

Each binary under `cmd/` is independent, so build them individually:

```sh
go build ./cmd/orchestrator
go build ./cmd/nodeagent
go build ./cmd/cli
```

## Run

```sh
go run ./cmd/orchestrator
go run ./cmd/nodeagent
go run ./cmd/cli
```

## Module

```
github.com/hashim715/mini-container-orchestrator
```
