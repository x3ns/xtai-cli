package domain

// Skill describes a reusable capability that can be invoked by the planner
// or executor. Concrete skills will be added in later phases.
type Skill interface {
	Name() string
	Description() string
}

