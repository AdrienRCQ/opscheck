package cmd

import (
	"github.com/spf13/cobra"
)

func newCheckCmd() *cobra.Command {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Permet d'effectuer un test",
	}
	checkCmd.AddCommand(newCheckTCPCmd())
	checkCmd.AddCommand(newCheckHTTPCmd())
	checkCmd.AddCommand(newCheckBackupCmd())
	return checkCmd
}
