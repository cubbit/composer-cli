package cmd_location

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service"
)

func TestLocationSubCmd_Describe_Integration_WithClusterName_Found(t *testing.T) {
	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "test-cluster",
					Type:      api.ClusterTypePhysical,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
			}, nil
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
	})
	err = locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
Cluster Information
─────────────────────────────────────────────────
╭──────────────────────────────────────┬──────────────┬──────────┬────────────────┬───────────────┬─────────────────────╮
│ Cluster ID                           │ Name         │ Type     │ Physical Nodes │ Virtual Nodes │ Last Update         │
├──────────────────────────────────────┼──────────────┼──────────┼────────────────┼───────────────┼─────────────────────┤
│ 550e8400-e29b-41d4-a716-446655440000 │ test-cluster │ physical │ 0              │ 0             │ 2024-01-15 10:30:00 │
╰──────────────────────────────────────┴──────────────┴──────────┴────────────────┴───────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected cluster details output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestLocationSubCmd_Describe_Integration_WithClusterID_Found(t *testing.T) {
	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440001",
					Name:      "another-cluster",
					Type:      api.ClusterTypeVirtual,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
			}, nil
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440001",
	})
	err = locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
Cluster Information
─────────────────────────────────────────────────
╭──────────────────────────────────────┬─────────────────┬─────────┬────────────────┬───────────────┬─────────────────────╮
│ Cluster ID                           │ Name            │ Type    │ Physical Nodes │ Virtual Nodes │ Last Update         │
├──────────────────────────────────────┼─────────────────┼─────────┼────────────────┼───────────────┼─────────────────────┤
│ 550e8400-e29b-41d4-a716-446655440001 │ another-cluster │ virtual │ 0              │ 0             │ 2024-01-15 10:30:00 │
╰──────────────────────────────────────┴─────────────────┴─────────┴────────────────┴───────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected cluster details output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestLocationSubCmd_Describe_Integration_WithClusterName_FullOutput(t *testing.T) {
	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	osImage := "Ubuntu 22.04 LTS"
	externalIP := "203.0.113.10"
	internalIP := "10.0.0.10"

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "test-cluster",
					Type:      api.ClusterTypePhysical,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate: baseTime,
						NextUpdate: baseTime.Add(time.Hour),
						IsUpdateOk: true,
						Nodes: []api.InfraAggregateNodeDetail{
							{
								NodeID:   "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
								NodeName: "node-1-host",
								Status: api.InfraAggregateStatus{
									Code:    string(api.StatusCodeOk),
									Details: "All systems operational",
								},
								OSName:     &osImage,
								CPU:        &api.InfraNodeCPUInfo{Cores: 16},
								RAM:        &api.InfraNodeRAMInfo{Available: 64.0},
								ExternalIP: &externalIP,
								InternalIP: &internalIP,
								Disks: []api.InfraAggregateDiskDetail{
									{
										DiskUUID:              "8e7d7a2e-3c5e-4f9a-9b8c-1d2e3f4a5b6c",
										Path:                  "/dev/sda",
										Used:                  true,
										PVRef:                 "f47ac10b-58cc-4372-a567-0e02b2c3d479",
										TotalStorageSizeBytes: 1099511627776,
										UsedStorageBytes:      549755813888,
										Status: api.InfraAggregateStatus{
											Code: string(api.StatusCodeOk),
										},
									},
								},
							},
						},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{
							{
								NodeID:   "4b6e8c2f-9a1d-4f3e-8b5c-7d4a6e5b3c2f",
								NodeName: "vnode-1-host",
								Status: api.InfraAggregateStatus{
									Code:    string(api.StatusCodeOk),
									Details: "S3 service running",
								},
								StorageType: api.VirtualStorageTypeS3,
								StorageConfiguration: map[string]any{
									"endpoint":      "https://s3.example.com",
									"bucket_name":   "cubbit-virtual",
									"region":        "us-east-1",
									"access_key_id": "AKIAIOSFODNN7EXAMPLE",
								},
							},
						},
					},
				},
			}, nil
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
	})
	err = locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
Cluster Information
─────────────────────────────────────────────────
╭──────────────────────────────────────┬──────────────┬──────────┬────────────────┬───────────────┬─────────────────────╮
│ Cluster ID                           │ Name         │ Type     │ Physical Nodes │ Virtual Nodes │ Last Update         │
├──────────────────────────────────────┼──────────────┼──────────┼────────────────┼───────────────┼─────────────────────┤
│ 550e8400-e29b-41d4-a716-446655440000 │ test-cluster │ physical │ 1              │ 1             │ 2024-01-15 10:30:00 │
╰──────────────────────────────────────┴──────────────┴──────────┴────────────────┴───────────────┴─────────────────────╯

