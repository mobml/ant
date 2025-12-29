package cmd

import (
	"github.com/spf13/cobra"
)

var areaCmd = &cobra.Command{
	Use:   "area",
	Short: "Manage areas",
}

func init() {
	rootCmd.AddCommand(areaCmd)
}
