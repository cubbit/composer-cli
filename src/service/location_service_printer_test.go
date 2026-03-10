package service

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/spf13/cobra"
)

func TestPrintClusterDetails_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	cluster := api.InfraAggregateCluster{
		ClusterID: "550e8400-e29b-41d4-a716-446655440000",
		Name:      "Test Cluster",
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
						Code:    api.StatusCodeOk,
						Details: "All systems operational",
					},
					OSName:     strPtr("Ubuntu 22.04 LTS"),
					CPU:        &api.InfraNodeCPUInfo{Cores: 16},
					RAM:        &api.InfraNodeRAMInfo{Available: 64.0},
					ExternalIP: strPtr("203.0.113.10"),
					InternalIP: strPtr("10.0.0.10"),
					Disks: []api.InfraAggregateDiskDetail{
						{
							DiskUUID:              "8e7d7a2e-3c5e-4f9a-9b8c-1d2e3f4a5b6c",
							Path:                  "/dev/sda",
							Used:                  true,
							PVRef:                 "f47ac10b-58cc-4372-a567-0e02b2c3d479",
							TotalStorageSizeBytes: 1099511627776,
							UsedStorageBytes:      549755813888,
							Status: api.InfraAggregateStatus{
								Code:    api.StatusCodeOk,
								Details: "Mounted and active",
							},
						},
						{
							DiskUUID:              "c40c6d3f-1e2a-4b8c-9d5e-6f7a8b9c0d1e",
							Path:                  "/dev/sdb",
							Used:                  false,
							PVRef:                 "",
							TotalStorageSizeBytes: 2199023255552,
							UsedStorageBytes:      0,
							Status: api.InfraAggregateStatus{
								Code:    api.StatusCodeOk,
								Details: "Available",
							},
						},
					},
				},
				{
					NodeID:   "f47ac10b-58cc-4372-a567-0e02b2c3d480",
					NodeName: "node-2-host",
					Status: api.InfraAggregateStatus{
						Code:    api.StatusCodeOk,
						Details: "All systems operational",
					},
					OSName:     strPtr("Ubuntu 22.04 LTS"),
					CPU:        &api.InfraNodeCPUInfo{Cores: 32},
					RAM:        &api.InfraNodeRAMInfo{Available: 128.0},
					ExternalIP: strPtr("203.0.113.11"),
					InternalIP: strPtr("10.0.0.11"),
					Disks: []api.InfraAggregateDiskDetail{
						{
							DiskUUID:              "123e4567-e89b-12d3-a456-426614174000",
							Path:                  "/dev/sdc",
							Used:                  true,
							PVRef:                 "9f3a3651-3362-4c1e-8d5a-7e8f9a0b1c2d",
							TotalStorageSizeBytes: 10995116277760,
							UsedStorageBytes:      3298534883328,
							Status: api.InfraAggregateStatus{
								Code:    api.StatusCodeOk,
								Details: "Mounted and active",
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
						Code:    api.StatusCodeOk,
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
				{
					NodeID:   "a8c5e3f1-2d4b-6e9a-8c7d-5f3a1b9c7e5f",
					NodeName: "vnode-2-host",
					Status: api.InfraAggregateStatus{
						Code:    api.StatusCodeOk,
						Details: "S3 service running",
					},
					StorageType: api.VirtualStorageTypeS3,
					StorageConfiguration: map[string]any{
						"endpoint":      "https://s3-backup.example.com",
						"bucket_name":   "cubbit-backup",
						"region":        "eu-west-1",
						"access_key_id": "AKIAIOSFODNN7BACKUP",
					},
				},
			},
		},
	}

	err = PrintClusterDetails(cmd, cluster)

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	expectedResult := strings.TrimSpace(`
Cluster Information
─────────────────────────────────────────────────
╭──────────────────────────────────────┬──────────────┬──────────┬────────────────┬───────────────┬─────────────────────╮
│ Cluster ID                           │ Name         │ Type     │ Physical Nodes │ Virtual Nodes │ Last Update         │
├──────────────────────────────────────┼──────────────┼──────────┼────────────────┼───────────────┼─────────────────────┤
│ 550e8400-e29b-41d4-a716-446655440000 │ Test Cluster │ physical │ 2              │ 2             │ 2024-01-15 10:30:00 │
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
    ├── Disk: 8e7d7a2e-3c5e-4f9a-9b8c-1d2e3f4a5b6c
    │   ├── Path: /dev/sda
    │   ├── Status: status_ok
    │   ├── Total: 1.00 TB
    │   ├── Used: 512.00 GB
    │   └── PV Reference: f47ac10b-58cc-4372-a567-0e02b2c3d479
    └── Disk: c40c6d3f-1e2a-4b8c-9d5e-6f7a8b9c0d1e
        ├── Path: /dev/sdb
        ├── Status: status_ok
        ├── Total: 2.00 TB
        └── Used: 0 bytes
node-2-host (f47ac10b-58cc-4372-a567-0e02b2c3d480)
├── Status: status_ok - All systems operational
├── Hardware
│   ├── OS: Ubuntu 22.04 LTS
│   ├── CPU Cores: 32
│   └── RAM: 128.00 GB
├── Network
│   ├── External IP: 203.0.113.11
│   └── Internal IP: 10.0.0.11
└── Disks
    └── Disk: 123e4567-e89b-12d3-a456-426614174000
        ├── Path: /dev/sdc
        ├── Status: status_ok
        ├── Total: 10.00 TB
        ├── Used: 3.00 TB
        └── PV Reference: 9f3a3651-3362-4c1e-8d5a-7e8f9a0b1c2d

Virtual Nodes
─────────────────────────────────────────────────
vnode-1-host (4b6e8c2f-9a1d-4f3e-8b5c-7d4a6e5b3c2f)
├── Status: status_ok - S3 service running
└── Storage: s3
    ├── access_key_id: AKIAIOSFODNN7EXAMPLE
    ├── bucket_name: cubbit-virtual
    ├── endpoint: https://s3.example.com
    └── region: us-east-1
vnode-2-host (a8c5e3f1-2d4b-6e9a-8c7d-5f3a1b9c7e5f)
├── Status: status_ok - S3 service running
└── Storage: s3
    ├── access_key_id: AKIAIOSFODNN7BACKUP
    ├── bucket_name: cubbit-backup
    ├── endpoint: https://s3-backup.example.com
    └── region: eu-west-1
`)

	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected cluster details output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func strPtr(s string) *string {
	return &s
}

func TestPrintClusterDetails_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	cluster := api.InfraAggregateCluster{
		ClusterID: "550e8400-e29b-41d4-a716-446655440000",
		Name:      "Test Cluster",
		Type:      api.ClusterTypePhysical,
		Details: api.InfraAggregateClusterDetail{
			LastUpdate: baseTime,
		},
	}

	err = PrintClusterDetails(cmd, cluster)

	if err != nil {
		t.Errorf("Expected no error in quiet mode, got: %v", err)
	}

	if out.Len() > 0 {
		t.Errorf("Expected no output in quiet mode, got: %s", out.String())
	}
}

