package cmd

import (
	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/mobml/ant/internal/models"
	"github.com/spf13/cobra"
)

var goalCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new goal",
	RunE:  runGoalCreateCmd,
}

func runGoalCreateCmd(cmd *cobra.Command, args []string) error {
	goal := &models.Goal{}

	if err := bindGoalFlags(cmd, goal); err != nil {
		return err
	}

	if wizard.NeedsGoalWizard(goal) {
		cmd.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.GoalWizard(goal); err != nil {
			return err
		}
	}

	return GoalService.CreateGoal(goal)
}

func bindGoalFlags(cmd *cobra.Command, goal *models.Goal) error {
	var err error

	if goal.Name, err = cmd.Flags().GetString("name"); err != nil {
		return err
	}

	if goal.AreaID, err = cmd.Flags().GetString("area-id"); err != nil {
		return err
	}

	if goal.Description, err = cmd.Flags().GetString("description"); err != nil {
		return err
	}
	return nil
}

func init() {
	goalCmd.AddCommand(goalCreateCmd)
	goalCreateCmd.Flags().StringP("name", "n", "", "Name of the goal")
	goalCreateCmd.Flags().StringP("area-id", "a", "", "ID of the associated area")
	goalCreateCmd.Flags().StringP("description", "d", "", "Description of the goal (optional)")
}
