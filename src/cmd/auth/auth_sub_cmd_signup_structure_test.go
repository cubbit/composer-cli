package cmd_auth

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAuthSubCmd_Structure_SignUp_WithAllRequiredFlags(t *testing.T) {
	mockService := service.NewAuthServiceMock()

	mockService.SignUpFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator signed up successfully")
		return nil
	}

	authCmd := NewAuthCmd(mockService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"signup",
		"--email", "test@test.com",
		"--username", "testuser",
		"--organization", "TestOrg",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Operator signed up successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAuthSubCmd_Structure_SignUp_WithMissingFlag(t *testing.T) {
	mockService := service.NewAuthServiceMock()

	mockService.SignUpFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator signed up successfully")
		return nil
	}

	flags := []string{
		"--email", "test@test.com",
		"--username", "testuser",
		"--organization", "TestOrg",
	}

	for i := 0; i < len(flags); i += 2 {
		authCmd := NewAuthCmd(mockService)

		commandOutput := new(bytes.Buffer)
		authCmd.SetOut(commandOutput)
		authCmd.SetErr(commandOutput)

		args := []string{"signup"}
		flagsCopy := append([]string(nil), flags...)
		missingFlag := flagsCopy[i][2:]
		filteredFlags := append(flagsCopy[:i], flagsCopy[i+2:]...)
		args = append(args, filteredFlags...)

		authCmd.SetArgs(args)
		err := authCmd.Execute()
		if err == nil {
			t.Fatalf("Expected error due to missing flag %s, got none", missingFlag)
		}

		if err != nil && err.Error() != fmt.Sprintf("required flag(s) \"%s\" not set", missingFlag) {
			t.Fatalf("error %s", err)
		}
	}
}

func TestAuthSubCmd_Structure_SignUp_WithOptionalFlags(t *testing.T) {
	mockService := service.NewAuthServiceMock()

	mockService.SignUpFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator signed up successfully")
		return nil
	}

	authCmd := NewAuthCmd(mockService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"signup",
		"--email", "test@test.com",
		"--username", "testuser",
		"--organization", "TestOrg",
		"--first-name", "Test",
		"--last-name", "User",
		"--password", "SecurePassword123",
		"--base-policy", `{"policyKey":"policyValue"}`,
		"--settings", `{"settingKey":"settingValue"}`,
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Operator signed up successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAuthSubCmd_Structure_SignUp_WithInvalidJSONStringFlags(t *testing.T) {
	mockService := service.NewAuthServiceMock()

	mockService.SignUpFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator signed up successfully")
		return nil
	}

	authCmd := NewAuthCmd(mockService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"signup",
		"--email", "test@test.com",
		"--username", "testuser",
		"--organization", "TestOrg",
		"--first-name", "Test",
		"--last-name", "User",
		"--password", "SecurePassword123",
		"--base-policy", `{"policyKey":}`,
		"--settings", `{"settingKey":"settingValue"}`,
	})

	err := authCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error due to invalid JSON string flag, got none")
	}

	if err != nil &&
		!strings.Contains(err.Error(), "invalid argument") &&
		!strings.Contains(err.Error(), "--base-policy") {
		t.Fatalf("Expected JSON parsing error for --base-policy, got %s", err)
	}

	authCmd.SetArgs([]string{
		"signup",
		"--email", "test@test.com",
		"--username", "testuser",
		"--organization", "TestOrg",
		"--first-name", "Test",
		"--last-name", "User",
		"--password", "SecurePassword123",
		"--base-policy", `{"policyKey": "asd"}`,
		"--settings", `{"settingKe}`,
	})

	err = authCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error due to invalid JSON string flag, got none")
	}

	if err != nil &&
		!strings.Contains(err.Error(), "invalid argument") &&
		!strings.Contains(err.Error(), "--settings") {
		t.Fatalf("Expected JSON parsing error for --settings, got %s", err)
	}
}
