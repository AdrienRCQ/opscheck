package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

func newCheckTCPCmd() *cobra.Command {
	var host string
	var port string
	var timeout time.Duration
	tcpchkCmd := &cobra.Command{
		Use:   "tcp",
		Short: "test la connexion en TCP vers une cible",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if timeout < 0 {
				// Vérification de la valeur de timeout
				return fmt.Errorf("Timeout ne peut pas être inférieur à 0")
			}
			start_time := time.Now()
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
			duration := time.Since(start_time)

			if err != nil {
				return fmt.Errorf("Connecting error: %w", err)
			}

			defer conn.Close()

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"OK - %s accessible en %.2f ms\n",
				net.JoinHostPort(host, port),
				(duration.Seconds())*1000,
			)
			return nil
		},
	}
	tcpchkCmd.Flags().StringVar(&host, "host", "localhost", " hôte cible")
	tcpchkCmd.MarkFlagRequired("host")
	tcpchkCmd.Flags().StringVar(&port, "port", "80", "port cible")
	tcpchkCmd.MarkFlagRequired("port")
	tcpchkCmd.Flags().DurationVar(&timeout, "timeout", 4*time.Second, "timeout connexion")

	return tcpchkCmd
}
