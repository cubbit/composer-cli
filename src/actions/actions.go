package actions

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"

	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_PATH = "./"

func CreateOperatorInteractive(cCtx *cli.Context) error {
	var err error
	apiServerUrl := input.TextPrompt("Enter the api server url: (default https://api.cubbit.eu/iam)")
	if apiServerUrl == "" {
		apiServerUrl = "https://api.cubbit.eu/iam"
	}

	firstName := input.TextPrompt("Enter first name:")
	lastName := input.TextPrompt("Enter last name:")
	email := input.TextPrompt("Enter email:")
	password := input.PasswordPrompt("Enter password:")

	if err = api.CreateOperator(apiServerUrl, firstName, lastName, email, password); err != nil {
		return fmt.Errorf("error while creating operator: %w", err)
	}

	fmt.Printf("Operator %s created successfully\n", email)
	return nil
}

func CreateOperator(cCtx *cli.Context) error {
	var err error
	var apiServerUrl, email, password, firstName, lastName string

	if cCtx.Bool("interactive") {
		return CreateOperatorInteractive(cCtx)
	}

	email = cCtx.String("email")
	password = cCtx.String("password")
	firstName = cCtx.String("first-name")
	lastName = cCtx.String("last-name")
	apiServerUrl = cCtx.String("api-server-url")

	if err = api.CreateOperator(apiServerUrl, firstName, lastName, email, password); err != nil {
		return fmt.Errorf("error while creating operator: %w", err)
	}

	fmt.Printf("Operator %s created successfully\n", email)
	return nil
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

	profile := input.TextPrompt("Enter the configuration profile (default: default)")
	if profile == "" {
		profile = "default"
	}

	if challenge, err = api.GenerateOperatorChallenge(apiServerUrl, email); err != nil {
		return fmt.Errorf("error while generating operator challenge: %w", err)
	}

	if refreshToken, err = api.PerformOperatorSignin(apiServerUrl, email, password, challenge, code); err != nil {
		return fmt.Errorf("error while performing operator signin: %w", err)
	}

	var conf = configuration.NewConfig(profile, apiServerUrl, refreshToken)

	conf.Store(configPath)

	fmt.Printf("User %s signed in successfully\n", email)

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

	profile := cCtx.String("profile")
	if profile == "" {
		profile = "default"
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

	var conf = configuration.NewConfig(profile, apiServerUrl, refreshToken)

	if err = conf.Store(configPath); err != nil {
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
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	if conf, configPath, err = readConfiguration(); err != nil {
		return fmt.Errorf("error while loading file path configuration: %w", err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, *conf); err != nil {
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
	if accessToken, err = rehydrateTokenConfig(configPath, *conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if operator, err = api.GetOperatorSelf(conf.ApiServerUrl, *accessToken); err != nil {
		return fmt.Errorf("error while retrieving operator id: %w", err)
	}
	if tenants, err = api.ListTenant(conf.ApiServerUrl, *accessToken, operator.ID); err != nil {
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

<<<<<<< HEAD
<<<<<<< HEAD
=======
// func RemoveTenant(cCtx *cli.Context) error {
// 	var err error
// 	var accessToken *string
// 	var configPath string
// 	var conf *configuration.Config
// 	var tenant *api.Tenant
// 	var tenants *api.TenantList

// 	if conf, configPath, err = readConfiguration(); err != nil {
// 		return fmt.Errorf("error while loading file path configuration: %w", err)
// 	}

// 	if accessToken, err = rehydrateTokenConfig(configPath, *conf); err != nil {
// 		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
// 	}

// 	if tenants, err = api.RemoveTenant(conf.ApiServerUrl, *accessToken, tenant.ID); err != nil {
// 		return fmt.Errorf("error while retrieving tenant list: %w", err)
// 	}

// 	fmt.Printf("Tenant removed: %s %s ", tenant.ID, tenant.Name)

// 	return nil
// }

>>>>>>> 09ed954 (feat(tenant): gives information abpout tenant)
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
	if accessToken, err = rehydrateTokenConfig(configPath, *conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

<<<<<<< HEAD
	id := cCtx.String("id")
	name := cCtx.String("name")
	format := cCtx.String("format")
	if operator, err = api.GetOperatorSelf(conf.ApiServerUrl, *accessToken); err != nil {
		return fmt.Errorf("error while retrieving operator id: %w", err)
	}
	if tenants, err = api.ListTenant(conf.ApiServerUrl, *accessToken, operator.ID); err != nil {
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
=======
func RemoveTenant(cCtx *cli.Context) error {
	id := cCtx.String("id")
	fmt.Println(id, "tenant removed")
>>>>>>> a796b82 (feat(tenant): remove tenants command)
=======
	id := cCtx.String("id")
	name := cCtx.String("name")
	format := cCtx.String("format")
	if operator, err = api.GetOperatorSelf(conf.ApiServerUrl, *accessToken); err != nil {
		return fmt.Errorf("error while retrieving operator id: %w", err)
	}
	if tenants, err = api.ListTenant(conf.ApiServerUrl, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("error while retrieving tenant list: %w", err)
	}

	switch {
	case id == "" && name == "":
		return fmt.Errorf("error while retrieving tenant description: %w", err)
	case name == "":
		for _, tenant := range tenants.Tenants {
			if id == tenant.ID {
				fmt.Printf("%s\n %s ", tenant.ID, *tenant.Description)
			}
		}
	case id == "":
		for _, tenant := range tenants.Tenants {
			if name == tenant.Name {
				FormatTenant(format, tenant)
			}
		}
	}

>>>>>>> 09ed954 (feat(tenant): gives information abpout tenant)
	return nil
}

func FormatTenant(format string, tenant *api.Tenant) {
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
			panic(err)
		}

		fmt.Println(string(formatJson))

	case format == "csv":
		fmt.Printf("ID,Name,")
		fmt.Printf("Description,")
		fmt.Printf("OwnerID,CreatedAt,")
		fmt.Printf("DeletedAt,")
		fmt.Printf("ImageUrl")
		for key := range tenant.Settings {
			fmt.Printf(",%s", key)
		}

		fmt.Println()

		fmt.Printf("%s,", tenant.ID)
		fmt.Printf("%s,", tenant.Name)

		if tenant.Description != nil { // invoca funzione

			fmt.Printf("\"%s\",", *tenant.Description) //replaceAll
		} else {
			fmt.Printf(",")
		}
	}
	fmt.Printf("%s,", tenant.OwnerID)
	fmt.Printf("%s,", tenant.CreatedAt)

	if tenant.DeletedAt != nil {
		fmt.Printf("%s,", tenant.DeletedAt)
	} else {
		fmt.Printf(",")
	}

	if tenant.ImageUrl != nil && *tenant.ImageUrl != "" {
		fmt.Printf("%s", *tenant.ImageUrl)
	}

	for _, value := range tenant.Settings {
		fmt.Printf(",%s", value)
	}
	fmt.Println()
}

type Writer struct {
	Comma   rune
	UseCRLF bool
	w       *bufio.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Comma: ',',
		w:     bufio.NewWriter(w),
	}
}

func validDelim(r rune) bool {
	return r != 0 && r != '"' && r != '\r' && r != '\n' && utf8.ValidRune(r) && r != utf8.RuneError
}

func (w *Writer) Write(record []string) error {
	var errInvalidDelim = errors.New("csv: invalid field or comment delimiter")
	if !validDelim(w.Comma) {
		return errInvalidDelim
	}

	for n, field := range record {
		if n > 0 {
			if _, err := w.w.WriteRune(w.Comma); err != nil {
				return err
			}
		}

		if err := w.w.WriteByte('"'); err != nil {
			return err
		}
		for len(field) > 0 {
			i := strings.IndexAny(field, "\"\r\n")
			if i < 0 {
				i = len(field)
			}

			if _, err := w.w.WriteString(field[:i]); err != nil {
				return err
			}
			field = field[i:]

			if len(field) > 0 {
				var err error
				switch field[0] {
				case '"':
					_, err = w.w.WriteString(`""`)
				case '\r':
					if !w.UseCRLF {
						err = w.w.WriteByte('\r')
					}
				case '\n':
					if w.UseCRLF {
						_, err = w.w.WriteString("\r\n")
					} else {
						err = w.w.WriteByte('\n')
					}
				}
				field = field[1:]
				if err != nil {
					return err
				}
			}
		}
		if err := w.w.WriteByte('"'); err != nil {
			return err
		}
	}
	var err error
	if w.UseCRLF {
		_, err = w.w.WriteString("\r\n")
	} else {
		err = w.w.WriteByte('\n')
	}
	return err
}
