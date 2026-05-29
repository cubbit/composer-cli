package describe

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/spf13/cobra"
)

func TestPrintGatewayDetails_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	gateway := api.GatewayV5GetResponse{
		ID:     "gateway-123",
		Name:   "test-gateway",
		Slug:   "test-gateway",
		Type:   "singlecluster",
		Status: "ready",
		RedundancyClasses: []api.GatewayV5GetRedundancyClass{
			{
				ID:   "rc-123",
				Name: "standard",
			},
			{
				ID:   "rc-456",
				Name: "archive",
			},
		},
	}

	err := PrintGatewayDetails(cmd, gateway)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Gateway: test-gateway
Status: ● Ready

Metadata:
  ID: gateway-123
  Slug: test-gateway
  Type: singlecluster

Redundancy Classes:
  - standard (rc-123)
  - archive (rc-456)`)

	if strings.TrimSpace(out.String()) != expected {
		t.Fatalf("Expected output:\n%s\nactual:\n%s", expected, strings.TrimSpace(out.String()))
	}
}

func TestPrintGatewayDetails_Human_WithoutRedundancyClasses(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	gateway := api.GatewayV5GetResponse{
		ID:     "gateway-123",
		Name:   "test-gateway",
		Slug:   "test-gateway",
		Type:   "manual",
		Status: "not-ready",
	}

	err := PrintGatewayDetails(cmd, gateway)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := out.String()
	if !strings.Contains(output, "Redundancy Classes:\n  N/A") {
		t.Fatalf("Expected empty redundancy classes output, got %q", output)
	}
}
