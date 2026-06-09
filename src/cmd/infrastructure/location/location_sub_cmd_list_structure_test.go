package cmd_location

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestLocationSubCmd_Structure_ListWithAllFlags(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Locations retrieved successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"list",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Locations retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}
