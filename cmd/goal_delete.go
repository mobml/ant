package cmd

import (
	"github.com/spf13/cobra"
)

var goalDeleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a goal by ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runGoalDeleteCmd,
}

func runGoalDeleteCmd(cmd *cobra.Command, args []string) error {
	goalID := args[0]

	return GoalService.DeleteGoal(goalID)
}

func init() {
	goalCmd.AddCommand(goalDeleteCmd)
}
