package wizard

import (
	"fmt"
	"strconv"

	"github.com/mobml/ant/internal/models"
)

func HabitLogWizard(habitLog *models.HabitLog) error {
	fields := []Field{
		{
			Label: "Habit ID",
			Value: func() string {
				return habitLog.HabitID
			},
			SetValue: func(v string) error {
				habitLog.HabitID = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Value",
			Value: func() string {
				return fmt.Sprintf("%f", habitLog.Value)
			},
			SetValue: func(v string) error {
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				habitLog.Value = val
				return nil
			},
			Optional: false,
		},
		{
			Label: "Note",
			Value: func() string {
				return habitLog.Note
			},
			SetValue: func(v string) error {
				habitLog.Note = v
				return nil
			},
			Optional: true,
		},
	}

	return Run(fields)
}

func NeedsHabitLogWizard(habitLog *models.HabitLog) bool {
	return habitLog.ID == "" || habitLog.Value == 0
}
