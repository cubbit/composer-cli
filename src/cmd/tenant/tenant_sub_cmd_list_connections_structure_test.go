package cmd_tenant

import (
	"bytes"
	"testing"

	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func TestTenantSubCmd_Structure_ListConnections(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		actualTenantID, _ := cmd.Flags().GetString("tenant-id")
		if actualTenantID != "test-tenant" {
			t.Fatalf("Expected tenant-id 'test-tenant', got %q", actualTenantID)
		}
		cmd.Println("Mock: Connections listed successfully")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list-connections", "--tenant-id", "test-tenant"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connections listed successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_ListConnections_TenantName(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		actualTenantName, _ := cmd.Flags().GetString("tenant-name")
		if actualTenantName != "my-tenant" {
			t.Fatalf("Expected tenant-name 'my-tenant', got %q", actualTenantName)
		}
		cmd.Println("Mock: Connections listed by tenant name")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list-connections", "--tenant-name", "my-tenant"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connections listed by tenant name\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_ListConnections_WithPage(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		page, _ := cmd.Flags().GetInt("page")
		items, _ := cmd.Flags().GetInt("items")

		if page != 2 {
			t.Fatalf("Expected page 2, got %d", page)
		}
		if items != 50 {
			t.Fatalf("Expected items 50, got %d", items)
		}

		cmd.Println("Mock: Connections listed with pagination")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list-connections", "--tenant-id", "test", "--page", "2", "--items", "50"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connections listed with pagination\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_ListConnections_SkipVerifyFlag(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		skipVerify, _ := cmd.Flags().GetBool("skip-verify")
		if !skipVerify {
			t.Fatalf("Expected skip-verify to be true, got false")
		}
		cmd.Println("Mock: Connections listed with skip-verify")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list-connections", "--tenant-id", "test", "--skip-verify"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connections listed with skip-verify\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_ListConnections_Alias(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Connections listed via alias")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"ls-connections", "--tenant-id", "test"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connections listed via alias\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
