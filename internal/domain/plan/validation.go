package plan

import (
	"time"

	"github.com/mobml/ant/internal/domain/common"
	"github.com/mobml/ant/internal/models"
)

func ValidatePlan(p *models.Plan) error {
	if p.Name == "" {
		return common.ErrIDRequired
	}
	if p.Duration <= 0 {
		return ErrDurationInvalid
	}
	if p.StartDate.IsZero() {
		return ErrStartDateInvalid
	}
	if p.StartDate.Before(time.Now()) {
		return ErrStartDatePast
	}

	return nil
}
