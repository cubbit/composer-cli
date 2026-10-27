package configuration_handler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func TestNewConfigurationHandler(t *testing.T) {
	h := NewConfigurationHandler()
	if h.loadType != LoadTypeNotLoaded {
		t.Errorf("expected loadType %q, got %q", LoadTypeNotLoaded, h.loadType)
	}
	if h.configFile != nil {
		t.Error("expected configFile to be nil")
	}
	if h.configPath != "" {
		t.Errorf("expected empty configPath, got %q", h.configPath)
	}
	if h.requestedProfile != nil {
		t.Error("expected requestedProfile to be nil")
	}
	if h.endpoints != nil {
		t.Error("expected endpoints to be nil")
	}
	if h.apiKey != nil {
		t.Error("expected apiKey to be nil")
	}
}

// --- getConfigurationPath ---

func TestGetConfigurationPath_FromFlag(t *testing.T) {
	h := NewConfigurationHandler()
	cmd := &cobra.Command{}
	cmd.Flags().String("config-path", "", "")
	cmd.Flags().Set("config-path", "/custom/path/config.toml")

	path, err := h.getConfigurationPath(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if path != "/custom/path/config.toml" {
		t.Errorf("expected /custom/path/config.toml, got %q", path)
	}
}

func TestGetConfigurationPath_FromEnv(t *testing.T) {
	h := NewConfigurationHandler()
	cmd := &cobra.Command{}
	cmd.Flags().String("config-path", "", "")
	cmd.Flags().Set("config-path", "")

	t.Setenv(configuration_models.ConfigurationPathEnvVariable, "/env/path")

	path, err := h.getConfigurationPath(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := filepath.Join("/env/path", filepath.Base(configuration_models.ConfigurationDefaultDirName), configuration_models.ConfigurationFileName)
	if path != expected {
		t.Errorf("expected %q, got %q", expected, path)
	}
}

func TestGetConfigurationPath_Default(t *testing.T) {
	h := NewConfigurationHandler()
	cmd := &cobra.Command{}
	cmd.Flags().String("config-path", "", "")

	// The default path is only reached when the env override is absent, and CI
	// runners export it. Clear it so the test measures the default, not the host.
	t.Setenv(configuration_models.ConfigurationPathEnvVariable, "")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to get home dir: %v", err)
	}

	path, err := h.getConfigurationPath(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := filepath.Join(homeDir, configuration_models.ConfigurationDefaultDirName, configuration_models.ConfigurationFileName)
	if path != expected {
		t.Errorf("expected %q, got %q", expected, path)
	}
}

// --- Load helpers ---

func newRootCmd() *cobra.Command {
	root := &cobra.Command{Use: "cubbit"}
	root.Flags().String("config-path", "", "")
	root.Flags().String("profile", "", "")
	root.Flags().String("endpoints", "", "")
	return root
}

func addSubCommand(parent *cobra.Command, use string) *cobra.Command {
	child := &cobra.Command{Use: use}
	parent.AddCommand(child)
	return child
}

func writeTempConfig(t *testing.T, dir string) string {
	t.Helper()
	cfg := configuration_models.ConfigV2{
		Version: configuration_models.ConfigurationV2VersionValue,
		Active:  configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				Output:         configuration_models.OutputHuman,
				APIKey:         "key-123",
				OrganizationID: "org-abc",
				Endpoints: configuration_models.EndpointsV2{
					IAM:  "https://iam.example.com",
					Dash: "https://dash.example.com",
					CH:   "https://ch.example.com",
				},
			},
		},
	}
	path := filepath.Join(dir, "config.toml")
	if err := cfg.Save(path); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}
	return path
}

// --- Load ---

func TestLoad_AlreadyLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard

	err := h.Load(nil, nil)
	if err == nil {
		t.Fatal("expected error for already loaded, got nil")
	}
	if !strings.Contains(err.Error(), "already loaded") {
		t.Errorf("expected already loaded error, got: %v", err)
	}
}

func TestLoad_Standard(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir)

	h := NewConfigurationHandler()
	root := newRootCmd()
	cmd := addSubCommand(root, "tenant")
	cmd.AddCommand(&cobra.Command{Use: "list"})
	child := &cobra.Command{Use: "list"}
	cmd.AddCommand(child)
	child.Flags().String("profile", "", "")
	child.Flags().String("config-path", "", "")
	child.Flags().Set("config-path", configPath)

	err := child.Execute()
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	err = h.Load(child, []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.loadType != LoadTypeStandard {
		t.Errorf("expected LoadTypeStandard, got %q", h.loadType)
	}
	if h.configFile == nil {
		t.Fatal("expected configFile to be set")
	}
	if h.endpoints == nil {
		t.Fatal("expected endpoints to be set")
	}
	if h.apiKey == nil || *h.apiKey != "key-123" {
		t.Errorf("expected apiKey 'key-123', got %v", h.apiKey)
	}
}

