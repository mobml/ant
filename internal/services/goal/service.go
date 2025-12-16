package goal

import (
	dc "github.com/mobml/ant/internal/domain/common"
	dg "github.com/mobml/ant/internal/domain/goal"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type GoalService interface {
	CreateGoal(goal *models.Goal) error
	ListGoalsByArea(areaID string) ([]*models.Goal, error)
	UpdateGoal(goal *models.Goal) error
	DeleteGoal(id string) error
}

type goalService struct {
	goalRepo repo.GoalRepository
}

func NewGoalService(goalRepo repo.GoalRepository) GoalService {
	return &goalService{
		goalRepo: goalRepo,
	}
}

func (s *goalService) CreateGoal(goal *models.Goal) error {
	if err := dg.ValidateGoal(goal); err != nil {
		return err
	}
	return s.goalRepo.Create(goal)
}

func (s *goalService) ListGoalsByArea(areaID string) ([]*models.Goal, error) {
	if err := dc.ValidateID(areaID); err != nil {
		return nil, err
	}
	return s.goalRepo.ListByArea(areaID)
}

func (s *goalService) UpdateGoal(goal *models.Goal) error {
	if err := dg.ValidateGoal(goal); err != nil {
		return err
	}
	if err := dc.ValidateID(goal.ID); err != nil {
		return err
	}
	return s.goalRepo.Update(goal)
}

func (s *goalService) DeleteGoal(id string) error {
	if err := dc.ValidateID(id); err != nil {
		return err
	}
	return s.goalRepo.Delete(id)
}
