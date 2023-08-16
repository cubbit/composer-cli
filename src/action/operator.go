package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/spf13/cobra"
)

func CreateOperatorInteractive(cmd *cobra.Command) error {
	var urls *configuration.Url
	var err error

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu)")

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	firstName := input.TextPrompt("Enter first name:")
	lastName := input.TextPrompt("Enter last name:")
	email := input.TextPrompt("Enter email:")
	password := input.PasswordPrompt("Enter password:")
	secret := input.PasswordPrompt("Enter secret:")

	if err = api.CreateOperator(*urls, firstName, lastName, email, password, secret); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingOperator, err)
	}

	fmt.Printf("Operator %s created successfully\n", email)
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

	fmt.Printf("Operator %s created successfully\n", email)
	return nil
}
