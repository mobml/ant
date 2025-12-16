package goal

//wirte validation for goal service

import (
	"errors"
	"strings"

	"github.com/mobml/ant/internal/models"
)

func ValidateGoal(goal *models.Goal) error {
	if goal == nil {
		return errors.New("goal cannot be nil")
	}
	if strings.TrimSpace(goal.Name) == "" {
		return errors.New("goal name cannot be empty")
	}
	if strings.TrimSpace(goal.AreaID) == "" {
		return errors.New("area ID cannot be empty")
	}
	return nil
}
