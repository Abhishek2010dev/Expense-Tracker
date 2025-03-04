package cmd

import (
	"Abhishek2010DevSingh/expenseTracker/utils"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func SummaryCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "summary",
		Short: "Show a summary of expenses",
		Run: func(cmd *cobra.Command, args []string) {
			record, err := utils.Read("data.csv")
			if err != nil {
				fmt.Println("Error reading expenses:", err)
				os.Exit(1)
			}

			month, err := cmd.Flags().GetInt16("month")
			if err != nil {
				fmt.Println("Error retrieving month flag:", err)
				os.Exit(1)
			}

			// Filter expenses by month
			if month != 0 {
				if month < 0 || month > 12 {
					fmt.Println("Invalid month")
					os.Exit(1)
				}
				summary := 0.0
				for _, record := range record {
					if int16(record.Date.Month()) == month {
						summary += record.Amount
					}
				}
				fmt.Printf("Total expenses in %s: $%.2f\n", time.Month(month), summary)

				return
			}

			summary := 0.0
			for _, record := range record {
				summary += record.Amount
			}
			fmt.Printf("Total expenses: $%.2f\n", summary)
		},
	}

	command.Flags().Int16P("month", "m", 0, "Filter expenses by date")

	return command
}
