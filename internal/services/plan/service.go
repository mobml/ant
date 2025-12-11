package plan

import (
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type PlanService interface {
	CreatePlan(plan *models.Plan) error
}

type planService struct {
	planRepo repo.PlanRepository
}

func NewPlanService(planRepo repo.PlanRepository) PlanService {
	return &planService{
		planRepo: planRepo,
	}
}

func (ps *planService) CreatePlan(plan *models.Plan) error {

	if err := validatePlan(plan); err != nil {
		return err
	}

	return ps.planRepo.Create(plan)
}

func (ps *planService) ListPlans() ([]*models.Plan, error) {
	return ps.planRepo.List()
}

func (ps *planService) UpdatePlan(plan *models.Plan) error {

	if err := validatePlan(plan); err != nil {
		return err
	}

	if err := validateID(plan.ID); err != nil {
		return err
	}

	return ps.planRepo.Update(plan)
}

func (ps *planService) DeletePlan(id string) error {

	if err := validateID(id); err != nil {
		return err
	}
	return ps.planRepo.Delete(id)
}
