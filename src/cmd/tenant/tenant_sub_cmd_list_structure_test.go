package cmd_tenant

import (
	"bytes"
	"testing"

	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func TestTenantSubCmd_Structure_List(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Tenants listed successfully")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Tenants listed successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_List_Alias(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Tenants listed via alias")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"ls"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Tenants listed via alias\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_List_WithQueryFlag(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		actualQuery, _ := cmd.Flags().GetStringArray("query")
		if len(actualQuery) != 1 || actualQuery[0] != "name:eq(test)" {
			t.Fatalf("Expected query 'name:eq(test)', got %v", actualQuery)
		}
		cmd.Println("Mock: Tenants listed with query")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list", "--query", "name:eq(test)"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Tenants listed with query\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_List_WithMultipleQueryFlags(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		actualQuery, _ := cmd.Flags().GetStringArray("query")
		if len(actualQuery) != 2 {
			t.Fatalf("Expected 2 query values, got %d: %v", len(actualQuery), actualQuery)
		}
		if actualQuery[0] != "name:eq(test)" {
			t.Fatalf("Expected first query 'name:eq(test)', got %q", actualQuery[0])
		}
		if actualQuery[1] != "status:eq(active)" {
			t.Fatalf("Expected second query 'status:eq(active)', got %q", actualQuery[1])
		}
		cmd.Println("Mock: Tenants listed with multiple queries")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list", "--query", "name:eq(test)", "--query", "status:eq(active)"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Tenants listed with multiple queries\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_List_WithPageAndItems(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		page, _ := cmd.Flags().GetInt("page")
		items, _ := cmd.Flags().GetInt("items")

		if page != 2 {
			t.Fatalf("Expected page 2, got %d", page)
		}
		if items != 50 {
			t.Fatalf("Expected items 50, got %d", items)
		}

		cmd.Println("Mock: Tenants listed with pagination")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"list", "--page", "2", "--items", "50"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Tenants listed with pagination\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
