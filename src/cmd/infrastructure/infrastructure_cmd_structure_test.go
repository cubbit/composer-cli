package cmd_infrastructure

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestInfrastructureSubCmd_Structure_LocationList(t *testing.T) {
	mockLocationService := service.NewLocationServiceMock()

	mockLocationService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Locations retrieved successfully")
		return nil
	}

	infrastructureCmd := NewInfrastructureCmd(mockLocationService)

	commandOutput := new(bytes.Buffer)
	infrastructureCmd.SetOut(commandOutput)
	infrastructureCmd.SetErr(commandOutput)
	infrastructureCmd.SetArgs([]string{
		"location",
		"list",
	})
	err := infrastructureCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Locations retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestInfrastructureSubCmd_Structure_LocationDescribe(t *testing.T) {
	mockLocationService := service.NewLocationServiceMock()

	mockLocationService.ListAggregatedFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Aggregated locations retrieved successfully")
		return nil
	}

	infrastructureCmd := NewInfrastructureCmd(mockLocationService)

	commandOutput := new(bytes.Buffer)
	infrastructureCmd.SetOut(commandOutput)
	infrastructureCmd.SetErr(commandOutput)
	infrastructureCmd.SetArgs([]string{
		"location",
		"describe",
		"--cluster-id", "test-cluster-id",
	})
	err := infrastructureCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Aggregated locations retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}
