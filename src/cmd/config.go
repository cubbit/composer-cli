// Package cmd provides CLI commands for managing configuration.
package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Cubbit CLI configuration",
	Long:  `Manage Cubbit CLI configuration including profiles, endpoints, and authentication settings.`,
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.NewConfig()
		if err := config.SaveConfig(); err != nil {
			utils.PrintError(err)
			return
		}
	},
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display current configuration",
	Long:  `Display the current configuration. Can show entire config or specific profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ConfigView(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit configuration file",
	Long:  `Open the configuration file in your default editor ($EDITOR).`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ConfigEdit(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var configProfilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List all configuration profiles",
	Long:  `List all available configuration profiles with their details.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ConfigProfiles(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var configSwitchProfileCmd = &cobra.Command{
	Use:   "switch-profile <profile-name>",
	Short: "Switch to a different profile",
	Long:  `Switch the active profile. This will set the default profile for subsequent commands.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ConfigSwitchProfile(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var configValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configuration",
	Long:  `Validate the configuration file and all profiles for correctness.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ConfigValidate(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configViewCmd)
	configCmd.AddCommand(configEditCmd)
	configCmd.AddCommand(configProfilesCmd)
	configCmd.AddCommand(configSwitchProfileCmd)
	configCmd.AddCommand(configValidateCmd)

	rootCmd.AddCommand(configCmd)
}
