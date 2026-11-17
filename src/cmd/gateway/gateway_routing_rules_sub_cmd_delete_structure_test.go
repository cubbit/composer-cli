package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayRoutingRulesDeleteCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.DeleteRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
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

		cmd.Println("Mock: routing rule deleted")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"routing-rules", "delete", "--gateway-id", testRoutingRuleGatewayID, "--id", testRoutingRuleID})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "Mock: routing rule deleted"
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesDeleteCmd_Structure_RmAliasWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "rm", "--gateway-id", testRoutingRuleGatewayID, "--id", testRoutingRuleID})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleteCalled {
		t.Fatal("DeleteRoutingRule should be called through the rm alias")
	}
}

func TestGatewayRoutingRulesDeleteCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "delete", "--id", testRoutingRuleID})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "gateway-id" not set`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteRoutingRule should not be called without --gateway-id")
	}
}

func TestGatewayRoutingRulesDeleteCmd_Structure_RequiresID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "delete", "--gateway-id", testRoutingRuleGatewayID})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "id" not set`) {
		t.Fatalf("Expected required id flag error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteRoutingRule should not be called without --id")
	}
}

func TestGatewayRoutingRulesDeleteCmd_Structure_RejectsPositionalRuleID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "delete", "--gateway-id", testRoutingRuleGatewayID, testRoutingRuleID})

	if err := gatewayCmd.Execute(); err == nil {
		t.Fatal("Expected error for positional rule ID argument")
	}
	if deleteCalled {
		t.Fatal("DeleteRoutingRule should not be called with positional arguments")
	}
}

func TestGatewayRoutingRulesDeleteCmd_Structure_HelpMentionsListAndID(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	deleteCmd, _, err := gatewayCmd.Find([]string{"routing-rules", "delete"})
	if err != nil {
		t.Fatalf("Expected routing-rules delete command, got %v", err)
	}

	for _, expected := range []string{"--gateway-id", "--id", "list"} {
		if !strings.Contains(deleteCmd.Long, expected) {
			t.Fatalf("Expected delete help to document %q, got %q", expected, deleteCmd.Long)
		}
	}
}
