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

	if _, err = tui.TextInputs("Enter your API server URL", false, tui.Input{Placeholder: "Enter the api server url: (default https://api.cubbit.eu)", IsPassword: false, Value: &apiServerUrl}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if _, err = tui.TextInputs("Fill in the form bellow", false, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if twoFa, err = tui.ChooseOne("Do you want to add a 2fa code?", false, []string{"Yes", "No"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if twoFa == "Yes" {
		if _, err = tui.TextInputs("Insert the two factor authentication", false, tui.Input{Placeholder: "Insert the 2fa code*", IsPassword: false, Value: &code}); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
	}

	if _, err = tui.TextInputs("Enter the config file path and name", true, tui.Input{Placeholder: "Enter the config file to load (default: ./)", IsPassword: false, Value: &configPath}, tui.Input{Placeholder: "Enter the configuration profile (default: default)", IsPassword: true, Value: &profile}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if configPath == "" {
		configPath = constants.DefaultFilePath
	}

	if profile == "" {
		profile = constants.DefaultProfile
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorOperatorSignInRequest, err)
	}

	conf = configuration.NewConfig(profile, *urls, refreshToken)
	conf.StoreSession(configPath)

	utils.PrintSuccess(fmt.Sprintf("user %s signed in successfully", email))

	return nil
}
