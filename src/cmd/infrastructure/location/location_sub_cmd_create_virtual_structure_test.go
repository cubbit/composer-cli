package cmd_location

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestLocationSubCmd_Structure_CreateVirtualWithAllFlags(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.CreateVirtualFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Virtual location created successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"create-virtual",
		"--name", "test-virtual-location",
		"--description", "This is a test virtual location",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Virtual location created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestLocationSubCmd_Structure_CreateVirtualWithMissingName(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.CreateVirtualFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Virtual location created successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"create-virtual",
	})
	err := locationCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestLocationSubCmd_Structure_CreateVirtualFailure(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.CreateVirtualFunc = func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("mock error")
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"create-virtual",
		"--name", "test-virtual-location",
		"--description", "This is a test virtual location",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "ERR mock error\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}
