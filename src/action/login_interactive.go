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
	var email, password, code, refreshToken, apiServerUrl, configPath, twoFa, profile, defaultConfigPath string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url
	var conf = configuration.NewConfig(configuration.SessionTypeOperator, "", configuration.Url{}, "")

	if _, err = tui.TextInputs(
		"Enter your API server URL",
		false, tui.Input{Placeholder: "Enter the api server url: (default https://api.cubbit.eu)", IsPassword: false, Value: &apiServerUrl}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if urls, err = configuration.ConfigureAPIServerURL(configuration.SessionTypeOperator, apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if _, err = tui.TextInputs("Fill in the form bellow", false, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if twoFa, err = tui.ChooseOne("Do you want to add a 2fa code?", false, false, []string{"Yes", "No"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if twoFa == "Yes" {
		if _, err = tui.TextInputs("Insert the two factor authentication", false, tui.Input{Placeholder: "Insert the 2fa code*", IsPassword: false, Value: &code}); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
	}

	if defaultConfigPath, err = configuration.GetDefaultConfigPath(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if _, err = tui.TextInputs("", true, tui.Input{Placeholder: fmt.Sprintf("Enter the config file path to load (default: %s)", defaultConfigPath), Value: &configPath}, tui.Input{Placeholder: "Enter the configuration profile (default: default)", Value: &profile}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if configPath == "" {
		configPath = defaultConfigPath
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

	conf = configuration.NewConfig(configuration.SessionTypeOperator, profile, *urls, refreshToken)
	conf.StoreSession(configPath)

	utils.PrintSuccess(fmt.Sprintf("user %s signed in successfully", email))

	return nil
}

func SignInAccountInteractive(cmd *cobra.Command) error {
	var err error
	var email, password, code, refreshToken, apiServerUrl, tenantID, configPath, twoFa, profile, defaultConfigPath string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url
	var conf = configuration.NewConfig(configuration.SessionTypeAccount, "", configuration.Url{}, "")

	if _, err = tui.TextInputs("Enter your API server URL", false, tui.Input{Placeholder: "Enter the api server url: (default https://api.cubbit.eu)", IsPassword: false, Value: &apiServerUrl}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if urls, err = configuration.ConfigureAPIServerURL(configuration.SessionTypeAccount, apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if _, err = tui.TextInputs("Enter your tenant id", false, tui.Input{Placeholder: "Enter the tenant id", IsPassword: false, Value: &tenantID}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if _, err = tui.TextInputs(
		"Fill in the form bellow",
		false,
		tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email},
		tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if twoFa, err = tui.ChooseOne("Do you want to add a 2fa code?", false, false, []string{"Yes", "No"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if twoFa == "Yes" {
		if _, err = tui.TextInputs("Insert the two factor authentication", false, tui.Input{Placeholder: "Insert the 2fa code*", IsPassword: false, Value: &code}); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
	}

	if defaultConfigPath, err = configuration.GetDefaultConfigPath(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if _, err = tui.TextInputs(
		"",
		true,
		tui.Input{Placeholder: fmt.Sprintf("Enter the config file path to load (default: %s)", defaultConfigPath), Value: &configPath},
		tui.Input{Placeholder: "Enter the configuration profile (default: default)", Value: &profile}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if configPath == "" {
		configPath = defaultConfigPath
	}

	if profile == "" {
		profile = constants.DefaultProfile
	}

	if challenge, err = api.GenerateAccountChallenge(*urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if refreshToken, err = api.PerformAccountSignin(*urls, tenantID, email, password, challenge, code); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorOperatorSignInRequest, err)
	}

	conf = configuration.NewConfig(configuration.SessionTypeAccount, profile, *urls, refreshToken)
	conf.StoreSession(configPath)

	utils.PrintSuccess(fmt.Sprintf("account %s signed in successfully", email))

	return nil
}
