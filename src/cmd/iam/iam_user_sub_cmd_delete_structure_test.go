package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Delete(t *testing.T) {
	mockService := user.NewUserServiceMock()
	mockService.DeleteUserFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: User deleted successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"delete",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: User deleted successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Delete_PositionalArg(t *testing.T) {
	mockService := user.NewUserServiceMock()
	deleteCalled := false
	mockService.DeleteUserFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		cmd.Println("Mock: User deleted successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"delete",
		"550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleteCalled {
		t.Fatal("DeleteUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Delete_Alias(t *testing.T) {
	mockService := user.NewUserServiceMock()
	deleteCalled := false
	mockService.DeleteUserFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"rm",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleteCalled {
		t.Fatal("DeleteUser should be called via 'rm' alias")
	}
}

func TestIAMUserSubCmd_Structure_Delete_WithUsername(t *testing.T) {
	mockService := user.NewUserServiceMock()
	deleteCalled := false
	mockService.DeleteUserFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"delete",
		"--username", "jdoe",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !deleteCalled {
		t.Fatal("DeleteUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Delete_TooManyArgs(t *testing.T) {
	mockService := user.NewUserServiceMock()
	deleteCalled := false
	mockService.DeleteUserFunc = func(cmd *cobra.Command, args []string) error {
		deleteCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"delete",
		"user-001",
		"user-002",
	})

	err := iamCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for too many arguments, got nil")
	}
	if !strings.Contains(err.Error(), "accepts at most 1 arg(s)") {
		t.Fatalf("Expected 'accepts at most 1 arg' error, got %v", err)
	}
	if deleteCalled {
		t.Fatal("DeleteUser should not be called when too many arguments are provided")
	}
}
