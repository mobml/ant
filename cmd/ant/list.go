package ant

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ListTaskCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "Lists all tasks",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "All tasks here")
		},
	}
}
