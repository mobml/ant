package cmd

import (
	"github.com/spf13/cobra"
)

var goalCmd = &cobra.Command{
	Use:   "goal",
	Short: "Manage goals",
}

func init() {
	rootCmd.AddCommand(goalCmd)
}
