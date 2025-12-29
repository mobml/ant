package cmd

import (
	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/spf13/cobra"
)

var areaUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an area",
	Args:  cobra.ExactArgs(1),
	RunE:  runAreaUpdateCmd,
}

func runAreaUpdateCmd(cmd *cobra.Command, args []string) error {
	areaID := args[0]

	area, err := AreaService.GetAreaByID(areaID)

	if err != nil {
		return err
	}

	if err := bindAreaFlags(cmd, area); err != nil {
		return err
	}

	if wizard.NeedsAreaWizard(area) {
		cmd.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.AreaWizard(area); err != nil {
			return err
		}
	}

	return AreaService.UpdateArea(area)
}

func init() {
	areaCmd.AddCommand(areaUpdateCmd)
	areaUpdateCmd.Flags().StringP("name", "n", "", "Name of the area")
	areaUpdateCmd.Flags().StringP("description", "d", "", "Description of the area (optional)")
	areaUpdateCmd.Flags().StringP("plan-id", "p", "", "ID of the associated plan")
}
