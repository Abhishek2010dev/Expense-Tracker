package cmd

import (
	"Abhishek2010DevSingh/expenseTracker/handler"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func DeleteCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete an expense",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetInt("id")
			if err != nil {
				fmt.Println("Error retrieving id flag:", err)
				os.Exit(1)
			}

			err = handler.DeleteRecord("data.csv", id)
			if err != nil {
				fmt.Println("Error deleting record:", err)
				os.Exit(1)
			}
			fmt.Println("Record deleted successfully")
		},
	}
	command.Flags().IntP("id", "i", 0, "ID of the expense to delete")
	command.MarkFlagRequired("id")
	return command
}
