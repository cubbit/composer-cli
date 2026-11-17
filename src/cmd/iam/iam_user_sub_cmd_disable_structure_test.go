package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Disable(t *testing.T) {
	mockService := user.NewUserServiceMock()
	disableCalled := false
	mockService.DisableUserFunc = func(cmd *cobra.Command, args []string) error {
		disableCalled = true
		cmd.Println("Mock: User disabled successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"disable",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !disableCalled {
		t.Fatal("DisableUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Disable_PositionalArg(t *testing.T) {
	mockService := user.NewUserServiceMock()
	disableCalled := false
	mockService.DisableUserFunc = func(cmd *cobra.Command, args []string) error {
		disableCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"disable",
		"550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !disableCalled {
		t.Fatal("DisableUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Disable_WithUsername(t *testing.T) {
	mockService := user.NewUserServiceMock()
	disableCalled := false
	mockService.DisableUserFunc = func(cmd *cobra.Command, args []string) error {
		disableCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"disable",
		"--username", "jdoe",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !disableCalled {
		t.Fatal("DisableUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Disable_TooManyArgs(t *testing.T) {
	mockService := user.NewUserServiceMock()
	disableCalled := false
	mockService.DisableUserFunc = func(cmd *cobra.Command, args []string) error {
		disableCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"disable",
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
	if disableCalled {
		t.Fatal("DisableUser should not be called when too many arguments are provided")
	}
}