Physical Nodes
─────────────────────────────────────────────────
node-1-host (6ba7b810-9dad-11d1-80b4-00c04fd430c8)
├── Status: status_ok - All systems operational
├── Hardware
│   ├── OS: Ubuntu 22.04 LTS
│   ├── CPU Cores: 16
│   └── RAM: 64.00 GB
├── Network
│   ├── External IP: 203.0.113.10
│   └── Internal IP: 10.0.0.10
└── Disks
    └── Disk: 8e7d7a2e-3c5e-4f9a-9b8c-1d2e3f4a5b6c
        ├── Path: /dev/sda
        ├── Status: status_ok
        ├── Total: 1.00 TB
        ├── Used: 512.00 GB
        └── PV Reference: f47ac10b-58cc-4372-a567-0e02b2c3d479

Virtual Nodes
─────────────────────────────────────────────────
vnode-1-host (4b6e8c2f-9a1d-4f3e-8b5c-7d4a6e5b3c2f)
├── Status: status_ok - S3 service running
└── Storage: s3
    ├── access_key_id: AKIAIOSFODNN7EXAMPLE
    ├── bucket_name: cubbit-virtual
    ├── endpoint: https://s3.example.com
    └── region: us-east-1
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected cluster details output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestLocationSubCmd_Describe_Integration_FilterByClusterName_WithMultipleClusters(t *testing.T) {
	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "production-cluster",
					Type:      api.ClusterTypePhysical,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440001",
					Name:      "staging-cluster",
					Type:      api.ClusterTypePhysical,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440002",
					Name:      "development-cluster",
					Type:      api.ClusterTypeVirtual,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
			}, nil
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "staging-cluster",
	})
	err = locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
Cluster Information
─────────────────────────────────────────────────
╭──────────────────────────────────────┬─────────────────┬──────────┬────────────────┬───────────────┬─────────────────────╮
│ Cluster ID                           │ Name            │ Type     │ Physical Nodes │ Virtual Nodes │ Last Update         │
├──────────────────────────────────────┼─────────────────┼──────────┼────────────────┼───────────────┼─────────────────────┤
│ 550e8400-e29b-41d4-a716-446655440001 │ staging-cluster │ physical │ 0              │ 0             │ 2024-01-15 10:30:00 │
╰──────────────────────────────────────┴─────────────────┴──────────┴────────────────┴───────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected cluster details output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestLocationSubCmd_Describe_Integration_FilterByClusterID_WithMultipleClusters(t *testing.T) {
	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "production-cluster",
					Type:      api.ClusterTypePhysical,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440001",
					Name:      "staging-cluster",
					Type:      api.ClusterTypePhysical,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440002",
					Name:      "development-cluster",
					Type:      api.ClusterTypeVirtual,
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   baseTime,
						NextUpdate:   baseTime.Add(time.Hour),
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
			}, nil
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440002",
	})
	err = locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
Cluster Information
─────────────────────────────────────────────────
╭──────────────────────────────────────┬─────────────────────┬─────────┬────────────────┬───────────────┬─────────────────────╮
│ Cluster ID                           │ Name                │ Type    │ Physical Nodes │ Virtual Nodes │ Last Update         │
├──────────────────────────────────────┼─────────────────────┼─────────┼────────────────┼───────────────┼─────────────────────┤
│ 550e8400-e29b-41d4-a716-446655440002 │ development-cluster │ virtual │ 0              │ 0             │ 2024-01-15 10:30:00 │
╰──────────────────────────────────────┴─────────────────────┴─────────┴────────────────┴───────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected cluster details output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestLocationSubCmd_Describe_Integration_WithClusterName_NotFound(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{}, nil
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "non-existent-cluster",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	expectedError := "cluster with name 'non-existent-cluster' not found"
	if !strings.Contains(output, expectedError) {
		t.Errorf("Expected output to contain error message %q, got %q", expectedError, output)
	}
}

func TestLocationSubCmd_Describe_Integration_WithClusterID_NotFound(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{}, nil
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-id", "00000000-0000-0000-0000-000000000000",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	expectedError := "cluster with ID '00000000-0000-0000-0000-000000000000' not found"
	if !strings.Contains(output, expectedError) {
		t.Errorf("Expected output to contain error message %q, got %q", expectedError, output)
	}
}

func TestLocationSubCmd_Describe_API_ReturnsError(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfraAggregateCluster, error) {
			return nil, fmt.Errorf("failed to connect to API")
		},
	}

	mockConfig := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	expectedError := "failed to list aggregated locations"
	if !strings.Contains(output, expectedError) {
		t.Errorf("Expected output to contain error message %q, got %q", expectedError, output)
	}
}
