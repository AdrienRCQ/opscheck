package cmd

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "opscheck",
		Short:   "Boite à outil sysops",
		Version: "1.0.1",
	}

	rootCmd.AddCommand(newInfoCmd())
	rootCmd.AddCommand(newCheckCmd())
	rootCmd.AddCommand(newLogsCmd())
	return rootCmd
}

func Execute() error {
	return newRootCmd().Execute()
}
