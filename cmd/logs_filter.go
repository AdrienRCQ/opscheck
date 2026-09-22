package cmd

import (
	"github.com/spf13/cobra"
)

func newLogsFilterCmd() *cobra.Command {
	var file string
	var contains string
	var ignoreCase bool
	filterLogsCmd := &cobra.Command{
		Use:   "filter",
		Short: "filtre des logs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	filterLogsCmd.Flags().StringVar(&file, "file", "", "fichier de logs")
	filterLogsCmd.MarkFlagRequired("file")
	filterLogsCmd.Flags().StringVar(&contains, "contains", "", "filtre")
	filterLogsCmd.MarkFlagRequired("contains")
	filterLogsCmd.Flags().BoolVar(&ignoreCase, "ignore-case", false, "Ignore case")

	return filterLogsCmd
}
