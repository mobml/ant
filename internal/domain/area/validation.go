package area

import (
	"github.com/mobml/ant/internal/domain/common"
	"github.com/mobml/ant/internal/models"
)

func ValidateArea(a *models.Area) error {
	if a.PlanID == "" {
		return common.ErrIDRequired
	}
	if a.Name == "" {
		return common.ErrNameRequired
	}
	return nil
}
