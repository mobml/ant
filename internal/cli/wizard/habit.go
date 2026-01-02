package wizard

import (
	"github.com/mobml/ant/internal/models"
)

func HabitWizard(habit *models.Habit, days *string) error {
	fields := []Field{
		{
			Label: "Habit name",
			Value: func() string {
				return habit.Name
			},
			SetValue: func(v string) error {
				habit.Name = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Description",
			Value: func() string {
				return habit.Description
			},
			SetValue: func(v string) error {
				habit.Description = v
				return nil
			},
			Optional: true,
		},
		{
			Label: "Goal ID",
			Value: func() string {
				return habit.GoalID
			},
			SetValue: func(v string) error {
				habit.GoalID = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Measure Type (e.g., count, duration)",
			Value: func() string {
				return string(habit.MeasureType)
			},
			SetValue: func(v string) error {
				habit.MeasureType = models.MeasureType(v)
				return nil
			},
			Optional: false,
		},
		{
			Label: "Days",
			Value: func() string {
				return *days
			},
			SetValue: func(v string) error {
				*days = v
				return nil
			},
			Optional: false,
		},
	}

	return Run(fields)
}

func NeedsHabitWizard(habit *models.Habit, days *string) bool {
	return habit.Name == "" || *days == "" || habit.MeasureType == "" || habit.GoalID == ""
}
