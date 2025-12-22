package cmd

import (
	"fmt"
	"time"

	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/mobml/ant/internal/models"
	"github.com/spf13/cobra"
)

var planCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new plan",
	RunE:  runPlanCreatecmd,
}

func runPlanCreatecmd(cmd *cobra.Command, args []string) error {
	plan := &models.Plan{}

	if err := bindPlanFlags(cmd, plan); err != nil {
		return err
	}

	if wizard.NeedsPlanWizard(plan) {
		fmt.Println("Some required fields are missing. Launching interactive wizard...")

		return wizard.PlanWizard(plan)
	}

	return PlanService.CreatePlan(plan)
}

func bindPlanFlags(cmd *cobra.Command, plan *models.Plan) error {
	var err error

	if plan.Name, err = cmd.Flags().GetString("name"); err != nil {
		return err
	}

	if plan.Duration, err = cmd.Flags().GetInt("duration"); err != nil {
		return err
	}

	if plan.Description, err = cmd.Flags().GetString("description"); err != nil {
		return err
	}

	start, err := cmd.Flags().GetString("start")
	if err != nil {
		return err
	}
	if start != "" {
		plan.StartDate, err = parseDate(start)
		return err
	}

	return nil
}

func parseDate(value string) (time.Time, error) {
	d, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format (use YYYY-MM-DD)")
	}
	return d, nil
}

func init() {
	planCreateCmd.Flags().StringP("name", "n", "", "Name of the plan")
	planCreateCmd.Flags().IntP("duration", "l", 0, "Duration of the plan in weeks")
	planCreateCmd.Flags().StringP("description", "d", "", "Description of the plan")
	planCreateCmd.Flags().StringP("start", "s", "", "Start date of the plan (YYYY-MM-DD)")
}
