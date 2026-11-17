package cmd_tenant

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
)

func TestTenantSubCmd_VerifyConnection_Integration_Success(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockConnectionAPI := &api.MockConnectionAPI{
		VerifyConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			connectionID string,
		) (*api.ConnectionV5DTO, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-001" {
				t.Fatalf("Expected tenant ID tnt-001, got %q", tenantID)
			}
			if connectionID != "conn-001" {
				t.Fatalf("Expected connection ID conn-001, got %q", connectionID)
			}

			return &api.ConnectionV5DTO{
				ID: "conn-001",
				Domain: api.DomainDTO{
					DomainName: "example.com",
				},
				Subdomain: strPtr("sub1"),
				Gateways: []api.GatewayConnectionV5DTO{
					{
						Name:      "gw-1",
						GatewayID: "gw-001",
						VerificationState: &api.VerificationStateV5DTO{
							RecordType: "gateway",
							Console: api.RecordInfoV5DTO{
								FQDN:   "console.gw-1.sub1.example.com",
								Status: "certificate_ok",
								Certificates: &api.CertificatesV5DTO{
									Hostname:          "console.gw-1.sub1.example.com",
									Issuer:            "Let's Encrypt Authority X3",
									Subject:           "CN=*.cubbit.io",
									ValidFrom:         "2024-01-01",
									ValidTo:           "2024-03-31",
									DaysRemaining:     30,
									FingerprintSHA256: "3a:5b:6c:7d:8e:9f:0a:1b:2c:3d:4e:5f:6a:7b:8c:9d:0e:1f:2a:3b:4c:5d:6e:7f:8a:9b:0c:1d:2e:3f:4a:5b",
									San:               []string{"*.cubbit.io", "cubbit.io"},
								},
							},
							S3: api.RecordInfoV5DTO{
								FQDN:   "s3.gw-1.sub1.example.com",
								Status: "dns_ok",
							},
							WildcardS3: api.RecordInfoV5DTO{
								FQDN:   "*.s3.gw-1.sub1.example.com",
								Status: "unset",
							},
						},
					},
				},
				DomainVerificationState: &[]api.DomainVerificationStateV5DTO{
					{
						GatewayID: "gw-001",
						VerificationState: &api.VerificationStateV5DTO{
							RecordType: "domain",
							Console: api.RecordInfoV5DTO{
								FQDN:   "console.sub1.example.com",
								Status: "certificate_ok",
							},
							S3: api.RecordInfoV5DTO{
								FQDN:   "s3.sub1.example.com",
								Status: "dns_ok",
							},
							WildcardS3: api.RecordInfoV5DTO{
								FQDN:   "*.s3.sub1.example.com",
								Status: "unset",
							},
						},
					},
				},
			}, nil
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		&api.MockTenantAPI{},
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
		mockConnectionAPI,
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
		"verify-connection",
		"--tenant-id", "tnt-001",
		"--connection-id", "conn-001",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := strings.TrimSpace(`
sub1.example.com
├── Domain Records
│   ├── Console (console.sub1.example.com)
│   │   └── Status: certificate_ok
│   ├── S3 (s3.sub1.example.com)
│   │   └── Status: dns_ok
│   └── Wildcard S3 (*.s3.sub1.example.com)
│       └── Status: unset
└── Gateway: gw-1 (gw-001)
    ├── Console (console.gw-1.sub1.example.com)
    │   ├── Status: certificate_ok
    │   └── Certificate
    │       ├── Issuer: Let's Encrypt Authority X3
    │       ├── Subject: CN=*.cubbit.io
    │       ├── Valid from: 2024-01-01
    │       ├── Valid to: 2024-03-31 (30 days remaining)
    │       ├── Fingerprint SHA256: 3a:5b:6c:7d:8e:9f:0a:1b:2c:3d:4e:5f:6a:7b:8c:9d:0e:1f:2a:3b:4c:5d:6e:7f:8a:9b:0c:1d:2e:3f:4a:5b
    │       └── SAN: *.cubbit.io, cubbit.io
    ├── S3 (s3.gw-1.sub1.example.com)
    │   └── Status: dns_ok
    └── Wildcard S3 (*.s3.gw-1.sub1.example.com)
        └── Status: unset
`)

	actualOutput := strings.TrimSpace(commandOutput.String())
	if actualOutput != expectedOutput {
		t.Errorf("Expected verify-connection output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}

func TestTenantSubCmd_VerifyConnection_Integration_Error(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockConnectionAPI := &api.MockConnectionAPI{
		VerifyConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			connectionID string,
		) (*api.ConnectionV5DTO, error) {
			return nil, errors.New("network error")
		},
	}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		&api.MockTenantAPI{},
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
		mockConnectionAPI,
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
		"verify-connection",
		"--tenant-id", "tnt-001",
		"--connection-id", "conn-001",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed to verify connection") {
		t.Errorf("Expected output to contain error message, got %q", output)
	}
}

func TestTenantSubCmd_VerifyConnection_Integration_MissingConnectionID(t *testing.T) {
	mockCfg := newTestTenantConfig()
	mockConnectionAPI := &api.MockConnectionAPI{}

	tenantService := servicetenant.NewTenantService(
		mockCfg,
		&api.MockTenantAPI{},
		&api.MockDomainAPI{},
		&api.MockGatewayAPI{},
		&api.MockProcessAPI{},
		mockConnectionAPI,
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
		"verify-connection",
		"--tenant-id", "tnt-001",
	})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatal("Expected error from Execute when --connection-id is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"connection-id\" not set") {
		t.Errorf("Expected error about missing connection-id flag, got %q", err.Error())
	}
}
