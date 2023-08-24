package cmd

import (
	"os"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/spf13/cobra"
)

var interactive bool

var rootCmd = &cobra.Command{
	Use:   "cubbit-operator-cli",
	Short: "The official Cubbit CLI (Command-Line Interface) for operators",
	Long:  "The CLI for managing operators, tenants and swarms in Cubbit distributed datacenter",
}

func Execute() {
	var err error
	if err = rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Run in interactive mode")
	rootCmd.PersistentFlags().String("profile", constants.DefaultProfile, "Profile Configuration")
	rootCmd.PersistentFlags().String("config", constants.DefaultFilePath, "Configuration path for file")
}
