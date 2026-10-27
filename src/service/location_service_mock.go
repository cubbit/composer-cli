package service

import (
	"fmt"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/spf13/cobra"
)

type LocationServiceMock struct {
	ListFunc              func(cmd *cobra.Command, args []string) error
	DescribeFunc          func(cmd *cobra.Command, args []string) error
	UpdateFunc            func(cmd *cobra.Command, args []string) error
	DeleteFunc            func(cmd *cobra.Command, args []string) error
	CreateVirtualFunc     func(cmd *cobra.Command, args []string) error
	CreateVirtualNodeFunc func(cmd *cobra.Command, args []string) error
}

func NewLocationServiceMock() *LocationServiceMock {
	return &LocationServiceMock{
		ListFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Locations list: [location1, location2, location3]")
			return nil
		},
		DescribeFunc: func(cmd *cobra.Command, args []string) error {
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
		CreateVirtualFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Virtual location created")
			return nil
		},
		CreateVirtualNodeFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Virtual node created")
			return nil
		},
		UpdateFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Location updated")
			return nil
		},
		DeleteFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Location deleted")
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

func (m *LocationServiceMock) Describe(cmd *cobra.Command, args []string) error {
	if m.DescribeFunc != nil {
		return m.DescribeFunc(cmd, args)
	}
	return nil
}

func (m *LocationServiceMock) Update(cmd *cobra.Command, args []string) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(cmd, args)
	}
	return nil
}

func (m *LocationServiceMock) Delete(cmd *cobra.Command, args []string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(cmd, args)
	}
	return nil
}

func (m *LocationServiceMock) CreateVirtual(cmd *cobra.Command, args []string) error {
	if m.CreateVirtualFunc != nil {
		return m.CreateVirtualFunc(cmd, args)
	}
	return nil
}

func (m *LocationServiceMock) CreateVirtualNode(cmd *cobra.Command, args []string) error {
	if m.CreateVirtualNodeFunc != nil {
		return m.CreateVirtualNodeFunc(cmd, args)
	}
	return nil
}
