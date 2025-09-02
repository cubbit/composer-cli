// Package configuration provides configuration management for the CLI.
package configuration

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type ProfileType string

const (
	ProfileTypeComposer ProfileType = "composer"
	ProfileTypeConsole  ProfileType = "console"
)

type OutputFormat string

const (
	OutputJSON  OutputFormat = "json"
	OutputXML   OutputFormat = "xml"
	OutputYAML  OutputFormat = "yaml"
	OutputHuman OutputFormat = "human"
)

type BaseConfig struct {
	Endpoint string       `toml:"endpoint"`
	Output   OutputFormat `toml:"output"`
}

type Profile struct {
	Inherits  string       `toml:"inherits,omitempty"`
	Type      ProfileType  `toml:"type"`
	Endpoint  string       `toml:"endpoint,omitempty"`
	Output    OutputFormat `toml:"output,omitempty"`
	APIKey    string       `toml:"api_key,omitempty"`
	UpdatedAt time.Time    `toml:"updated_at,omitempty"`
}

type ActiveConfig struct {
	Profile string `toml:"profile"`
}

type Config struct {
	Default BaseConfig          `toml:"default"`
	Active  ActiveConfig        `toml:"active"`
	Profile map[string]*Profile `toml:"-"`

	ConfigPath string `toml:"-"`
	URLs       *URLs  `toml:"-"`
}

type ResolvedProfile struct {
	Name      string
	Type      ProfileType
	Endpoint  string
	Output    OutputFormat
	APIKey    string
	UpdatedAt time.Time
}

type URLs struct {
	BaseURL string `yaml:"base_url"`
	IamURL  string `yaml:"iam"`
	DashURL string `yaml:"dash"`
	ChURL   string `yaml:"ch"`
}

func NewConfig() *Config {
	return &Config{
		Default: BaseConfig{
			Endpoint: constants.BaseAPIURL,
			Output:   OutputHuman,
		},
		Profile: make(map[string]*Profile),
	}
}

func LoadConfig() (*Config, error) {
	var err error
	var configPath string

	if configPath, err = GetDefaultConfigPath(); err != nil {
		return nil, fmt.Errorf("failed to get default config path: %w", err)
	}

	configFile := filepath.Join(configPath, "config.toml")

	if err := os.MkdirAll(configPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	config := NewConfig()
	config.ConfigPath = configPath

	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		return config, nil
	}

	var rawConfig map[string]interface{}
	if _, err := toml.DecodeFile(configFile, &rawConfig); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	if defaultSection, ok := rawConfig["default"]; ok {
		if defaultMap, ok := defaultSection.(map[string]interface{}); ok {
			if endpoint, ok := defaultMap["endpoint"].(string); ok {
				config.Default.Endpoint = endpoint
			}
			if output, ok := defaultMap["output"].(string); ok {
				config.Default.Output = OutputFormat(output)
			}

		}
	}

	if activeSection, ok := rawConfig["active"]; ok {
		if activeMap, ok := activeSection.(map[string]interface{}); ok {
			if profile, ok := activeMap["profile"].(string); ok {
				config.Active.Profile = profile
			}
		}
	}

	_, ok := rawConfig["profile"].(map[string]interface{})
	if ok {

		profiles, ok := rawConfig["profile"].(map[string]interface{})
		if !ok {
			return config, fmt.Errorf("invalid profile section in config file")
		}

		for key, value := range profiles {
			profileName := key
			if profileMap, ok := value.(map[string]interface{}); ok {
				profile := &Profile{}

				if inherits, ok := profileMap["inherits"].(string); ok {
					profile.Inherits = inherits
				}
				if profileType, ok := profileMap["type"].(string); ok {
					profile.Type = ProfileType(profileType)
				}
				if endpoint, ok := profileMap["endpoint"].(string); ok {
					profile.Endpoint = endpoint
				}
				if output, ok := profileMap["output"].(string); ok {
					profile.Output = OutputFormat(output)
				}
				if apiKey, ok := profileMap["api_key"].(string); ok {
					profile.APIKey = apiKey
				}
				if updatedAt, ok := profileMap["updated_at"].(time.Time); ok {
					profile.UpdatedAt = updatedAt
				}

				config.Profile[profileName] = profile
			}

		}
	}

	return config, nil
}

