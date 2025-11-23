package cmd

import (
	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all candidate frogs",
	Long:  `Clear all tasks from your list of candidate frogs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return db.DeleteAllCandidates()
	},
}
