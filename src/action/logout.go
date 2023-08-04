package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

func SignOutOperatorInteractive(ctx *cli.Context) error {
	var err error

	configPath := input.TextPrompt("Enter the config file to load (default: ./)")
	if configPath == "" {
		configPath = constants.DefaultFilePath
	}

	profile := input.TextPrompt("Enter the configuration profile (default: default)")
	if profile == "" {
		profile = constants.DefaultProfile
	}

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

	if err = conf.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	fmt.Printf("Configuration %s signed out successfully\n", profile)

	return nil
}

func SignOutOperator(ctx *cli.Context) error {
	var err error

	if ctx.Bool("interactive") {
		return SignOutOperatorInteractive(ctx)
	}

	profile := ctx.String("profile")
	if profile == "" {
		profile = constants.DefaultProfile
	}

	configPath := ctx.String("config")
	if configPath == "" {
		configPath = constants.DefaultFilePath
	}

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

	if err = conf.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	fmt.Printf("Configuration %s signed out successfully\n", profile)

	return nil
}
