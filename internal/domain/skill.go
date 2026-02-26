package domain

// Skill describes a reusable capability that can be invoked by the planner
// or executor. Concrete skills are provided by individual plugins.
type Skill interface {
	Name() string
	Description() string
}

