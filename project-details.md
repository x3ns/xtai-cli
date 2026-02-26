# xtai (`xtai`)
## A Model-Agnostic Natural-Language CLI Agent for Real Developer Workflows (Go-first)

---

# 1. Executive Summary

**xtai (`xtai`)** is a developer-first, natural-language CLI agent that plans and executes real terminal commands to accomplish complex engineering tasks—**without being tied to a single LLM provider**.

You type:

```bash
xtai "Find which dependabot alerts already have JIRA tickets and generate a coverage report"
````

The agent:

1. Produces a transparent execution plan
2. Executes required CLI/API commands safely
3. Correlates results using deterministic + AI reasoning
4. Generates reproducible artifacts
5. Logs the full session for audit and replay

**Key design goals:**

* **LLM-agnostic**: swap between OpenAI, Claude, Ollama, or any compatible provider via configuration
* **Go-first implementation**: robust, distributable binary + strong concurrency primitives
* **Optional Python “assist” modules**: only where it materially improves developer velocity (e.g., parsing, ML eval harness), invoked as managed subprocesses

---

# 2. What Problem It Solves

Modern developer workflows are fragmented across:

* GitHub (PRs, alerts, CI)
* JIRA (tickets)
* Cloud APIs
* Logs (DataDog, CloudWatch, etc.)
* Local repositories
* Terraform / Kubernetes state
* CI/CD pipelines

High-value engineering tasks require:

* Multiple CLI calls
* Parsing JSON
* Searching across systems
* Correlating structured + unstructured data
* Manual decision-making

These tasks are repetitive, error-prone, context-heavy, and time-consuming.

`xtai` reduces cognitive overhead by acting as a **tool-orchestrating agent** that:

* plans steps
* runs tools
* captures outputs
* derives results
* produces artifacts

---

# 3. Core Philosophy

## 3.1 Tool-First Architecture

The agent does not hallucinate answers it can fetch via tools. If the truth is available via:

* `gh`, `jira`, `kubectl`, `terraform`, `git`, `curl`
* local file parsing
* cloud APIs

…it must retrieve it.

## 3.2 Deterministic Where Possible

* Parsing lockfiles → deterministic
* Matching CVE IDs → deterministic
* Building dependency graphs → deterministic

LLMs are used for:

* planning and decomposition
* ambiguity resolution
* summarization
* hypothesis ranking

## 3.3 Transparent Execution

The agent shows:

* the plan
* commands it will run
* the output of each command
* the final derived reasoning

## 3.4 Safe-by-Default

* read-only by default
* explicit confirmation for writes
* command risk scoring
* allow/deny policies
* secrets redaction in logs

## 3.5 Observable & Reproducible

Each session:

* is logged
* is replayable
* produces artifacts
* can be regression tested

---

# 4. LLM-Agnostic Provider Model

## 4.1 Provider Interface (Conceptual)

`xtai` treats LLMs as interchangeable backends behind a stable interface:

* **Generate**: produce a plan or analysis from prompts + tool schema
* **Embed** (optional): create embeddings for clustering/search
* **Stream** (optional): incremental output for better CLI UX

### Supported backends (initial)

* OpenAI (API)
* Anthropic Claude (API)
* Ollama (local)
* “Custom HTTP” (any model behind an OpenAI-compatible or bespoke endpoint)

## 4.2 Provider Configuration

LLM choice is part of setup:

```bash
xtai configure llm \
  --provider ollama \
  --model qwen2.5:14b \
  --endpoint http://localhost:11434
```

Or:

```bash
xtai configure llm \
  --provider openai \
  --model gpt-5 \
  --env OPENAI_API_KEY
```

Or:

```bash
xtai configure llm \
  --provider anthropic \
  --model claude-3.5-sonnet \
  --env ANTHROPIC_API_KEY
```

**Routing rules** can be configured:

* “planner model” (higher reasoning)
* “summarizer model” (cheaper)
* “embedder” (optional)

## 4.3 Model Routing Strategy (Recommended)

* **Planner**: best reasoning model available (Claude/OpenAI)
* **Extractor/Summarizer**: cheaper model
* **Offline/local**: Ollama for privacy-sensitive environments, with reduced capability expectations

---

# 5. High-Level Architecture

```text
User
  ↓
CLI (`xtai`) — Go binary
  ↓
Context Collector (Go)
  ↓
Planner (LLM via Provider Interface)
  ↓
Structured Plan (JSON)
  ↓
Policy Engine + Risk Scorer (Go)
  ↓
Tool Executor (Go)
  ↓
Outputs + Extracted Facts (Go)
  ↓
Reasoning/Correlation (Go + optional Python)
  ↓
