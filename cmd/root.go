package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// Flag definitions for global access throughout application
	rootCmd.AddCommand(samazonServerCmd)
}

var rootCmd = &cobra.Command{
	Use:              "samazon",
	PersistentPreRun: initRoot,
	Short:            "samazon handles the simulation of an Amazon/UPS type delivery system",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initRoot(cmd *cobra.Command, args []string) {
	// initRoot is called before any command gets executed
}
