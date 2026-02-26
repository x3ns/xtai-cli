package domain

import "context"

// Plan represents a structured execution plan produced by the planner.
// The concrete shape may evolve over time, but this placeholder captures the
// core concept for Clean Architecture boundaries.
type Plan struct {
	ID          string
	Description string
	Steps       []PlanStep
}

// PlanStep describes a single step within a Plan.
type PlanStep struct {
	ID         string
	Summary    string
	Command    string
	AllowWrite bool
}

// PlanRequest captures the inputs required to build a Plan from a natural
// language goal and environment context.
type PlanRequest struct {
	Goal   string
	DryRun bool
}

// Planner defines the behavior required to turn a natural-language goal into
// a structured Plan. Concrete implementations live in usecase/infrastructure
// layers and depend on LLM providers via the Provider interface.
type Planner interface {
	CreatePlan(ctx context.Context, request PlanRequest) (Plan, error)
}

