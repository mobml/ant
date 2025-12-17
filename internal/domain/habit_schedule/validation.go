package habitschedule

import (
	"strings"

	"github.com/mobml/ant/internal/domain/common"
	"github.com/mobml/ant/internal/models"
)

func ValidateHabitSchedule(hs *models.HabitSchedule) error {
	if strings.TrimSpace(hs.HabitID) == "" {
		return common.ErrIDRequired
	}
	if hs.DayOfWeek < 0 || hs.DayOfWeek > 6 {
		return ErrInvalidDayOfWeek
	}

	return nil
}
