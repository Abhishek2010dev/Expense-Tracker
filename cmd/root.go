package cmd

import "github.com/spf13/cobra"

func Init() {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(AddCommand())
	rootCmd.AddCommand(ListCommand())
	rootCmd.AddCommand(SummaryCommand())
	rootCmd.AddCommand(DeleteCommand())
	rootCmd.Execute()
}
