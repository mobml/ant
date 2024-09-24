package ant

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RemoveTaskCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Short: "Remove a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStderr(), "Task '%s' removed\n", args[0])
		},
	}
}
