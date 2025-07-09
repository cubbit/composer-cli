package cmd

import (
	"encoding/json"
	"os"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/spf13/cobra"
)

const ENABLE_ACCOUNT_SECTION = false

var interactive bool
var packageJSON []byte

type PackageData struct {
	Version string `json:"version"`
}

var rootCmd = &cobra.Command{
	Use:   "cubbit-operator-cli",
	Short: "The official Cubbit CLI (Command-Line Interface) for operators",
	Long:  "The CLI for managing operators, tenants and swarms in Cubbit distributed datacenter",
}

func Execute(packageJSON []byte) {
	var pkg PackageData
	if err := json.Unmarshal(packageJSON, &pkg); err != nil {
		os.Exit(1)
	}

	rootCmd.Version = pkg.Version
	rootCmd.SetVersionTemplate("{{.Use}} version {{.Version}}\n")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	var defaultConfigPath string
	var err error
	if defaultConfigPath, err = configuration.GetDefaultConfigPath(); err != nil {
		os.Exit(1)
	}

	rootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Run in interactive mode")
	rootCmd.PersistentFlags().String("profile", constants.DefaultProfile, "Profile Configuration")
	rootCmd.PersistentFlags().String("config", defaultConfigPath, "Configuration path for file")
	rootCmd.PersistentFlags().Bool("human", false, "Output in human-readable format")
}
