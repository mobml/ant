package cmd

import (
	"fmt"

	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/mobml/ant/internal/models"
	"github.com/spf13/cobra"
)

var habitMarkCmd = &cobra.Command{
	Use:   "mark [habit id]",
	Short: "Mark a habit as completed for today",
	Args:  cobra.ExactArgs(1),
	RunE:  runHabitMark,
}

func runHabitMark(cmd *cobra.Command, args []string) error {
	habitLog := models.HabitLog{}
	habitLog.HabitID = args[0]

	if err := bindHabitLogFlags(cmd, &habitLog); err != nil {
		return err
	}

	fmt.Println("Creating habit log for Habit ID:", habitLog.HabitID)

	if wizard.NeedsHabitLogWizard(&habitLog) {
		cmd.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.HabitLogWizard(&habitLog); err != nil {
			return err
		}
	}

	return HabitLogService.CreateHabitLog(&habitLog)
}

func bindHabitLogFlags(cmd *cobra.Command, habitLog *models.HabitLog) error {
	var err error

	if habitLog.Value, err = cmd.Flags().GetFloat64("value"); err != nil {
		return err
	}

	if habitLog.Note, err = cmd.Flags().GetString("note"); err != nil {
		return err
	}

	return nil
}

func init() {
	habitCmd.AddCommand(habitMarkCmd)
	habitMarkCmd.Flags().Float64("value", 0, "Value associated with the habit log (required)")
	habitMarkCmd.Flags().String("note", "", "Optional note for the habit log")
}
