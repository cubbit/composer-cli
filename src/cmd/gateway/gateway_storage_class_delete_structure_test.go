package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayStorageClassDeleteCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.DeleteStorageClassFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != "gateway-123" {
			t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
		}

		storageClassID, _ := cmd.Flags().GetString("id")
		if storageClassID != "sc-001" {
			t.Fatalf("Expected storage class ID sc-001, got %q", storageClassID)
		}

		cmd.Println("Mock: Storage class deleted successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"storage-class", "delete",
		"--gateway-id", "gateway-123",
		"--id", "sc-001",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Storage class deleted successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayStorageClassDeleteCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteStorageClassFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"storage-class", "delete", "--id", "sc-001"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `"gateway-id"`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteStorageClass should not be called without --gateway-id")
	}
}

func TestGatewayStorageClassDeleteCmd_Structure_RequiresID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	deleteCalled := false
	mockService.DeleteStorageClassFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"storage-class", "delete", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --id flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "id" not set`) {
		t.Fatalf("Expected required id flag error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteStorageClass should not be called without --id")
	}
}

func TestGatewayStorageClassDeleteCmd_Structure_HelpDoesNotClaimKeptRules(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	storageClassCmd, _, err := gatewayCmd.Find([]string{"storage-class", "delete"})
	if err != nil {
		t.Fatalf("Expected storage-class delete command, got %v", err)
	}

	for _, stale := range []string{
		"Routing rules",
		"are kept",
		"uploads using",
	} {
		if strings.Contains(storageClassCmd.Long, stale) {
			t.Fatalf("Expected delete help not to contain %q, got %q", stale, storageClassCmd.Long)
		}
	}
	if !strings.Contains(storageClassCmd.Long, "Default storage classes cannot be deleted") {
		t.Fatalf("Expected delete help to mention default storage class deletion limits, got %q", storageClassCmd.Long)
	}
}
