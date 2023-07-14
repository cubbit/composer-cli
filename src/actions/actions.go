package actions

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_PATH = "./"

func convertUrls(apiServerUrl string) *configuration.Url {
	if apiServerUrl == "" {
		apiServerUrl = "https://api.cubbit.eu"
	}

	urls := &configuration.Url{
		IamUrl:  apiServerUrl + "/iam",
		HiveUrl: apiServerUrl + "/hive",
	}
	return urls
}

func apiServerUrlConfiguration(apiServerUrl string) (*configuration.Url, error) {
	var urls *configuration.Url
	var conf = configuration.NewConfig("", configuration.Url{}, "")

	devPath := DEFAULT_FILE_PATH

	if _, err := url.ParseRequestURI(apiServerUrl); err == nil || apiServerUrl == "" {
		urls = convertUrls(apiServerUrl)
	} else {
		fmt.Printf("configuring endpoint for %s\n", apiServerUrl)

		if urls, err = conf.LoadUrl(devPath, apiServerUrl); err != nil {
			return nil, fmt.Errorf("error while loading dev path: %w", err)
		}
	}

	return urls, nil
}

func CreateOperatorInteractive(cCtx *cli.Context) error {
	var err error
	var urls *configuration.Url

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu)")

	if urls, err = apiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	firstName := input.TextPrompt("Enter first name:")
	lastName := input.TextPrompt("Enter last name:")
	email := input.TextPrompt("Enter email:")
	password := input.PasswordPrompt("Enter password:")

	if err = api.CreateOperator(*urls, firstName, lastName, email, password); err != nil {
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

	if urls, err = apiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	if err = api.CreateOperator(*urls, firstName, lastName, email, password); err != nil {
		return fmt.Errorf("error while creating operator: %w", err)
	}

	fmt.Printf("Operator %s created successfully\n", email)
	return nil
}

func SignInOperatorInteractive(cCtx *cli.Context) error {
	var err error
	var code, refreshToken string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url
	var conf = configuration.NewConfig("", configuration.Url{}, "")

	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu)")

	if urls, err = apiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
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

	profile := input.TextPrompt("Enter the configuration profile (default: default)")
	if profile == "" {
		profile = "default"
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator signin: %w", err)
	}

	conf = configuration.NewConfig(profile, *urls, refreshToken)

	conf.Store(configPath)

	fmt.Printf("User %s signed in successfully\n", email)

	return nil
}

func SignInOperator(cCtx *cli.Context) error {
	var err error
	var apiServerUrl, email, password, code, refreshToken string
	var challenge *api.ChallengeResponseModel
	var urls *configuration.Url

	if cCtx.Bool("interactive") {
		return SignInOperatorInteractive(cCtx)
	}

	configPath := cCtx.String("config")
	if configPath == "" {
		configPath = DEFAULT_FILE_PATH
	}

	profile := cCtx.String("profile")
	if profile == "" {
		profile = "default"
	}

	email = cCtx.String("email")
	password = cCtx.String("password")
	code = cCtx.String("code")
	apiServerUrl = cCtx.String("api-server-url")

	if urls, err = apiServerUrlConfiguration(apiServerUrl); err != nil {
		return fmt.Errorf("error while configuri api server url %w", err)
	}

	if challenge, err = api.GenerateOperatorChallenge(*urls, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(*urls, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator singin: %w", err)
	}

	var confs = configuration.NewConfig(profile, *urls, refreshToken)

	if err = confs.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("User %s signed in successfully\n", email)

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

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

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

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

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
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if err = conf.Store(configPath); err != nil {
		return fmt.Errorf("error while storing file path configuration: %w", err)
	}

	fmt.Printf("Access token: %s\n", *accessToken)

	return nil
}

func CreateTenant(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	name := cCtx.String("name")
	description := cCtx.Args().First()
	if len(description) > 200 {
		return fmt.Errorf("tenant description is over 200 characters: %w", err)
	}

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
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	var settings map[string]interface{}

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("error while parsing json settings: %w", err)
	}

	if response, err = api.CreateTenant(conf.Urls, *accessToken, name, &description, &imageUrl, settings); err != nil {
		return fmt.Errorf("error while creating the tenant: %w", err)
	}

	fmt.Printf("Successfully created tenant: %s\n", response.ID)
	return nil
}

