package cmd

import (
	"github.com/mobml/ant/database"
	"github.com/mobml/ant/internal/repositories"

	pa "github.com/mobml/ant/internal/services/area"
	ps "github.com/mobml/ant/internal/services/plan"
)

var (
	PlanService ps.PlanService
	AreaService pa.AreaService
)

func InitServices() {
	db := database.DB()

	PlanService = ps.NewPlanService(
		repositories.NewPlanRepository(db),
	)

	AreaService = pa.NewAreaService(
		repositories.NewAreaRepository(db),
	)
}