func TestLoad_StandardWithProfileFlag(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir)

	h := NewConfigurationHandler()
	root := newRootCmd()
	cmd := addSubCommand(root, "tenant")
	child := &cobra.Command{Use: "list"}
	cmd.AddCommand(child)
	child.Flags().String("profile", "", "")
	child.Flags().String("config-path", "", "")
	child.Flags().Set("config-path", configPath)
	child.Flags().Set("profile", "prod")

	err := child.Execute()
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	err = h.Load(child, []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.loadType != LoadTypeStandard {
		t.Errorf("expected LoadTypeStandard, got %q", h.loadType)
	}
}

func TestLoad_StandardMissingConfig(t *testing.T) {
	dir := t.TempDir()

	h := NewConfigurationHandler()
	root := newRootCmd()
	cmd := addSubCommand(root, "tenant")
	child := &cobra.Command{Use: "list"}
	cmd.AddCommand(child)
	child.Flags().String("config-path", "", "")
	child.Flags().Set("config-path", filepath.Join(dir, "nonexistent.toml"))

	err := h.Load(child, []string{})
	if err == nil {
		t.Fatal("expected error for missing config, got nil")
	}
	if !strings.Contains(err.Error(), "failed to load configuration") {
		t.Errorf("expected load error, got: %v", err)
	}
}

