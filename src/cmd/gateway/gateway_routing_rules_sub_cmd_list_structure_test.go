package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayRoutingRulesListCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.ListRoutingRulesFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != testRoutingRuleGatewayID {
			t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
		}

		cmd.Println("Mock: routing rules listed")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"routing-rules", "list", "--gateway-id", testRoutingRuleGatewayID})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "Mock: routing rules listed"
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesListCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	listCalled := false
	mockService.ListRoutingRulesFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "list"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "gateway-id" not set`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if listCalled {
		t.Fatal("ListRoutingRules should not be called without --gateway-id")
	}
}

func TestGatewayRoutingRulesListCmd_Structure_RejectsPositionalGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	listCalled := false
	mockService.ListRoutingRulesFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "list", testRoutingRuleGatewayID})

	if err := gatewayCmd.Execute(); err == nil {
		t.Fatal("Expected error for positional gateway ID argument")
	}
	if listCalled {
		t.Fatal("ListRoutingRules should not be called with positional arguments")
	}
}

func TestGatewayRoutingRulesListCmd_Structure_HelpDocumentsMetadataOnly(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	listCmd, _, err := gatewayCmd.Find([]string{"routing-rules", "list"})
	if err != nil {
		t.Fatalf("Expected routing-rules list command, got %v", err)
	}

	for _, expected := range []string{
		"--gateway-id <gateway-id>",
		"--output json",
		"metadata only",
		"get",
	} {
		if !strings.Contains(listCmd.Long, expected) {
			t.Fatalf("Expected list help to document %q, got %q", expected, listCmd.Long)
		}
	}
	if strings.Contains(listCmd.Long, "Action") {
		t.Fatalf("Expected list help not to mention the action column, got %q", listCmd.Long)
	}
}
