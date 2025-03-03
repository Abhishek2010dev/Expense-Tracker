package cmd

import "github.com/spf13/cobra"

func Init() {
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(AddCommand())
	rootCmd.AddCommand(ListCommand())
	rootCmd.AddCommand(SummaryCommand())
	rootCmd.Execute()
}
