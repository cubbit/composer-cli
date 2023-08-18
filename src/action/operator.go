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

func CreateOperatorInteractive(cmd *cobra.Command) error {
	var urls *configuration.Url
	var apiServerUrl, firstName, lastName, email, password, secret string
	var err error

	if apiServerUrl, err = tui.TextInput("Enter the api server url: (default https://api.cubbit.eu)", false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	outs := tui.Inputs("", true, tui.Input{Placeholder: "First Name", IsPassword: false}, tui.Input{Placeholder: "Last Name", IsPassword: false}, tui.Input{Placeholder: "Email", IsPassword: false}, tui.Input{Placeholder: "Password", IsPassword: true}, tui.Input{Placeholder: "Secret", IsPassword: true})
	firstName = outs[0]
	lastName = outs[1]
	email = outs[2]
	password = outs[3]
	secret = outs[4]

	if err = api.CreateOperator(*urls, firstName, lastName, email, password, secret); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingOperator, err)
	}
	utils.PrintSuccess(fmt.Sprintf("operator %s created successfully\n", email))
	return nil
}

func CreateOperator(cmd *cobra.Command) error {
	var url *configuration.Url
	var err error
	var email, password, firstName, lastName, apiServerUrl, secret string

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if password, err = cmd.Flags().GetString("password"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if firstName, err = cmd.Flags().GetString("first-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if lastName, err = cmd.Flags().GetString("last-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if apiServerUrl, err = cmd.Flags().GetString("api-server-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if secret, err = cmd.Flags().GetString("secret"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if url, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if err = api.CreateOperator(*url, firstName, lastName, email, password, secret); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingOperator, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s created successfully\n", email))
	return nil
}
