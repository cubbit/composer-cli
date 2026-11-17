package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayStorageClassCreateCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.CreateStorageClassFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != "gateway-123" {
			t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
		}

		name, _ := cmd.Flags().GetString("name")
		if name != "cold-data" {
			t.Fatalf("Expected storage class name cold-data, got %q", name)
		}

		cmd.Println("Mock: Storage class created successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"storage-class", "create",
		"--gateway-id", "gateway-123",
		"--name", "cold-data",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Storage class created successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassCreateCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateStorageClassFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"storage-class", "create", "--name", "cold-data"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `"gateway-id"`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateStorageClass should not be called without --gateway-id")
	}
}

func TestGatewayStorageClassCreateCmd_Structure_RequiresName(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateStorageClassFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"storage-class", "create", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --name flag")
	}
	if !strings.Contains(err.Error(), `"name"`) {
		t.Fatalf("Expected required name flag error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateStorageClass should not be called without --name")
	}
}

func TestGatewayStorageClassCreateCmd_Structure_HelpUsesGatewayIDFlag(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	storageClassCmd, _, err := gatewayCmd.Find([]string{"storage-class", "create"})
	if err != nil {
		t.Fatalf("Expected storage-class create command, got %v", err)
	}

	if !strings.Contains(storageClassCmd.Long, "--gateway-id <gateway-id>") {
		t.Fatalf("Expected create help to mention the --gateway-id flag, got %q", storageClassCmd.Long)
	}
	if strings.Contains(storageClassCmd.Long, "Pool storage class names") {
		t.Fatalf("Expected create help not to mention pool storage class names, got %q", storageClassCmd.Long)
	}
}
