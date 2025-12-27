package cmd

import "github.com/spf13/cobra"

var planDeleteCmd = &cobra.Command{
	Use:   "del [plan ID]",
	Short: "Delete a plan by its ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runPlanDeleteCmd,
}

func runPlanDeleteCmd(cmd *cobra.Command, args []string) error {
	planID := args[0]
	return PlanService.DeletePlan(planID)
}

func init() {
	planCmd.AddCommand(planDeleteCmd)
}
