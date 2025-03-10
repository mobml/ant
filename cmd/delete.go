package cmd

import (
	"fmt"
	"strconv"

	"github.com/mobml/ant/tasks"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [id]",
	Short: "Remove a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Please provide a valid ID")
			return
		}
		tasks.DeleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