func TestPrintClusters(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	baseTime, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	clusters := []api.InfraAggregateCluster{
		{
			ClusterID: "550e8400-e29b-41d4-a716-446655440000",
			Name:      "Cluster One",
			Type:      api.ClusterTypePhysical,
			Details: api.InfraAggregateClusterDetail{
				LastUpdate: baseTime,
				Nodes: []api.InfraAggregateNodeDetail{
					{
						NodeID:   "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
						NodeName: "node-1",
						Status: api.InfraAggregateStatus{
							Code: api.StatusCodeOk,
						},
					},
				},
				VirtualNodes: []api.InfraAggregateVirtualNodeDetail{
					{
						NodeID:   "4b6e8c2f-9a1d-4f3e-8b5c-7d4a6e5b3c2f",
						NodeName: "vnode-1",
						Status: api.InfraAggregateStatus{
							Code: api.StatusCodeOk,
						},
					},
				},
			},
		},
		{
			ClusterID: "6ba7b810-9dad-11d1-80b4-00c04fd430c9",
			Name:      "Cluster Two",
			Type:      api.ClusterTypeVirtual,
			Details: api.InfraAggregateClusterDetail{
				LastUpdate: baseTime.Add(time.Hour),
				Nodes: []api.InfraAggregateNodeDetail{
					{
						NodeID:   "f47ac10b-58cc-4372-a567-0e02b2c3d480",
						NodeName: "node-2",
						Status: api.InfraAggregateStatus{
							Code: api.StatusCodeOk,
						},
					},
					{
						NodeID:   "f47ac10b-58cc-4372-a567-0e02b2c3d481",
						NodeName: "node-3",
						Status: api.InfraAggregateStatus{
							Code:    api.StatusCodeWarning,
							Details: "Some issues detected",
						},
					},
				},
				VirtualNodes: []api.InfraAggregateVirtualNodeDetail{
					{
						NodeID:   "a8c5e3f1-2d4b-6e9a-8c7d-5f3a1b9c7e5f",
						NodeName: "vnode-2",
						Status: api.InfraAggregateStatus{
							Code: api.StatusCodeOk,
						},
					},
				},
			},
		},
	}

	err = PrintClusters(cmd, clusters)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────────────────────────────┬─────────────┬──────────┬────────────────┬───────────────┬─────────────────────╮
│ Cluster ID                           │ Name        │ Type     │ Physical Nodes │ Virtual Nodes │ Last Update         │
├──────────────────────────────────────┼─────────────┼──────────┼────────────────┼───────────────┼─────────────────────┤
│ 550e8400-e29b-41d4-a716-446655440000 │ Cluster One │ physical │ 1              │ 1             │ 2024-01-15 10:30:00 │
│ 6ba7b810-9dad-11d1-80b4-00c04fd430c9 │ Cluster Two │ virtual  │ 2              │ 1             │ 2024-01-15 11:30:00 │
╰──────────────────────────────────────┴─────────────┴──────────┴────────────────┴───────────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected clusters output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}
