package habit

import (
	"strings"

	"github.com/mobml/ant/internal/domain/common"
	"github.com/mobml/ant/internal/models"
)

func ValidateHabit(h *models.Habit) error {
	if strings.TrimSpace(h.Name) == "" {
		return common.ErrNameRequired
	}
	if !isValidMeasureType(h.MeasureType) {
		return ErrInvalidMeasureType
	}

	if strings.TrimSpace(h.GoalID) == "" {
		return ErrGoalIDRequired
	}

	return nil
}

func isValidMeasureType(measureType models.MeasureType) bool {
	switch measureType {
	case models.MeasureBoolean, models.MeasureInteger, models.MeasureFloat,
		models.MeasureHours, models.MeasureDistance, models.MeasureCustom:
		return true
	default:
		return false
	}
}
