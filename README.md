# xtai CLI (`xtai`)

xtai is a developer-first, natural-language CLI agent that plans and executes
real terminal commands to accomplish complex engineering tasks, without being
tied to a single LLM provider.

This repository contains the Go implementation of the `xtai` CLI tool,
following the engineering and architecture guidelines defined in:

- `project-details.md` — high-level product and architecture specification
- `AGENT.md` — Clean Code & Clean Architecture guidelines for xtai

## Project Structure

The project follows a Clean Architecture layout:

- `cmd/xtai` — CLI entrypoint
- `internal/domain` — core business entities and interfaces
- `internal/usecase` — application services coordinating domain logic
- `internal/delivery` — CLI command handlers (Cobra)
- `internal/infrastructure` — adapters for LLMs, executors, storage, config, and external CLIs
- `pkg/logger` — shared logging helpers

## Prerequisites

- Go 1.26 or newer

## Building

```bash
go build ./cmd/xtai
```

This will produce an `xtai` (or `xtai.exe` on Windows) binary.

## Running

Show CLI help:

```bash
go run ./cmd/xtai --help
```

The initial Phase 1 commands are stubbed and print placeholder messages:

- `xtai plan`
- `xtai execute`
- `xtai replay`
- `xtai sessions`

## Testing

Run the Go test suite:

```bash
go test ./...
```

## Next Phases

Subsequent phases will implement:

- Provider-agnostic LLM interface and adapters
- Safe execution engine and policy system
- Planner subsystem and structured Plan format
- Session logging, replay, and artifacts
- Plugin/skill system and evaluation harness

All phases must continue to respect the constraints and design principles
described in `project-details.md` and `AGENT.md`.

# xtai

**xtai** is a model-agnostic natural-language CLI agent for real developer workflows. It plans and executes terminal commands to accomplish complex engineering tasks across GitHub, JIRA, cloud APIs, and local repos—without being tied to a single LLM provider.

## Specification

The full product specification, architecture, and design goals are in **[project-details.md](project-details.md)**. Use that document as the source of truth for:

- Executive summary and problem statement  
- LLM-agnostic provider model and configuration  
- High-level architecture and system components  
- MVP and V1 scope  
- Security model and use cases  

## Building

```bash
go build -o xtai .
```

## Usage (skeleton)

- `xtai` — show help  
- `xtai plan [goal]` — produce an execution plan (no execution)  
- `xtai execute [goal]` — execute a plan or goal with safe command execution  
- `xtai replay <session_id>` — replay a stored session  
- `xtai sessions` — list or inspect stored sessions  

## Configuration

Configuration can be loaded from YAML or JSON. See `internal/config` for the loader. Example locations (to be wired in): `~/.xtai/config.yaml` or `./.xtai/config.yaml`.

## Tests

```bash
go test ./...
```

## License

See repository license.
