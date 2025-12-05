package cmd

import (
	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/ryanmalonzo/eat-that-frog/internal/utils"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all candidate frogs",
	Long:  `Clear all tasks from your list of candidate frogs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		confirmDelete := utils.AskForConfirmation(cmd, "Do you really want to clear all candidate frogs?", false)
		if confirmDelete {
			return db.DeleteAllCandidates()
		}
		return nil
	},
}
