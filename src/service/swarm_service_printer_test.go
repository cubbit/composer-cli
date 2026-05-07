package service

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/spf13/cobra"
)

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
