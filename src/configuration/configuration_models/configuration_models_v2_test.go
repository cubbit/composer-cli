package configuration_models

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestConfigV2Save(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")

	cfg := ConfigV2{
		Version: ConfigurationV2VersionValue,
		Active:  ActiveConfigV2{Profile: "prod"},
		Profile: map[string]ProfileV2{
			"prod": {
				Output:         OutputHuman,
				APIKey:         "test-key",
				OrganizationID: "org-123",
				Endpoints: EndpointsV2{
					IAM:  "https://iam.example.com",
					Dash: "https://dash.example.com",
					CH:   "https://ch.example.com",
				},
			},
		},
	}

	if err := cfg.Save(path); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read saved file: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, `version = "v2"`) {
		t.Errorf("expected version in output, got: %s", content)
	}
	if !strings.Contains(content, `profile`) {
		t.Errorf("expected profile in output, got: %s", content)
	}
}

func TestConfigV2Save_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")

	ts := time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)
	original := ConfigV2{
		Version: ConfigurationV2VersionValue,
		Active:  ActiveConfigV2{Profile: "prod"},
		Profile: map[string]ProfileV2{
			"prod": {
				Output:         OutputHuman,
				APIKey:         "key-123",
				UpdatedAt:      ts,
				OrganizationID: "org-abc",
				Endpoints: EndpointsV2{
					IAM:  "https://iam.example.com",
					Dash: "https://dash.example.com",
					CH:   "https://ch.example.com",
				},
			},
		},
	}

	if err := original.Save(path); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := ParseAndValidateConfigV2(path)
	if err != nil {
		t.Fatalf("ParseAndValidateConfigV2 failed: %v", err)
	}

	if loaded.Version != original.Version {
		t.Errorf("version mismatch: %q vs %q", loaded.Version, original.Version)
	}
	if loaded.Active.Profile != original.Active.Profile {
		t.Errorf("active profile mismatch: %q vs %q", loaded.Active.Profile, original.Active.Profile)
	}

	lp := loaded.Profile["prod"]
	op := original.Profile["prod"]
	if lp.APIKey != op.APIKey {
		t.Errorf("api_key mismatch: %q vs %q", lp.APIKey, op.APIKey)
	}
	if lp.OrganizationID != op.OrganizationID {
		t.Errorf("organization_id mismatch: %q vs %q", lp.OrganizationID, op.OrganizationID)
	}
	if lp.Endpoints.IAM != op.Endpoints.IAM {
		t.Errorf("endpoints.iam mismatch: %q vs %q", lp.Endpoints.IAM, op.Endpoints.IAM)
	}
	if lp.Endpoints.Dash != op.Endpoints.Dash {
		t.Errorf("endpoints.dash mismatch: %q vs %q", lp.Endpoints.Dash, op.Endpoints.Dash)
	}
	if lp.Endpoints.CH != op.Endpoints.CH {
		t.Errorf("endpoints.ch mismatch: %q vs %q", lp.Endpoints.CH, op.Endpoints.CH)
	}
}

