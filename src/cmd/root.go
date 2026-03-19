package cmd

import (
	"encoding/json"
	"os"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/spf13/cobra"
)

var interactive bool
var packageJSON []byte
var devNull *os.File

type PackageData struct {
	Version string `json:"version"`
}

func setupSilentMode(cmd *cobra.Command) {
	silent, _ := cmd.Flags().GetBool("silent")
	if silent {
		f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
		if err == nil {
			devNull = f
			os.Stdout = f
			os.Stderr = f
		}
	}
}

func cleanupSilentMode() {
	if devNull != nil {
		devNull.Close()
		devNull = nil
	}
}

var rootCmd = &cobra.Command{
	Use:   "cubbit",
	Short: "The official Cubbit CLI (Command-Line Interface) for operators",
	Long:  "The CLI for managing operators, tenants and swarms in Cubbit distributed datacenter",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setupSilentMode(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, _ := cmd.Flags().GetString("endpoint")
		if endpoint != "" {
			err := configuration.SetAPIEndpoint(endpoint)
			if err != nil {
				os.Exit(1)
			}
			return
		}

		cmd.Help()
	},
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
	cleanupSilentMode()
}

func init() {
	// Persistent flags (available to all subcommands)
	rootCmd.PersistentFlags().String("profile", "", "Profile Configuration")
	rootCmd.PersistentFlags().String("output", "human", "Output format: human (default), json, yaml, xml")
	rootCmd.PersistentFlags().Bool("no-headers", false, "Suppress table headers in human output (for easier scripting)")
	rootCmd.PersistentFlags().Bool("quiet", false, "Minimize stdout for CI/CD workflows (no table output, just essentials)")
	rootCmd.PersistentFlags().Bool("silent", false, "Redirect all output to /dev/null")

	// Local flags (only available to root command)
	rootCmd.Flags().String("endpoint", "", "Override the default API endpoint URL")
}
