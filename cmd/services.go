package cmd

import (
	"github.com/mobml/ant/database"
	"github.com/mobml/ant/internal/repositories"

	ps "github.com/mobml/ant/internal/services/plan"
)

var (
	PlanService ps.PlanService
)

func InitServices() {
	db := database.DB()

	PlanService = ps.NewPlanService(
		repositories.NewPlanRepository(db),
	)
}
