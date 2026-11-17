package configuration_handler

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

type ConfigurationHandlerInterface interface {
	Load(cmd *cobra.Command, args []string) error
	Save() error

	CreateProfile(
		profileName string,
		apiKey string,
		organizationID string,
	) error

	SwitchProfile(profileName string) error

	DeleteProfile(profileName string) error
	DeleteAllProfiles() error

	GetConfigFile() (configuration_models.ConfigV2, error)
	GetConfigFilePath() (string, error)
	GetProfiles() (map[string]configuration_models.ProfileV2, error)
	GetActiveProfile() (configuration_models.ProfileV2, error)
	GetActiveProfileName() (string, error)
	// returns endpoints for active profile
	GetEndpoints() (configuration_models.EndpointsV2, error)
	// returns API key for active profile
	GetAPIKey() (string, error)
}

type LoadType string

const (
	LoadTypeNotLoaded         LoadType = "not-loaded"
	LoadTypeStandard          LoadType = "standard"
	LoadTypeForAuthCommands   LoadType = "for-auth-commands"
	LoadTypeForConfigCommands LoadType = "for-config-commands"
	LoadTypeConfigNotRequired LoadType = "config-not-required"
)

var pathLoadTypeMapping = map[string]LoadType{
	"auth activate":   LoadTypeForAuthCommands,
	"auth login":      LoadTypeForAuthCommands,
	"auth signup":     LoadTypeForAuthCommands,
	"config edit":     LoadTypeForConfigCommands,
	"config validate": LoadTypeForConfigCommands,
	"config init":     LoadTypeForConfigCommands,

	"version":       LoadTypeConfigNotRequired,
	"docs":          LoadTypeConfigNotRequired,
	"docs markdown": LoadTypeConfigNotRequired,
	"docs man":      LoadTypeConfigNotRequired,
	"docs rst":      LoadTypeConfigNotRequired,
	"docs yaml":     LoadTypeConfigNotRequired,
	"docs tree":     LoadTypeConfigNotRequired,
	"help":          LoadTypeConfigNotRequired,
}

type ConfigurationHandler struct {
	loadType LoadType

	configPath string
	configFile *configuration_models.ConfigV2

	requestedProfile *string
	endpoints        *configuration_models.EndpointsV2
	apiKey           *string
}

func NewConfigurationHandler() *ConfigurationHandler {
	return &ConfigurationHandler{
		configFile:       nil,
		configPath:       "",
		loadType:         LoadTypeNotLoaded,
		requestedProfile: nil,
		endpoints:        nil,
		apiKey:           nil,
	}
}
