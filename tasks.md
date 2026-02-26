Below are the **final refined prompts for each phase** of the **xtai** project — now ready for you to feed into your coding agent. Each prompt explicitly references the required design principles from your `project-details.md` spec *and* the `AGENT.md` guidelines (Clean Code, Clean Architecture, service-based design) so that generated code fully conforms to your engineering expectations.

These prompts are structured to be **incremental, testable, and maintainable**, and explicitly call out acceptance criteria, tests, folder structure, and interface boundaries — consistent with industry best practices for Go Clean Architecture projects.([Medium][1])

---

## 🧱 Phase 1 — Project Skeleton & CLI Foundation

```
Below are two reference files to include verbatim in context:

<contents of project-details.md>
<contents of AGENT.md>

Also include the following **template project structure** exactly as shown:
```
xtai-cli/
├── cmd/
│   └── xtai/                # CLI entrypoint(s)
│       └── main.go
├── internal/
│   ├── domain/              # Core business entities and interfaces
│   │   ├── plan.go
│   │   ├── session.go
│   │   ├── skill.go
│   │   └── provider.go
│   ├── usecase/             # Application logic / services
│   │   ├── planner.go
│   │   ├── executor.go
│   │   ├── session_service.go
│   │   └── skill_service.go
│   ├── delivery/            # CLI handlers (interaction layer)
│   │   ├── plan_cmd.go
│   │   ├── execute_cmd.go
│   │   ├── replay_cmd.go
│   │   └── sessions_cmd.go
│   ├── infrastructure/      # Adapters for external systems
│   │   ├── llm/             # LLM provider adapters (OpenAI, Claude, Ollama)
│   │   │   ├── openai.go
│   │   │   └── ollama.go
│   │   ├── executor/        # Shell/CLI execution adapters
│   │   │   └── shell_exec.go
│   │   ├── storage/         # Session persistence (SQLite/JSONL)
│   │   │   └── session_store.go
│   │   ├── cli/             # Third-party CLI integrations (GitHub, Jira)
│   │   │   ├── github.go
│   │   │   └── jira.go
│   │   └── config/          # Configuration loader (YAML/JSON)
│   │       └── config.go
├── pkg/                     # Shared helpers (optional reusable packages)
│   └── logger/
├── scripts/                 # Dev scripts (build, format, lint)
├── testdata/                # Fixtures and sample outputs for tests
├── .github/                 # CI workflows (tests, lint, vet)
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

Generate the **Go project skeleton** for the `xtai` CLI tool.

### Requirements

1. Follow **Clean Code** principles and **Clean Architecture** layering as defined in `AGENT.md`.
2. Generate a **`go.mod`** for module `github.com/x3ns/xtai`.
3. Create stubbed Go files in the directory tree above:
   - Each package should have appropriate `package` declarations.
   - Include doc comments explaining purpose and responsibilities.
4. Implement CLI using **Cobra** (or equivalent) with stubbed command handlers for:
   - `xtai plan`
   - `xtai execute`
   - `xtai replay`
   - `xtai sessions`
5. Configuration loading code under `internal/infrastructure/config`.
6. Placeholder code in domain/usecase layers with interfaces defined clearly.
7. Add **integration tests** covering at least:
   - CLI usage (`xtai --help`)
   - Each stub command prints a placeholder message
8. README.md must reference both `project-details.md` and `AGENT.md`, and outline how to build/test.

### Deliverables

✔ Directory tree  
✔ Go source files  
✔ Minimal test files under a `*_test.go` convention
✔ README.md  

Write output as full file contents in code blocks, with tests included.
```

---

## 🤖 Phase 2 — Provider-Agnostic LLM Interface

```
Include in context:
<project-details.md>
<AGENT.md>

Using these as engineering rules:

Implement a provider-agnostic LLM interface in Go that:
- Defines an `llm.Provider` interface in domain/usecases.
- Provides concrete adapters for:
  - OpenAI
  - Claude/Anthropic
  - Ollama (local)
- Supports config-based selection via YAML.
- Injects provider dependencies into usecases via constructors.

Deliverables:
- Go interface and adapter implementations
- Config structs + loader code
- Unit tests with mocks demonstrating provider switching
- README explaining how to configure and select providers