func TestLoad_ForConfigCommands(t *testing.T) {
	h := NewConfigurationHandler()
	root := newRootCmd()
	cmd := addSubCommand(root, "config")
	child := &cobra.Command{Use: "init"}
	cmd.AddCommand(child)
	child.Flags().String("config-path", "", "")

	err := child.Execute()
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	err = h.Load(child, []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.loadType != LoadTypeForConfigCommands {
		t.Errorf("expected LoadTypeForConfigCommands, got %q", h.loadType)
	}
}

func TestLoad_ForAuthCommandsNoExistingConfig(t *testing.T) {
	h := NewConfigurationHandler()
	root := newRootCmd()
	cmd := addSubCommand(root, "auth")
	child := &cobra.Command{Use: "login"}
	cmd.AddCommand(child)
	child.Flags().String("config-path", "", "")
	child.Flags().String("endpoints", "", "")
	child.Flags().Set("config-path", filepath.Join(t.TempDir(), "nonexistent.toml"))

	err := child.Execute()
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	err = h.Load(child, []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.loadType != LoadTypeForAuthCommands {
		t.Errorf("expected LoadTypeForAuthCommands, got %q", h.loadType)
	}
	if h.endpoints == nil {
		t.Fatal("expected endpoints to be set from defaults")
	}
	if h.endpoints.IAM != constants.BaseAPIURL+constants.BaseIamURI {
		t.Errorf("expected default IAM endpoint, got %q", h.endpoints.IAM)
	}
	if h.configFile != nil {
		t.Error("expected configFile to be nil when no config exists")
	}
}

func TestLoad_ForAuthCommandsWithExistingConfig(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir)

	h := NewConfigurationHandler()
	root := newRootCmd()
	cmd := addSubCommand(root, "auth")
	child := &cobra.Command{Use: "login"}
	cmd.AddCommand(child)
	child.Flags().String("config-path", "", "")
	child.Flags().String("endpoints", "", "")
	child.Flags().Set("config-path", configPath)

	err := child.Execute()
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	err = h.Load(child, []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.loadType != LoadTypeForAuthCommands {
		t.Errorf("expected LoadTypeForAuthCommands, got %q", h.loadType)
	}
	if h.configFile == nil {
		t.Fatal("expected configFile to be set")
	}
	if h.endpoints == nil {
		t.Fatal("expected endpoints to be set")
	}
	if h.endpoints.IAM != "https://iam.example.com" {
		t.Errorf("expected IAM from config, got %q", h.endpoints.IAM)
	}
}

// --- loadForConfigCommands ---

func TestLoadForConfigCommands(t *testing.T) {
	h := NewConfigurationHandler()
	err := h.loadForConfigCommands(nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.loadType != LoadTypeForConfigCommands {
		t.Errorf("expected LoadTypeForConfigCommands, got %q", h.loadType)
	}
}

// --- getEndpointsFromDefaults ---

func TestGetEndpointsFromDefaults(t *testing.T) {
	ep := getEndpointsFromDefaults()
	if ep.IAM != constants.BaseAPIURL+constants.BaseIamURI {
		t.Errorf("expected IAM %s, got %q", constants.BaseAPIURL+constants.BaseIamURI, ep.IAM)
	}
	if ep.Dash != constants.BaseAPIURL+constants.BaseDashURI {
		t.Errorf("expected Dash %s, got %q", constants.BaseAPIURL+constants.BaseDashURI, ep.Dash)
	}
	if ep.CH != constants.BaseAPIURL+constants.BaseChURI {
		t.Errorf("expected CH %s, got %q", constants.BaseAPIURL+constants.BaseChURI, ep.CH)
	}
}

// --- GetConfigFilePath ---

func TestGetConfigFilePath_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	_, err := h.GetConfigFilePath()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
	if !strings.Contains(err.Error(), "not loaded") {
		t.Errorf("expected not loaded error, got: %v", err)
	}
}

func TestGetConfigFilePath_Success(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeForConfigCommands
	h.configPath = "/test/path/config.toml"

	path, err := h.GetConfigFilePath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if path != "/test/path/config.toml" {
		t.Errorf("expected /test/path/config.toml, got %q", path)
	}
}

// --- GetConfigFile ---

func TestGetConfigFile_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	_, err := h.GetConfigFile()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestGetConfigFile_Success(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	apiKey := "test-key"
	h.configFile = &configuration_models.ConfigV2{
		Version: configuration_models.ConfigurationV2VersionValue,
		Active:  configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				APIKey: apiKey,
				Endpoints: configuration_models.EndpointsV2{
					IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
				},
				OrganizationID: "org",
				Output:         configuration_models.OutputHuman,
			},
		},
	}
	h.apiKey = &apiKey
	ep := configuration_models.EndpointsV2{IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com"}
	h.endpoints = &ep

	cfg, err := h.GetConfigFile()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Version != configuration_models.ConfigurationV2VersionValue {
		t.Errorf("version mismatch")
	}
}

func TestGetConfigFile_NotStandardLoadType(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeForConfigCommands
	_, err := h.GetConfigFile()
	if err == nil {
		t.Fatal("expected error for LoadTypeForConfigCommands, got nil")
	}
}

// --- GetAPIKey ---

func TestGetAPIKey_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	_, err := h.GetAPIKey()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestGetAPIKey_Success(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	key := "my-key"
	h.apiKey = &key

	k, err := h.GetAPIKey()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if k != "my-key" {
		t.Errorf("expected 'my-key', got %q", k)
	}
}

// --- GetEndpoints ---

func TestGetEndpoints_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	_, err := h.GetEndpoints()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestGetEndpoints_SuccessStandard(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	ep := configuration_models.EndpointsV2{IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com"}
	h.endpoints = &ep

	e, err := h.GetEndpoints()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.IAM != "https://iam.com" {
		t.Errorf("expected IAM https://iam.com, got %q", e.IAM)
	}
}

func TestGetEndpoints_SuccessAuthCommands(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeForAuthCommands
	ep := configuration_models.EndpointsV2{IAM: "https://iam.auth.com", Dash: "https://dash.auth.com", CH: "https://ch.auth.com"}
	h.endpoints = &ep

	e, err := h.GetEndpoints()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.IAM != "https://iam.auth.com" {
		t.Errorf("expected IAM https://iam.auth.com, got %q", e.IAM)
	}
}

// --- GetProfiles ---

func TestGetProfiles_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	_, err := h.GetProfiles()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestGetProfiles_Success(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configFile = &configuration_models.ConfigV2{
		Version: configuration_models.ConfigurationV2VersionValue,
		Active:  configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				APIKey: "key", OrganizationID: "org", Output: configuration_models.OutputHuman,
				Endpoints: configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"},
			},
			"staging": {
				APIKey: "key2", OrganizationID: "org2", Output: configuration_models.OutputHuman,
				Endpoints: configuration_models.EndpointsV2{IAM: "iam2", Dash: "dash2", CH: "ch2"},
			},
		},
	}

	profiles, err := h.GetProfiles()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(profiles) != 2 {
		t.Errorf("expected 2 profiles, got %d", len(profiles))
	}
	if _, ok := profiles["staging"]; !ok {
		t.Error("expected 'staging' profile")
	}
}

// --- GetActiveProfile ---

