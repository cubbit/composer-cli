package cmd_tenant

import (
	"bytes"
	"strings"
	"testing"

	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func TestTenantSubCmd_Structure_Create_WithAllFlags(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Tenant created successfully")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1,gatewayID2:mysub",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Tenant created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestTenantSubCmd_Structure_Create_WithMultipleConnections(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Tenant created successfully")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1,gatewayID2",
		"--connection", "domainID2:gatewayID3",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Tenant created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestTenantSubCmd_Structure_Create_WithDescription(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Tenant created successfully")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
		"--description", "Production tenant",
		"--connection", "domainID:gatewayID1,gatewayID2:mysub",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Tenant created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestTenantSubCmd_Structure_Create_InteractiveMode(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Tenant created in interactive mode")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--interactive",
	})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Tenant created in interactive mode\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestTenantSubCmd_Structure_Create_MissingName(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--slug", "cubbit-1",
		"--connection", "domainID:gatewayID1",
	})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --name is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s)") || !strings.Contains(err.Error(), "\"name\"") {
		t.Fatalf("Expected error message to indicate missing required flag 'name', got: %v", err)
	}
}

func TestTenantSubCmd_Structure_Create_MissingSlug(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--connection", "domainID:gatewayID1",
	})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --slug is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s)") || !strings.Contains(err.Error(), "\"slug\"") {
		t.Fatalf("Expected error message to indicate missing required flag 'slug', got: %v", err)
	}
}

func TestTenantSubCmd_Structure_Create_MissingConnection(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
		"--name", "Tenant Cubbit 1",
		"--slug", "cubbit-1",
	})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --connection is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s)") || !strings.Contains(err.Error(), "\"connection\"") {
		t.Fatalf("Expected error message to indicate missing required flag 'connection', got: %v", err)
	}
}

func TestTenantSubCmd_Structure_Create_MissingAllFlags(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{
		"create",
	})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when required flags are missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s)") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}
