package cmd

import (
	"Abhishek2010DevSingh/expenseTracker/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func SummaryCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "summary",
		Short: "Show a summary of expenses",
		Run: func(cmd *cobra.Command, args []string) {
			record, err := utils.Read("data.csv")
			if err != nil {
				fmt.Println("Error reading expenses:", err)
				os.Exit(1)
			}
			summary := 0.0
			for _, record := range record {
				summary += record.Amount
			}
			fmt.Printf("Total expenses: $%.2f\n", summary)
		},
	}
}
