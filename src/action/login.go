package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/spf13/cobra"
)

func SignInOperatorInteractive(cmd *cobra.Command) error {
	var err error
	var code, refreshToken string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url
	var conf = configuration.NewConfig("", configuration.Url{}, "")

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu)")

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
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
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallenge, err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorOperatorSignIn, err)
	}

	conf = configuration.NewConfig(profile, *urls, refreshToken)

	conf.StoreSession(configPath)

	fmt.Printf("User %s signed in successfully\n", email)

	return nil
}

func SignInOperator(cmd *cobra.Command) error {
	var err error
	var apiServerUrl, email, password, code, refreshToken, profile, configPath string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if password, err = cmd.Flags().GetString("password"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if apiServerUrl, err = cmd.Flags().GetString("api-server-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if code, err = cmd.Flags().GetString("code"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if configPath, err = cmd.Flags().GetString("config"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if profile, err = cmd.Flags().GetString("profile"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallenge, err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorOperatorSignIn, err)
	}

	var confs = configuration.NewConfig(profile, *urls, refreshToken)

	if err = confs.StoreSession(configPath); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorStoringSession, err)
	}

	fmt.Printf("User %s signed in successfully\n", email)

	return nil
}
