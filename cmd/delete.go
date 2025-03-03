package cmd

import "github.com/spf13/cobra"

func DeleteCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete an expense",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return command
}
