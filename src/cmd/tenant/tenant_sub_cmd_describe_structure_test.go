package cmd_tenant

import (
	"bytes"
	"testing"

	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func TestTenantSubCmd_Structure_Describe(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || args[0] != "tnt-001" {
			t.Fatalf("Expected args ['tnt-001'], got %v", args)
		}
		cmd.Println("Mock: Tenant described successfully")
		return nil
	}

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"describe", "tnt-001"})

	err := tenantCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Tenant described successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestTenantSubCmd_Structure_Describe_NoArgs_Error(t *testing.T) {
	mockService := servicetenant.NewTenantServiceMock()

	tenantCmd := NewTenantCmd(mockService)

	commandOutput := new(bytes.Buffer)
	tenantCmd.SetOut(commandOutput)
	tenantCmd.SetErr(commandOutput)
	tenantCmd.SetArgs([]string{"describe"})

	err := tenantCmd.Execute()
	if err == nil {
		t.Fatal("Expected error from Execute, got nil")
	}

	expectedError := "accepts 1 arg(s), received 0"
	if err.Error() != expectedError {
		t.Fatalf("Expected error %q, got %q", expectedError, err.Error())
	}
}
