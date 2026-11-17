package cmd_location

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestLocationSubCmd_Describe_UsesSingleLocationEndpoint(t *testing.T) {
	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	mockAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
			t.Error("Expected the aggregated list to not be used when --cluster-id is provided")
			return nil, nil
		},
		GetAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _, clusterID string) (*api.InfraAggregateCluster, error) {
			return &api.InfraAggregateCluster{
				ClusterID: clusterID,
				Name:      "single-cluster",
				Type:      api.ClusterTypePhysical,
				Details: api.InfraAggregateClusterDetail{
					LastUpdate:   baseTime,
					NextUpdate:   baseTime.Add(time.Hour),
					IsUpdateOk:   true,
					Nodes:        []api.InfraAggregateNodeDetail{},
					VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
				},
			}, nil
		},
	}

	locationService := service.NewLocationService(setupDescribeMockConfig(), mockAPI, nil, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440001",
	})

	if err := locationCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.Contains(commandOutput.String(), "single-cluster") {
		t.Fatalf("Expected the cluster from the single endpoint, got %q", commandOutput.String())
	}
}
