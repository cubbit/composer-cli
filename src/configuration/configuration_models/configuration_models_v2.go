package configuration_models

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type ConfigV2 struct {
	Version ConfigurationV2Version `toml:"version" json:"version" yaml:"version"`
	Active  ActiveConfigV2         `toml:"active" json:"active" yaml:"active"`
	Profile map[string]ProfileV2   `toml:"profile" json:"profile" yaml:"profile"`
}

func ParseAndValidateConfigV2(path string) (*ConfigV2, error) {
	var cfg ConfigV2
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("failed to decode configuration v2 file %q: %w", path, err)
	}
	if cfg.Version != ConfigurationV2VersionValue {
		return nil, fmt.Errorf("%w: expected version %q but got %q", ErrConfigurationVersionNotSupported, ConfigurationV2VersionValue, cfg.Version)
	}

	if cfg.Active.Profile == "" {
		return nil, fmt.Errorf("%w: no active profile set in configuration", ErrActiveProfileNotSet)
	}

	// check if active profile exists among configured profiles
	if _, ok := cfg.Profile[cfg.Active.Profile]; !ok {
		return nil, fmt.Errorf(
			"%w: active profile %q does not exist in configured profiles",
			ErrActiveProfileNotFound,
			cfg.Active.Profile,
		)
	}

	// check if each profile is valid
	for profileName, profile := range cfg.Profile {
		if err := profile.Validate(profileName); err != nil {
			return nil, fmt.Errorf("invalid profile %q: %w", profileName, err)
		}
	}

	return &cfg, nil
}

func CreateEmptyConfigV2() ConfigV2 {
	return ConfigV2{
		Version: ConfigurationV2VersionValue,
		Active:  ActiveConfigV2{},
		Profile: make(map[string]ProfileV2),
	}
}

func CreateConfigV2(
	profileName string,
	outputFormat OutputFormat,
	apiKey string,
	organizationID string,
	endpoints EndpointsV2,
	updatedAt time.Time,
) ConfigV2 {
	return ConfigV2{
		Version: ConfigurationV2VersionValue,
		Active: ActiveConfigV2{
			Profile: profileName,
		},
		Profile: map[string]ProfileV2{
			profileName: {
				Output:         outputFormat,
				APIKey:         apiKey,
				OrganizationID: organizationID,
				Endpoints:      endpoints,
				UpdatedAt:      updatedAt,
			},
		},
	}
}

func (c *ConfigV2) CreateProfile(
	profileName string,
	outputFormat OutputFormat,
	apiKey string,
	organizationID string,
	endpoints EndpointsV2,
	updatedAt time.Time,
) {
	c.Profile[profileName] = ProfileV2{
		Output:         outputFormat,
		APIKey:         apiKey,
		OrganizationID: organizationID,
		Endpoints:      endpoints,
		UpdatedAt:      updatedAt,
	}
}

func (c *ConfigV2) DeleteProfile(profileName string) error {
	if _, ok := c.Profile[profileName]; !ok {
		return fmt.Errorf("%w: profile %q does not exist", ErrProfileNotFound, profileName)
	}

	delete(c.Profile, profileName)

	if c.Active.Profile == profileName {
		c.Active.Profile = ""
	}

	return nil
}

func (c *ConfigV2) DeleteAllProfiles() error {
	c.Profile = make(map[string]ProfileV2)
	c.Active.Profile = ""

	return nil
}

func (c *ConfigV2) SetActiveProfile(profileName string) error {
	if _, ok := c.Profile[profileName]; !ok {
		return fmt.Errorf("%w: profile %q does not exist", ErrProfileNotFound, profileName)
	}

	c.Active.Profile = profileName

	return nil
}

func (c ConfigV2) GetProfile(profileName string) (ProfileV2, error) {
	profile, ok := c.Profile[profileName]
	if !ok {
		return ProfileV2{}, fmt.Errorf("%w: profile %q does not exist", ErrProfileNotFound, profileName)
	}
	return profile, nil
}

func (c ConfigV2) GetActiveProfileName() string {
	return c.Active.Profile
}

func (c ConfigV2) GetActiveProfile() ProfileV2 {
	return c.Profile[c.Active.Profile]
}

func (c ConfigV2) GetEndpoints() EndpointsV2 {
	return c.GetActiveProfile().Endpoints
}

func (c ConfigV2) GetAPIKey() string {
	return c.GetActiveProfile().APIKey
}

func (c ConfigV2) GetOrganizationID() string {
	return c.GetActiveProfile().OrganizationID
}

func (c ConfigV2) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	return toml.NewEncoder(file).Encode(c)
}

type ConfigurationV2Version string

const (
	ConfigurationV2VersionValue ConfigurationV2Version = "v2"
)

type ActiveConfigV2 struct {
	Profile string `toml:"profile" json:"profile" yaml:"profile"`
}

type ProfileV2 struct {
	Output         OutputFormat `toml:"output" json:"output" yaml:"output"`
	APIKey         string       `toml:"api_key" json:"api_key" yaml:"api_key"`
	UpdatedAt      time.Time    `toml:"updated_at" json:"updated_at" yaml:"updated_at"`
	OrganizationID string       `toml:"organization_id" json:"organization_id" yaml:"organization_id"`
	Endpoints      EndpointsV2  `toml:"endpoints" json:"endpoints" yaml:"endpoints"`
}

func (p ProfileV2) Validate(profileName string) error {
	if p.Output == "" {
		return fmt.Errorf("%w: %s, output format not set", ErrInvalidProfile, profileName)
	}
	if p.APIKey == "" {
		return fmt.Errorf("%w: %s, API key not set", ErrInvalidProfile, profileName)
	}
	if p.OrganizationID == "" {
		return fmt.Errorf("%w: %s, organization ID not set", ErrInvalidProfile, profileName)
	}
	if err := p.Endpoints.Validate(); err != nil {
		return fmt.Errorf("%w: %s, invalid endpoints configuration: %w", ErrInvalidProfile, profileName, err)
	}
	return nil
}

type OutputFormat string

const (
	OutputHuman OutputFormat = "human"
)

type EndpointsV2 struct {
	IAM  string `toml:"iam" json:"iam" yaml:"iam"`
	Dash string `toml:"dash" json:"dash" yaml:"dash"`
	CH   string `toml:"ch" json:"ch" yaml:"ch"`
}

func (e EndpointsV2) Validate() error {
	if e.IAM == "" {
		return fmt.Errorf("IAM endpoint not set")
	}
	if e.Dash == "" {
		return fmt.Errorf("dashboard endpoint not set")
	}
	if e.CH == "" {
		return fmt.Errorf("CH endpoint not set")
	}
	return nil
}
