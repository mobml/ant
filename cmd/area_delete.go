package cmd

import (
	"github.com/spf13/cobra"
)

var areaDeleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an area",
	Args:  cobra.ExactArgs(1),
	RunE:  runAreaDeleteCmd,
}

func runAreaDeleteCmd(cmd *cobra.Command, args []string) error {
	areaID, err := cmd.Flags().GetString("id")
	if err != nil {
		return err
	}

	return AreaService.DeleteArea(areaID)
}

func init() {
	areaCmd.AddCommand(areaDeleteCmd)
}
