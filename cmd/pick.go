package cmd

import (
	"fmt"
	"strconv"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/spf13/cobra"
)

var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Pick a candidate frog to eat",
	Long:  `Pick a candidate frog from your list to eat.`,
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var index string
		if len(args) == 1 {
			index = args[0]
		}

		count, err := db.CountCandidates()
		if err != nil {
			return err
		}

		if count == 0 {
			return fmt.Errorf("no candidate frogs available to pick")
		}

		if index == "" {
			cmd.Println("Please choose from the following candidate frogs:")
			PrintAllCandidates(cmd)

			var choice int
			_, err := fmt.Scan(&choice)
			if err != nil {
				return err
			}

			if err != nil {
				return err
			}

			if choice < 1 || choice > count {
				return fmt.Errorf("invalid choice")
			}

			return db.PickCandidate(choice - 1)
		} else {
			parsedIndex, err := strconv.Atoi(index)
			if err != nil {
				return err
			}
			return db.PickCandidate(parsedIndex - 1)
		}
	},
}
