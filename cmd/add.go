package cmd

import (
	"Abhishek2010DevSingh/expenseTracker/handler"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	addExpenseCommand := &cobra.Command{
		Use:   "add",
		Short: "Add a new expense",
		Run: func(cmd *cobra.Command, args []string) {
			amount, err := cmd.Flags().GetFloat64("amount")
			if err != nil {
				fmt.Println("Error retrieving amount flag:", err)
				os.Exit(1)
			}

			description, err := cmd.Flags().GetString("description")
			if err != nil {
				fmt.Println("Error retrieving description flag:", err)
				os.Exit(1)
			}

			handler.WriteCsv(description, amount)
		},
	}

	addExpenseCommand.Flags().Float64P("amount", "a", 0.0, "Expense amount (required)")
	addExpenseCommand.Flags().StringP("description", "d", "", "Expense description (required)")
	addExpenseCommand.MarkFlagRequired("amount")
	addExpenseCommand.MarkFlagRequired("description")

	return addExpenseCommand
}