The generated code should clearly separate domain interfaces from infra adapters and respect SOLID principles.
```

---

## 🛡 Phase 3 — Safe Execution Engine & Policy System

```
Include in context:
<project-details.md>
<AGENT.md>

Implement the **safe command execution engine** and **policy system** in Go according to Clean Architecture.

Goals:
- Define `executor.CommandExecutor` interface in domain/usecases.
- Provide infrastructure implementation using `os/exec` with:
  - Timeouts
  - Env sanitization
  - Structured logging
- Implement a policy engine that:
  - Loads allow/deny patterns from config
  - Enforces safety before execution
  - Prompts the user in the delivery CLI for approval

Deliverables:
- domain/usecases interfaces
- infrastructure/executor and policy code
- CLI bindings in delivery layer
- Tests verifying:
  - policy enforcement
  - safe vs unsafe commands
  - timeout behavior

Ensure architecture strictly separates interface from implementation.
```

---

## 📑 Phase 4 — Planner + Structured Plan Format

```
Include in context:
<project-details.md>
<AGENT.md>

Implement the **planner subsystem** using Clean Architecture.

Requirements:
- Define a JSON schema for a `Plan` artifact.
- Create Planner service interfaces in domain/usecases.
- Planner should:
  - Accept natural language instructions
  - Use llm.Provider abstraction
  - Emit validated JSON plans
- Prompt templates must be stored separately from business logic.

Deliverables:
- Go plan models + schema
- Planner usecase implementation
- Prompt template resources
- Unit tests verifying Plan structure given canned LLM output

All logic must be documented and testable without UI or infrastructure dependencies.
```

---

## 🔁 Phase 5 — Session Logging, Replay & Artifacts

```
Include in context:
<project-details.md>
<AGENT.md>

Implement **session persistence, replay, and artifact management**.

Requirements:
- Define session models in domain/usecases
- Implement a storage adapter (SQLite or JSONL) under infrastructure
- CLI commands:
  - `xtai sessions list`
  - `xtai replay <id>`
- Store:
  - plan JSON
  - executed steps with outputs
  - artifacts

Deliverables:
- session store adapter implementation
- delivery CLI integration
- tests validating:
  - session CRUD
  - replay fidelity (reproduces all outputs)

Ensure separation of storage concerns from core logic.
```

---

## 🔌 Phase 6 — Plugin/Skill System & First Skill

```
Include in context:
<project-details.md>
<AGENT.md>

Build the **plugin/skill system** and first skill according to Clean Architecture.

Goals:
- Define a `Skill` interface in domain/usecases.
- Implement plugin registry and loader in infrastructure.
- First plugin: **Dependabot ↔ JIRA Reconciler** with:
  - Deterministic lookup via GitHub/JIRA CLIs
  - LLM fallback for ambiguous matches
  - Generates reconciliation report (JSON + Markdown)

Deliverables:
- Plugin interface + registry
- `dependabot_jira` skill code
- Parsers for GitHub and JIRA output
- Tests mocking dependencies and verifying report results

Ensure skills don’t depend on external adapters directly — only via interfaces.
```

---

## 🧪 Phase 7 — Baseline Test & Evaluation Harness

```
Include in context:
<project-details.md>
<AGENT.md>

Implement the **baseline test and evaluation harness** that supports regression testing across all phases.

Requirements:
- Fixture directories for:
  - planner outputs
  - executor mocks
  - session logs
- Test helpers:
  - assertPlanMatches
  - assertExecutionPolicy
  - assertReplayConsistency
- CI config (GitHub Actions) with:
  - unit tests
  - snapshot tests
  - static analysis
- Documentation on adding new fixtures and tests

Deliverables:
- Go test suite + fixtures
- CI pipeline configuration
- Test helper library

Tests must validate correctness and architecture boundaries.
```

---

## 🔁 Prompt Engineering Best Practices (to include at top of every prompt)

Always begin each phase prompt with:

```
Here are the reference specs:
<contents of project-details.md>
<contents of AGENT.md>
```

This ensures your coding agent has the full architectural and style context. A clean architecture design for Go prefers domain/usecases isolation from infrastructure/cli. Models and interfaces live in the core; adapters implement them externally. Use tests to validate every layer’s behavior.([Medium][1])
