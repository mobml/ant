package cmd

import (
	"github.com/spf13/cobra"
)

var goalListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all goals",
	Args:  cobra.ExactArgs(1),
	RunE:  runGoalListCmd,
}

func runGoalListCmd(cmd *cobra.Command, args []string) error {
	goals, err := GoalService.ListGoalsByArea(args[0])
	if err != nil {
		return err
	}

	if len(goals) == 0 {
		cmd.Println("No goals found.")
		return nil
	}

	for _, goal := range goals {
		cmd.Printf("ID: %s\nName: %s\nArea ID: %s\nDescription: %s\n\n", goal.ID, goal.Name, goal.AreaID, goal.Description)
	}

	return nil
}

func init() {
	goalCmd.AddCommand(goalListCmd)
}
