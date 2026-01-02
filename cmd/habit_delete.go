package cmd

import (
	"github.com/spf13/cobra"
)

var habitDeleteCmd = &cobra.Command{
	Use:   "del [habit_id]",
	Short: "Delete a habit by its ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runHabitDeleteCmd,
}

func runHabitDeleteCmd(cmd *cobra.Command, args []string) error {
	habitID := args[0]

	err := HabitService.DeleteHabit(habitID)
	if err != nil {
		return err
	}

	cmd.Printf("Habit with ID %s has been deleted.\n", habitID)
	return nil
}

func init() {
	habitCmd.AddCommand(habitDeleteCmd)
}
