// Package utils exports utilitary functions used by various frog commands.
package utils

import (
	"bufio"
	"strings"

	"github.com/spf13/cobra"
)

func AskForConfirmation(cmd *cobra.Command, message string, defaultValue bool) bool {
	defaultValueString := "[y/N]"
	if defaultValue {
		defaultValueString = "[Y/n]"
	}

	cmd.Printf("%s %s: ", message, defaultValueString)

	scanner := bufio.NewScanner(cmd.InOrStdin())
	if !scanner.Scan() {
		return false
	}

	response := strings.ToLower(strings.TrimSpace(scanner.Text()))
	switch response {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	case "":
		return defaultValue
	default:
		return AskForConfirmation(cmd, message, defaultValue)
	}
}