Artifacts + Session Store (Go)
```

---

# 6. System Components (Go-first)

## 6.1 CLI Front-End (Go)

Responsibilities:

* parse natural language request
* load config + policies
* render plan and diffs
* prompt user confirmations
* present final report and artifacts

Recommended libs:

* `cobra` or `urfave/cli`
* `charmbracelet` (bubbletea/lipgloss) OR simpler TUI with plain output

## 6.2 Context Collector (Go)

Collects local signals:

* git repo + remote
* current branch + diff stats
* environment metadaxtai (sanitized)
* tool availability (`gh`, `jira`, `kubectl`, etc.)

## 6.3 Planner (LLM, provider-agnostic)

Input:

* user goal
* tool schema (available tools + parameters)
* context snapshot
* policy constraints

Output:

* structured plan (validated JSON)

## 6.4 Tool/Command Executor (Go)

Executes:

* shell commands
* structured HTTP calls
* “virtual tools” like `read_file`, `search_repo`, `parse_json`

Key features:

* timeouts
* stdout/stderr capture
* exit code tracking
* streaming output support
* pagination helpers
* output size caps + summarization hooks

## 6.5 Plugin System (Go)

Plugins encapsulate skills.

Each plugin defines:

* trigger conditions / capabilities
* deterministic parsers/matchers
* report formatters
* optional schemas for tool use

Examples:

* dependabot ↔ jira reconciler
* ci flake analyzer
* terraform plan explainer
* k8s rca helper

## 6.6 Policy Engine (Go)

Defines:

* allowed commands patterns
* blocked commands patterns
* risk scoring
* write permissions
* CI non-interactive behavior

Supports:

* YAML policy file (`policy.yaml`)
* local overrides

## 6.7 Session Store (Go)

Stores:

* plan
* tool calls
* outputs
* extracted facts
* artifacts

Recommended:

* SQLite (via `modernc.org/sqlite` or `mattn/go-sqlite3`)
* JSONL append-only logs for portability

Enables:

```bash
xtai replay <session_id>
xtai explain <session_id>
```

---

# 7. Optional Python Integration (Only Where It Helps)

Go remains the orchestrator and source of truth.

Python is used selectively for:

* advanced parsing of complex formats (if Go libs are painful)
* evaluation harness / metrics scripts
* optional embeddings clustering experiments
* quick iteration on fuzzy matching logic

Integration approach:

* Python invoked as a subprocess with strict I/O contracts:

  * JSON in / JSON out
* sandboxed execution
* versioned “shim” modules

Example:

```bash
xtai py run parse_terraform_plan --input plan.txt --output facts.json
```

---

# 8. Real-World Use Cases

## 8.1 Dependabot ↔ JIRA Coverage (Example)

```bash
xtai "In this repo, find all high/critical dependabot alerts and tell me which have JIRA tickets. Output a table."
```

Agent flow:

1. discover org/repo from git remote
2. pull dependabot alerts via `gh api ... --paginate`
3. normalize advisories (CVE/GHSA/package/ecosystem)
4. query JIRA via JQL
5. strict matching first, semantic fallback if needed
6. output report + artifacts

Artifacts:

* `report.md`
* `report.json`

## 8.2 CI Flakiness Analysis

```bash
xtai "Analyze last 100 CI runs and detect flaky tests."
```

## 8.3 Terraform Cost Delxtai Explanation

```bash
xtai "Explain cost delxtai between this Terraform branch and main."
```

## 8.4 Kubernetes Pod Failure RCA

```bash
xtai "Why is api-service restarting?"
```

---

# 9. Reliability & Evaluation

## Offline Evaluation

Create fixtures:

* sample GitHub API responses
* sample JIRA issue sets
* sample CI logs
* sample Terraform plans

Measure:

* correct tool selection
* matching precision/recall
* false positive rate
* safety compliance (blocked commands never run)

## Regression Testing

* golden sessions
* replay in CI
* compare artifacts + key metrics

## Failure Handling

* bounded retries
* clear error summaries
* stop and ask user when ambiguity persists

---

# 10. Security Model

* secrets redaction in session logs
* never display env vars by default
* no implicit write operations
* explicit flags for mutating actions
* command blocking for destructive operations
* optional “restricted mode” for enterprise environments

---

# 13. Why This Is Architecturally Interesting

xtai demonstrates:

* agent planning and orchestration
* provider-agnostic model routing
* safe command execution systems
* deterministic + generative hybrid design
* observability and reproducibility
* a plugin architecture for evolving capabilities

It is not:

* a generic code assistant
* an LLM wrapper
* a toy terminal chatbot

---

# 14. Positioning

xtai is a safe, observable, tool-first natural-language CLI agent built as a Go binary, with optional Python helpers, and a fully LLM-agnostic provider layer.

It reduces cognitive overhead, automates cross-system reasoning, and serves as a strong showcase of advanced AI + infrastructure engineering.
