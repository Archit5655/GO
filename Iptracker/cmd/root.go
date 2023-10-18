package cmd

import (

	"github.com/spf13/cobra"

)

var (
	// Used for flags.


	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IPtracker CLI app",
		Long: `IPtracking.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