func TestGetActiveProfile_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	_, err := h.GetActiveProfile()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestGetActiveProfile_Success(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configFile = &configuration_models.ConfigV2{
		Version: configuration_models.ConfigurationV2VersionValue,
		Active:  configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				APIKey: "key-123", OrganizationID: "org-abc", Output: configuration_models.OutputHuman,
				Endpoints: configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"},
			},
		},
	}

	p, err := h.GetActiveProfile()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.APIKey != "key-123" || p.OrganizationID != "org-abc" {
		t.Errorf("profile fields mismatch: %+v", p)
	}
}

// --- GetActiveProfileName ---

func TestGetActiveProfileName_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	_, err := h.GetActiveProfileName()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestGetActiveProfileName_Success(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configFile = &configuration_models.ConfigV2{
		Active: configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				APIKey: "key", OrganizationID: "org", Output: configuration_models.OutputHuman,
				Endpoints: configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"},
			},
		},
	}

	name, err := h.GetActiveProfileName()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "prod" {
		t.Errorf("expected 'prod', got %q", name)
	}
}

// --- Save ---

func TestSave_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	err := h.Save()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
	if !strings.Contains(err.Error(), "not loaded") {
		t.Errorf("expected not loaded error, got: %v", err)
	}
}

func TestSave_NoPath(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard

	err := h.Save()
	if err == nil {
		t.Fatal("expected error for empty path, got nil")
	}
	if !strings.Contains(err.Error(), "not loaded") {
		t.Errorf("expected not loaded error, got: %v", err)
	}
}

func TestSave_NilConfig(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configPath = "/some/path.toml"

	err := h.Save()
	if err == nil {
		t.Fatal("expected error for nil config, got nil")
	}
	if !strings.Contains(err.Error(), "no target configuration entity") {
		t.Errorf("expected no config entity error, got: %v", err)
	}
}

func TestSave_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")

	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configPath = path
	h.configFile = &configuration_models.ConfigV2{
		Version: configuration_models.ConfigurationV2VersionValue,
		Active:  configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				APIKey: "key", OrganizationID: "org", Output: configuration_models.OutputHuman,
				Endpoints: configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"},
			},
		},
	}

	if err := h.Save(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("expected file to be saved")
	}
}

// --- CreateProfile ---

func TestCreateProfile_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeNotLoaded
	err := h.CreateProfile("test", "key", "org")
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestCreateProfile_ConfigCommandsNotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	h.loadType = LoadTypeForConfigCommands
	err := h.CreateProfile("test", "key", "org")
	if err == nil {
		t.Fatal("expected error for config commands load type, got nil")
	}
}

func TestCreateProfile_WithoutExistingConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")

	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configPath = path
	ep := configuration_models.EndpointsV2{IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com"}
	h.endpoints = &ep

	if err := h.CreateProfile("new-profile", "key-456", "org-xyz"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if h.configFile == nil {
		t.Fatal("expected configFile to be set")
	}
	p, ok := h.configFile.Profile["new-profile"]
	if !ok {
		t.Fatal("expected profile 'new-profile' to exist")
	}
	if p.APIKey != "key-456" || p.OrganizationID != "org-xyz" {
		t.Errorf("profile fields mismatch: %+v", p)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("expected config file to be saved")
	}
}

func TestCreateProfile_WithExistingConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")
	existingConfig := configuration_models.CreateConfigV2(
		"existing", configuration_models.OutputHuman, "key-1", "org-1",
		configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"}, time.Now(),
	)
	if err := existingConfig.Save(path); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}
	parsed, err := configuration_models.ParseAndValidateConfigV2(path)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	ep := configuration_models.EndpointsV2{IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com"}
	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configPath = path
	h.configFile = parsed
	h.endpoints = &ep

	if err := h.CreateProfile("new-profile", "key-456", "org-xyz"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := h.configFile.Profile["existing"]; !ok {
		t.Error("expected existing profile to still exist")
	}
	if _, ok := h.configFile.Profile["new-profile"]; !ok {
		t.Error("expected new profile to exist")
	}
}

// --- DeleteProfile ---

func TestDeleteProfile_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	err := h.DeleteProfile("test")
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestDeleteProfile_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")

	cfg := &configuration_models.ConfigV2{
		Version: configuration_models.ConfigurationV2VersionValue,
		Active:  configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				APIKey: "key", OrganizationID: "org", Output: configuration_models.OutputHuman,
				Endpoints: configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"},
			},
		},
	}
	if err := cfg.Save(path); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configPath = path
	h.configFile = cfg

	if err := h.DeleteProfile("prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := h.configFile.Profile["prod"]; ok {
		t.Error("expected profile to be deleted")
	}
	var reloaded configuration_models.ConfigV2
	if _, err := toml.DecodeFile(path, &reloaded); err != nil {
		t.Fatalf("failed to reload: %v", err)
	}
	if _, ok := reloaded.Profile["prod"]; ok {
		t.Error("expected profile to be deleted from saved file")
	}
}

