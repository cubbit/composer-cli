package configuration

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type SessionType string

const (
	SessionTypeUnauthenticated SessionType = "unauthenticated"
	SessionTypeUser            SessionType = "user"
	SessionTypeAccount         SessionType = "account"
	SessionTypeOperator        SessionType = "operator"
)

type Url struct {
	IamUrl  string `yaml:"iam"`
	HiveUrl string `yaml:"hive"`
	DashUrl string `yaml:"dash"`
	ChUrl   string `yaml:"ch"`
}

type Config struct {
	Name         string      `yaml:"name"`
	Urls         Url         `yaml:"servers_url"`
	RefreshToken string      `yaml:"refresh_token"`
	UpdatedAt    time.Time   `yaml:"updated_at"`
	SessionType  SessionType `yaml:"session_type"`
}

type Session struct {
	Session map[string]Config `yaml:"session"`
}

func NewConfig(sessionType SessionType, name string, urls Url, refreshToken string) *Config {
	return &Config{
		Name:         name,
		Urls:         urls,
		RefreshToken: refreshToken,
		UpdatedAt:    time.Now(),
		SessionType:  sessionType,
	}
}

func (c *Config) LoadUrl(filePath string, envName string) (*Url, error) {
	var err error
	var data []byte
	var urls Url

	if data, err = os.ReadFile(filePath + ".env." + envName); err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &urls); err != nil {
		return nil, err
	}

	if urls.IamUrl == "" {
		return nil, fmt.Errorf(constants.ErrorIamConfigNotFound)
	}

	if urls.HiveUrl == "" {
		return nil, fmt.Errorf(constants.ErrorHiveConfigNotFound)
	}

	if urls.DashUrl == "" {
		return nil, fmt.Errorf(constants.ErrorDashConfigNotFound)
	}

	if urls.ChUrl == "" {
		return nil, fmt.Errorf(constants.ErrorChConfigNotFound)
	}

	return &urls, nil
}

func (c *Config) LoadAndCheckSession(filePath string, name string, expectedSessionType SessionType) error {
	var err error
	var session *Session
	var config Config
	var ok bool

	if session, err = c.loadSession(filePath); err != nil {
		return err
	}
	if config, ok = session.Session[name]; !ok {
		return fmt.Errorf("Session for %s not found in configuration file", name)
	}

	if config.Urls.IamUrl == "" {
		return fmt.Errorf(constants.ErrorIamConfigNotFound)
	}

	if config.Urls.HiveUrl == "" {
		return fmt.Errorf(constants.ErrorHiveConfigNotFound)
	}

	if config.Urls.DashUrl == "" {
		return fmt.Errorf(constants.ErrorDashConfigNotFound)
	}

	if config.Urls.ChUrl == "" {
		return fmt.Errorf(constants.ErrorChConfigNotFound)
	}

	if config.RefreshToken == "" {
		return fmt.Errorf(constants.ErrorTokenNotFound)
	}

	if config.Name == "" {
		return fmt.Errorf(constants.ErrorNameConfigNotFound)
	}

	if config.SessionType == "" {
		return fmt.Errorf(constants.ErrorSessionTypeConfigNotFound)
	}

	if config.SessionType != expectedSessionType {
		return fmt.Errorf(constants.ErrorSessionTypeConfigInvalid)
	}

	*c = config

	return nil
}

func (c *Config) StoreSession(path string) error {
	var err error
	var file *os.File
	var data []byte
	var session *Session

	if session, err = c.loadSession(path); err != nil {
		return err
	}

	if file, err = os.Create(filepath.Join(path, constants.DefaultConfigFileName)); err != nil {
		return err
	}

	defer file.Close()

	if session == nil || session.Session == nil {
		session = &Session{
			Session: make(map[string]Config),
		}
	}

	session.Session[c.Name] = *c

	if data, err = yaml.Marshal(session); err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

func ReadConfig(cmd *cobra.Command, expectedSessionType SessionType, isLastStep ...bool) (*Config, string, error) {
	var configPath string
	var err error
	var profile string
	var interactive bool

	if interactive, err = cmd.Flags().GetBool("interactive"); err != nil {
		return nil, configPath, fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if interactive {
		if configPath, profile, err = promptForConfigFile(isLastStep[0]); err != nil {
			return nil, configPath, fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
	} else {
		if profile, err = cmd.Flags().GetString("profile"); err != nil {
			return nil, configPath, fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
		}

		if configPath, err = cmd.Flags().GetString("config"); err != nil {
			return nil, configPath, fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)

		}
	}

	var conf = NewConfig("", profile, Url{}, "")
	if err = conf.LoadAndCheckSession(configPath, profile, expectedSessionType); err != nil {
		return nil, "", fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return conf, configPath, nil
}

func ConfigureAPIServerURL(sessionType SessionType, apiServerUrl string) (*Url, error) {
	var urls *Url
	var conf = NewConfig(sessionType, "", Url{}, "")

	devPath := constants.DefaultFilePath

	if _, err := url.ParseRequestURI(apiServerUrl); err == nil || apiServerUrl == "" {
		urls = composeURL(apiServerUrl)
	} else {
		if urls, err = conf.LoadUrl(devPath, apiServerUrl); err != nil {
			return nil, fmt.Errorf("error while loading dev path: %w", err)
		}
	}

	return urls, nil
}

func (c *Config) loadSession(path string) (*Session, error) {
	var err error
	var data []byte
	var session Session

	if data, err = os.ReadFile(filepath.Join(path, constants.DefaultConfigFileName)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Session{
				Session: make(map[string]Config),
			}, nil
		}

		return nil, err
	}

	if err = yaml.Unmarshal(data, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func composeURL(apiServerUrl string) *Url {
	if apiServerUrl == "" {
		apiServerUrl = constants.BaseAPIURL
	}

	url := &Url{
		IamUrl:  apiServerUrl + constants.BaseIamURI,
		HiveUrl: apiServerUrl + constants.BaseHiveURI,
		DashUrl: constants.BaseDashURL,
		ChUrl:   apiServerUrl + constants.BaseChURI,
	}
	return url
}

func promptForConfigFile(isLastStep ...bool) (string, string, error) {
	var configPath, name, defaultConfigPath string
	var err error

	if defaultConfigPath, err = GetDefaultConfigPath(); err != nil {
		return defaultConfigPath, name, fmt.Errorf("error while getting default config path: %w", err)
	}

	if _, err := tui.TextInputs("Enter your config file path and name", isLastStep[0], tui.Input{Placeholder: fmt.Sprintf("Enter the config file path to load (default: %s)", defaultConfigPath), IsPassword: false, Value: &configPath}, tui.Input{Placeholder: "Enter the configuration name (default: default)", IsPassword: false, Value: &name}); err != nil {
		return configPath, name, fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if configPath == "" {
		configPath = defaultConfigPath
	}
	if name == "" {
		name = constants.DefaultProfile
	}

	return configPath, name, err
}

func GetDefaultConfigPath() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}
	configFolder := ".config"
	configPath := filepath.Join(homeDir, configFolder)

	return configPath, nil
}
