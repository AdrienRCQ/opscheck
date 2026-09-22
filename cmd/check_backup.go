package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type backupCheckResult struct {
	Path         string     `json:"path"`
	Size         int64      `json:"size_bytes"`
	LastModified *time.Time `json:"last_modified,omitempty"`
	AgeSeconds   float64    `json:"age_seconds"`
	Exists       bool       `json:"exists"`
	RegularFile  bool       `json:"regular_file"`
	Success      bool       `json:"success"`
	Error        string     `json:"error,omitempty"`
}

func controlBackupFile(filePath string, maxAge time.Duration) (backupCheckResult, error) {
	result := backupCheckResult{
		Path: filePath,
	}

	info, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return result, fmt.Errorf("le fichier de sauvegarde n'existe pas : %s", filePath)
		}

		return result, fmt.Errorf(
			"impossible d'inspecter %s : %w",
			filePath,
			err,
		)
	}

	result.Exists = true
	result.RegularFile = info.Mode().IsRegular()
	if !result.RegularFile {
		return result, fmt.Errorf(
			"le chemin ne désigne pas un fichier ordinaire : %s",
			filePath,
		)
	}

	lastModified := info.ModTime()
	age := time.Since(lastModified)

	result.Size = info.Size()
	result.LastModified = &lastModified
	result.AgeSeconds = age.Seconds()

	if result.Size == 0 {
		return result, fmt.Errorf(
			"le fichier de sauvegarde est vide : %s",
			filePath,
		)
	}

	if age < 0 {
		return result, fmt.Errorf(
			"la date de modification est dans le futur : %s",
			filePath,
		)
	}

	if age > maxAge {
		return result, fmt.Errorf(
			"sauvegarde trop ancienne : âge %s, maximum autorisé %s",
			age.Round(time.Second),
			maxAge,
		)
	}
	result.Success = true
	return result, nil
}

func newCheckBackupCmd() *cobra.Command {
	var targetPath string
	var maxAge time.Duration
	var jsonOutput bool

	backupCheckCmd := &cobra.Command{
		Use:   "backup",
		Short: "contrôle l'existance d'un fichier de backup et sa durée de vie",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if targetPath == "" {
				return fmt.Errorf("le chemin de sauvegarde ne peut pas être vide")
			}
			if maxAge <= 0 {
				return fmt.Errorf("max-age doit être strictement positif")
			}

			result, checkErr := controlBackupFile(targetPath, maxAge)
			if checkErr != nil {
				result.Error = checkErr.Error()
			}

			if jsonOutput {
				encoder := json.NewEncoder(cmd.OutOrStdout())
				encoder.SetIndent("", "  ")
				if err := encoder.Encode(result); err != nil {
					return fmt.Errorf("encodage JSON impossible : %w", err)
				}

			} else if result.Success {
				_, err := fmt.Fprintf(
					cmd.OutOrStdout(),
					"OK - sauvegarde présente : %s\n"+
						"Taille : %d octets\n"+
						"Dernière modification : %s\n",
					result.Path,
					result.Size,
					result.LastModified.Format(time.RFC3339),
				)
				if err != nil {
					return fmt.Errorf("affichage du résultat impossible : %w", err)
				}
			}
			return checkErr
		},
	}

	// Inscription des flags
	backupCheckCmd.Flags().StringVar(&targetPath, "path", "", "Chemin du fichier de sauvegarde")
	if err := backupCheckCmd.MarkFlagRequired("path"); err != nil {
		panic(err)
	}
	backupCheckCmd.Flags().DurationVar(&maxAge, "max-age", 24*time.Hour, "Ancienneté maximal de la sauvegarde")
	if err := backupCheckCmd.MarkFlagRequired("max-age"); err != nil {
		panic(err)
	}

	backupCheckCmd.Flags().BoolVar(&jsonOutput, "json", false, "Afficher le résultat au format JSON")

	return backupCheckCmd
}
