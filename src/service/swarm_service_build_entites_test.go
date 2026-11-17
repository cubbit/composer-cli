package service

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
)

func TestBuildNexuses_PhysicalClusters(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{
			ClusterID: "cluster-1",
			Name:      "Location 1",
			Type:      api.ClusterTypePhysical,
			Details: api.InfraAggregateClusterDetail{
				Nodes: []api.InfraAggregateNodeDetail{
					{
						NodeID:   "node-1",
						NodeName: "Server A",
						Disks: []api.InfraAggregateDiskDetail{
							{DiskUUID: "disk-1", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeOk)}},
							{DiskUUID: "disk-2", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeOk)}},
						},
					},
					{
						NodeID:   "node-2",
						NodeName: "Server B",
						Disks: []api.InfraAggregateDiskDetail{
							{DiskUUID: "disk-3", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeOk)}},
						},
					},
				},
			},
		},
	}

	nodeIDsByCluster := map[string][]string{
		"cluster-1": {"node-1", "node-2"},
	}

	nexuses := buildNexuses(clusters, nodeIDsByCluster)

	if len(nexuses) != 1 {
		t.Fatalf("Expected 1 nexus, got %d", len(nexuses))
	}

	nexus := nexuses[0]
	if nexus.ClusterID != "cluster-1" {
		t.Errorf("Expected ClusterID 'cluster-1', got '%s'", nexus.ClusterID)
	}
	if nexus.ClusterType != api.ClusterTypePhysical {
		t.Errorf("Expected ClusterType 'physical', got '%s'", nexus.ClusterType)
	}
	if len(nexus.Nodes) != 2 {
		t.Fatalf("Expected 2 nodes, got %d", len(nexus.Nodes))
	}
	if nexus.Nodes[0].ServerID != "node-1" {
		t.Errorf("Expected first node ID 'node-1', got '%s'", nexus.Nodes[0].ServerID)
	}
	if len(nexus.Nodes[0].Volumes) != 2 {
		t.Errorf("Expected node-1 to have 2 volumes, got %d", len(nexus.Nodes[0].Volumes))
	}
	if nexus.Nodes[1].ServerID != "node-2" {
		t.Errorf("Expected second node ID 'node-2', got '%s'", nexus.Nodes[1].ServerID)
	}
	if len(nexus.Nodes[1].Volumes) != 1 {
		t.Errorf("Expected node-2 to have 1 volume, got %d", len(nexus.Nodes[1].Volumes))
	}
}

func TestBuildNexuses_VirtualClusters(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{
			ClusterID: "vcluster-1",
			Name:      "Virtual Location",
			Type:      api.ClusterTypeVirtual,
			Details: api.InfraAggregateClusterDetail{
				VirtualNodes: []api.InfraAggregateVirtualNodeDetail{
					{NodeID: "vnode-1", NodeName: "VNode A"},
					{NodeID: "vnode-2", NodeName: "VNode B"},
				},
			},
		},
	}

	nodeIDsByCluster := map[string][]string{
		"vcluster-1": {"vnode-1"},
	}

	nexuses := buildNexuses(clusters, nodeIDsByCluster)

	if len(nexuses) != 1 {
		t.Fatalf("Expected 1 nexus, got %d", len(nexuses))
	}

	nexus := nexuses[0]
	if nexus.ClusterID != "vcluster-1" {
		t.Errorf("Expected ClusterID 'vcluster-1', got '%s'", nexus.ClusterID)
	}
	if nexus.ClusterType != api.ClusterTypeVirtual {
		t.Errorf("Expected ClusterType 'virtual', got '%s'", nexus.ClusterType)
	}
	if len(nexus.VirtualNodes) != 1 {
		t.Fatalf("Expected 1 virtual node, got %d", len(nexus.VirtualNodes))
	}
	if nexus.VirtualNodes[0].ServerID != "vnode-1" {
		t.Errorf("Expected virtual node ID 'vnode-1', got '%s'", nexus.VirtualNodes[0].ServerID)
	}
}

