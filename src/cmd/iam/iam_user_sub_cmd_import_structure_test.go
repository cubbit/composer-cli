package cmd_iam

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Import_WithFile(t *testing.T) {
	mockService := service.NewUserServiceMock()
	mockService.ImportUsersFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Users imported successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"import",
		"--file", "users.json",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Users imported successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Import_WithSample(t *testing.T) {
	mockService := service.NewUserServiceMock()
	importCalled := false
	mockService.ImportUsersFunc = func(cmd *cobra.Command, args []string) error {
		importCalled = true
		cmd.Println("Mock: Sample generated")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{"user", "import", "--sample", "json"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !importCalled {
		t.Fatal("ImportUsers should be called for sample generation")
	}
	expectedOutput := "Mock: Sample generated\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
