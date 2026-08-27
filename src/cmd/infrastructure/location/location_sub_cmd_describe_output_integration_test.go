package cmd_location

import (
	"bytes"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestLocationSubCmd_Describe_Output_JSON(t *testing.T) {
	baseTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
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

	mockConfig := setupDescribeMockConfig()
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
		"--output", "json",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "cluster_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "test-cluster",
  "details": {
    "last_update": "2024-01-15T10:30:00Z",
    "next_update": "2024-01-15T11:30:00Z",
    "is_update_ok": true,
    "nodes": [],
    "virtual_nodes": []
  },
  "type": "physical"
}
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_Describe_Output_YAML(t *testing.T) {
	baseTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
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

	mockConfig := setupDescribeMockConfig()
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
		"--output", "yaml",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `cluster_id: 550e8400-e29b-41d4-a716-446655440000
name: test-cluster
details:
    last_update: 2024-01-15T10:30:00Z
    next_update: 2024-01-15T11:30:00Z
    is_update_ok: true
    nodes: []
    virtual_nodes: []
type: physical
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_Describe_Output_Full_JSON(t *testing.T) {
	baseTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	osImage := "Ubuntu 22.04 LTS"
	externalIP := "203.0.113.10"
	internalIP := "10.0.0.10"

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
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

	mockConfig := setupDescribeMockConfig()
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
		"--output", "json",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "cluster_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "test-cluster",
  "details": {
    "last_update": "2024-01-15T10:30:00Z",
    "next_update": "2024-01-15T11:30:00Z",
    "is_update_ok": true,
    "nodes": [
      {
        "node_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
        "node_name": "node-1-host",
        "status": {
          "code": "status_ok",
          "details": "All systems operational"
        },
        "os_name": "Ubuntu 22.04 LTS",
        "cpu": {
          "cores": 16
        },
        "ram": {
          "available": 64
        },
        "external_ip": "203.0.113.10",
        "internal_ip": "10.0.0.10",
        "disks": [
          {
            "disk_uuid": "8e7d7a2e-3c5e-4f9a-9b8c-1d2e3f4a5b6c",
            "path": "/dev/sda",
            "used": true,
            "pv_ref": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
            "total_storage_size_bytes": 1099511627776,
            "used_storage_bytes": 549755813888,
            "status": {
              "code": "status_ok"
            }
          }
        ]
      }
    ],
    "virtual_nodes": [
      {
        "node_id": "4b6e8c2f-9a1d-4f3e-8b5c-7d4a6e5b3c2f",
        "node_name": "vnode-1-host",
        "status": {
          "code": "status_ok",
          "details": "S3 service running"
        },
        "storage_type": "s3",
        "storage_configuration": {
          "access_key_id": "AKIAIOSFODNN7EXAMPLE",
          "bucket_name": "cubbit-virtual",
          "endpoint": "https://s3.example.com",
          "region": "us-east-1"
        }
      }
    ]
  },
  "type": "physical"
}
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_Describe_Output_JSON_FromProfile(t *testing.T) {
	baseTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
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

	mockConfig := setupDescribeMockConfig()

		mockConfig.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
			}, nil
	}
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "cluster_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "test-cluster",
  "details": {
    "last_update": "2024-01-15T10:30:00Z",
    "next_update": "2024-01-15T11:30:00Z",
    "is_update_ok": true,
    "nodes": [],
    "virtual_nodes": []
  },
  "type": "physical"
}
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_Describe_Output_YAML_FromProfile(t *testing.T) {
	baseTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
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

	mockConfig := setupDescribeMockConfig()

		mockConfig.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
			}, nil
	}
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name", "test-cluster",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `cluster_id: 550e8400-e29b-41d4-a716-446655440000
name: test-cluster
details:
    last_update: 2024-01-15T10:30:00Z
    next_update: 2024-01-15T11:30:00Z
    is_update_ok: true
    nodes: []
    virtual_nodes: []
type: physical
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}
