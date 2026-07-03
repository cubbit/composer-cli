package service

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/spf13/cobra"
)

func TestPrintSwarmList_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Bool("no-headers", false, "no headers")
	cmd.Flags().Set("quiet", "false")
	cmd.Flags().Set("no-headers", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	status := api.EvaluatedStatusType("online")
	firstSync := createdAt.Add(30 * time.Minute)
	secondSync := createdAt.Add(90 * time.Minute)

	swarms := []api.ListSwarmV5ItemPresentation{
		{
			ListSwarmV5Item: api.ListSwarmV5Item{
				ID:                   "swarm-001",
				Name:                 "prod-swarm",
				TotalStorageBytes:    1099511627776,
				UsedStorageBytes:     549755813888,
				CreatedAt:            createdAt,
				NexusCount:           3,
				RedundancyClassCount: 2,
			},
			SummaryStatusNullable: api.SummaryStatusNullable{
				EvaluatedStatus:              &status,
				EvaluatedStatusLastUpdatedAt: &firstSync,
			},
		},
		{
			ListSwarmV5Item: api.ListSwarmV5Item{
				ID:                   "swarm-002",
				Name:                 "dev-swarm",
				TotalStorageBytes:    107374182400,
				UsedStorageBytes:     0,
				CreatedAt:            createdAt,
				NexusCount:           1,
				RedundancyClassCount: 1,
			},
			SummaryStatusNullable: api.SummaryStatusNullable{
				EvaluatedStatusLastUpdatedAt: &secondSync,
			},
		},
	}

	err = PrintSwarmList(cmd, swarms)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
╭────────────┬───────────────────┬─────────┬───────────┬─────────────────────┬─────────────────────┬────────╮
│ Name       │ Used Capacity     │ % Usage │ Locations │ Created On          │ Last Sync           │ Status │
├────────────┼───────────────────┼─────────┼───────────┼─────────────────────┼─────────────────────┼────────┤
│ prod-swarm │ 512.00 GB/1.00 TB │ 50%     │ 3         │ 2024-01-15 10:30:00 │ 2024-01-15 11:00:00 │ Online │
│ dev-swarm  │ 0 bytes/100.00 GB │ 0%      │ 1         │ 2024-01-15 10:30:00 │ 2024-01-15 12:00:00 │ N/A    │
╰────────────┴───────────────────┴─────────┴───────────┴─────────────────────┴─────────────────────┴────────╯
`)

	if strings.TrimSpace(out.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(out.String()))
	}
}

func TestPrintSwarmList_Human_Empty(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Bool("no-headers", false, "no headers")
	cmd.Flags().Set("quiet", "false")
	cmd.Flags().Set("no-headers", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintSwarmList(cmd, []api.ListSwarmV5ItemPresentation{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := out.String()
	if !strings.Contains(output, "No swarms found") {
		t.Fatalf("Expected empty state message, got %q", output)
	}
}

func TestPrintSwarmDetails_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	statusUpdatedAt := createdAt.Add(2 * time.Hour)
	description := "Production swarm"
	evaluatedStatus := api.EvaluatedStatusType("online")

	swarm := api.SwarmV5Presentation{
		SwarmV5: api.SwarmV5{
			ListSwarmV5Item: api.ListSwarmV5Item{
				ID:                   "swarm-123",
				Name:                 "test-swarm",
				TotalStorageBytes:    1099511627776,
				UsedStorageBytes:     549755813888,
				CreatedAt:            createdAt,
				NexusCount:           3,
				RedundancyClassCount: 2,
			},
			OrganizationID: "org-123",
			OwnerID:        "owner-123",
			Description:    &description,
			Configuration: map[string]interface{}{
				"features": []interface{}{"tiering", "compression"},
				"foo":      "bar",
				"nested": map[string]interface{}{
					"enabled": true,
					"mode":    "strict",
				},
			},
			CreationStatus: "created",
		},
		SummaryStatusNullable: api.SummaryStatusNullable{
			EvaluatedStatus:              &evaluatedStatus,
			EvaluatedStatusLastUpdatedAt: &statusUpdatedAt,
		},
	}

	err = PrintSwarmDetails(cmd, swarm)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Swarm: test-swarm
Status: ● Online
Last Update: 2024-01-15 12:30:00
Organization: org-123
Owner: owner-123

Storage Usage:
  Usage: [█████░░░░░] 50%
  Total Used: 512.00 GB
  Total Assigned: 1.00 TB
  Total Unused: 512.00 GB

Metadata:
  ID: swarm-123
  Description: Production swarm
  Created At: 2024-01-15 10:30:00
  Creation Status: created

Composition:
  Nexus Count: 3
  Redundancy Class Count: 2

Configuration:
  features: [tiering, compression]
  foo: bar
  nested: {enabled: true, mode: strict}`)

	if strings.TrimSpace(out.String()) != expected {
		t.Fatalf("Expected output:\n%s\nactual:\n%s", expected, strings.TrimSpace(out.String()))
	}
}
