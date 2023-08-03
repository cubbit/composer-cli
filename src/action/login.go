package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

func SignInOperatorInteractive(ctx *cli.Context) error {
	var err error
	var code, refreshToken string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url
	var conf = configuration.NewConfig("", configuration.Url{}, "")

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu)")

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	email := input.TextPrompt("Enter email:")
	password := input.PasswordPrompt("Enter password:")

	if input.YesNoPrompt("Do you want to add a 2fa code?", false) {
		code = input.TextPrompt("Please insert the 2fa code:")
	}

	configPath := input.TextPrompt("Enter the config file to load (default: ./)")
	if configPath == "" {
		configPath = constants.DefaultFilePath
	}

	profile := input.TextPrompt("Enter the configuration profile (default: default)")
	if profile == "" {
		profile = constants.DefaultProfile
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator signin: %w", err)
	}

	conf = configuration.NewConfig(profile, *urls, refreshToken)

	conf.StoreSession(configPath)

	fmt.Printf("User %s signed in successfully\n", email)

	return nil
}

func SignInOperator(ctx *cli.Context) error {
	var err error
	var apiServerUrl, email, password, code, refreshToken string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url

	if ctx.Bool("interactive") {
		return SignInOperatorInteractive(ctx)
	}

	configPath := ctx.String("config")
	if configPath == "" {
		configPath = constants.DefaultFilePath
	}

	profile := ctx.String("profile")
	if profile == "" {
		profile = "default"
	}

	email = ctx.String("email")
	password = ctx.String("password")
	code = ctx.String("code")
	apiServerUrl = ctx.String("api-server-url")

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator singin: %w", err)
	}

	var confs = configuration.NewConfig(profile, *urls, refreshToken)

	if err = confs.StoreSession(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("User %s signed in successfully\n", email)

	return nil
}
