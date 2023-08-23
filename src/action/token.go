package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func GenerateAccessToken(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if err = conf.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	utils.PrintSuccess(fmt.Sprintf("access token: %s\n", *accessToken))

	return nil
}

func rehydrateTokenConfig(configPath string, conf *configuration.Config) (*string, error) {
	var accessToken, refreshToken string
	var err error

	if accessToken, refreshToken, err = api.RefreshOperatorAccessToken(conf.Urls, conf.RefreshToken); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	conf.RefreshToken = refreshToken

	if err = conf.StoreSession(configPath); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	return &accessToken, nil
}
