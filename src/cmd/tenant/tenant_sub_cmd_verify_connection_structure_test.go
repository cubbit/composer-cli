package cmd_tenant

import (
	"bytes"
	"testing"

	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func TestTenantSubCmd_Structure_VerifyConnection(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.VerifyConnectionFunc = func(cmd *cobra.Command, args []string) error {
		actualTenantID, _ := cmd.Flags().GetString("tenant-id")
		actualConnID, _ := cmd.Flags().GetString("connection-id")

		if actualTenantID != "test-tenant" {
			t.Fatalf("Expected tenant-id 'test-tenant', got %q", actualTenantID)
		}
		if actualConnID != "test-conn" {
			t.Fatalf("Expected connection-id 'test-conn', got %q", actualConnID)
		}

		cmd.Println("Mock: Connection verified successfully")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"verify-connection", "--tenant-id", "test-tenant", "--connection-id", "test-conn"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connection verified successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_VerifyConnection_TenantName(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.VerifyConnectionFunc = func(cmd *cobra.Command, args []string) error {
		actualTenantName, _ := cmd.Flags().GetString("tenant-name")
		actualConnID, _ := cmd.Flags().GetString("connection-id")

		if actualTenantName != "my-tenant" {
			t.Fatalf("Expected tenant-name 'my-tenant', got %q", actualTenantName)
		}
		if actualConnID != "test-conn" {
			t.Fatalf("Expected connection-id 'test-conn', got %q", actualConnID)
		}

		cmd.Println("Mock: Connection verified by tenant name")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"verify-connection", "--tenant-name", "my-tenant", "--connection-id", "test-conn"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connection verified by tenant name\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_VerifyConnection_Alias(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.VerifyConnectionFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Connection verified via alias")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)

	aliases := []string{"verify-conn", "vconn"}
	for _, alias := range aliases {
		commandOutput.Reset()
		tenantCmd.SetArgs([]string{alias, "--tenant-id", "test", "--connection-id", "test-conn"})

		err := tenantCmd.Execute()
		if err != nil {
			t.Fatalf("Expected no error for alias %q, got %v", alias, err)
		}

		expectedOutput := "Mock: Connection verified via alias\n"
		if commandOutput.String() != expectedOutput {
			t.Fatalf("For alias %q, expected output %q, got %q", alias, expectedOutput, commandOutput.String())
		}
	}
}
