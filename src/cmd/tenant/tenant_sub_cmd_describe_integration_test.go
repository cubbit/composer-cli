package cmd_tenant

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
)

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func TestTenantSubCmd_Describe_Integration_Success(t *testing.T) {
	mockCfg := newTestTenantConfig()

	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-001" {
				t.Fatalf("Expected tenant ID tnt-001, got %q", tenantID)
			}

			return &api.TenantV5DTO{
				ID:          "tnt-001",
				Name:        "prod-tenant",
				Slug:        "prod-tenant",
				Description: strPtr("Production tenant"),
				CreatedAt:   frozen,
				Storage:     &api.UsageDTO{Consumed: 1073741824, Reserved: 2147483648, Percentage: 0.5},
				Bandwidth:   &api.UsageDTO{Consumed: 524288000, Reserved: 1073741824, Percentage: 0.25},
				Settings: &api.TenantSettings{
					ConsoleUrl:        strPtr("https://console.example.com"),
					GatewayUrl:        strPtr("https://gw.example.com"),
					SignupDisabled:    boolPtr(false),
					WhitelabelEnabled: boolPtr(true),
				},
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"describe",
		"tnt-001",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := strings.TrimSpace(`
prod-tenant
├── ID: tnt-001
├── Slug: prod-tenant
├── Description: Production tenant
├── Created at: 2026-06-24 20:24:13 UTC
├── ZK Enabled: no
├── Resource Usage
│   ├── Storage
│   │   ├── Consumed: 1.0 GB
│   │   ├── Reserved: 2.0 GB
│   │   └── Used: 50.0%
│   └── Bandwidth
│       ├── Consumed: 500.0 MB
│       ├── Reserved: 1.0 GB
│       └── Used: 25.0%
└── Settings
    ├── Console URL: https://console.example.com
    ├── Gateway URL: https://gw.example.com
    ├── Signup: enabled
    ├── Display name: none
    ├── Support link: none
    ├── Allowed domains: none
    ├── Blocked domains: none
    ├── Whitelabel: yes
    ├── Project
    │   ├── Max storage: none
    │   └── Max egress bandwidth: none
    ├── Account
    │   ├── Max projects: none
    │   ├── Auth providers: none
    │   ├── Auth Providers Config
    │   │   ├── Google client ID: none
    │   │   ├── Microsoft application ID: none
    │   │   └── Microsoft directory ID: none
    │   └── OAuth Providers: none
    ├── Notifications
    │   └── none
    └── White Label
        ├── DNS
        │   ├── Value: none
        │   ├── Challenge: none
        │   └── Verified: no
        └── Email Domain
            ├── Value: none
            └── Verified: no
`)

	actualOutput := strings.TrimSpace(commandOutput.String())
	if actualOutput != expectedOutput {
		t.Errorf("Expected output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}

func TestTenantSubCmd_Describe_Integration_Minimal(t *testing.T) {
	mockCfg := newTestTenantConfig()

	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-min" {
				t.Fatalf("Expected tenant ID tnt-min, got %q", tenantID)
			}

			return &api.TenantV5DTO{
				ID:        "tnt-min",
				Name:      "minimal-tenant",
				Slug:      "minimal-tenant",
				CreatedAt: frozen,
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"describe",
		"tnt-min",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := strings.TrimSpace(`
minimal-tenant
├── ID: tnt-min
├── Slug: minimal-tenant
├── Description: none
├── Created at: 2026-06-24 20:24:13 UTC
├── ZK Enabled: no
├── Resource Usage
│   ├── Storage: N/A
│   └── Bandwidth: N/A
└── Settings
    ├── Console URL: none
    ├── Gateway URL: none
    ├── Signup: enabled
    ├── Display name: none
    ├── Support link: none
    ├── Allowed domains: none
    ├── Blocked domains: none
    ├── Whitelabel: no
    ├── Project
    │   ├── Max storage: none
    │   └── Max egress bandwidth: none
    ├── Account
    │   ├── Max projects: none
    │   ├── Auth providers: none
    │   ├── Auth Providers Config
    │   │   ├── Google client ID: none
    │   │   ├── Microsoft application ID: none
    │   │   └── Microsoft directory ID: none
    │   └── OAuth Providers: none
    ├── Notifications
    │   └── none
    └── White Label
        ├── DNS
        │   ├── Value: none
        │   ├── Challenge: none
        │   └── Verified: no
        └── Email Domain
            ├── Value: none
            └── Verified: no
`)

	actualOutput := strings.TrimSpace(commandOutput.String())
	if actualOutput != expectedOutput {
		t.Errorf("Expected output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}

func TestTenantSubCmd_Describe_Integration_ShowSecrets(t *testing.T) {
	mockCfg := newTestTenantConfig()

	frozen := time.Date(2026, 6, 24, 20, 24, 13, 0, time.UTC)
	googleSecret := "GOCSPX-google-client-secret-xyz"
	microsoftSecret := "microsoft~client~secret~abc"
	customSecret := "custom-oidc-secret-789"
	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			return &api.TenantV5DTO{
				ID:          "tnt-sec",
				Name:        "secrets-tenant",
				Slug:        "secrets-tenant",
				Description: strPtr("Tenant with secrets"),
				CreatedAt:   frozen,
				Settings: &api.TenantSettings{
					Account: &api.TenantSettingsAccount{
						OAuthProviders: &[]api.OAuthProvider{
							{
								Name:         "google",
								ClientID:     "google-client-123.apps.googleusercontent.com",
								ClientSecret: googleSecret,
								IssuerURL:    "https://accounts.google.com",
								RedirectURL:  "https://console.example.com/auth/callback",
							},
							{
								Name:         "microsoft",
								ClientID:     "microsoft-app-456",
								ClientSecret: microsoftSecret,
								IssuerURL:    "https://login.microsoftonline.com/tenant/v2.0",
								RedirectURL:  "https://console.example.com/auth/callback",
							},
							{
								Name:         "custom-oidc",
								ClientID:     "custom-client-789",
								ClientSecret: customSecret,
								IssuerURL:    "https://auth.example.com",
								RedirectURL:  "https://console.example.com/auth/callback",
							},
						},
					},
				},
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
	)

	t.Run("default hides secrets", func(t *testing.T) {
		tenantCmd := NewTenantCmd(tenantService)
		tenantCmd.PersistentFlags().String("profile", "", "Profile")
		tenantCmd.PersistentFlags().String("output", "human", "Output format")
		tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
		tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

		commandOutput := new(bytes.Buffer)
		tenantCmd.SetOut(commandOutput)
		tenantCmd.SetErr(commandOutput)
		tenantCmd.SetArgs([]string{"describe", "tnt-sec"})

		err := tenantCmd.Execute()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expectedOutput := strings.TrimSpace(`
secrets-tenant
├── ID: tnt-sec
├── Slug: secrets-tenant
├── Description: Tenant with secrets
├── Created at: 2026-06-24 20:24:13 UTC
├── ZK Enabled: no
├── Resource Usage
│   ├── Storage: N/A
│   └── Bandwidth: N/A
└── Settings
    ├── Console URL: none
    ├── Gateway URL: none
    ├── Signup: enabled
    ├── Display name: none
    ├── Support link: none
    ├── Allowed domains: none
    ├── Blocked domains: none
    ├── Whitelabel: no
    ├── Project
    │   ├── Max storage: none
    │   └── Max egress bandwidth: none
    ├── Account
    │   ├── Max projects: none
    │   ├── Auth providers: none
    │   ├── Auth Providers Config
    │   │   ├── Google client ID: none
    │   │   ├── Microsoft application ID: none
    │   │   └── Microsoft directory ID: none
    │   └── OAuth Providers
    │       ├── google
    │       │   ├── Client ID: google-client-123.apps.googleusercontent.com
    │       │   ├── Client Secret: ************
    │       │   ├── Issuer URL: https://accounts.google.com
    │       │   └── Redirect URL: https://console.example.com/auth/callback
    │       ├── microsoft
    │       │   ├── Client ID: microsoft-app-456
    │       │   ├── Client Secret: ************
    │       │   ├── Issuer URL: https://login.microsoftonline.com/tenant/v2.0
    │       │   └── Redirect URL: https://console.example.com/auth/callback
    │       └── custom-oidc
    │           ├── Client ID: custom-client-789
    │           ├── Client Secret: ************
    │           ├── Issuer URL: https://auth.example.com
    │           └── Redirect URL: https://console.example.com/auth/callback
    ├── Notifications
    │   └── none
    └── White Label
        ├── DNS
        │   ├── Value: none
        │   ├── Challenge: none
        │   └── Verified: no
        └── Email Domain
            ├── Value: none
            └── Verified: no
`)

		actualOutput := strings.TrimSpace(commandOutput.String())
		if actualOutput != expectedOutput {
			t.Errorf("Expected output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
		}
	})

	t.Run("show-secrets reveals values", func(t *testing.T) {
		tenantCmd := NewTenantCmd(tenantService)
		tenantCmd.PersistentFlags().String("profile", "", "Profile")
		tenantCmd.PersistentFlags().String("output", "human", "Output format")
		tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
		tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

		commandOutput := new(bytes.Buffer)
		tenantCmd.SetOut(commandOutput)
		tenantCmd.SetErr(commandOutput)
		tenantCmd.SetArgs([]string{"describe", "tnt-sec", "--show-secrets"})

		err := tenantCmd.Execute()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expectedOutput := strings.TrimSpace(`
secrets-tenant
├── ID: tnt-sec
├── Slug: secrets-tenant
├── Description: Tenant with secrets
├── Created at: 2026-06-24 20:24:13 UTC
├── ZK Enabled: no
├── Resource Usage
│   ├── Storage: N/A
│   └── Bandwidth: N/A
└── Settings
    ├── Console URL: none
    ├── Gateway URL: none
    ├── Signup: enabled
    ├── Display name: none
    ├── Support link: none
    ├── Allowed domains: none
    ├── Blocked domains: none
    ├── Whitelabel: no
    ├── Project
    │   ├── Max storage: none
    │   └── Max egress bandwidth: none
    ├── Account
    │   ├── Max projects: none
    │   ├── Auth providers: none
    │   ├── Auth Providers Config
    │   │   ├── Google client ID: none
    │   │   ├── Microsoft application ID: none
    │   │   └── Microsoft directory ID: none
    │   └── OAuth Providers
    │       ├── google
    │       │   ├── Client ID: google-client-123.apps.googleusercontent.com
    │       │   ├── Client Secret: GOCSPX-google-client-secret-xyz
    │       │   ├── Issuer URL: https://accounts.google.com
    │       │   └── Redirect URL: https://console.example.com/auth/callback
    │       ├── microsoft
    │       │   ├── Client ID: microsoft-app-456
    │       │   ├── Client Secret: microsoft~client~secret~abc
    │       │   ├── Issuer URL: https://login.microsoftonline.com/tenant/v2.0
    │       │   └── Redirect URL: https://console.example.com/auth/callback
    │       └── custom-oidc
    │           ├── Client ID: custom-client-789
    │           ├── Client Secret: custom-oidc-secret-789
    │           ├── Issuer URL: https://auth.example.com
    │           └── Redirect URL: https://console.example.com/auth/callback
    ├── Notifications
    │   └── none
    └── White Label
        ├── DNS
        │   ├── Value: none
        │   ├── Challenge: none
        │   └── Verified: no
        └── Email Domain
            ├── Value: none
            └── Verified: no
`)

		actualOutput := strings.TrimSpace(commandOutput.String())
		if actualOutput != expectedOutput {
			t.Errorf("Expected output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
		}
	})
}

func TestTenantSubCmd_Describe_Integration_Error(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{
		GetTenantV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
		) (*api.TenantV5DTO, error) {
			return nil, errors.New("network error")
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	stderrBuffer := new(bytes.Buffer)
	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(stderrBuffer)
	tenantCmd.SetArgs([]string{
		"describe",
		"tnt-err",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	expectedStderr := "ERR failed to describe tenant: network error\n"
	actualStderr := stderrBuffer.String()
	if actualStderr != expectedStderr {
		t.Errorf("Expected stderr %q, got %q", expectedStderr, actualStderr)
	}
}

func TestTenantSubCmd_Describe_Integration_NoArgs(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockTenantAPI := &api.MockTenantAPI{}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		mockTenantAPI,
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
	)

	tenantCmd := NewTenantCmd(tenantService)
	tenantCmd.PersistentFlags().String("profile", "", "Profile")
	tenantCmd.PersistentFlags().String("output", "human", "Output format")
	tenantCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	tenantCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	stderrBuffer := new(bytes.Buffer)
	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(stderrBuffer)
	tenantCmd.SetArgs([]string{
		"describe",
	})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatal("Expected error from Execute, got nil")
	}

	expectedError := "accepts 1 arg(s), received 0"
	if err.Error() != expectedError {
		t.Errorf("Expected error %q, got %q", expectedError, err.Error())
	}
}
