package cmd

import (
	"fmt"
	"github.com/mobml/ant/tasks"
	"github.com/spf13/cobra"
	"strconv"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Complete a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Please provide a valid id")
		}
		tasks.CompleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