func TestParseAndValidateConfigV2_Valid(t *testing.T) {
	path := writeTestConfigV2(t, `
version = "v2"

[active]
profile = "prod"

[profile.prod]
output = "human"
api_key = "key-123"
organization_id = "org-abc"

[profile.prod.endpoints]
iam = "https://iam.example.com"
dash = "https://dash.example.com"
ch = "https://ch.example.com"
`)

	cfg, err := ParseAndValidateConfigV2(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Version != ConfigurationV2VersionValue {
		t.Errorf("expected version %q, got %q", ConfigurationV2VersionValue, cfg.Version)
	}
	if cfg.Active.Profile != "prod" {
		t.Errorf("expected active profile 'prod', got %q", cfg.Active.Profile)
	}
	if cfg.Profile["prod"].APIKey != "key-123" {
		t.Errorf("expected api_key 'key-123', got %q", cfg.Profile["prod"].APIKey)
	}
}

func TestParseAndValidateConfigV2_WrongVersion(t *testing.T) {
	path := writeTestConfigV2(t, `
version = "v1"

[active]
profile = "prod"

[profile.prod]
output = "human"
api_key = "key"
organization_id = "org"
`)

	_, err := ParseAndValidateConfigV2(path)
	if err == nil {
		t.Fatal("expected error for wrong version, got nil")
	}
	if !strings.Contains(err.Error(), "configuration version not supported") {
		t.Errorf("expected version not supported error, got: %v", err)
	}
}

func TestParseAndValidateConfigV2_MissingActiveProfile(t *testing.T) {
	path := writeTestConfigV2(t, `
version = "v2"

[active]
profile = ""

[profile.prod]
output = "human"
api_key = "key"
organization_id = "org"
`)

	_, err := ParseAndValidateConfigV2(path)
	if err == nil {
		t.Fatal("expected error for missing active profile, got nil")
	}
	if !strings.Contains(err.Error(), "no active profile set") {
		t.Errorf("expected no active profile error, got: %v", err)
	}
}

func TestParseAndValidateConfigV2_ActiveProfileNotFound(t *testing.T) {
	path := writeTestConfigV2(t, `
version = "v2"

[active]
profile = "nonexistent"

[profile.prod]
output = "human"
api_key = "key"
organization_id = "org"
`)

	_, err := ParseAndValidateConfigV2(path)
	if err == nil {
		t.Fatal("expected error for active profile not found, got nil")
	}
	if !strings.Contains(err.Error(), "active profile") || !strings.Contains(err.Error(), "does not exist") {
		t.Errorf("expected active profile not found error, got: %v", err)
	}
}

func TestParseAndValidateConfigV2_InvalidProfile(t *testing.T) {
	path := writeTestConfigV2(t, `
version = "v2"

[active]
profile = "prod"

[profile.prod]
output = "human"
api_key = ""
organization_id = "org"
`)

	_, err := ParseAndValidateConfigV2(path)
	if err == nil {
		t.Fatal("expected error for invalid profile, got nil")
	}
	if !strings.Contains(err.Error(), "invalid profile") {
		t.Errorf("expected invalid profile error, got: %v", err)
	}
}

func TestProfileV2Validate_Valid(t *testing.T) {
	p := ProfileV2{
		Output:         OutputHuman,
		APIKey:         "key",
		OrganizationID: "org",
		Endpoints: EndpointsV2{
			IAM:  "https://iam.example.com",
			Dash: "https://dash.example.com",
			CH:   "https://ch.example.com",
		},
	}
	if err := p.Validate("test"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestProfileV2Validate_MissingOutput(t *testing.T) {
	p := ProfileV2{
		APIKey:         "key",
		OrganizationID: "org",
		Endpoints: EndpointsV2{
			IAM:  "https://iam.example.com",
			Dash: "https://dash.example.com",
			CH:   "https://ch.example.com",
		},
	}
	err := p.Validate("test")
	if err == nil {
		t.Fatal("expected error for missing output, got nil")
	}
	if !strings.Contains(err.Error(), "output format not set") {
		t.Errorf("expected output format error, got: %v", err)
	}
}

func TestProfileV2Validate_MissingAPIKey(t *testing.T) {
	p := ProfileV2{
		Output:         OutputHuman,
		OrganizationID: "org",
		Endpoints: EndpointsV2{
			IAM:  "https://iam.example.com",
			Dash: "https://dash.example.com",
			CH:   "https://ch.example.com",
		},
	}
	err := p.Validate("test")
	if err == nil {
		t.Fatal("expected error for missing api_key, got nil")
	}
	if !strings.Contains(err.Error(), "API key not set") {
		t.Errorf("expected API key error, got: %v", err)
	}
}

func TestProfileV2Validate_MissingOrganizationID(t *testing.T) {
	p := ProfileV2{
		Output: OutputHuman,
		APIKey: "key",
		Endpoints: EndpointsV2{
			IAM:  "https://iam.example.com",
			Dash: "https://dash.example.com",
			CH:   "https://ch.example.com",
		},
	}
	err := p.Validate("test")
	if err == nil {
		t.Fatal("expected error for missing organization_id, got nil")
	}
	if !strings.Contains(err.Error(), "organization ID not set") {
		t.Errorf("expected organization ID error, got: %v", err)
	}
}

func TestProfileV2Validate_InvalidEndpoints(t *testing.T) {
	p := ProfileV2{
		Output:         OutputHuman,
		APIKey:         "key",
		OrganizationID: "org",
		Endpoints:      EndpointsV2{},
	}
	err := p.Validate("test")
	if err == nil {
		t.Fatal("expected error for invalid endpoints, got nil")
	}
	if !strings.Contains(err.Error(), "invalid endpoints") {
		t.Errorf("expected invalid endpoints error, got: %v", err)
	}
}

func TestEndpointsV2Validate_Valid(t *testing.T) {
	e := EndpointsV2{
		IAM:  "https://iam.example.com",
		Dash: "https://dash.example.com",
		CH:   "https://ch.example.com",
	}
	if err := e.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEndpointsV2Validate_MissingIAM(t *testing.T) {
	e := EndpointsV2{
		Dash: "https://dash.example.com",
		CH:   "https://ch.example.com",
	}
	err := e.Validate()
	if err == nil {
		t.Fatal("expected error for missing IAM endpoint")
	}
}

func TestEndpointsV2Validate_MissingDash(t *testing.T) {
	e := EndpointsV2{
		IAM: "https://iam.example.com",
		CH:  "https://ch.example.com",
	}
	err := e.Validate()
	if err == nil {
		t.Fatal("expected error for missing dashboard endpoint")
	}
}

func TestEndpointsV2Validate_MissingCH(t *testing.T) {
	e := EndpointsV2{
		IAM:  "https://iam.example.com",
		Dash: "https://dash.example.com",
	}
	err := e.Validate()
	if err == nil {
		t.Fatal("expected error for missing CH endpoint")
	}
}

func TestErrors_AreNonNil(t *testing.T) {
	if ErrConfigurationVersionNotSupported == nil {
		t.Error("ErrConfigurationVersionNotSupported should not be nil")
	}
	if ErrConfigurationIsAlreadyLoaded == nil {
		t.Error("ErrConfigurationIsAlreadyLoaded should not be nil")
	}
	if ErrConfigurationIsNotLoaded == nil {
		t.Error("ErrConfigraitonIsNotLoaded should not be nil")
	}
	if ErrActiveProfileNotFound == nil {
		t.Error("ErrActiveProfileNotFound should not be nil")
	}
	if ErrActiveProfileNotSet == nil {
		t.Error("ErrActiveProfileNotSet should not be nil")
	}
	if ErrInvalidProfile == nil {
		t.Error("ErrInvalidProfile should not be nil")
	}
}

func TestConstants_AreNonEmpty(t *testing.T) {
	if ConfigurationPathEnvVariable == "" {
		t.Error("ConfigurationPathEnvVariable should not be empty")
	}
	if ConfigurationDefaultDirName == "" {
		t.Error("ConfigurationDefaultDirName should not be empty")
	}
	if ConfigurationFileName == "" {
		t.Error("ConfigurationFileName should not be empty")
	}
}

func TestCreateEmptyConfigV2(t *testing.T) {
	cfg := CreateEmptyConfigV2()
	if cfg.Version != ConfigurationV2VersionValue {
		t.Errorf("expected version %q, got %q", ConfigurationV2VersionValue, cfg.Version)
	}
	if cfg.Active.Profile != "" {
		t.Errorf("expected empty active profile, got %q", cfg.Active.Profile)
	}
	if len(cfg.Profile) != 0 {
		t.Errorf("expected empty profiles map, got %d entries", len(cfg.Profile))
	}
}

func TestCreateConfigV2(t *testing.T) {
	ts := time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)
	endpoints := EndpointsV2{
		IAM:  "https://iam.example.com",
		Dash: "https://dash.example.com",
		CH:   "https://ch.example.com",
	}
	cfg := CreateConfigV2("prod", OutputHuman, "key-123", "org-abc", endpoints, ts)
	if cfg.Version != ConfigurationV2VersionValue {
		t.Errorf("expected version %q, got %q", ConfigurationV2VersionValue, cfg.Version)
	}
	if cfg.Active.Profile != "prod" {
		t.Errorf("expected active profile 'prod', got %q", cfg.Active.Profile)
	}
	p, ok := cfg.Profile["prod"]
	if !ok {
		t.Fatal("expected profile 'prod' to exist")
	}
	if p.APIKey != "key-123" || p.OrganizationID != "org-abc" || p.Output != OutputHuman {
		t.Errorf("profile fields mismatch: %+v", p)
	}
	if !p.UpdatedAt.Equal(ts) {
		t.Errorf("updated_at mismatch: %v vs %v", p.UpdatedAt, ts)
	}
}

func TestConfigV2CreateProfile(t *testing.T) {
	cfg := CreateEmptyConfigV2()
	ts := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	cfg.CreateProfile("staging", OutputHuman, "key-staging", "org-staging", EndpointsV2{
		IAM: "https://iam.staging.com", Dash: "https://dash.staging.com", CH: "https://ch.staging.com",
	}, ts)

	p, ok := cfg.Profile["staging"]
	if !ok {
		t.Fatal("expected profile 'staging' to exist after CreateProfile")
	}
	if p.APIKey != "key-staging" || p.OrganizationID != "org-staging" {
		t.Errorf("profile fields mismatch: %+v", p)
	}
	if !p.UpdatedAt.Equal(ts) {
		t.Errorf("updated_at mismatch")
	}
	if p.Endpoints.IAM != "https://iam.staging.com" {
		t.Errorf("endpoint IAM mismatch")
	}
}

func TestConfigV2DeleteProfile(t *testing.T) {
	cfg := CreateConfigV2("prod", OutputHuman, "key", "org", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())

	if err := cfg.DeleteProfile("prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := cfg.Profile["prod"]; ok {
		t.Error("expected profile 'prod' to be deleted")
	}
}

func TestConfigV2DeleteProfile_NotExists(t *testing.T) {
	cfg := CreateEmptyConfigV2()
	err := cfg.DeleteProfile("nonexistent")
	if err == nil {
		t.Fatal("expected error for deleting nonexistent profile, got nil")
	}
	if !strings.Contains(err.Error(), "profile not found") {
		t.Errorf("expected profile not found error, got: %v", err)
	}
}

func TestConfigV2DeleteProfile_ClearsActive(t *testing.T) {
	cfg := CreateConfigV2("prod", OutputHuman, "key", "org", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())

	if err := cfg.DeleteProfile("prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Active.Profile != "" {
		t.Errorf("expected active profile to be cleared, got %q", cfg.Active.Profile)
	}
}

func TestConfigV2DeleteProfile_KeepsActiveIfDifferent(t *testing.T) {
	cfg := CreateConfigV2("prod", OutputHuman, "key", "org", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())
	cfg.CreateProfile("staging", OutputHuman, "key2", "org2", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())
	cfg.Active.Profile = "prod"

	if err := cfg.DeleteProfile("staging"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Active.Profile != "prod" {
		t.Errorf("expected active profile to remain 'prod', got %q", cfg.Active.Profile)
	}
}

func TestConfigV2DeleteAllProfiles(t *testing.T) {
	cfg := CreateConfigV2("prod", OutputHuman, "key", "org", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())
	cfg.CreateProfile("staging", OutputHuman, "key2", "org2", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())

	if err := cfg.DeleteAllProfiles(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Profile) != 0 {
		t.Errorf("expected empty profiles, got %d", len(cfg.Profile))
	}
	if cfg.Active.Profile != "" {
		t.Errorf("expected active profile to be cleared, got %q", cfg.Active.Profile)
	}
}

func TestConfigV2SetActiveProfile(t *testing.T) {
	cfg := CreateConfigV2("prod", OutputHuman, "key", "org", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())
	cfg.Active.Profile = ""

	if err := cfg.SetActiveProfile("prod"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestConfigV2SetActiveProfile_NotExists(t *testing.T) {
	cfg := CreateEmptyConfigV2()
	err := cfg.SetActiveProfile("nonexistent")
	if err == nil {
		t.Fatal("expected error for setting nonexistent active profile, got nil")
	}
	if !strings.Contains(err.Error(), "profile not found") {
		t.Errorf("expected profile not found error, got: %v", err)
	}
}

func TestConfigV2GetProfile(t *testing.T) {
	cfg := CreateConfigV2("prod", OutputHuman, "key-123", "org-abc", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())

	p, err := cfg.GetProfile("prod")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.APIKey != "key-123" {
		t.Errorf("expected API key 'key-123', got %q", p.APIKey)
	}
	if p.OrganizationID != "org-abc" {
		t.Errorf("expected org ID 'org-abc', got %q", p.OrganizationID)
	}
}

func TestConfigV2GetProfile_NotExists(t *testing.T) {
	cfg := CreateEmptyConfigV2()
	_, err := cfg.GetProfile("nonexistent")
	if err == nil {
		t.Fatal("expected error for getting nonexistent profile, got nil")
	}
	if !strings.Contains(err.Error(), "profile not found") {
		t.Errorf("expected profile not found error, got: %v", err)
	}
}

func TestConfigV2Getters(t *testing.T) {
	cfg := CreateConfigV2("prod", OutputHuman, "key-123", "org-abc", EndpointsV2{
		IAM: "https://iam.com", Dash: "https://dash.com", CH: "https://ch.com",
	}, time.Now())

	if name := cfg.GetActiveProfileName(); name != "prod" {
		t.Errorf("GetActiveProfileName: expected 'prod', got %q", name)
	}

	p := cfg.GetActiveProfile()
	if p.APIKey != "key-123" || p.OrganizationID != "org-abc" {
		t.Errorf("GetActiveProfile: fields mismatch: %+v", p)
	}

	if ep := cfg.GetEndpoints(); ep.IAM != "https://iam.com" || ep.Dash != "https://dash.com" || ep.CH != "https://ch.com" {
		t.Errorf("GetEndpoints: mismatch: %+v", ep)
	}

	if key := cfg.GetAPIKey(); key != "key-123" {
		t.Errorf("GetAPIKey: expected 'key-123', got %q", key)
	}

	if org := cfg.GetOrganizationID(); org != "org-abc" {
		t.Errorf("GetOrganizationID: expected 'org-abc', got %q", org)
	}
}

func TestEndpointsV2InitsValidate_ValidBase(t *testing.T) {
	e := EndpointsV2Inits{Base: "https://api.example.com"}
	if err := e.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEndpointsV2InitsValidate_ValidIndividual(t *testing.T) {
	e := EndpointsV2Inits{
		IAM:  "https://iam.example.com",
		Dash: "https://dash.example.com",
		CH:   "https://ch.example.com",
	}
	if err := e.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEndpointsV2InitsValidate_InvalidNoEndpoints(t *testing.T) {
	e := EndpointsV2Inits{}
	err := e.Validate()
	if err == nil {
		t.Fatal("expected error for empty endpoints, got nil")
	}
	if !strings.Contains(err.Error(), "either base endpoint or all individual endpoints") {
		t.Errorf("expected validation error, got: %v", err)
	}
}

func TestEndpointsV2InitsValidate_InvalidPartialIndividual(t *testing.T) {
	e := EndpointsV2Inits{IAM: "https://iam.example.com"}
	err := e.Validate()
	if err == nil {
		t.Fatal("expected error for partial individual endpoints, got nil")
	}
	if !strings.Contains(err.Error(), "either base endpoint or all individual endpoints") {
		t.Errorf("expected validation error, got: %v", err)
	}
}

func TestEndpointsV2InitsValidate_BaseWithPartial(t *testing.T) {
	e := EndpointsV2Inits{Base: "https://api.example.com", IAM: "https://custom-iam.example.com"}
	if err := e.Validate(); err != nil {
		t.Fatalf("expected base to skip individual validation, got: %v", err)
	}
}

func TestEndpointsV2InitsToEndpointsV2_Base(t *testing.T) {
	e := EndpointsV2Inits{Base: "https://api.example.com"}
	ep := e.ToEndpointsV2()
	if ep.IAM != "https://api.example.com/iam" {
		t.Errorf("IAM: expected https://api.example.com/iam, got %q", ep.IAM)
	}
	if ep.Dash != "https://api.example.com/api" {
		t.Errorf("Dash: expected https://api.example.com/api, got %q", ep.Dash)
	}
	if ep.CH != "https://api.example.com/composer-hub" {
		t.Errorf("CH: expected https://api.example.com/composer-hub, got %q", ep.CH)
	}
}

func TestEndpointsV2InitsToEndpointsV2_Individual(t *testing.T) {
	e := EndpointsV2Inits{
		IAM:  "https://iam.example.com",
		Dash: "https://dash.example.com",
		CH:   "https://ch.example.com",
	}
	ep := e.ToEndpointsV2()
	if ep.IAM != "https://iam.example.com" {
		t.Errorf("IAM mismatch: got %q", ep.IAM)
	}
	if ep.Dash != "https://dash.example.com" {
		t.Errorf("Dash mismatch: got %q", ep.Dash)
	}
	if ep.CH != "https://ch.example.com" {
		t.Errorf("CH mismatch: got %q", ep.CH)
	}
}

func TestEndpointsV2InitsToEndpointsV2_Mixed(t *testing.T) {
	e := EndpointsV2Inits{
		Base: "https://api.example.com",
		IAM:  "https://custom-iam.example.com",
	}
	ep := e.ToEndpointsV2()
	if ep.IAM != "https://custom-iam.example.com" {
		t.Errorf("IAM: expected custom override, got %q", ep.IAM)
	}
	if ep.Dash != "https://api.example.com/api" {
		t.Errorf("Dash: expected base expansion, got %q", ep.Dash)
	}
	if ep.CH != "https://api.example.com/composer-hub" {
		t.Errorf("CH: expected base expansion, got %q", ep.CH)
	}
}

func TestParseAndValidateEndpointsV2Inits_Valid(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "endpoints.toml")
	content := `base = "https://api.example.com"`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write: %v", err)
	}

	ep, err := ParseAndValidateEndpointsV2Inits(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ep.Base != "https://api.example.com" {
		t.Errorf("expected base, got %q", ep.Base)
	}
}

func TestParseAndValidateEndpointsV2Inits_Invalid(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "endpoints.toml")
	content := `base = ""`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write: %v", err)
	}

	_, err := ParseAndValidateEndpointsV2Inits(path)
	if err == nil {
		t.Fatal("expected error for invalid endpoints, got nil")
	}
}

func TestParseAndValidateEndpointsV2Inits_FileNotFound(t *testing.T) {
	_, err := ParseAndValidateEndpointsV2Inits("/nonexistent/path.toml")
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}

func TestConfigV2Save_InvalidPath(t *testing.T) {
	cfg := CreateEmptyConfigV2()
	err := cfg.Save("/nonexistent/dir/config.toml")
	if err == nil {
		t.Fatal("expected error for invalid path, got nil")
	}
}

// --- helpers ---

func writeTestConfigV2(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}
	return path
}
