# AGENT.md — Engineering & Architecture Guidelines for xtai

This file contains the core engineering principles and style guidelines that must be referenced when generating, refactoring, reviewing, or evaluating code for **xtai** (the natural language CLI agent tool).

---

## 1. Clean Code Principles

The codebase **must follow Clean Code guidelines** — code written here is intended for long-term maintenance, readability, and collaboration.

### Fundamental Clean Code Rules

1. **Separation of Concerns (SoC)**  
   Each module, package, or function should settle one concern and not mix unrelated logic.  
   Changes in one area should not ripple into unrelated parts of the system.

2. **Descriptive Naming**  
   Variables, functions, packages, and modules must use meaningful names that describe intent.

3. **No Duplication**  
   Avoid repeated logic. Abstract shared behavior or logic into a single place.

4. **Keep It Simple and Focused (KISS)**  
   Favor simple solutions over clever ones. Do not over-engineer.

5. **Document for Humans**  
   Comments and documentation must explain *why*, not *what*.  
   Code should be primarily self-descriptive.

6. **Test-First Mindset**  
   Write tests early and maintain them. Tests should cover:
   - unit behavior
   - component boundaries
   - expected failures  
   Tests serve as *executable documentation*.

7. **You Ain’t Gonna Need It (YAGNI)**  
   Implement only what is required now, not speculative features.

---

## 2. Clean Architecture Principles

xtai must follow **Clean Architecture**, where dependencies always point *inward* and business logic is isolated from external concerns (frameworks, UI, databases, CLI, service providers).

### Core Rules

1. **Dependency Rule**  
   Source code dependencies should only point toward the domain/business logic.  
   External adapters should depend on interfaces defined in core packages.

2. **Layer Separation**  
   Organize the project into conceptual layers:
   - **Domain (Entities, Business Rules)**
   - **Usecases/Services (Application Logic)**
   - **Adapters/Interfaces (External APIs/Clients)**
   - **Delivery/UI (CLI entrypoints)**

3. **No Framework Coupling**  
   The architecture should not depend on frameworks or libraries at the core. They belong only in outer layers.

4. **Testable Business Logic**  
   Domain and usecases must be testable without running infrastructure or CLI.

---

## 3. Service-Based Architecture Guidelines

Although xtai is not a distributed system, we still adopt **service-based principles** for modularity, reuse, and composability.

### Guidelines

1. **Service Granularity**  
   Services (packages/modules) should encapsulate *single, cohesive responsibilities*.  
   Avoid over-granular services that increase complexity without benefit.

2. **Service Reusability**  
   Design modules so that they can be reused across different workflows and plugins.  
   Logic should not be tied to a specific command or workflow.

3. **Interface-Driven Services**  
   Services expose only minimal and necessary interfaces; consumers should depend on abstractions (interfaces) not concrete types.

4. **Loose Coupling**  
   Components should communicate through well-defined contracts (interfaces), not through direct calls to implementation.

5. **Stateless Where Possible**  
   Design services, especially core logic, to avoid shared mutable state to improve testability and predictability.

---

## 4. SOLID Design Principles

Apply SOLID principles in all object and package design:

- **S**ingle Responsibility — one reason to change.
- **O**pen/Closed — open for extension, closed for modification.
- **L**iskov Substitution — subtypes should be usable in place of base types.
- **I**nterface Segregation — small, specific interfaces rather than large ones.
- **D**ependency Inversion — depend on abstractions, not concretions.

---

## 5. Error Handling & Logging

- All errors must be explicitly handled with context.
- CLI errors must be user-friendly, actionable, and not leak internal state.
- Logs must include traceable facts without secrets.

---

## 6. Testing & Quality

- Unit tests for every module, especially service and domain layers.
- Integration tests for workflows (planner → executor → session store).
- Fixtures for repeatable test data.
- CI must fail on regressions and enforce tests.

---

## 7. Prompting Best Practices (for Agent Code Generation)

When generating code with an LLM:

- Provide **project-details.md** + **AGENT.md** in context.
- Ask for explicit:
  - folder structure
  - file headers
  - comments and docstrings
  - SOLID compliance
  - Clean Architecture layering
- Avoid ambiguous prompts; specify examples and acceptance criteria.

---

## 8. Naming & Style

- Use idiomatic Go naming conventions.
- Package names should reflect role, not implementation detail.
- Functions must be short, cohesive, and single-purpose.

---

## 9. Versioning & Compatibility

- Use semantic versioning.
- Maintain backward compatibility in public interfaces.
- Prefer to use Go 1.26

---

## 10. Documentation

- Each module must include README or inline docs explaining intent and design rationale.
- Design decisions must be documented where architecture deviates from standard patterns.
- Functions should be self-explanatory and not require comments unless really needed to explain complex tasks