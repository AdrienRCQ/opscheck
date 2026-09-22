package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

type httpCheckResult struct {
	Url          string  `json:"url"`
	StatusCode   int     `json:"status_code"`
	ExpectStatus int     `json:"expected_status"`
	Duration     float64 `json:"duration_ms"`
	Success      bool    `json:"success"`
}

func newCheckHTTPCmd() *cobra.Command {
	var targetURL string
	var expectStatus int
	var timeout time.Duration
	var jsonOutput bool
	httpchkCmd := &cobra.Command{
		Use:   "http",
		Short: "test la connexion en HTTP vers une cible",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Vérification de la valeur de timeout
			if timeout <= 0 {
				return fmt.Errorf("Timeout ne peut pas être inférieur à 0")
			}
			// Validation du expectStatus
			if expectStatus < 100 || expectStatus > 599 {
				return fmt.Errorf("Expect status doit être compris entre 100 et 599")
			}
			// Vérification du format de l'url
			parsedURL, err := url.Parse(targetURL)
			if err != nil {
				return fmt.Errorf("URL invalide: %w", err)
			}
			if (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") ||
				parsedURL.Hostname() == "" {
				return fmt.Errorf("L'URL doit contenir un schéma http ou https et un hôte")
			}

			startTime := time.Now()
			client := &http.Client{Timeout: timeout}
			resp, err := client.Get(targetURL)
			duration := time.Since(startTime)
			if err != nil {
				return fmt.Errorf("REequête vers %s impossible : %w", targetURL, err)
			}

			defer resp.Body.Close() // fermeture corps de réponse

			result := httpCheckResult{
				Url:          targetURL,
				StatusCode:   resp.StatusCode,
				ExpectStatus: expectStatus,
				Duration:     duration.Seconds() * 1000,
				Success:      resp.StatusCode == expectStatus,
			}

			if jsonOutput { // JSON Output
				encoder := json.NewEncoder(cmd.OutOrStdout())
				encoder.SetIndent("", " ")
				if err := encoder.Encode(result); err != nil {
					return fmt.Errorf("Encodage JSON impossible : %w", err)
				}
			} else { // Affichage classique
				label := "OK"
				if !result.Success {
					label = "FAIL"
				}

				_, err := fmt.Fprintf(
					cmd.OutOrStdout(),
					"%s - %s retourne %d (attendu %d) en %.2f ms \n",
					label,
					result.Url,
					result.StatusCode,
					result.ExpectStatus,
					result.Duration,
				)
				if err != nil {
					return fmt.Errorf("affichage du résultat impossible %w", err)
				}

				return nil
			}
			if !result.Success {
				return fmt.Errorf(
					"Statut inattendu : attendu %d, reçu %d",
					result.ExpectStatus,
					result.StatusCode,
				)
			}
			return nil
		},
	}

	// Inscription des flags
	httpchkCmd.Flags().StringVar(&targetURL, "url", "localhost", "url cible")
	if err := httpchkCmd.MarkFlagRequired("url"); err != nil {
		panic(err)
	}
	httpchkCmd.Flags().IntVar(&expectStatus, "expect-status", 200, "code http désiré")
	if err := httpchkCmd.MarkFlagRequired("expect-status"); err != nil {
		panic(err)
	}
	httpchkCmd.Flags().DurationVar(&timeout, "timeout", 4*time.Second, "timeout connexion")
	httpchkCmd.Flags().BoolVar(&jsonOutput, "json", false, "Afficher le résultat au format JSON")

	return httpchkCmd
}
