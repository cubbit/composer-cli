package action

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/spf13/cobra"
)

var server *http.Server

func Server(cmd *cobra.Command, args []string) error {
	var err error
	var configPath, port string
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if port, err = cmd.Flags().GetString("port"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	go StartServer(conf, configPath, port)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	fmt.Println("Shutting down server")
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}
	return nil
}

func StartServer(conf *configuration.Config, configPath string, port string) error {
	server = &http.Server{Addr: ":" + port}

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		sessionType := r.URL.Query().Get("session_type")
		if sessionType == "" {
			http.Error(w, "SessionType not provided", http.StatusBadRequest)
			return
		}

		if configuration.SessionType(sessionType) == configuration.SessionTypeOperator {
			if accessToken, err := rehydrateTokenConfig(configPath, conf); err != nil {
				http.Error(w, constants.ErrorGeneratingToken, http.StatusInternalServerError)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(*accessToken))
			}
		} else {
			http.Error(w, "SessionType not supported", http.StatusBadRequest)
			return
		}
	})

	fmt.Printf("Server starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("Server error: %v", err)
	}

	return nil
}
