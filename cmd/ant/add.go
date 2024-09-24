package ant

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddTaskCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "add a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Task '%s' added\n", args[0])
		},
	}
}
