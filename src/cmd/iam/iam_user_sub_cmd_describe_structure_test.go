package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Describe(t *testing.T) {
	mockService := user.NewUserServiceMock()
	mockService.DescribeUserFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: User described successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"describe",
		"--user-id", "user-001",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: User described successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Describe_PositionalArg(t *testing.T) {
	mockService := user.NewUserServiceMock()
	var capturedUserID string
	mockService.DescribeUserFunc = func(cmd *cobra.Command, args []string) error {
		capturedUserID = args[0]
		cmd.Println("Mock: User described successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"describe",
		"user-002",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if capturedUserID != "user-002" {
		t.Fatalf("Expected captured user ID %q, got %q", "user-002", capturedUserID)
	}

	expectedOutput := "Mock: User described successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Describe_WithUsername(t *testing.T) {
	mockService := user.NewUserServiceMock()
	mockService.DescribeUserFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: User described successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"describe",
		"--username", "jdoe",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: User described successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Describe_CurrentUser(t *testing.T) {
	mockService := user.NewUserServiceMock()
	describeCalled := false
	mockService.DescribeUserFunc = func(cmd *cobra.Command, args []string) error {
		describeCalled = true
		if len(args) != 0 {
			t.Fatalf("Expected no args, got %v", args)
		}
		cmd.Println("Mock: User described successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"describe",
		"--self",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !describeCalled {
		t.Fatal("DescribeUser should be called")
	}

	expectedOutput := "Mock: User described successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Describe_TooManyArgs(t *testing.T) {
	mockService := user.NewUserServiceMock()

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"describe",
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
}
