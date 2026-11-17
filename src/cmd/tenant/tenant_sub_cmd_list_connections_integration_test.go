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

func TestTenantSubCmd_ListConnections_Integration_Success(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockConnectionAPI := &api.MockConnectionAPI{
		ListConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			opts ...api.ListConnectionsV5Option,
		) (*api.GenericPaginatedResponse[api.ConnectionV5DTO], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if tenantID != "tnt-001" {
				t.Fatalf("Expected tenant ID tnt-001, got %q", tenantID)
			}

			return &api.GenericPaginatedResponse[api.ConnectionV5DTO]{
				Data: []api.ConnectionV5DTO{
					{
						ID: "conn-001",
						Domain: api.DomainDTO{
							DomainName: "example.com",
						},
						Subdomain: strPtr("sub1"),
						Gateways: []api.GatewayConnectionV5DTO{
							{
								Name: "gw-1",
								VerificationState: &api.VerificationStateV5DTO{
									Console: api.RecordInfoV5DTO{FQDN: "console.gw-1.example.com", Status: "dns_ok"},
									S3:      api.RecordInfoV5DTO{FQDN: "s3.gw-1.example.com", Status: "dns_ok"},
								},
							},
						},
					},
					{
						ID: "conn-002",
						Domain: api.DomainDTO{
							DomainName: "test.org",
						},
						Gateways: []api.GatewayConnectionV5DTO{
							{
								Name: "gw-2",
								VerificationState: &api.VerificationStateV5DTO{
									Console:    api.RecordInfoV5DTO{FQDN: "console.gw-2.test.org", Status: "cert_ok"},
									S3:         api.RecordInfoV5DTO{FQDN: "s3.gw-2.test.org", Status: "dns_ok"},
									WildcardS3: api.RecordInfoV5DTO{FQDN: "*.s3.gw-2.test.org", Status: "unset"},
								},
							},
						},
					},
				},
				NextPage: nil,
				Count:    2,
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
		"list-connections",
		"--tenant-id", "tnt-001",
		"--skip-verify",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := strings.TrimSpace(`
sub1.example.com
├── ID: conn-001
└── Gateways
    └── gw-1
        └── Gateway verification
            ├── Console: dns_ok
            ├── S3: dns_ok
            └── Wildcard S3: 
test.org
├── ID: conn-002
└── Gateways
    └── gw-2
        └── Gateway verification
            ├── Console: cert_ok
            ├── S3: dns_ok
            └── Wildcard S3: unset
`)

	actualOutput := strings.TrimSpace(commandOutput.String())
	if actualOutput != expectedOutput {
		t.Errorf("Expected list-connections output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}

func TestTenantSubCmd_ListConnections_Integration_WithDomainValidation(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockConnectionAPI := &api.MockConnectionAPI{
		ListConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			opts ...api.ListConnectionsV5Option,
		) (*api.GenericPaginatedResponse[api.ConnectionV5DTO], error) {
			return &api.GenericPaginatedResponse[api.ConnectionV5DTO]{
				Data: []api.ConnectionV5DTO{
					{
						ID: "conn-003",
						Domain: api.DomainDTO{
							DomainName: "example.com",
						},
						Subdomain: strPtr("www"),
						Gateways: []api.GatewayConnectionV5DTO{
							{
								Name: "gw-main",
								GatewayID: "gw-main-id",
								VerificationState: &api.VerificationStateV5DTO{
									Console:    api.RecordInfoV5DTO{FQDN: "console.gw-main.example.com", Status: "dns_ok"},
									S3:         api.RecordInfoV5DTO{FQDN: "s3.gw-main.example.com", Status: "dns_ok"},
									WildcardS3: api.RecordInfoV5DTO{FQDN: "*.s3.gw-main.example.com", Status: "unset"},
								},
							},
						},
						DomainVerificationState: &[]api.DomainVerificationStateV5DTO{
							{
								GatewayID: "gw-main-id",
								VerificationState: &api.VerificationStateV5DTO{
									Console:    api.RecordInfoV5DTO{FQDN: "console.gw-main.example.com", Status: "cert_ok"},
									S3:         api.RecordInfoV5DTO{FQDN: "s3.gw-main.example.com", Status: "dns_ok"},
									WildcardS3: api.RecordInfoV5DTO{FQDN: "*.s3.gw-main.example.com", Status: "unset"},
								},
							},
						},
					},
				},
				NextPage: nil,
				Count:    1,
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
		"list-connections",
		"--tenant-id", "tnt-001",
		"--skip-verify",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := strings.TrimSpace(`
www.example.com
├── ID: conn-003
└── Gateways
    └── gw-main
        ├── Gateway verification
        │   ├── Console: dns_ok
        │   ├── S3: dns_ok
        │   └── Wildcard S3: unset
        └── Domain verification
            ├── Console: cert_ok
            ├── S3: dns_ok
            └── Wildcard S3: unset
`)

	actualOutput := strings.TrimSpace(commandOutput.String())
	if actualOutput != expectedOutput {
		t.Errorf("Expected list-connections output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}

func TestTenantSubCmd_ListConnections_Integration_Empty(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockConnectionAPI := &api.MockConnectionAPI{
		ListConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			opts ...api.ListConnectionsV5Option,
		) (*api.GenericPaginatedResponse[api.ConnectionV5DTO], error) {
			return &api.GenericPaginatedResponse[api.ConnectionV5DTO]{
				Data:     []api.ConnectionV5DTO{},
				NextPage: nil,
				Count:    0,
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
		"list-connections",
		"--tenant-id", "tnt-001",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "No connections found.\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Errorf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestTenantSubCmd_ListConnections_Integration_Error(t *testing.T) {
	mockCfg := newTestTenantConfig()

	mockConnectionAPI := &api.MockConnectionAPI{
		ListConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			opts ...api.ListConnectionsV5Option,
		) (*api.GenericPaginatedResponse[api.ConnectionV5DTO], error) {
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
		"list-connections",
		"--tenant-id", "tnt-001",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed to list tenant connections") {
		t.Errorf("Expected output to contain error message, got %q", output)
	}
}

func TestTenantSubCmd_ListConnections_Integration_WithVerify(t *testing.T) {
	mockCfg := newTestTenantConfig()

	var listCallCount int
	var verifyCallCount int

	mockConnectionAPI := &api.MockConnectionAPI{
		ListConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			opts ...api.ListConnectionsV5Option,
		) (*api.GenericPaginatedResponse[api.ConnectionV5DTO], error) {
			listCallCount++

			data := []api.ConnectionV5DTO{
				{
					ID: "conn-001",
					Domain: api.DomainDTO{
						DomainName: "example.com",
					},
					Gateways: []api.GatewayConnectionV5DTO{
						{
							Name: "gw-1",
							VerificationState: &api.VerificationStateV5DTO{
								Console: api.RecordInfoV5DTO{FQDN: "console.gw-1.example.com", Status: "dns_ok"},
								S3:      api.RecordInfoV5DTO{FQDN: "s3.gw-1.example.com", Status: "dns_ok"},
							},
						},
					},
				},
			}

			if listCallCount == 2 {
				data[0].Gateways[0].VerificationState.Console.Status = "certificate_ok"
				data[0].Gateways[0].VerificationState.S3.Status = "certificate_ok"
			}

			return &api.GenericPaginatedResponse[api.ConnectionV5DTO]{
				Data:     data,
				NextPage: nil,
				Count:    1,
			}, nil
		},
		VerifyConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			tenantID string,
			connectionID string,
		) (*api.ConnectionV5DTO, error) {
			verifyCallCount++

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
				ID: connectionID,
				Domain: api.DomainDTO{
					DomainName: "example.com",
				},
				Gateways: []api.GatewayConnectionV5DTO{
					{
						Name: "gw-1",
						VerificationState: &api.VerificationStateV5DTO{
							Console: api.RecordInfoV5DTO{FQDN: "console.gw-1.example.com", Status: "certificate_ok"},
							S3:      api.RecordInfoV5DTO{FQDN: "s3.gw-1.example.com", Status: "certificate_ok"},
						},
					},
				},
				DomainVerificationState: &[]api.DomainVerificationStateV5DTO{
					{
						GatewayID: "gw-001",
						VerificationState: &api.VerificationStateV5DTO{
							RecordType: "domain",
							Console:    api.RecordInfoV5DTO{FQDN: "console.example.com", Status: "certificate_ok"},
							S3:         api.RecordInfoV5DTO{FQDN: "s3.example.com", Status: "certificate_ok"},
							WildcardS3: api.RecordInfoV5DTO{FQDN: "*.s3.example.com", Status: "unset"},
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
		"list-connections",
		"--tenant-id", "tnt-001",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if listCallCount != 1 {
		t.Errorf("Expected ListConnectionsV5 to be called 1 time (verify results merged in-memory), got %d", listCallCount)
	}

	if verifyCallCount != 1 {
		t.Errorf("Expected VerifyConnectionV5 to be called 1 time, got %d", verifyCallCount)
	}

	expectedOutput := strings.TrimSpace(`
example.com
├── ID: conn-001
└── Gateways
    └── gw-1
        └── Gateway verification
            ├── Console: certificate_ok
            ├── S3: certificate_ok
            └── Wildcard S3: 
`)

	actualOutput := strings.TrimSpace(commandOutput.String())
	if actualOutput != expectedOutput {
		t.Errorf("Expected list-connections output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}
