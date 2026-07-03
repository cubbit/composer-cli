package cmd_auth

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAuthSubCmd_Structure_LogoutWithAllRequiredFlags(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LogoutFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Logout successful")
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"logout",
	})
	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Logout successful\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}
