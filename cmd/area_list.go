package cmd

import (
	"github.com/spf13/cobra"
)

var areaListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all areas",
	Args:  cobra.ExactArgs(1),
	RunE:  runAreaListCmd,
}

func runAreaListCmd(cmd *cobra.Command, args []string) error {
	areas, err := AreaService.ListAreasByPlan(args[0])
	if err != nil {
		return err
	}

	for _, area := range areas {
		cmd.Println(area.ID, area.Name, area.Description)
	}

	return nil
}

func init() {
	areaCmd.AddCommand(areaListCmd)
}
