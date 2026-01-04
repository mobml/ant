package cmd

import (
	"github.com/mobml/ant/database"
	"github.com/mobml/ant/internal/repositories"

	as "github.com/mobml/ant/internal/services/area"
	gs "github.com/mobml/ant/internal/services/goal"
	hs "github.com/mobml/ant/internal/services/habit"
	hls "github.com/mobml/ant/internal/services/habit_log"
	ps "github.com/mobml/ant/internal/services/plan"
)

var (
	PlanService     ps.PlanService
	AreaService     as.AreaService
	GoalService     gs.GoalService
	HabitService    hs.HabitService
	HabitLogService hls.HabitLogService
)

func InitServices() {
	db := database.DB()

	PlanService = ps.NewPlanService(
		repositories.NewPlanRepository(db),
	)

	AreaService = as.NewAreaService(
		repositories.NewAreaRepository(db),
	)

	GoalService = gs.NewGoalService(
		repositories.NewGoalRepository(db),
	)

	HabitService = hs.NewHabitService(
		repositories.NewHabitRepository(db),
	)

	HabitLogService = hls.NewHabitLogService(
		repositories.NewHabitLogRepository(db),
	)
}