func rehydrateTokenConfig(configPath string, conf *configuration.Config) (*string, error) {
	var accessToken, refreshToken string
	var err error

	if accessToken, refreshToken, err = api.RefreshAccessToken(conf.Urls, conf.RefreshToken); err != nil {
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

	var conf = configuration.NewConfig(profile, configuration.Url{}, "")

	if err = conf.Load(configPath, profile); err != nil {
		return nil, "", fmt.Errorf("error while loading file path configuration: %w", err)
	}

	return &conf, configPath, nil
}

func ListTenant(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var operator *api.Operator
	var tenants *api.TenantList

	fmt.Println("these are your tenants")

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("error while retrieving operator id: %w", err)
	}
	if tenants, err = api.ListTenant(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("error while retrieving tenant list: %w", err)
	}

	verbose := cCtx.Bool("verbose")
	l := cCtx.Bool("l")

	for _, tenant := range tenants.Tenants {
		if verbose {
			fmt.Printf("%s %s %s ", tenant.ID, tenant.Name, *tenant.Description)
		} else {
			fmt.Printf("%s ", tenant.Name)
		}
		if l {
			fmt.Println()
		}
	}

	return nil
}

func RemoveTenant(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath, deleteTenantToken string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel

	id := cCtx.String("id")
	name := cCtx.String("name")
	email := cCtx.String("email")
	password := cCtx.String("password")
	code := cCtx.String("code")

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
	}

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id == "" {
		var operator *api.Operator
		var tenants *api.TenantList

		if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("error while retrieving operator id: %w", err)
		}
		if tenants, err = api.ListTenant(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("error while retrieving tenant list: %w", err)
		}
		for _, tenant := range tenants.Tenants {
			if name == tenant.Name {
				id = tenant.ID
			}
		}
		if id == "" {
			fmt.Printf("Tenant %s not found\n", name)
			return nil
		}
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if deleteTenantToken, err = api.ForgeOperatorDeleteTenantToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("error while forging operator delete token: %w", err)
	}

	if err = api.RemoveTenant(conf.Urls, *accessToken, id, deleteTenantToken); err != nil {
		return fmt.Errorf("error while deleting tenant: %w", err)
	}

	fmt.Printf("tenant %s removed successfully\n", id)
	return nil
}

func DescribeTenant(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var tenants *api.TenantList
	var operator *api.Operator

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	id := cCtx.String("id")
	name := cCtx.String("name")
	format := cCtx.String("format")
	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("error while retrieving operator id: %w", err)
	}
	if tenants, err = api.ListTenant(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("error while retrieving tenant list: %w", err)
	}

	switch {
	case id == "" && name == "":
		return fmt.Errorf("error while retrieving tenant description: %w", err)
	case name == "":
		for _, tenant := range tenants.Tenants {
			if id == tenant.ID {
				FormatTenant(format, tenant)
			}
		}
	case id == "":
		for _, tenant := range tenants.Tenants {
			if name == tenant.Name {
				FormatTenant(format, tenant)
			}
		}
	default:
		return fmt.Errorf("error, tenant name or id incorrect: %w", err)
	}

	return nil
}