func (c *Config) SaveConfig() error {
	if c.ConfigPath == "" {
		var err error
		if c.ConfigPath, err = GetDefaultConfigPath(); err != nil {
			return fmt.Errorf("failed to get default config path: %w", err)
		}
	}

	configFile := filepath.Join(c.ConfigPath, "config.toml")

	if err := os.MkdirAll(c.ConfigPath, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	file, err := os.Create(configFile)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)

	topLevel := struct {
		Default BaseConfig   `toml:"default"`
		Active  ActiveConfig `toml:"active"`
	}{
		Default: c.Default,
		Active:  c.Active,
	}

	if err := encoder.Encode(topLevel); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	for name, profile := range c.Profile {
		if _, err := fmt.Fprintf(file, "\n[profile.%s]\n", name); err != nil {
			return fmt.Errorf("failed to write profile header: %w", err)
		}

		profileData := make(map[string]interface{})

		if profile.Inherits != "" {
			profileData["inherits"] = profile.Inherits
		}
		if profile.Type != "" {
			profileData["type"] = string(profile.Type)
		}
		if profile.Endpoint != "" && profile.Endpoint != c.Default.Endpoint {
			profileData["endpoint"] = profile.Endpoint
		}
		if profile.Output != "" {
			profileData["output"] = string(profile.Output)
		}
		if profile.APIKey != "" {
			profileData["api_key"] = profile.APIKey
		}
		if !profile.UpdatedAt.IsZero() {
			profileData["updated_at"] = profile.UpdatedAt
		}

		for key, value := range profileData {
			switch v := value.(type) {
			case string:
				if _, err := fmt.Fprintf(file, "  %s = %q\n", key, v); err != nil {
					return fmt.Errorf("failed to write profile field: %w", err)
				}
			case int:
				if _, err := fmt.Fprintf(file, "  %s = %d\n", key, v); err != nil {
					return fmt.Errorf("failed to write profile field: %w", err)
				}
			case time.Time:
				if _, err := fmt.Fprintf(file, "  %s = %s\n", key, v.Format(time.RFC3339Nano)); err != nil {
					return fmt.Errorf("failed to write profile field: %w", err)
				}
			}
		}
	}

	return nil
}

func (c *Config) ResolveProfile(profileName string) (*ResolvedProfile, error) {
	if profileName == "" {
		return nil, errors.New("no profile name found")
	}

	profile, exists := c.Profile[profileName]
	if !exists {
		return nil, fmt.Errorf("profile '%s' not found", profileName)
	}

	resolved := &ResolvedProfile{
		Name:      profileName,
		Type:      profile.Type,
		UpdatedAt: profile.UpdatedAt,
		Endpoint:  profile.Endpoint,
		APIKey:    profile.APIKey,
	}

	if err := c.applyInheritance(resolved, profile, make(map[string]bool)); err != nil {
		return nil, err
	}

	return resolved, nil
}

func SetAPIEndpoint(endpoint string) error {
	var c *Config
	var err error

	if c, err = LoadConfig(); err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if endpoint == "" {
		return errors.New("endpoint cannot be empty")
	}

	c.Default.Endpoint = endpoint

	if err := c.SaveConfig(); err != nil {
		return fmt.Errorf("failed to save updated configuration: %w", err)
	}

	return nil
}

func (c *Config) CreateProfile(name string, profileType ProfileType, endpoint, apiKey string) error {
	if c.Profile == nil {
		c.Profile = make(map[string]*Profile)
	}

	inherits := "default"

	profile := &Profile{
		Inherits:  inherits,
		Type:      profileType,
		APIKey:    apiKey,
		UpdatedAt: time.Now(),
		Endpoint:  endpoint,
	}

	c.Profile[name] = profile
	c.Active.Profile = name
	return nil
}

func (c *Config) UpdateProfile(name, refreshToken string) error {
	profile, exists := c.Profile[name]
	if !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}

	profile.UpdatedAt = time.Now()

	c.Profile[name] = profile

	if err := c.SaveConfig(); err != nil {
		return fmt.Errorf("failed to save updated profile '%s': %w", name, err)
	}

	return nil
}

func (c *Config) DeleteProfile(name string) error {
	if name == "" {
		name = c.Active.Profile
	}

	if _, exists := c.Profile[name]; !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}

	delete(c.Profile, name)
	c.Active.Profile = ""
	return c.SaveConfig()
}

func (c *Config) ListProfiles() []string {
	profiles := []string{"default"}
	for name := range c.Profile {
		profiles = append(profiles, name)
	}
	return profiles
}

