package cmd

import (
	"fmt"

	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/mobml/ant/internal/models"
	"github.com/spf13/cobra"
)

var areaCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new area",
	RunE:  runAreaCreateCmd,
}

func runAreaCreateCmd(cmd *cobra.Command, args []string) error {
	area := &models.Area{}

	if err := bindAreaFlags(cmd, area); err != nil {
		return err
	}

	if wizard.NeedsAreaWizard(area) {
		fmt.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.AreaWizard(area); err != nil {
			return err
		}
	}

	return AreaService.CreateArea(area)
}

func bindAreaFlags(cmd *cobra.Command, area *models.Area) error {
	var err error

	if area.Name, err = cmd.Flags().GetString("name"); err != nil {
		return err
	}

	if area.PlanID, err = cmd.Flags().GetString("plan-id"); err != nil {
		return err
	}

	if area.Description, err = cmd.Flags().GetString("description"); err != nil {
		return err
	}
	return nil
}

func init() {
	areaCmd.AddCommand(areaCreateCmd)
	areaCreateCmd.Flags().StringP("name", "n", "", "Name of the area")
	areaCreateCmd.Flags().StringP("plan-id", "p", "", "ID of the associated plan")
	areaCreateCmd.Flags().StringP("description", "d", "", "Description of the area (optional)")
}
