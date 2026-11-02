package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayConnectionDeleteCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.DeleteConnectionFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != "gateway-123" {
			t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
		}

		connectionID, _ := cmd.Flags().GetString("id")
		if connectionID != testConnectionID {
			t.Fatalf("Expected connection ID %s, got %q", testConnectionID, connectionID)
		}

		cmd.Println("Mock: Connection deleted successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"connection", "delete",
		"--gateway-id", "gateway-123",
		"--id", testConnectionID,
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connection deleted successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionDeleteCmd_Structure_Alias(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteConnectionFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{
		"connection", "rm", "--gateway-id", "gateway-123", "--id", testConnectionID,
	})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleteCalled {
		t.Fatal("Expected the rm alias to invoke DeleteConnection")
	}
}

func TestGatewayConnectionDeleteCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteConnectionFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "delete", "--id", testConnectionID})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `"gateway-id"`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteConnection should not be called without --gateway-id")
	}
}

func TestGatewayConnectionDeleteCmd_Structure_RequiresID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteConnectionFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "delete", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "id" not set`) {
		t.Fatalf("Expected required id flag error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteConnection should not be called without --id")
	}
}

func TestGatewayConnectionDeleteCmd_Structure_RejectsPositionalGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteConnectionFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "delete", "gateway-123", "--id", testConnectionID})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for positional gateway ID argument")
	}
	if deleteCalled {
		t.Fatal("DeleteConnection should not be called with positional arguments")
	}
}

func TestGatewayConnectionDeleteCmd_Structure_Help(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	connectionCmd, _, err := gatewayCmd.Find([]string{"connection", "delete"})
	if err != nil {
		t.Fatalf("Expected connection delete command, got %v", err)
	}

	if !strings.Contains(connectionCmd.Long, "default connection cannot be deleted") {
		t.Fatalf("Expected delete help to document the default connection restriction, got %q", connectionCmd.Long)
	}
	if !strings.Contains(connectionCmd.Long, "number of impacted rules") ||
		!strings.Contains(connectionCmd.Long, "--output json|yaml") {
		t.Fatalf("Expected delete help to document how impacted rules are shown, got %q", connectionCmd.Long)
	}
	if !strings.Contains(connectionCmd.Long, "--gateway-id <gateway-id>") {
		t.Fatalf("Expected delete help to mention the --gateway-id flag, got %q", connectionCmd.Long)
	}
}
