package configuration

import (
	"errors"
	"fmt"
	"os"
	"time"

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

func (c *Config) load(filePath string) (*Session, error) {
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

func (c *Config) Load(filePath string, name string) error {
	var err error
	var session *Session
	var config Config
	var ok bool

	if session, err = c.load(filePath); err != nil {
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

func (c *Config) Store(filePath string) error {
	var err error
	var file *os.File
	var data []byte
	var session *Session

	if session, err = c.load(filePath); err != nil {
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
