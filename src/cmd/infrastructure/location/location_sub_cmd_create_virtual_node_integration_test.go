package cmd_location

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestLocationSubCmd_CreateVirtualNode_Integration_Success(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID, name, storageType string, config map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				Status:      api.InfraAggregateStatus{Code: string(api.StatusCodeOk), Details: "Node created"},
				StorageType: api.InfraVirtualStorageType(storageType),
				StorageConfiguration: map[string]any{
					"bucket": config["bucket"],
				},
			}, nil
		},
	}

	mockConfig := setupDescribeMockConfig()
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

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

	expected := `
Virtual Nodes
─────────────────────────────────────────────────
test-virtual-node (new-node-id)
├── Status: status_ok - Node created
└── Storage: s3
    └── bucket: test-bucket
`
	if buf.String() != expected {
		t.Errorf("snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_CreateVirtualNode_Integration_InvalidJSON(t *testing.T) {
	mockAPI := &api.MockLocationAPI{}
	mockConfig := setupDescribeMockConfig()
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"create-virtual-node",
		"--name", "test-virtual-node",
		"--cluster-id", "test-cluster-id",
		"--storage-type", "s3",
		"--configuration", "invalid json{",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute, got %v", err)
	}

	expected := `ERR error while parsing json configuration configuration: invalid character 'i' looking for beginning of value
`
	if buf.String() != expected {
		t.Errorf("snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_CreateVirtualNode_Integration_APIError(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(_ configuration_models.EndpointsV2, _, _, _, _, _ string, _ map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			return nil, fmt.Errorf("failed to create virtual node")
		},
	}

	mockConfig := setupDescribeMockConfig()
	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

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
		t.Fatalf("Expected no error from Execute, got %v", err)
	}

	expected := `ERR failed to create virtual node: failed to create virtual node
`
	if buf.String() != expected {
		t.Errorf("snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}
