package plan

import (
	dc "github.com/mobml/ant/internal/domain/common"
	dp "github.com/mobml/ant/internal/domain/plan"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type PlanService interface {
	CreatePlan(plan *models.Plan) error
	ListPlans() ([]*models.Plan, error)
	UpdatePlan(plan *models.Plan) error
	DeletePlan(id string) error
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

	if err := dp.ValidatePlan(plan); err != nil {
		return err
	}

	return ps.planRepo.Create(plan)
}

func (ps *planService) ListPlans() ([]*models.Plan, error) {
	return ps.planRepo.List()
}

func (ps *planService) UpdatePlan(plan *models.Plan) error {

	if err := dp.ValidatePlan(plan); err != nil {
		return err
	}

	if err := dc.ValidateID(plan.ID); err != nil {
		return err
	}

	return ps.planRepo.Update(plan)
}

func (ps *planService) DeletePlan(id string) error {

	if err := dc.ValidateID(id); err != nil {
		return err
	}
	return ps.planRepo.Delete(id)
}
