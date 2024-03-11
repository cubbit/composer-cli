//go:build exclude

package cmd

import (
	"fmt"
	"github.com/cubbit/cubbit/client/cli/src/action"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server to handle the token requests",
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.Server(cmd, args); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "7373", "Port for the server to listen on")
}
