package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var habitTodayCmd = &cobra.Command{
	Use:   "today",
	Short: "List habits for today",
	RunE:  runHabitToday,
}

func runHabitToday(cmd *cobra.Command, args []string) error {
	habits, err := HabitService.HabitsForToday()
	if err != nil {
		return fmt.Errorf("failed to get habits for today: %w", err)
	}

	if len(habits) == 0 {
		cmd.Println("No habits found for today.")
		return nil
	}

	cmd.Println("Habits for today:")
	for _, habit := range habits {
		status := "Not completed"
		if habit.WorkedToday {
			status = "Completed"
		}
		cmd.Printf("- %s[%s]: %s [%s]\n", habit.Name, habit.ID, habit.Description, status)
	}

	return nil
}

func init() {
	habitCmd.AddCommand(habitTodayCmd)
}
