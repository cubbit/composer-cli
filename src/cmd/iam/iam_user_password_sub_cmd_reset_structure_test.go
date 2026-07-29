package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func TestIAMUserPasswordSubCmd_Structure_Reset(t *testing.T) {
	mockService := user.NewUserServiceMock()
	resetCalled := false
	mockService.ResetUserPasswordFunc = func(cmd *cobra.Command, args []string) error {
		resetCalled = true
		cmd.Println("Mock: User password reset successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
		"--new-password", "new-secret",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resetCalled {
		t.Fatal("ResetUserPassword should be called")
	}

	expectedOutput := "Mock: User password reset successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserPasswordSubCmd_Structure_Reset_PositionalArg(t *testing.T) {
	mockService := user.NewUserServiceMock()
	resetCalled := false
	mockService.ResetUserPasswordFunc = func(cmd *cobra.Command, args []string) error {
		resetCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"550e8400-e29b-41d4-a716-446655440000",
		"--new-password", "new-secret",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resetCalled {
		t.Fatal("ResetUserPassword should be called")
	}
}

func TestIAMUserPasswordSubCmd_Structure_Reset_WithUsername(t *testing.T) {
	mockService := user.NewUserServiceMock()
	resetCalled := false
	mockService.ResetUserPasswordFunc = func(cmd *cobra.Command, args []string) error {
		resetCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"--username", "jdoe",
		"--new-password", "new-secret",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resetCalled {
		t.Fatal("ResetUserPassword should be called")
	}
}

func TestIAMUserPasswordSubCmd_Structure_Reset_DefersTargetValidationToService(t *testing.T) {
	mockService := user.NewUserServiceMock()
	resetCalled := false
	mockService.ResetUserPasswordFunc = func(cmd *cobra.Command, args []string) error {
		resetCalled = true
		if len(args) != 0 {
			t.Fatalf("Expected no args, got %v", args)
		}
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"--new-password", "new-secret",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !resetCalled {
		t.Fatal("ResetUserPassword should be called")
	}
}

func TestIAMUserPasswordSubCmd_Structure_Reset_MissingNewPassword(t *testing.T) {
	mockService := user.NewUserServiceMock()
	resetCalled := false
	mockService.ResetUserPasswordFunc = func(cmd *cobra.Command, args []string) error {
		resetCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err == nil {
		t.Fatal("Expected error when --new-password is missing, got nil")
	}
	if !strings.Contains(err.Error(), `required flag(s) "new-password" not set`) {
		t.Fatalf("Expected missing new-password flag error, got %v", err)
	}
	if resetCalled {
		t.Fatal("ResetUserPassword should not be called when required flags are missing")
	}
}

func TestIAMUserPasswordSubCmd_Structure_Reset_TooManyArgs(t *testing.T) {
	mockService := user.NewUserServiceMock()
	resetCalled := false
	mockService.ResetUserPasswordFunc = func(cmd *cobra.Command, args []string) error {
		resetCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"user-001",
		"user-002",
		"--new-password", "new-secret",
	})

	err := iamCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for too many arguments, got nil")
	}
	if !strings.Contains(err.Error(), "accepts at most 1 arg(s)") {
		t.Fatalf("Expected 'accepts at most 1 arg' error, got %v", err)
	}
	if resetCalled {
		t.Fatal("ResetUserPassword should not be called when too many arguments are provided")
	}
}
