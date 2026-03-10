package service

import (
	"fmt"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/spf13/cobra"
)

type LocationServiceMock struct {
	ListFunc           func(cmd *cobra.Command, args []string) error
	ListAggregatedFunc func(cmd *cobra.Command, args []string) error
}

func NewLocationServiceMock() *LocationServiceMock {
	return &LocationServiceMock{
		ListFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Locations list: [location1, location2, location3]")
			return nil
		},
		ListAggregatedFunc: func(cmd *cobra.Command, args []string) error {
			mockClusters := []api.InfraAggregateCluster{
				{
					ClusterID: "cluster-001",
					Name:      "mock-cluster-1",
					Type:      "physical",
					Details: api.InfraAggregateClusterDetail{
						LastUpdate:   time.Time{},
						NextUpdate:   time.Time{},
						IsUpdateOk:   true,
						Nodes:        []api.InfraAggregateNodeDetail{},
						VirtualNodes: []api.InfraAggregateVirtualNodeDetail{},
					},
				},
			}
			views := ConvertAggregateClustersView(mockClusters)
			for _, v := range views {
				cmd.Println(fmt.Sprintf("%s\t%s\t%s\t%d\t%d\t%s",
					v.ClusterID, v.Name, v.ClusterType, v.PhysicalNodes, v.VirtualNodes, v.LastUpdate))
			}
			return nil
		},
	}
}

func (m *LocationServiceMock) List(cmd *cobra.Command, args []string) error {
	if m.ListFunc != nil {
		return m.ListFunc(cmd, args)
	}
	return nil
}

func (m *LocationServiceMock) ListAggregated(cmd *cobra.Command, args []string) error {
	if m.ListAggregatedFunc != nil {
		return m.ListAggregatedFunc(cmd, args)
	}
	return nil
}
