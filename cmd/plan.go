package cmd

import (
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Manage plans",
}

func init() {
	rootCmd.AddCommand(planCmd)
	planCmd.AddCommand(planCreateCmd)
}
