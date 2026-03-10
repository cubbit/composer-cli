package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/api"
)

type AggregateClusterView struct {
	ClusterID     string
	Name          string
	ClusterType   string
	PhysicalNodes int
	VirtualNodes  int
	LastUpdate    string
}

func NewAggregateClusterView(cluster api.InfraAggregateCluster) AggregateClusterView {
	return AggregateClusterView{
		ClusterID:     cluster.ClusterID,
		Name:          cluster.Name,
		ClusterType:   string(cluster.Type),
		PhysicalNodes: len(cluster.Details.Nodes),
		VirtualNodes:  len(cluster.Details.VirtualNodes),
		LastUpdate:    cluster.Details.LastUpdate.Format("2006-01-02 15:04:05"),
	}
}

func (v AggregateClusterView) GetRow() []string {
	return []string{
		v.ClusterID,
		v.Name,
		v.ClusterType,
		fmt.Sprintf("%d", v.PhysicalNodes),
		fmt.Sprintf("%d", v.VirtualNodes),
		v.LastUpdate,
	}
}

func ConvertAggregateClustersView(clusters []api.InfraAggregateCluster) []AggregateClusterView {
	views := make([]AggregateClusterView, len(clusters))
	for i, cluster := range clusters {
		views[i] = NewAggregateClusterView(cluster)
	}
	return views
}
