//go:build exclude

package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/spf13/cobra"
)

var serverStarted bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server to handle the token requests",
	Run: func(cmd *cobra.Command, args []string) {
		if serverStarted {
			fmt.Println("Server is already running")
			return
		}

		serverStarted = true

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			fmt.Println("Error getting port")
			return
		}

		profile, err := cmd.Flags().GetString("profile")
		if err != nil {
			fmt.Println("Error getting profile")
			return
		}

		config, err := cmd.Flags().GetString("config")
		if err != nil {
			fmt.Println("Error getting config")
			return
		}

		http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
			sessionType := r.URL.Query().Get("session_type")
			if sessionType == "" {
				http.Error(w, "SessionType not provided", http.StatusBadRequest)
				return
			}
			executablePath, err := os.Executable()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error locating executable: %v", err), http.StatusInternalServerError)
				return
			}

			executableFullPath := filepath.Join(filepath.Dir(executablePath), filepath.Base(executablePath))

			var output bytes.Buffer
			if configuration.SessionType(sessionType) == configuration.SessionTypeOperator {
				cmd := exec.Command(executableFullPath, "operator", "token", "--profile", profile, "--config", config)
				cmd.Stdout = &output
				if err := cmd.Run(); err != nil {
					http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
					return
				}
			}

			token := extractToken(output.String())
			if token == "" {
				http.Error(w, "Error extracting token", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(token))
		})

		addr := fmt.Sprintf(":%s", port)
		fmt.Printf("Server running at http://127.0.0.1%s\n", addr)
		http.ListenAndServe(addr, nil)
	},
}

func extractToken(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(output, "Token:") {
			tokenParts := strings.Split(line, ":")
			if len(tokenParts) > 1 {
				return strings.TrimSpace(tokenParts[1])
			}
		}
	}
	return ""
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "7373", "Port for the server to listen on")
}
