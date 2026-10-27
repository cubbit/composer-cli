package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewaySubCmd_Structure_Create_WithAllFlags(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway created successfully with all flags")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--description", "Test gateway description",
		"--ingress-type", "manual",
		"--swarm-rc", "550e8400-e29b-41d4-a716-446655440001:550e8400-e29b-41d4-a716-446655440002:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Gateway created successfully with all flags\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestGatewaySubCmd_Structure_Create_WithMultipleSwarmRCFlags(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway created with multiple swarm-rc flags")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "550e8400-e29b-41d4-a716-446655440001:550e8400-e29b-41d4-a716-446655440002:true",
		"--swarm-rc", "550e8400-e29b-41d4-a716-446655440003:550e8400-e29b-41d4-a716-446655440004:false",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Gateway created with multiple swarm-rc flags\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestGatewaySubCmd_Structure_Create_WithOptionalDescription(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway created with description")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--description", "This is a test gateway",
		"--ingress-type", "manual",
		"--swarm-rc", "550e8400-e29b-41d4-a716-446655440001:550e8400-e29b-41d4-a716-446655440002:true",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Gateway created with description\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestGatewaySubCmd_Structure_Create_InteractiveMode(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway created in interactive mode")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"create",
		"--interactive",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Gateway created in interactive mode\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestGatewaySubCmd_Structure_Create_MissingName(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"create",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
		"--swarm-rc", "550e8400-e29b-41d4-a716-446655440001:550e8400-e29b-41d4-a716-446655440002:true",
	})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --name is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"name\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}

func TestGatewaySubCmd_Structure_Create_MissingSwarmRC(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"create",
		"--name", "test-gateway",
		"--slug", "test-gateway",
		"--cluster-id", "550e8400-e29b-41d4-a716-446655440000",
		"--ingress-type", "manual",
	})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --swarm-rc is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"swarm-rc\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}
