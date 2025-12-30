package wizard

import (
	"github.com/mobml/ant/internal/models"
)

func NeedsGoalWizard(goal *models.Goal) bool {
	return goal.Name == "" || goal.AreaID == ""
}

func GoalWizard(goal *models.Goal) error {
	fields := []Field{
		{
			Label: "Goal name",
			Value: func() string {
				return goal.Name
			},
			SetValue: func(v string) error {
				goal.Name = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Area ID",
			Value: func() string {
				return goal.AreaID
			},
			SetValue: func(v string) error {
				goal.AreaID = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Description",
			Value: func() string {
				return goal.Description
			},
			SetValue: func(v string) error {
				goal.Description = v
				return nil
			},
			Optional: true,
		},
	}
	return Run(fields)
}
