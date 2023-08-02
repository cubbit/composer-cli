package actions

import (
	"fmt"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/urfave/cli/v2"
)



func GenerateAccessToken(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfiguration(cCtx); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("Access token: %s\n", *accessToken)

	return nil
}

func rehydrateTokenConfig(configPath string, conf *configuration.Config) (*string, error) {
	var accessToken, refreshToken string
	var err error

	if accessToken, refreshToken, err = api.RefreshAccessToken(conf.Urls, conf.RefreshToken); err != nil {
		return nil, fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	conf.RefreshToken = refreshToken

	if err = conf.Store(configPath); err != nil {
		return nil, fmt.Errorf("error while storing file path configuration: %w", err)
	}

	return &accessToken, nil
}
