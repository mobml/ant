package plan

import (
	"time"

	"github.com/mobml/ant/internal/models"
)

func validatePlan(p *models.Plan) error {
	if p.Name == "" {
		return ErrNameRequired
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

func validateID(id string) error {
	if id == "" {
		return ErrIDRequired
	}
	return nil
}
