/*
Package cmd implements the command-line interface for the Eat That Frog CLI.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "frog",
	Short: "Eat That Frog - Daily task prioritization CLI",
	Long:  `A minimal CLI to help you eat that frog every day 🐸`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
