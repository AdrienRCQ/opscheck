package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

type systemInfo struct {
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	CPUs     int    `json:"cups"`
}

func newInfoCmd() *cobra.Command {
	var jsonOutput bool
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "Afficher les informations de la machine",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			hostname, err := os.Hostname()
			if err != nil {
				return fmt.Errorf("Récupération du hostname impossible %w", err)
			}
			info := systemInfo{
				Hostname: hostname,
				OS:       runtime.GOOS,
				Arch:     runtime.GOARCH,
				CPUs:     runtime.NumCPU(),
			}
			if jsonOutput {
				encoder := json.NewEncoder(cmd.OutOrStdout())
				encoder.SetIndent("", " ")
				return encoder.Encode(info)
			}
			_, err = fmt.Fprintf(
				cmd.OutOrStdout(),
				"Machine : %s\nOS : %s\nArchitecture : %s\nCPU Logiques : %d\n",
				info.Hostname,
				info.OS,
				info.Arch,
				info.CPUs,
			)
			return err
		},
	}

	infoCmd.Flags().BoolVar(
		&jsonOutput,
		"json",
		false,
		"Afficher le résultat au format JSON",
	)
	return infoCmd
}
