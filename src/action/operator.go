package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

func CreateOperatorInteractive(ctx *cli.Context) error {
	var err error
	var urls *configuration.Url

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

func CreateOperator(ctx *cli.Context) error {
	var err error
	var email, password, firstName, lastName string
	var urls *configuration.Url

	if ctx.Bool("interactive") {
		return CreateOperatorInteractive(ctx)
	}

	email = ctx.String("email")
	password = ctx.String("password")
	firstName = ctx.String("first-name")
	lastName = ctx.String("last-name")
	apiServerUrl := ctx.String("api-server-url")
	secret := ctx.String("secret")

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if err = api.CreateOperator(*urls, firstName, lastName, email, password, secret); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingOperator, err)
	}

	fmt.Printf("Operator %s created successfully\n", email)
	return nil
}
