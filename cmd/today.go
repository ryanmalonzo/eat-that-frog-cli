package cmd

import (
	"errors"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/spf13/cobra"
)

func GetTodayFrog() (string, error) {
	frog, err := db.GetTodayFrog()
	if err != nil {
		return "", err
	}

	if frog == "" {
		return "", errors.New("no frog has been picked for today")
	}

	return frog, nil
}

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's frog",
	Long:  `Show the task that has been picked as today's frog to eat.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		frog, err := GetTodayFrog()
		if err != nil {
			return err
		}

		cmd.Println(frog)
		return nil
	},
}
