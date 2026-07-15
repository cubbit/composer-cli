package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Create(t *testing.T) {
	mockService := service.NewUserServiceMock()
	mockService.CreateUserFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: User created successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"create",
		"--username", "user1",
		"--password", "secret",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: User created successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Create_MissingUsername(t *testing.T) {
	mockService := service.NewUserServiceMock()
	createCalled := false
	mockService.CreateUserFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"create",
		"--password", "secret",
	})

	err := iamCmd.Execute()
	if err == nil {
		t.Fatal("Expected error when --username is missing, got nil")
	}

	if !strings.Contains(err.Error(), `required flag(s) "username" not set`) {
		t.Fatalf("Expected missing username flag error, got %v", err)
	}
	if createCalled {
		t.Fatal("CreateUser should not be called when required flags are missing")
	}
}
