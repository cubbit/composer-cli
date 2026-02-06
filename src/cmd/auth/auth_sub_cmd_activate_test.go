package cmd_auth

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAuthSubCmd_Structure_Activate_WithAllRequiredFlags(t *testing.T) {
	mockService := service.NewAuthServiceMock()

	mockService.ActivateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator activated successfully")
		return nil
	}

	authCmd := NewAuthCmd(mockService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"activate",
		"--token", "activation-token-123",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Operator activated successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAuthSubCmd_Structure_Activate_WithMissingFlag(t *testing.T) {
	mockService := service.NewAuthServiceMock()

	mockService.ActivateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator activated successfully")
		return nil
	}

	flags := []string{
		"--token", "activation-token-123",
	}

	for i := 0; i < len(flags); i += 2 {
		authCmd := NewAuthCmd(mockService)

		commandOutput := new(bytes.Buffer)
		authCmd.SetOut(commandOutput)
		authCmd.SetErr(commandOutput)

		args := []string{"activate"}
		flagsCopy := append([]string(nil), flags...)
		missingFlag := flagsCopy[i][2:]
		filteredFlags := append(flagsCopy[:i], flagsCopy[i+2:]...)
		args = append(args, filteredFlags...)

		authCmd.SetArgs(args)
		err := authCmd.Execute()
		if err == nil {
			t.Fatalf("Expected error due to missing flag %q, got none", missingFlag)
		}

		commandOutputString := commandOutput.String()
		expectedErrorSubstring := "required flag(s) \"" + missingFlag + "\" not set"
		if !strings.Contains(commandOutputString, expectedErrorSubstring) {
			t.Fatalf("Expected error message to contain %q, got %q", expectedErrorSubstring, commandOutputString)
		}
	}
}
