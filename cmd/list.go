package cmd

import (
	"Abhishek2010DevSingh/expenseTracker/utils"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		Run: func(cmd *cobra.Command, args []string) {
			records, err := utils.Read("data.csv")
			if err != nil {
				fmt.Println("Error reading expenses:", err)
				return
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tDate\tDescription\tAmount")
			for _, record := range records {
				date := record.Date.Format("2006-01-02")
				fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\n", record.ID, date, record.Description, record.Amount)
			}
			w.Flush()
		},
	}
}
