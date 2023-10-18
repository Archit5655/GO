package cmd

import (

	"fmt"
	
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getversion)
}


var getversion = &cobra.Command{
	Use:   "version",
	Short: "trace ip",
	Long:  `trace ip`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.1")
	},
}

