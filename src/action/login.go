package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func SignInOperatorInteractive(cmd *cobra.Command) error {
	var err error
	var email, password, code, refreshToken, apiServerUrl, configPath, twoFa, profile string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url
	var conf = configuration.NewConfig("", configuration.Url{}, "")
	var outs []string

	if apiServerUrl, err = tui.TextInput("Enter the api server url: (default https://api.cubbit.eu)", false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	outs = tui.Inputs("", false, tui.Input{Placeholder: "Email", IsPassword: false}, tui.Input{Placeholder: "Password", IsPassword: true})
	email = outs[0]
	password = outs[1]

	if twoFa, err = tui.ChooseOne("Do you want to add a 2fa code?", []string{"Yes", "No"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}
	if twoFa == "Yes" {
		if code, err = tui.TextInput("Insert the 2fa code", false); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
	}

	outs = tui.Inputs("", true, tui.Input{Placeholder: "Enter the config file to load (default: ./)", IsPassword: false}, tui.Input{Placeholder: "Enter the configuration profile (default: default)", IsPassword: true})
	configPath = outs[0]
	profile = outs[1]

	if configPath == "" {
		configPath = constants.DefaultFilePath
	}
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

	utils.PrintSuccess(fmt.Sprintf("user %s signed in successfully\n", email))
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

	utils.PrintSuccess(fmt.Sprintf("user %s signed in successfully\n", email))
	return nil
}
