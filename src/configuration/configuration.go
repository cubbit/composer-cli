package configuration

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/input"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type Url struct {
	IamUrl  string `yaml:"iam"`
	HiveUrl string `yaml:"hive"`
}

type Config struct {
	Name         string    `yaml:"name"`
	Urls         Url       `yaml:"servers_url"`
	RefreshToken string    `yaml:"refresh_token"`
	UpdatedAt    time.Time `yaml:"updated_at"`
}

type Session struct {
	Session map[string]Config `yaml:"session"`
}

func NewConfig(name string, urls Url, refreshToken string) Config {
	return Config{
		Name:         name,
		Urls:         urls,
		RefreshToken: refreshToken,
		UpdatedAt:    time.Now(),
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
		return nil, fmt.Errorf("iam api server url is not defined in configuration")
	}

	if urls.HiveUrl == "" {
		return nil, fmt.Errorf("hive api server url is not defined in configuration")
	}

	return &urls, nil
}

func (c *Config) LoadAndCheckSession(filePath string, name string) error {
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
		return fmt.Errorf("iam api server url is not defined in configuration")
	}

	if config.Urls.HiveUrl == "" {
		return fmt.Errorf("hive api server url is not defined in configuration")
	}

	if config.RefreshToken == "" {
		return fmt.Errorf("refresh token is not defined in configuration")
	}

	if config.Name == "" {
		return fmt.Errorf("name is not defined in configuration")
	}

	*c = config

	return nil
}

func (c *Config) StoreSession(filePath string) error {
	var err error
	var file *os.File
	var data []byte
	var session *Session

	if session, err = c.loadSession(filePath); err != nil {
		return err
	}

	if file, err = os.Create(filePath + ".config"); err != nil {
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

func ReadConfig(ctx *cli.Context) (*Config, string, error) {
	var configPath string
	var err error
	var profile string

	if ctx.Bool("interactive") {
		profile, configPath = promptForConfigFile()
	} else {
		profile = ctx.String("profile")
		if profile == "" {
			profile = "default"
		}
		configPath = ctx.String("config")
		if configPath == "" {
			configPath = constants.DefaultFilePath
		}
	}

	var conf = NewConfig(profile, Url{}, "")

	if err = conf.LoadAndCheckSession(configPath, profile); err != nil {
		return nil, "", fmt.Errorf("error while loading file path configuration: %w", err)
	}

	return &conf, configPath, nil
}

func ConfigureAPIServerURL(apiServerUrl string) (*Url, error) {
	var urls *Url
	var conf = NewConfig("", Url{}, "")

	devPath := constants.DefaultFilePath

	if _, err := url.ParseRequestURI(apiServerUrl); err == nil || apiServerUrl == "" {
		urls = composeURL(apiServerUrl)
	} else {
		fmt.Printf("configuring endpoint for %s\n", apiServerUrl)

		if urls, err = conf.LoadUrl(devPath, apiServerUrl); err != nil {
			return nil, fmt.Errorf("error while loading dev path: %w", err)
		}
	}

	return urls, nil
}

func (c *Config) loadSession(filePath string) (*Session, error) {
	var err error
	var data []byte
	var session Session

	if data, err = os.ReadFile(filePath + ".config"); err != nil {
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
	}
	return url
}

func promptForConfigFile() (string, string) {
	configPath := input.TextPrompt("Enter the config file to load (default: ./)")
	if configPath == "" {
		configPath = constants.DefaultFilePath
	}

	name := input.TextPrompt("Enter the configuration name (default: default)")
	if name == "" {
		name = "default"
	}

	return configPath, name
}
