package cmd

import (
	"fmt"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/spf13/cobra"
)

var skipCmd = &cobra.Command{
	Use:   "skip",
	Short: "Skip today's frog",
	Long:  `Skip the task that has been picked as today's frog to eat.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		frog, err := GetTodayFrog()
		if err != nil {
			return err
		}

		err = db.SkipTodayFrog(frog)
		if err != nil {
			return err
		}

		cmd.Println(fmt.Sprintf("❌ %s", frog))
		return nil
	},
}
