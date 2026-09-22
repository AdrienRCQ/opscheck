package cmd

import (
	"github.com/spf13/cobra"
)

func newLogsCmd() *cobra.Command {
	checkCmd := &cobra.Command{
		Use:   "logs",
		Short: "Actions ciblées sur les logs",
	}
	checkCmd.AddCommand(newLogsFilterCmd())
	return checkCmd
}
