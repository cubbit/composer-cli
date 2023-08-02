package actions

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_PATH = "./"

func SignInOperatorInteractive(cCtx *cli.Context) error {
	var err error
	var code, refreshToken string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url
	var conf = configuration.NewConfig("", configuration.Url{}, "")

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu)")

	if urls, err = configuration.ApiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	email := input.TextPrompt("Enter email:")
	password := input.PasswordPrompt("Enter password:")

	if input.YesNoPrompt("Do you want to add a 2fa code?", false) {
		code = input.TextPrompt("Please insert the 2fa code:")
	}

	configPath := input.TextPrompt("Enter the config file to load (default: ./)")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	profile := input.TextPrompt("Enter the configuration profile (default: default)")
	if profile == "" {
		profile = "default"
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator signin: %w", err)
	}

	conf = configuration.NewConfig(profile, *urls, refreshToken)

	conf.Store(configPath)

	fmt.Printf("User %s signed in successfully\n", email)

	return nil
}

func SignInOperator(cCtx *cli.Context) error {
	var err error
	var apiServerUrl, email, password, code, refreshToken string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url

	if cCtx.Bool("interactive") {
		return SignInOperatorInteractive(cCtx)
	}

	configPath := cCtx.String("config")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	profile := cCtx.String("profile")
	if profile == "" {
		profile = "default"
	}

	email = cCtx.String("email")
	password = cCtx.String("password")
	code = cCtx.String("code")
	apiServerUrl = cCtx.String("api-server-url")

	if urls, err = configuration.ApiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator singin: %w", err)
	}

	var confs = configuration.NewConfig(profile, *urls, refreshToken)

	if err = confs.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("User %s signed in successfully\n", email)

	return nil
}
