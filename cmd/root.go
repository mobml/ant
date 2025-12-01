package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ant",
	Short: "A CLI tool for personal productivity and growth written in Go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Ant! \nuse --help to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
