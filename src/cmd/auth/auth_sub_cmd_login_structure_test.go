package cmd_auth

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAuthSubCmd_Structure_LoginWithAllRequiredFlags(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Login successful")
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test-profile",
	})
	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Login successful\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAuthSubCmd_Structure_LoginWithMissingFlags(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Login successful")
		return nil
	}

	flags := []string{
		"--profile", "test-profile",
	}

	for i := 0; i < len(flags); i += 2 {
		authCmd := NewAuthCmd(mockAuthService)

		commandOutput := new(bytes.Buffer)
		authCmd.SetOut(commandOutput)
		authCmd.SetErr(commandOutput)

		args := []string{"login"}
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
