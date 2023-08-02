package actions

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

func SignOutOperatorInteractive(cCtx *cli.Context) error {
	var err error

	configPath := input.TextPrompt("Enter the config file to load (default: ./)")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	profile := input.TextPrompt("Enter the configuration profile (default: default)")
	if profile == "" {
		profile = "default"
	}

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("Configuration %s signed out successfully\n", profile)

	return nil
}

func SignOutOperator(cCtx *cli.Context) error {
	var err error

	if cCtx.Bool("interactive") {
		return SignOutOperatorInteractive(cCtx)
	}

	profile := cCtx.String("profile")
	if profile == "" {
		profile = "default"
	}

	configPath := cCtx.String("config")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("Configuration %s signed out successfully\n", profile)

	return nil
}
