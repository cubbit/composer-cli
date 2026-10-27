package cmd_gateway

import (
	"bytes"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewaySubCmd_Structure_Describe_WithPositionalID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway described successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"describe", "gateway-123"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateway described successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_Describe_WithGatewayName(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway described successfully by name")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"describe", "--gateway-name", "test-gateway"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateway described successfully by name\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_Describe_WithGatewayIDFlag(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway described successfully by ID flag")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"describe", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateway described successfully by ID flag\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_Describe_AliasInfo(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway described via alias")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"info", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateway described via alias\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
