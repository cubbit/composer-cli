package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayConnectionListCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != "gateway-123" {
			t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
		}

		cmd.Println("Mock: Connections listed successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"connection", "list", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connections listed successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionListCmd_Structure_Alias(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	listCalled := false
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "ls", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !listCalled {
		t.Fatal("Expected the ls alias to invoke ListConnections")
	}
}

func TestGatewayConnectionListCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	listCalled := false
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "list"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "gateway-id" not set`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if listCalled {
		t.Fatal("ListConnections should not be called without --gateway-id")
	}
}

func TestGatewayConnectionListCmd_Structure_RejectsPositionalGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	listCalled := false
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "list", "gateway-123"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for positional gateway ID argument")
	}
	if listCalled {
		t.Fatal("ListConnections should not be called with positional arguments")
	}
}

func TestGatewayConnectionListCmd_Structure_HelpMentionsGatewayIDFlag(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	connectionCmd, _, err := gatewayCmd.Find([]string{"connection", "list"})
	if err != nil {
		t.Fatalf("Expected connection list command, got %v", err)
	}

	if !strings.Contains(connectionCmd.Long, "--gateway-id <gateway-id>") {
		t.Fatalf("Expected list help to mention the --gateway-id flag, got %q", connectionCmd.Long)
	}
	if strings.Contains(connectionCmd.Long, "GATEWAY_ID as a positional") {
		t.Fatalf("Expected list help not to mention positional gateway IDs, got %q", connectionCmd.Long)
	}
}
