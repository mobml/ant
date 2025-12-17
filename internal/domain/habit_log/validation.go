package habitlog

import (
	"math"
	"strings"

	"github.com/mobml/ant/internal/domain/common"
	"github.com/mobml/ant/internal/models"
)

func ValidateHabitLog(h *models.HabitLog) error {
	if strings.TrimSpace(h.HabitID) == "" {
		return common.ErrIDRequired
	}
	if h.LogDate.IsZero() {
		return ErrLogDateRequired
	}

	if !isValidFloat(h.Value) {
		return ErrValueRequired
	}

	return nil
}

func isValidFloat(f float64) bool {
	return !math.IsNaN(f) && !math.IsInf(f, 0)
}