func (c *Config) ValidateProfile(profileName string) error {
	resolved, err := c.ResolveProfile(profileName)
	if err != nil {
		return err
	}

	if resolved.Endpoint != "" && resolved.Endpoint != "localhost" {
		if _, err := url.ParseRequestURI(resolved.Endpoint); err != nil {
			return fmt.Errorf("invalid endpoint URL '%s': %w", resolved.Endpoint, err)
		}
	}

	switch resolved.Type {
	case ProfileTypeComposer, ProfileTypeConsole:
		if resolved.APIKey == "" {
			return fmt.Errorf("api_key is required for profile type '%s'", resolved.Type)
		}
	default:
		return fmt.Errorf("invalid profile type '%s'", resolved.Type)
	}

	return nil
}

func (c *Config) GetActiveProfile() string {
	return c.Active.Profile
}

func (c *Config) SetActiveProfile(profileName string) error {
	if _, err := c.ResolveProfile(profileName); err != nil {
		return fmt.Errorf("cannot set active profile '%s': %w", profileName, err)
	}

	c.Active.Profile = profileName
	return c.SaveConfig()
}

func GetDefaultConfigPath() (string, error) {
	if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
		return filepath.Join(xdgConfig, "cubbit"), nil
	}

	// Fall back to ~/.config/cubbit
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	return filepath.Join(homeDir, ".config", "cubbit"), nil
}

