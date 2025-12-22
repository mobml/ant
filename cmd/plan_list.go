package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var planListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all plans",
	RunE:  runPlanListcmd,
}

func runPlanListcmd(cmd *cobra.Command, args []string) error {
	plans, err := PlanService.ListPlans()
	if err != nil {
		return err
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	fmt.Println("|    ID    |    Name     |  Start Date |   Duration   |     Description    |     Created At    |    Updated At   ")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	for _, p := range plans {
		startDate := "N/A"
		if !p.StartDate.IsZero() {
			startDate = p.StartDate.Format("2006-01-02")
		}
		fmt.Printf("| %s | %s  | %s | %d weeks | %s | %s | %s |\n",
			p.ID, p.Name, startDate, p.Duration, p.Description, p.CreatedAt, p.UpdatedAt)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}

	return nil
}

func init() {
	planCmd.AddCommand(planListCmd)
}
