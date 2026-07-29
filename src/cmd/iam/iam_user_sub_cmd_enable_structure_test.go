package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Enable(t *testing.T) {
	mockService := user.NewUserServiceMock()
	enableCalled := false
	mockService.EnableUserFunc = func(cmd *cobra.Command, args []string) error {
		enableCalled = true
		cmd.Println("Mock: User enabled successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"enable",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !enableCalled {
		t.Fatal("EnableUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Enable_PositionalArg(t *testing.T) {
	mockService := user.NewUserServiceMock()
	enableCalled := false
	mockService.EnableUserFunc = func(cmd *cobra.Command, args []string) error {
		enableCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"enable",
		"550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !enableCalled {
		t.Fatal("EnableUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Enable_WithUsername(t *testing.T) {
	mockService := user.NewUserServiceMock()
	enableCalled := false
	mockService.EnableUserFunc = func(cmd *cobra.Command, args []string) error {
		enableCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"enable",
		"--username", "jdoe",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !enableCalled {
		t.Fatal("EnableUser should be called")
	}
}

func TestIAMUserSubCmd_Structure_Enable_TooManyArgs(t *testing.T) {
	mockService := user.NewUserServiceMock()
	enableCalled := false
	mockService.EnableUserFunc = func(cmd *cobra.Command, args []string) error {
		enableCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"enable",
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
	if enableCalled {
		t.Fatal("EnableUser should not be called when too many arguments are provided")
	}
}
