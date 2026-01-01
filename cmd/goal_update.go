package cmd

import (
	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/spf13/cobra"
)

var goalUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a goal",
	Args:  cobra.ExactArgs(1),
	RunE:  runGoalUpdateCmd,
}

func runGoalUpdateCmd(cmd *cobra.Command, args []string) error {
	goalID := args[0]

	goal, err := GoalService.GetGoalByID(goalID)

	if err != nil {
		return err
	}

	if err := bindGoalFlags(cmd, goal); err != nil {
		return err
	}

	if wizard.NeedsGoalWizard(goal) {
		cmd.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.GoalWizard(goal); err != nil {
			return err
		}
	}

	return GoalService.UpdateGoal(goal)
}

func init() {
	goalCmd.AddCommand(goalUpdateCmd)
	goalUpdateCmd.Flags().StringP("name", "n", "", "Name of the goal")
	goalUpdateCmd.Flags().StringP("description", "d", "", "Description of the goal (optional)")
	goalUpdateCmd.Flags().StringP("area-id", "a", "", "ID of the associated area")
}