// --- DeleteAllProfiles ---

func TestDeleteAllProfiles_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	err := h.DeleteAllProfiles()
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestDeleteAllProfiles_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")

	h := NewConfigurationHandler()
	h.loadType = LoadTypeStandard
	h.configPath = path
	h.configFile = &configuration_models.ConfigV2{
		Version: configuration_models.ConfigurationV2VersionValue,
		Active:  configuration_models.ActiveConfigV2{Profile: "prod"},
		Profile: map[string]configuration_models.ProfileV2{
			"prod": {
				APIKey: "key", OrganizationID: "org", Output: configuration_models.OutputHuman,
				Endpoints: configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"},
			},
		},
	}
	if err := h.configFile.Save(path); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	if err := h.DeleteAllProfiles(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(h.configFile.Profile) != 0 {
		t.Errorf("expected empty profiles, got %d", len(h.configFile.Profile))
	}
	var reloaded configuration_models.ConfigV2
	if _, err := toml.DecodeFile(path, &reloaded); err != nil {
		t.Fatalf("failed to reload: %v", err)
	}
	if len(reloaded.Profile) != 0 {
		t.Errorf("expected empty profiles on disk, got %d", len(reloaded.Profile))
	}
}

// --- SwitchProfile ---

func TestSwitchProfile_NotLoaded(t *testing.T) {
	h := NewConfigurationHandler()
	err := h.SwitchProfile("test")
	if err == nil {
		t.Fatal("expected error for not loaded, got nil")
	}
}

func TestSwitchProfile_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")
	ts := time.Now()
	cfg := configuration_models.CreateConfigV2(
		"prod", configuration_models.OutputHuman, "key", "org",
		configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"}, ts,
	)
	cfg.CreateProfile("staging", configuration_models.OutputHuman, "key2", "org2",
		configuration_models.EndpointsV2{IAM: "iam2", Dash: "dash2", CH: "ch2"}, ts,
	)
	if err := cfg.Save(path); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
	parsed, _ := configuration_models.ParseAndValidateConfigV2(path)

	h := new(ConfigurationHandler)
	h.loadType = LoadTypeStandard
	h.configPath = path
	h.configFile = parsed

	if err := h.SwitchProfile("staging"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if h.configFile.Active.Profile != "staging" {
		t.Errorf("expected active profile 'staging', got %q", h.configFile.Active.Profile)
	}

	var reloaded configuration_models.ConfigV2
	if _, err := toml.DecodeFile(path, &reloaded); err != nil {
		t.Fatalf("failed to reload: %v", err)
	}
	if reloaded.Active.Profile != "staging" {
		t.Errorf("expected active profile 'staging' on disk, got %q", reloaded.Active.Profile)
	}
}

func TestSwitchProfile_NotExists(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")
	cfg := configuration_models.CreateConfigV2(
		"prod", configuration_models.OutputHuman, "key", "org",
		configuration_models.EndpointsV2{IAM: "iam", Dash: "dash", CH: "ch"}, time.Now(),
	)
	if err := cfg.Save(path); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
	parsed, _ := configuration_models.ParseAndValidateConfigV2(path)

	h := new(ConfigurationHandler)
	h.loadType = LoadTypeStandard
	h.configPath = path
	h.configFile = parsed

	err := h.SwitchProfile("nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent profile, got nil")
	}
	if !strings.Contains(err.Error(), "profile not found") {
		t.Errorf("expected profile not found error, got: %v", err)
	}
}

// --- MockConfigurationHandler ---

func TestNewMockConfigurationHandler(t *testing.T) {
	m := NewMockConfigurationHandler()
	if m.LoadFunc == nil {
		t.Error("expected LoadFunc to be set")
	}
	if m.SaveFunc == nil {
		t.Error("expected SaveFunc to be set")
	}
	if m.CreateProfileFunc == nil {
		t.Error("expected CreateProfileFunc to be set")
	}

	if err := m.Load(nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := m.Save(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := m.CreateProfile("test", "key", "org"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMockConfigurationHandler_OverrideFuncs(t *testing.T) {
	m := NewMockConfigurationHandler()
	m.LoadFunc = func(cmd *cobra.Command, args []string) error {
		return configuration_models.ErrConfigurationIsAlreadyLoaded
	}

	err := m.Load(nil, nil)
	if err != configuration_models.ErrConfigurationIsAlreadyLoaded {
		t.Errorf("expected ErrConfigurationIsAlreadyLoaded, got: %v", err)
	}
}
