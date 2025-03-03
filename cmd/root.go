package cmd

import "github.com/spf13/cobra"

func Init() {
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(addExpenseCommand())
	rootCmd.Execute()
}
