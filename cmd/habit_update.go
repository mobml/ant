package cmd

import (
	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/spf13/cobra"
)

var habitUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a habit",
	Args:  cobra.ExactArgs(1),
	RunE:  runHabitUpdateCmd,
}

func runHabitUpdateCmd(cmd *cobra.Command, args []string) error {
	habitID := args[0]

	habit, err := HabitService.GetHabitByID(habitID)

	if err != nil {
		return err
	}

	var daysRaw string

	if err := bindHabitFlags(cmd, habit, &daysRaw); err != nil {
		return err
	}

	if wizard.NeedsHabitWizard(habit, &daysRaw) {
		cmd.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.HabitWizard(habit, &daysRaw); err != nil {
			return err
		}
	}

	days, err := parseDays(daysRaw)
	if err != nil {
		return err
	}

	return HabitService.UpdateHabitWithSchedule(habit, days)
}

func init() {
	habitCmd.AddCommand(habitUpdateCmd)
	habitUpdateCmd.Flags().StringP("name", "n", "", "Name of the habit")
	habitUpdateCmd.Flags().StringP("description", "d", "", "Description of the habit (optional)")
	habitUpdateCmd.Flags().StringP("goal-id", "g", "", "ID of the associated goal")
	habitUpdateCmd.Flags().StringP("days", "", "", "Comma-separated list of days of the week for the habit schedule (0=Sunday, 6=Saturday)")
}
