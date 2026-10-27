package cmd_location

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestLocationSubCmd_CreateVirtualNode_Output_JSON(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID, name, storageType string, config map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				Status:      api.InfraAggregateStatus{Code: string(api.StatusCodeOk), Details: "Node created"},
				StorageType: api.VirtualStorageTypeS3,
				StorageConfiguration: map[string]any{
					"bucket": "test-bucket",
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
		"create-virtual-node",
		"--name", "test-virtual-node",
		"--cluster-id", "test-cluster-id",
		"--storage-type", "s3",
		"--configuration", `{"bucket": "test-bucket"}`,
		"--output", "json",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "node_id": "new-node-id",
    "node_name": "test-virtual-node",
    "status": {
      "code": "status_ok",
      "details": "Node created"
    },
    "storage_type": "s3",
    "storage_configuration": {
      "bucket": "test-bucket"
    }
  }
]
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_CreateVirtualNode_Output_YAML(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID, name, storageType string, config map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				Status:      api.InfraAggregateStatus{Code: string(api.StatusCodeOk), Details: "Node created"},
				StorageType: api.VirtualStorageTypeS3,
				StorageConfiguration: map[string]any{
					"bucket": "test-bucket",
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
		"create-virtual-node",
		"--name", "test-virtual-node",
		"--cluster-id", "test-cluster-id",
		"--storage-type", "s3",
		"--configuration", `{"bucket": "test-bucket"}`,
		"--output", "yaml",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- node_id: new-node-id
  node_name: test-virtual-node
  status:
    code: status_ok
    details: Node created
  storage_type: s3
  storage_configuration:
    bucket: test-bucket
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_CreateVirtualNode_Output_JSON_FromProfile(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID, name, storageType string, config map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				Status:      api.InfraAggregateStatus{Code: string(api.StatusCodeOk), Details: "Node created"},
				StorageType: api.VirtualStorageTypeS3,
				StorageConfiguration: map[string]any{
					"bucket": "test-bucket",
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
		"create-virtual-node",
		"--name", "test-virtual-node",
		"--cluster-id", "test-cluster-id",
		"--storage-type", "s3",
		"--configuration", `{"bucket": "test-bucket"}`,
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "node_id": "new-node-id",
    "node_name": "test-virtual-node",
    "status": {
      "code": "status_ok",
      "details": "Node created"
    },
    "storage_type": "s3",
    "storage_configuration": {
      "bucket": "test-bucket"
    }
  }
]
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_CreateVirtualNode_Output_YAML_FromProfile(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID, name, storageType string, config map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				Status:      api.InfraAggregateStatus{Code: string(api.StatusCodeOk), Details: "Node created"},
				StorageType: api.VirtualStorageTypeS3,
				StorageConfiguration: map[string]any{
					"bucket": "test-bucket",
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
		"create-virtual-node",
		"--name", "test-virtual-node",
		"--cluster-id", "test-cluster-id",
		"--storage-type", "s3",
		"--configuration", `{"bucket": "test-bucket"}`,
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- node_id: new-node-id
  node_name: test-virtual-node
  status:
    code: status_ok
    details: Node created
  storage_type: s3
  storage_configuration:
    bucket: test-bucket
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}