func ReadConfig(cmd *cobra.Command, expectedProfileType ProfileType, isLastStep ...bool) (*ResolvedProfile, *URLs, error) {
	var err error
	var profile string
	var interactive bool

	if interactive, err = cmd.Flags().GetBool("interactive"); err != nil {
		return nil, nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if interactive {
		if profile, err = promptForConfigFile(isLastStep...); err != nil {
			return nil, nil, fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
	} else {
		if profile, err = cmd.Flags().GetString("profile"); err != nil {
			return nil, nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
		}

		if profile == "" {
			config, err := LoadConfig()
			if err != nil {
				return nil, nil, fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
			}

			profile = config.GetActiveProfile()
			if profile == "" {
				return nil, nil, fmt.Errorf("no active profile set, please specify a profile using --profile or set an active profile")
			}
		}
	}

	config, err := LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	resolved, err := config.ResolveProfile(profile)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolved.Type != expectedProfileType {
		return nil, nil, fmt.Errorf("profile '%s' has type '%s', expected '%s'", profile, resolved.Type, expectedProfileType)
	}

	if config.URLs == nil {
		urls, err := config.ResolveURLs(profile)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve URLs for profile '%s': %w", profile, err)
		}
		config.URLs = urls
	}

	return resolved, config.URLs, nil
}

func (c *Config) ResolveProfileAndURLs(cmd *cobra.Command, expectedProfileType ProfileType) (*ResolvedProfile, *URLs, error) {
	var err error
	var profile string

	if profile, err = cmd.Flags().GetString("profile"); err != nil {
		return nil, nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if profile == "" {
		config, err := LoadConfig()
		if err != nil {
			return nil, nil, fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
		}

		profile = config.GetActiveProfile()
		if profile == "" {
			return nil, nil, fmt.Errorf("no active profile set, please specify a profile using --profile flag or set an active profile")
		}
	}

	config, err := LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	resolved, err := config.ResolveProfile(profile)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolved.Type != expectedProfileType {
		return nil, nil, fmt.Errorf("profile '%s' has type '%s', expected '%s'", profile, resolved.Type, expectedProfileType)
	}

	if config.URLs == nil {
		urls, err := config.ResolveURLs(profile)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve URLs for profile '%s': %w", profile, err)
		}
		config.URLs = urls
	}

	return resolved, config.URLs, nil
}

func (c *Config) LoadURL(filePath string, envName string) (*URLs, error) {
	var err error
	var data []byte
	var urls URLs

	envFile := filePath + ".env." + envName
	if data, err = os.ReadFile(envFile); err != nil {
		return nil, fmt.Errorf("failed to read environment file %s: %w", envFile, err)
	}

	if err = yaml.Unmarshal(data, &urls); err != nil {
		return nil, fmt.Errorf("failed to unmarshal URL configuration: %w", err)
	}

	if urls.IamURL == "" {
		return nil, fmt.Errorf(constants.ErrorIamConfigNotFound)
	}

	if urls.DashURL == "" {
		return nil, fmt.Errorf(constants.ErrorDashConfigNotFound)
	}

	if urls.ChURL == "" {
		return nil, fmt.Errorf(constants.ErrorChConfigNotFound)
	}

	return &urls, nil
}

func ConfigureAPIServerURL(endpoint string) (*URLs, error) {
	if _, err := url.ParseRequestURI(endpoint); err == nil || endpoint == "" {
		return composeURL(endpoint), nil
	} else {
		config := NewConfig()
		devPath := constants.DefaultFilePath

		urls, err := config.LoadURL(devPath, "local")
		if err != nil {
			return nil, fmt.Errorf("error while loading environment '%s': %w", endpoint, err)
		}
		return urls, nil
	}
}

func (c *Config) ResolveURLs(profileName string) (*URLs, error) {
	resolved, err := c.ResolveProfile(profileName)
	if err != nil {
		return nil, err
	}

	urls, err := ConfigureAPIServerURL(resolved.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to configure URLs for profile '%s': %w", profileName, err)
	}

	return &URLs{
		BaseURL: urls.BaseURL,
		IamURL:  urls.IamURL,
		DashURL: urls.DashURL,
		ChURL:   urls.ChURL,
	}, nil
}

func (c *Config) GetProfileWithUrls(profileName string) (*ResolvedProfile, *URLs, error) {
	profile, err := c.ResolveProfile(profileName)
	if err != nil {
		return nil, nil, err
	}

	urls, err := c.ResolveURLs(profileName)
	if err != nil {
		return nil, nil, err
	}

	return profile, urls, nil
}

func composeURL(apiServerUrl string) *URLs {
	if apiServerUrl == "" {
		apiServerUrl = constants.BaseAPIURL
	}

	return &URLs{
		BaseURL: apiServerUrl,
		IamURL:  apiServerUrl + constants.BaseIamURI,
		DashURL: constants.BaseDashURL,
		ChURL:   apiServerUrl + constants.BaseChURI,
	}
}

func promptForConfigFile(isLastStep ...bool) (string, error) {
	var name, defaultConfigPath string
	var err error

	if defaultConfigPath, err = GetDefaultConfigPath(); err != nil {
		return "", fmt.Errorf("error while getting default config path: %w", err)
	}

	lastStep := len(isLastStep) > 0 && isLastStep[0]

	if _, err := tui.TextInputs("Enter your profile name", lastStep,
		tui.Input{
			Placeholder: "Enter the profile name (default: use active profile)",
			IsPassword:  false,
			Value:       &name,
		}); err != nil {
		return "", fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	config, err := LoadConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config from '%s': %w", defaultConfigPath, err)
	}
	if name == "" {
		name = config.GetActiveProfile()
		if name == "" {
			return "", fmt.Errorf("no default active profile set, please input a profile name")
		}
	}

	if _, exists := config.Profile[name]; !exists {
		return "", fmt.Errorf("profile '%s' does not exist in config file '%s'", name, defaultConfigPath)
	}

	return name, nil
}

func (c *Config) applyInheritance(resolved *ResolvedProfile, profile *Profile, visited map[string]bool) error {
	if profile.Inherits != "" {
		if profile.Inherits == "default" {
			if profile.Output == "" {
				resolved.Output = c.Default.Output
			}

			if profile.Endpoint == "" {
				resolved.Endpoint = c.Default.Endpoint
			}

		} else {
			parentProfile, exists := c.Profile[profile.Inherits]
			if !exists {
				return fmt.Errorf("inherited profile '%s' not found", profile.Inherits)
			}
			if err := c.applyInheritance(resolved, parentProfile, visited); err != nil {
				return err
			}
		}
	}

	c.applyProfileValues(resolved, profile)

	return nil
}

func (c *Config) applyProfileValues(resolved *ResolvedProfile, profile *Profile) {
	if profile.Endpoint != "" {
		resolved.Endpoint = profile.Endpoint
	}
	if profile.Output != "" {
		resolved.Output = profile.Output
	}
	if profile.APIKey != "" {
		resolved.APIKey = profile.APIKey
	}
	if profile.Type != "" {
		resolved.Type = profile.Type
	}
	if !profile.UpdatedAt.IsZero() {
		resolved.UpdatedAt = profile.UpdatedAt
	}
}

func (c *Config) UpdateAPIKey(name, apiKey string) error {
	profile, exists := c.Profile[name]
	if !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}

	if apiKey != "" {
		profile.APIKey = apiKey
	}
	profile.UpdatedAt = time.Now()

	c.Profile[name] = profile

	if err := c.SaveConfig(); err != nil {
		return fmt.Errorf("failed to save updated profile '%s': %w", name, err)
	}

	return nil
}
