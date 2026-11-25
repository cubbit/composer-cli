package configuration_handler

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func (h *ConfigurationHandler) Load(cmd *cobra.Command, args []string) error {
	if h.loadType != LoadTypeNotLoaded {
		return configuration_models.ErrConfigurationIsAlreadyLoaded
	}

	configPath, err := h.getConfigurationPath(cmd, args)
	if err != nil {
		return fmt.Errorf("failed to determine configuration path: %w", err)
	}

	h.configPath = configPath

	_, internalPath, found := strings.Cut(cmd.CommandPath(), " ")

	if !found {
		// Root command (ds3-composer with no subcommand — shows help): no config needed
		return h.loadForConfigNotRequired()
	}

	loadType, ok := pathLoadTypeMapping[internalPath]
	if !ok {
		return h.loadStandard(cmd, args)
	}

	switch loadType {
	case LoadTypeForAuthCommands:
		return h.loadForAuthCommands(cmd, args)
	case LoadTypeForConfigCommands:
		h.loadType = LoadTypeForConfigCommands
		return h.loadForConfigCommands(cmd, args)
	case LoadTypeConfigNotRequired:
		return h.loadForConfigNotRequired()
	default:
		return h.loadStandard(cmd, args)
	}
}

func (h *ConfigurationHandler) loadForAuthCommands(cmd *cobra.Command, args []string) error {
	initEndpointsPath, err := cmd.Flags().GetString("endpoints")
	if err != nil {
		return fmt.Errorf("failed to get 'endpoints' flag: %w", err)
	}

	requestedProfileName, err := cmd.Flags().GetString("profile")
	if err != nil {
		return fmt.Errorf("failed to get 'profile' flag: %w", err)
	}

	var initEndpoints *configuration_models.EndpointsV2Inits
	if initEndpointsPath != "" {
		initEndpoints, err = configuration_models.ParseAndValidateEndpointsV2Inits(initEndpointsPath)
		if err != nil {
			return fmt.Errorf("failed to parse endpoints file: %w", err)
		}
	}

	// there is no error check, config might not exist
	// for signup and login, it's ok to not have a config already, we will create one with the inits
	config, err := configuration_models.ParseAndValidateConfigV2(h.configPath)
	if err != nil {
		config = nil
	}

	if config != nil {
		h.configFile = config

		if initEndpoints != nil {
			h.endpoints = initEndpoints.ToEndpointsV2()
		} else {
			endpoints := getEndpointsForProfile(config, requestedProfileName)
			h.endpoints = &endpoints
		}

		if requestedProfileName != "" {
			h.requestedProfile = &requestedProfileName
		}

		h.loadType = LoadTypeForAuthCommands
		return nil
	}

	if initEndpoints != nil {
		h.endpoints = initEndpoints.ToEndpointsV2()
		h.loadType = LoadTypeForAuthCommands
		return nil
	}

	h.endpoints = getEndpointsFromDefaults()
	h.loadType = LoadTypeForAuthCommands

	return nil
}

func (h *ConfigurationHandler) loadForConfigCommands(cmd *cobra.Command, args []string) error {
	h.loadType = LoadTypeForConfigCommands

	return nil
}

func (h *ConfigurationHandler) loadForConfigNotRequired() error {
	h.loadType = LoadTypeConfigNotRequired

	return nil
}

func (h *ConfigurationHandler) loadStandard(cmd *cobra.Command, _ []string) error {
	config, err := configuration_models.ParseAndValidateConfigV2(h.configPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration from path %q: %w", h.configPath, err)
	}

	requestedProfileName, err := cmd.Flags().GetString("profile")
	if err != nil {
		return fmt.Errorf("failed to get 'profile' flag: %w", err)
	}

	selectedProfile := config.GetActiveProfile()
	if requestedProfileName != "" {
		// the requested profile must exist, we are not creating a new one here
		selectedProfile, err = config.GetProfile(requestedProfileName)
		if err != nil {
			return err
		}
	}

	h.configFile = config
	h.loadType = LoadTypeStandard
	h.requestedProfile = &requestedProfileName
	h.endpoints = &selectedProfile.Endpoints
	h.apiKey = &selectedProfile.APIKey

	return nil
}

// getEndpointsForProfile returns the endpoints of the requested profile, falling back
// to the active profile when the requested one is not configured (it might be a profile
// that is going to be created by the running command).
func getEndpointsForProfile(
	config *configuration_models.ConfigV2,
	requestedProfileName string,
) configuration_models.EndpointsV2 {
	if requestedProfileName != "" {
		if requestedProfile, err := config.GetProfile(requestedProfileName); err == nil {
			return requestedProfile.Endpoints
		}
	}

	return config.GetActiveProfile().Endpoints
}

func getEndpointsFromDefaults() *configuration_models.EndpointsV2 {
	return &configuration_models.EndpointsV2{
		IAM:  fmt.Sprintf("%s%s", constants.BaseAPIURL, constants.BaseIamURI),
		Dash: fmt.Sprintf("%s%s", constants.BaseAPIURL, constants.BaseDashURI),
		CH:   fmt.Sprintf("%s%s", constants.BaseAPIURL, constants.BaseChURI),
	}
}
