# mini-container-orchestrator

A small container orchestrator written in Go, built as a learning project. Status: early scaffold — core packages are stubbed out and not yet implemented.

## Project layout

```
cmd/
  orchestrator/   entry point for the orchestrator service
  nodeagent/      entry point for the node agent
  cli/            entry point for the CLI

internal/
  scheduler/      placement/scheduling logic
  runtime/        container runtime integration
  registry/       node/task registry

pkg/
  api/            shared API types
  api/cli/        CLI-facing API helpers
  api/nodeagent/  node agent-facing API helpers
```

`cmd/` holds the buildable binaries; `internal/` holds implementation details private to this module; `pkg/` holds code intended to be reusable/importable.

## Requirements

- Go 1.24.3+

## Build

```sh
go build ./...
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
