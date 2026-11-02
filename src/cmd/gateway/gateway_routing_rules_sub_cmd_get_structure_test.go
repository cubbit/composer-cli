package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayRoutingRulesGetCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.GetRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != testRoutingRuleGatewayID {
			t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
		}
		ruleID, _ := cmd.Flags().GetString("id")
		if ruleID != testRoutingRuleID {
			t.Fatalf("Expected rule ID %s, got %q", testRoutingRuleID, ruleID)
		}

		cmd.Println("Mock: routing rule shown")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"routing-rules", "get", "--gateway-id", testRoutingRuleGatewayID, "--id", testRoutingRuleID})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expectedResult := "Mock: routing rule shown"
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected mock output, got %q", actualResult)
	}
}

func TestGatewayRoutingRulesGetCmd_Structure_RejectsPositionalGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	serviceCalled := false
	mockService.GetRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		serviceCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "get", testRoutingRuleGatewayID, "--id", testRoutingRuleID})

	if err := gatewayCmd.Execute(); err == nil {
		t.Fatal("Expected error for positional gateway ID argument")
	}
	if serviceCalled {
		t.Fatal("GetRoutingRule should not be called with positional arguments")
	}
}

func TestGatewayRoutingRulesGetCmd_Structure_HelpMentionsGatewayAndRuleIDFlags(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	getCmd, _, err := gatewayCmd.Find([]string{"routing-rules", "get"})
	if err != nil {
		t.Fatalf("Expected routing-rules get command, got %v", err)
	}

	for _, expected := range []string{
		"--id",
		"--output json",
		"--output yaml",
		"full policy document",
		"--gateway-id",
		"list",
		"update --file-path",
	} {
		if !strings.Contains(getCmd.Long, expected) {
			t.Fatalf("Expected get help to document %q, got %q", expected, getCmd.Long)
		}
	}
	if strings.Contains(getCmd.Long, "more than one") {
		t.Fatalf("Expected get help not to mention the old single/multiple rule model, got %q", getCmd.Long)
	}
}
