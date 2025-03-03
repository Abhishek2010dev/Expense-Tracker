package cmd

import "github.com/spf13/cobra"

func Init() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(addExpenseCommand())
	rootCmd.Execute()
}
