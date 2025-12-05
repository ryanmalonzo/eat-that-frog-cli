package cmd

import (
	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/ryanmalonzo/eat-that-frog/internal/utils"
	"github.com/spf13/cobra"
)

var force bool

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all candidate frogs",
	Long:  `Clear all tasks from your list of candidate frogs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !force {
			confirmDelete := utils.AskForConfirmation(cmd, "Do you really want to clear all candidate frogs?", false)
			if !confirmDelete {
				return nil
			}
		}
		return db.DeleteAllCandidates()
	},
}

func init() {
	clearCmd.Flags().BoolVarP(&force, "yes", "y", false, "Skip confirmation prompt")
}
