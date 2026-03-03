package cmd_config

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestConfigSubCmd_Structure_SwitchProfile_WithRequiredArgument(t *testing.T) {
	mockConfigService := service.NewConfigServiceMock()

	mockConfigService.SwitchProfileFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Profile switched successfully")
		return nil
	}

	configCmd := NewConfigCmd(mockConfigService)

	commandOutput := new(bytes.Buffer)
	configCmd.SetOut(commandOutput)
	configCmd.SetErr(commandOutput)
	configCmd.SetArgs([]string{
		"switch-profile",
		"test-profile",
	})

	err := configCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Profile switched successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestConfigSubCmd_Structure_SwitchProfile_WithMissingArgument(t *testing.T) {
	mockConfigService := service.NewConfigServiceMock()

	mockConfigService.SwitchProfileFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Profile switched successfully")
		return nil
	}

	configCmd := NewConfigCmd(mockConfigService)

	commandOutput := new(bytes.Buffer)
	configCmd.SetOut(commandOutput)
	configCmd.SetErr(commandOutput)
	configCmd.SetArgs([]string{
		"switch-profile",
	})

	err := configCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error due to missing argument, got none")
	}

	if err.Error() != "accepts 1 arg(s), received 0" {
		t.Fatalf("Expected argument count error, got %s", err)
	}
}
