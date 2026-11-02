package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayRoutingRulesCreateCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.CreateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != testRoutingRuleGatewayID {
			t.Fatalf("Expected gateway ID %s, got %q", testRoutingRuleGatewayID, gatewayID)
		}
		filePath, _ := cmd.Flags().GetString("file-path")
		if filePath != "rule.json" {
			t.Fatalf("Expected file path rule.json, got %q", filePath)
		}

		cmd.Println("Mock: routing rule created")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"routing-rules", "create",
		"--gateway-id", testRoutingRuleGatewayID,
		"--file-path", "rule.json",
	})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "Mock: routing rule created"
	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestGatewayRoutingRulesCreateCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "create", "--file-path", "rule.json"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "gateway-id" not set`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateRoutingRule should not be called without --gateway-id")
	}
}

func TestGatewayRoutingRulesCreateCmd_Structure_RequiresFilePath(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "create", "--gateway-id", testRoutingRuleGatewayID})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --file-path flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "file-path" not set`) {
		t.Fatalf("Expected required file-path flag error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateRoutingRule should not be called without --file-path")
	}
}

func TestGatewayRoutingRulesCreateCmd_Structure_RejectsPositionalDocumentPath(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateRoutingRuleFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"routing-rules", "create", "--gateway-id", testRoutingRuleGatewayID, "rule.json"})

	if err := gatewayCmd.Execute(); err == nil {
		t.Fatal("Expected error for positional document path argument")
	}
	if createCalled {
		t.Fatal("CreateRoutingRule should not be called with positional arguments")
	}
}

func TestGatewayRoutingRulesCreateCmd_Structure_HelpDocumentsFilePathAndStdin(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	createCmd, _, err := gatewayCmd.Find([]string{"routing-rules", "create"})
	if err != nil {
		t.Fatalf("Expected routing-rules create command, got %v", err)
	}

	for _, expected := range []string{
		"--file-path",
		`"-"`,
		"stdin",
		"confirmation message",
		"get --id",
	} {
		if !strings.Contains(createCmd.Long, expected) {
			t.Fatalf("Expected create help to document %q, got %q", expected, createCmd.Long)
		}
	}
	if strings.Contains(createCmd.Long, "indented JSON") {
		t.Fatalf("Expected create help not to claim the document is printed, got %q", createCmd.Long)
	}
}
