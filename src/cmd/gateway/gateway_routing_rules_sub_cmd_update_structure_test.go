package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayRoutingRulesUpdateCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.UpdateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
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
		filePath, _ := cmd.Flags().GetString("file-path")
		if filePath != "rule.json" {
			t.Fatalf("Expected file path rule.json, got %q", filePath)
		}

		cmd.Println("Mock: routing rule updated")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"routing-rules", "update",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"--file-path", "rule.json",
	})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "Mock: routing rule updated"
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesUpdateCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	updateCalled := false
	mockService.UpdateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		updateCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "update", "--id", testRoutingRuleID, "--file-path", "rule.json"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "gateway-id" not set`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if updateCalled {
		t.Fatal("UpdateRoutingRule should not be called without --gateway-id")
	}
}

func TestGatewayRoutingRulesUpdateCmd_Structure_RequiresID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	updateCalled := false
	mockService.UpdateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		updateCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "update", "--gateway-id", testRoutingRuleGatewayID, "--file-path", "rule.json"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "id" not set`) {
		t.Fatalf("Expected required id flag error, got %q", err.Error())
	}
	if updateCalled {
		t.Fatal("UpdateRoutingRule should not be called without --id")
	}
}

func TestGatewayRoutingRulesUpdateCmd_Structure_RequiresFilePath(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	updateCalled := false
	mockService.UpdateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		updateCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "update", "--gateway-id", testRoutingRuleGatewayID, "--id", testRoutingRuleID})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --file-path flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "file-path" not set`) {
		t.Fatalf("Expected required file-path flag error, got %q", err.Error())
	}
	if updateCalled {
		t.Fatal("UpdateRoutingRule should not be called without --file-path")
	}
}

func TestGatewayRoutingRulesUpdateCmd_Structure_RejectsPositionalDocumentPath(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	updateCalled := false
	mockService.UpdateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		updateCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{
		"routing-rules", "update",
		"--gateway-id", testRoutingRuleGatewayID,
		"--id", testRoutingRuleID,
		"rule.json",
	})

	if err := gatewayCmd.Execute(); err == nil {
		t.Fatal("Expected error for positional document path argument")
	}
	if updateCalled {
		t.Fatal("UpdateRoutingRule should not be called with positional arguments")
	}
}

func TestGatewayRoutingRulesUpdateCmd_Structure_HelpDocumentsFlagsAndReplaceSemantics(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	updateCmd, _, err := gatewayCmd.Find([]string{"routing-rules", "update"})
	if err != nil {
		t.Fatalf("Expected routing-rules update command, got %v", err)
	}

	for _, expected := range []string{
		"--gateway-id",
		"--id",
		"--file-path",
		"replaces the whole policy",
		"confirmation message",
		"get --id",
	} {
		if !strings.Contains(updateCmd.Long, expected) {
			t.Fatalf("Expected update help to document %q, got %q", expected, updateCmd.Long)
		}
	}
	if strings.Contains(updateCmd.Long, "indented JSON") {
		t.Fatalf("Expected update help not to claim the document is printed, got %q", updateCmd.Long)
	}
}