func FormatTenant(format string, tenant *api.Tenant) error {
	switch {
	case format == "default":
		fmt.Printf("ID: %s\n", tenant.ID)
		fmt.Printf("Name: %s\n", tenant.Name)

		if tenant.Description != nil {
			fmt.Printf("Description: %s\n", *tenant.Description)
		}

		fmt.Printf("OwnerID: %s\n", tenant.OwnerID)
		fmt.Printf("CreatedAt: %s\n", tenant.CreatedAt)

		if tenant.DeletedAt != nil {
			fmt.Printf("DeletedAt: %s\n", tenant.DeletedAt)
		}

		if tenant.ImageUrl != nil && *tenant.ImageUrl != "" {
			fmt.Printf("ImageUrl: %s\n", *tenant.ImageUrl)
		}

		fmt.Printf("Settings:\n")
		for key, value := range tenant.Settings {
			fmt.Printf(" - %s: %s\n", key, value)
		}

	case format == "json":
		formatJson, err := json.Marshal(api.Tenant{ID: tenant.ID, Name: tenant.Name, Description: tenant.Description, OwnerID: tenant.OwnerID, CreatedAt: tenant.CreatedAt, DeletedAt: tenant.DeletedAt, ImageUrl: tenant.ImageUrl, Settings: tenant.Settings})

		if err != nil {
			return fmt.Errorf("error while opening json file: %w", err)
		}

		fmt.Println(string(formatJson))

	case format == "csv":
		records := [][]string{
			{"ID", "Name", "Description", "OwnerID", "CreatedAt", "DeletedAt", "ImageUrl"},
		}
		for key := range tenant.Settings {
			fmt.Printf(",%s", key)
		}
		var values []string
		values = append(values, tenant.ID, tenant.Name, *tenant.Description, tenant.OwnerID, tenant.CreatedAt.String(), *tenant.ImageUrl)

		if tenant.DeletedAt != nil {
			values = append(values, tenant.DeletedAt.String())
		} else {
			values = append(values, "")
		}

		records = append(records, values)

		w := csv.NewWriter(os.Stdout)

		for _, record := range records {
			if err := w.Write(record); err != nil {
				return fmt.Errorf("error writing record to csv: %w", err)
			}
		}

		w.Flush()

		if err := w.Error(); err != nil {
			return fmt.Errorf("error occurred during the Flush: %w", err)
		}
	}

	fmt.Println()
	return nil
}

func EditTenantDescription(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	name := cCtx.String("name")
	id := cCtx.String("id")

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
	}

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if cCtx.Args().Len() != 1 {
		return fmt.Errorf("invalid image url: %w", err)
	}

	description := cCtx.Args().First()
	if len(description) > 200 {
		fmt.Println(description)
		return fmt.Errorf("tenant description is over 200 characters: %w", err)
	}

	if err = api.EditTenantDescription(conf.Urls, *accessToken, id, description); err != nil {
		return fmt.Errorf("error while retrieving tenant list: %w", err)
	}

	fmt.Printf("tenant %s description updated successfully\n", id)
	return nil
}

func EditTenantImage(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	name := cCtx.String("name")
	id := cCtx.String("id")

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
	}

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if cCtx.Args().Len() != 1 {
		return fmt.Errorf("invalid umage url: %w", err)
	}

	image := cCtx.Args().First()
	if image != "" {
		if _, err := url.ParseRequestURI(image); err != nil {
			return fmt.Errorf("image url is not adequate: %w", err)
		}
	}

	if err = api.EditTenantImage(conf.Urls, *accessToken, id, image); err != nil {
		return fmt.Errorf("error while retrieving tenant list: %w", err)
	}

	fmt.Printf("tenant %s image updated successfully\n", id)
	return nil
}

func CreateSwarm(cCtx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config
	var operator *api.Operator

	name := cCtx.String("name")
	description := cCtx.String("description")

	configurationString := cCtx.String("configuration")
	if configurationString == "" {
		configurationString = "{}"
	}

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	var configuration map[string]interface{}

	if err = json.Unmarshal([]byte(configurationString), &configuration); err != nil {
		return fmt.Errorf("error while parsing json configuration: %w", err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("error while retrieving operator id: %w", err)
	}

	if response, err = api.CreateSwarm(conf.Urls, *accessToken, operator.ID, name, &description, configuration); err != nil {
		return fmt.Errorf("error while creating the swarm: %w", err)
	}

	fmt.Printf("Successfully created swarm: %s\n", response.ID)
	return nil
}