func TestBuildNexuses_EmptyNodeSelection(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{
			ClusterID: "cluster-1",
			Name:      "Location 1",
			Type:      api.ClusterTypePhysical,
			Details: api.InfraAggregateClusterDetail{
				Nodes: []api.InfraAggregateNodeDetail{
					{
						NodeID:   "node-1",
						NodeName: "Server A",
						Disks: []api.InfraAggregateDiskDetail{
							{DiskUUID: "disk-1", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeOk)}},
						},
					},
				},
			},
		},
	}

	nodeIDsByCluster := map[string][]string{
		"cluster-1": {},
	}

	nexuses := buildNexuses(clusters, nodeIDsByCluster)

	if len(nexuses) != 1 {
		t.Fatalf("Expected 1 nexus, got %d", len(nexuses))
	}

	if len(nexuses[0].Nodes) != 0 {
		t.Errorf("Expected 0 nodes for empty selection, got %d", len(nexuses[0].Nodes))
	}
}

func TestBuildNexuses_SkipsUnhealthyDisks(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{
			ClusterID: "cluster-1",
			Name:      "Location 1",
			Type:      api.ClusterTypePhysical,
			Details: api.InfraAggregateClusterDetail{
				Nodes: []api.InfraAggregateNodeDetail{
					{
						NodeID:   "node-1",
						NodeName: "Server A",
						Disks: []api.InfraAggregateDiskDetail{
							{DiskUUID: "disk-ok", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeOk)}},
							{DiskUUID: "disk-warn", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeWarning)}},
							{DiskUUID: "disk-err", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeError)}},
							{DiskUUID: "disk-inuse", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeInUse)}},
						},
					},
				},
			},
		},
	}

	nodeIDsByCluster := map[string][]string{
		"cluster-1": {"node-1"},
	}

	nexuses := buildNexuses(clusters, nodeIDsByCluster)

	if len(nexuses) != 1 {
		t.Fatalf("Expected 1 nexus, got %d", len(nexuses))
	}

	volumes := nexuses[0].Nodes[0].Volumes
	if len(volumes) != 1 {
		t.Fatalf("Expected 1 volume (only ok), got %d", len(volumes))
	}
	if volumes[0].VolumeID != "disk-ok" {
		t.Errorf("Expected volume ID 'disk-ok', got '%s'", volumes[0].VolumeID)
	}
}

func TestBuildNexuses_MultipleClusters(t *testing.T) {
	clusters := []api.InfraAggregateCluster{
		{
			ClusterID: "phys-cluster",
			Name:      "Physical",
			Type:      api.ClusterTypePhysical,
			Details: api.InfraAggregateClusterDetail{
				Nodes: []api.InfraAggregateNodeDetail{
					{
						NodeID:   "phys-node",
						NodeName: "Phys Server",
						Disks: []api.InfraAggregateDiskDetail{
							{DiskUUID: "disk-1", Status: api.InfraAggregateStatus{Code: string(api.StatusCodeOk)}},
						},
					},
				},
			},
		},
		{
			ClusterID: "virt-cluster",
			Name:      "Virtual",
			Type:      api.ClusterTypeVirtual,
			Details: api.InfraAggregateClusterDetail{
				VirtualNodes: []api.InfraAggregateVirtualNodeDetail{
					{NodeID: "virt-node", NodeName: "Virt Server"},
				},
			},
		},
	}

	nodeIDsByCluster := map[string][]string{
		"phys-cluster": {"phys-node"},
		"virt-cluster": {"virt-node"},
	}

	nexuses := buildNexuses(clusters, nodeIDsByCluster)

	if len(nexuses) != 2 {
		t.Fatalf("Expected 2 nexuses, got %d", len(nexuses))
	}

	if nexuses[0].ClusterType != api.ClusterTypePhysical || len(nexuses[0].Nodes) == 0 {
		t.Error("Expected first nexus to be physical with nodes")
	}
	if nexuses[1].ClusterType != api.ClusterTypeVirtual || len(nexuses[1].VirtualNodes) == 0 {
		t.Error("Expected second nexus to be virtual with virtual nodes")
	}
}
