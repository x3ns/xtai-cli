package usecase

import (
	"context"

	"github.com/x3ns/xtai-cli/internal/domain"
)

// PlannerService coordinates building plans using a domain.Planner and other
// dependencies. It lives in the application layer and is wired from the
// delivery/CLI layer.
type PlannerService struct {
	planner domain.Planner
}

// NewPlannerService constructs a PlannerService.
func NewPlannerService(planner domain.Planner) *PlannerService {
	return &PlannerService{
		planner: planner,
	}
}

// CreatePlan delegates to the underlying domain.Planner. Additional
// orchestration responsibilities will be added in later phases.
func (s *PlannerService) CreatePlan(ctx context.Context, req domain.PlanRequest) (domain.Plan, error) {
	return s.planner.CreatePlan(ctx, req)
}

