package cmd

import (
	"context"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func Init() {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "Expense Tracker",
		Short: "📊 A simple CLI to track your expenses",
		Long:  `Track, list, and manage your daily expenses from the command line.`,
		Example: `
# ➕ Add an expense
go run . add --description "Lunch" --amount 20

# 📋 List all expenses
go run . list

# 📈 View total expense summary
go run . summary

# ❌ Delete an expense
go run . delete --id 2

# 📅 Monthly summary for August
go run . summary --month 8

# 🧪 Mix and match arguments
DESCRIPTION="Dinner" AMOUNT=10 go run . add --description "$DESCRIPTION" --amount $AMOUNT

# 🧵 Multi-line environment setup with execution
ENV_DB_PATH="./db.json" \
  USER_ID=42 \
  go run . summary --month 6
`,
	}
	rootCmd.AddCommand(AddCommand())
	rootCmd.AddCommand(ListCommand())
	rootCmd.AddCommand(SummaryCommand())
	rootCmd.AddCommand(DeleteCommand())

	fang.Execute(context.TODO(), rootCmd)
}
