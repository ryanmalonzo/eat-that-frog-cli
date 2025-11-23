package cmd

import (
	"fmt"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/spf13/cobra"
)

func PrintAllCandidates(cmd *cobra.Command) error {
	candidates, err := db.GetAllCandidates()
	if err != nil {
		return err
	}

	for i, candidate := range candidates {
		cmd.Println(fmt.Sprintf("%v: %v", i+1, candidate))
	}
	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all candidate frogs",
	Long:  `List all tasks that are candidates to be eaten as frogs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return PrintAllCandidates(cmd)
	},
}
