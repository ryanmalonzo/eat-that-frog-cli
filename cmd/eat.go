package cmd

import (
	"fmt"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/ryanmalonzo/eat-that-frog/internal/frog"
	"github.com/spf13/cobra"
)

var eatCmd = &cobra.Command{
	Use:   "eat [task]",
	Short: "Set a task as today's frog directly",
	Long:  `Set a task as today's frog without adding it to candidates first. If a frog already exists for today and is still pending, it will be moved back to candidates.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		task := args[0]
		err := db.SetTodayFrog(task)
		if err != nil {
			return err
		}

		// Get and display the newly set frog
		frogTask, frogStatus, err := frog.GetTodayFrog()
		if err != nil {
			return err
		}

		cmd.Println(fmt.Sprintf("%s %s", frog.GetStatusEmoji(frogStatus), frogTask))

		return nil
	},
}
