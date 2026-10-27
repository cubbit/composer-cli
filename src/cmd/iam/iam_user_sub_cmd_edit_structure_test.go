package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Edit(t *testing.T) {
	mockService := user.NewUserServiceMock()
	mockService.EditUserFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: User edited successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
		"--first-name", "Alice",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: User edited successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Edit_PositionalArg(t *testing.T) {
	mockService := user.NewUserServiceMock()
	editCalled := false
	mockService.EditUserFunc = func(cmd *cobra.Command, args []string) error {
		editCalled = true
		cmd.Println("Mock: User edited successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"550e8400-e29b-41d4-a716-446655440000",
		"--last-name", "Doe",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !editCalled {
		t.Fatal("EditUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Edit_UpdateAlias(t *testing.T) {
	mockService := user.NewUserServiceMock()
	editCalled := false
	mockService.EditUserFunc = func(cmd *cobra.Command, args []string) error {
		editCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"update",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
		"--email", "alice@example.com",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !editCalled {
		t.Fatal("EditUser should be called via 'update' alias")
	}
}

func TestIAMUserSubCmd_Structure_Edit_WithUsername(t *testing.T) {
	mockService := user.NewUserServiceMock()
	editCalled := false
	mockService.EditUserFunc = func(cmd *cobra.Command, args []string) error {
		editCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"--username", "jdoe",
		"--first-name", "John",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !editCalled {
		t.Fatal("EditUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Edit_WithoutFields(t *testing.T) {
	mockService := user.NewUserServiceMock()
	editCalled := false
	mockService.EditUserFunc = func(cmd *cobra.Command, args []string) error {
		editCalled = true
		return cmd.Help()
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no structure-level error, got %v", err)
	}
	if !editCalled {
		t.Fatal("EditUser should be called so the service can validate missing update fields")
	}
}

func TestIAMUserSubCmd_Structure_Edit_TooManyArgs(t *testing.T) {
	mockService := user.NewUserServiceMock()
	editCalled := false
	mockService.EditUserFunc = func(cmd *cobra.Command, args []string) error {
		editCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"user-001",
		"user-002",
		"--first-name", "Alice",
	})

	err := iamCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for too many arguments, got nil")
	}
	if !strings.Contains(err.Error(), "accepts at most 1 arg(s)") {
		t.Fatalf("Expected 'accepts at most 1 arg' error, got %v", err)
	}
	if editCalled {
		t.Fatal("EditUser should not be called when too many arguments are provided")
	}
}
