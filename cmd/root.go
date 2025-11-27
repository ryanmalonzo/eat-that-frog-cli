/*
Package cmd implements the command-line interface for the Eat That Frog CLI.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
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

func initCommands() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(pickCmd)
	rootCmd.AddCommand(eatCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(skipCmd)
	rootCmd.AddCommand(todayCmd)
	rootCmd.AddCommand(clearCmd)
}

func Execute() {
	// Initialize the database
	if err := db.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "Error initializing database:", err)
		os.Exit(1)
	}
	defer db.Close()

	initCommands()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
