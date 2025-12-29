package cmd

import (
	"fmt"

	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/spf13/cobra"
)

var planUpdateCmd = &cobra.Command{
	Use:   "update [plan ID]",
	Short: "Update a plan by its ID",
	RunE:  runPlanUpdateCmd,
}

func runPlanUpdateCmd(cmd *cobra.Command, args []string) error {
	planID := args[0]

	plan, err := PlanService.GetPlanByID(planID)
	if err != nil {
		return err
	}

	if err := bindPlanFlags(cmd, plan); err != nil {
		return err
	}

	if wizard.NeedsPlanWizard(plan) {
		fmt.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.PlanWizard(plan); err != nil {
			return err
		}
	}

	return PlanService.UpdatePlan(plan)
}

func init() {
	planCmd.AddCommand(planUpdateCmd)
	planUpdateCmd.Flags().StringP("name", "n", "", "Name of the plan")
	planUpdateCmd.Flags().IntP("duration", "l", 0, "Duration of the plan in weeks")
	planUpdateCmd.Flags().StringP("description", "d", "", "Description of the plan")
	planUpdateCmd.Flags().StringP("start", "s", "", "Start date of the plan (YYYY-MM-DD)")
}
