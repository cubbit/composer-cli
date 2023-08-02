package actions

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

func CreateOperatorInteractive(cCtx *cli.Context) error {
	var err error
	var urls *configuration.Url

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu)")

	if urls, err = configuration.ApiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	firstName := input.TextPrompt("Enter first name:")
	lastName := input.TextPrompt("Enter last name:")
	email := input.TextPrompt("Enter email:")
	password := input.PasswordPrompt("Enter password:")
	secret := input.PasswordPrompt("Enter secret:")

	if err = api.CreateOperator(*urls, firstName, lastName, email, password, secret); err != nil {
		return fmt.Errorf("error while creating operator: %w", err)
	}

	fmt.Printf("Operator %s created successfully\n", email)
	return nil
}

func CreateOperator(cCtx *cli.Context) error {
	var err error
	var email, password, firstName, lastName string
	var urls *configuration.Url

	if cCtx.Bool("interactive") {
		return CreateOperatorInteractive(cCtx)
	}

	email = cCtx.String("email")
	password = cCtx.String("password")
	firstName = cCtx.String("first-name")
	lastName = cCtx.String("last-name")
	apiServerUrl := cCtx.String("api-server-url")
	secret := cCtx.String("secret")

	if urls, err = configuration.ApiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	if err = api.CreateOperator(*urls, firstName, lastName, email, password, secret); err != nil {
		return fmt.Errorf("error while creating operator: %w", err)
	}

	fmt.Printf("Operator %s created successfully\n", email)
	return nil
}
