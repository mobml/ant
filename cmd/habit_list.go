package cmd

import (
	"github.com/spf13/cobra"
)

var habitListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all habits",
	RunE:  runHabitListCmd,
}

func runHabitListCmd(cmd *cobra.Command, args []string) error {
	habits, err := HabitService.ListHabits()
	if err != nil {
		return err
	}

	if len(habits) == 0 {
		cmd.Println("No habits found.")
		return nil
	}

	for _, habit := range habits {
		cmd.Printf(
			"ID: %s\n"+
				"Goal ID: %s\n"+
				"Name: %s\n"+
				"Description: %s\n"+
				"Measure Type: %s\n"+
				"Measure Unit: %s\n"+
				"Created At: %s\n"+
				"Updated At: %s\n\n",
			habit.ID,
			habit.GoalID,
			habit.Name,
			habit.Description,
			habit.MeasureType,
			habit.MeasureUnit,
			habit.CreatedAt,
			habit.UpdatedAt,
		)
	}

	return nil
}

func init() {
	habitCmd.AddCommand(habitListCmd)
}
