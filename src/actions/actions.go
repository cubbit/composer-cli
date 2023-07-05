package actions

import (
	"encoding/json"
	"fmt"
	"net/url"

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
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(apiServerUrl, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator signin: %w", err)
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
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(apiServerUrl, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator singin: %w", err)
	}

	var conf = configuration.NewConfig(name, apiServerUrl, refreshToken)

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
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

	profile := input.TextPrompt("Enter the configuration profile (default: default)")
	if profile == "" {
		profile = "default"
	}

	var conf = configuration.NewConfig(profile, "", "")

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("Configuration %s signed out successfully\n", profile)

	return nil
}

func SignOutOperator(cCtx *cli.Context) error {
	var err error

	if cCtx.Bool("interactive") {
		return SignOutOperatorInteractive(cCtx)
	}

	profile := cCtx.String("profile")
	if profile == "" {
		profile = "default"
	}

	configPath := cCtx.String("config")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	var conf = configuration.NewConfig(profile, "", "")

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("Configuration %s signed out successfully\n", profile)

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
	var accessToken, configPath string
	var conf configuration.Config

	readConfiguration()
	rehydrateTokenConfig(configPath, conf)

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("Access token: %s\n", accessToken)

	return nil
}

func CreateTenant(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	name := cCtx.String("name")
	description := cCtx.String("description")
	imageUrl := cCtx.String("image-url")
	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("image url is not adequate: %w", err)
		}
	}

	settingsString := cCtx.String("settings")
	if settingsString == "" {
		settingsString = "{}"
	}

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, *conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	var settings map[string]interface{}

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("error while parsing json settings: %w", err)
	}

	if response, err = api.CreateTenant(conf.ApiServerUrl, *accessToken, name, &description, &imageUrl, settings); err != nil {
		return fmt.Errorf("error while creating the tenant: %w", err)
	}

	fmt.Printf("Successfully created tenant: %s\n", response.ID)
	return nil
}

func rehydrateTokenConfig(configPath string, conf configuration.Config) (*string, error) {
	var accessToken, refreshToken string
	var err error

	if accessToken, refreshToken, err = api.RefreshAccessToken(conf.ApiServerUrl, conf.RefreshToken); err != nil {
		return nil, fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	conf.RefreshToken = refreshToken

	if err = conf.Store(configPath); err != nil {
		return nil, fmt.Errorf("error while storing file path configuration: %w", err)
	}

	return &accessToken, nil
}

func readConfiguration() (*configuration.Config, string, error) {
	var cCtx *cli.Context
	var configPath string
	var err error
	var profile string

	if cCtx.Bool("interactive") {
		profile, configPath = readConfigFile()
	} else {
		profile = cCtx.String("profile")
		if profile == "" {
			profile = "default"
		}
		configPath = cCtx.String("config")
		if configPath == "" {
			configPath = DEFAULT_FILE_PATH
		}
	}

	var conf = configuration.NewConfig(profile, "", "")

	if err = conf.Load(configPath, profile); err != nil {
		return nil, "", fmt.Errorf("error while loading file path configuration: %w", err)
	}

	return &conf, configPath, nil
}
