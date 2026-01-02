package cmd

import (
	"github.com/spf13/cobra"
)

var habitCmd = &cobra.Command{
	Use:   "habit",
	Short: "Manage habits",
}

func init() {
	rootCmd.AddCommand(habitCmd)
}
