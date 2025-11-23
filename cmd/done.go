package cmd

import (
	"fmt"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark today's frog as done",
	Long:  `Mark the task that has been picked as today's frog to eat as done.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		frog, err := GetTodayFrog()
		if err != nil {
			return err
		}

		err = db.MarkFrogAsDone(frog)
		if err != nil {
			return err
		}

		cmd.Println(fmt.Sprintf("✅ %s", frog))
		return nil
	},
}
