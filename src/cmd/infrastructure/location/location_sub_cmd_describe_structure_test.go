package cmd_location

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestLocationSubCmd_Structure_DescribeWithAllFlags(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.ListAggregatedFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Aggregated locations retrieved successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Aggregated locations retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestLocationSubCmd_Structure_DescribeWithClusterName(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.ListAggregatedFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Aggregated locations retrieved successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name",
		"test-cluster",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Aggregated locations retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestLocationSubCmd_Structure_DescribeWithClusterID(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.ListAggregatedFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Aggregated locations retrieved successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-id",
		"test-cluster-id",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Aggregated locations retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestLocationSubCmd_Structure_DescribeMutuallyExclusiveFlags(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"describe",
		"--cluster-name",
		"test-cluster",
		"--cluster-id",
		"test-cluster-id",
	})
	err := locationCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error for mutually exclusive flags, got nil")
	}
}
