package cmd_config

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestConfigSubCmd_Structure_Profiles(t *testing.T) {
	mockConfigService := service.NewConfigServiceMock()

	mockConfigService.ProfilesFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Profiles listed successfully")
		return nil
	}

	configCmd := NewConfigCmd(mockConfigService)

	commandOutput := new(bytes.Buffer)
	configCmd.SetOut(commandOutput)
	configCmd.SetErr(commandOutput)
	configCmd.SetArgs([]string{
		"profiles",
	})

	err := configCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Profiles listed successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}
