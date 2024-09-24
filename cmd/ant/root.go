package ant

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ant",
		Short: "Ant a minimalist task manager",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Welcome to Ant!")
		},
	}

	cmd.AddCommand(AddTaskCommand())
	cmd.AddCommand(RemoveTaskCommand())
	cmd.AddCommand(ListTaskCommand())
	return cmd
}
