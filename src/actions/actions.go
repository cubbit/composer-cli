package actions

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_PATH = "./"

func CreateOperatorInteractive(cCtx *cli.Context) error {
	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu/iam)")
	if apiServerUrl == "" {
		apiServerUrl = "https://api.cubbit.eu/iam"
	}

	firstName := input.TextPrompt("Enter first name:")
	lastName := input.TextPrompt("Enter last name:")
	email := input.TextPrompt("Enter email:")
	password := input.PasswordPrompt("Enter password:")

	return api.CreateOperator(apiServerUrl, firstName, lastName, email, password)
}

func CreateOperator(cCtx *cli.Context) error {
	var apiServerUrl, email, password, firstName, lastName string

	if cCtx.Bool("interactive") {
		return CreateOperatorInteractive(cCtx)
	}

	email = cCtx.String("email")
	password = cCtx.String("password")
	firstName = cCtx.String("first-name")
	lastName = cCtx.String("last-name")
	apiServerUrl = cCtx.String("api-server-url")

	return api.CreateOperator(apiServerUrl, firstName, lastName, email, password)
}

func SignInOperatorInteractive(cCtx *cli.Context) error {
	var err error
	var code, refreshToken string
	var challenge *api.ChallengeResponseModel

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu/iam)")
	if apiServerUrl == "" {
		apiServerUrl = "https://api.cubbit.eu/iam"
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

	name := input.TextPrompt("Enter the configuration name (default: default)")
	if name == "" {
		name = "default"
	}

	if challenge, err = api.GenerateOperatorChallenge(apiServerUrl, email); err != nil {
		return err
	}

	if refreshToken, err = api.PerformOperatorSignin(apiServerUrl, email, password, challenge, code); err != nil {
		return err
	}

	var conf = configuration.NewConfig(name, apiServerUrl, refreshToken)

	conf.Store(configPath)

	fmt.Printf("User %s signed in in successfully\n", email)

	return nil
}

func SignInOperator(cCtx *cli.Context) error {
	var err error
	var apiServerUrl, email, password, code, refreshToken string
	var challenge *api.ChallengeResponseModel

	if cCtx.Bool("interactive") {
		return SignInOperatorInteractive(cCtx)
	}

	configPath := cCtx.String("config")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	name := cCtx.String("name")
	if name == "" {
		name = "default"
	}

	email = cCtx.String("email")
	password = cCtx.String("password")
	code = cCtx.String("code")
	apiServerUrl = cCtx.String("api-server-url")

	if challenge, err = api.GenerateOperatorChallenge(apiServerUrl, email); err != nil {
		return err
	}

	if refreshToken, err = api.PerformOperatorSignin(apiServerUrl, email, password, challenge, code); err != nil {
		return err
	}

	var conf = configuration.NewConfig(name, apiServerUrl, refreshToken)

	if err = conf.Store(configPath); err != nil {
		return err
	}

	fmt.Printf("User %s signed in in successfully\n", email)

	return nil
}

func SignOutOperatorInteractive(cCtx *cli.Context) error {
	var err error
	configPath := input.TextPrompt("Enter the config file to load (default: ./)")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	name := input.TextPrompt("Enter the configuration name (default: default)")
	if name == "" {
		name = "default"
	}

	var conf = configuration.NewConfig(name, "", "")

	if err = conf.Store(configPath); err != nil {
		return err
	}

	fmt.Printf("Configuration %s signed out successfully\n", name)

	return nil
}

func SignOutOperator(cCtx *cli.Context) error {
	var err error

	if cCtx.Bool("interactive") {
		return SignOutOperatorInteractive(cCtx)
	}

	name := cCtx.String("name")
	if name == "" {
		name = "default"
	}

	configPath := cCtx.String("config")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	var conf = configuration.NewConfig(name, "", "")

	if err = conf.Store(configPath); err != nil {
		return err
	}

	fmt.Printf("Configuration %s signed out successfully\n", name)

	return nil
}

func readConfigFile() (string, string) {
	configPath := input.TextPrompt("Enter the config file to load (default: ./)")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	name := input.TextPrompt("Enter the configuration name (default: default)")
	if name == "" {
		name = "default"
	}

	return configPath, name
}

func GenerateAccessToken(cCtx *cli.Context) error {
	var err error
	var refreshToken, accessToken, name, configPath string

	if cCtx.Bool("interactive") {
		name, configPath = readConfigFile()
	} else {
		name = cCtx.String("name")
		if name == "" {
			name = "default"
		}

		configPath = cCtx.String("config")
		if configPath == "" {
			configPath = DEFAULT_FILE_PATH
		}
	}

	var conf = configuration.NewConfig(name, "", "")

	if err = conf.Load(configPath, name); err != nil {
		return err
	}

	if accessToken, refreshToken, err = api.RefreshAccessToken(conf.ApiServerUrl, conf.RefreshToken); err != nil {
		return err
	}

	conf.RefreshToken = refreshToken

	if err = conf.Store(configPath); err != nil {
		return err
	}

	fmt.Printf("Access token: %s\n", accessToken)

	return nil
}

func CreateTenant(cCtx *cli.Context) error {
	fmt.Printf("Costruisco un tenant")
	return nil
}
