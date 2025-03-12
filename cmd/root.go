package cmd

import (
	"fmt"
	"github.com/mobml/ant/tasks"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ant",
	Short: "A CLI task manager tool written in Go",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := tasks.LoadTasks()

		if err != nil {
			fmt.Println("Error loading tasks: ", err)
		}
	},
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
