package cmd_location

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestLocationSubCmd_Structure_CreateVirtualNodeWithAllFlags(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.CreateVirtualNodeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Virtual node created successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"create-virtual-node",
		"--name", "test-virtual-node",
		"--cluster-id", "test-cluster-id",
		"--storage-type", "test-storage-type",
		"--configuration", `{"endpoint": "https://s3.example.com", "bucket": "dev-bucket-1", "prefix": "folder", "region": "eu-west-1", "access_key": "my_access_key", "secret_key": "my_secret_key"}`,
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Virtual node created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestLocationSubCmd_Structure_CreateVirtualNodeWithMissingFlags(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.CreateVirtualNodeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Virtual node created successfully")
		return nil
	}

	missingFlags := [][]string{
		{"--name", "test-virtual-node"},
		{"--cluster-id", "test-cluster-id"},
		{"--storage-type", "test-storage-type"},
		{"--configuration", `{"endpoint": "https://s3.example.com", "bucket": "dev-bucket-1", "prefix": "folder", "region": "eu-west-1", "access_key": "my_access_key", "secret_key": "my_secret_key"}`},
	}
	for i := range missingFlags {
		locationCmd := NewLocationCmd(mockService)

		commandOutput := new(bytes.Buffer)
		locationCmd.SetOut(commandOutput)
		locationCmd.SetErr(commandOutput)

		flags := []string{
			"create-virtual-node",
		}
		for j := range missingFlags {
			if j != i {
				flags = append(flags, missingFlags[j]...)
			}
		}
		locationCmd.SetArgs(flags)
		err := locationCmd.Execute()
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	}
}

func TestLocationSubCmd_Structure_CreateVirtualNodeFailure(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.CreateVirtualNodeFunc = func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("mock error")
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"create-virtual-node",
		"--name", "test-virtual-node",
		"--cluster-id", "test-cluster-id",
		"--storage-type", "test-storage-type",
		"--configuration", `{"endpoint": "https://s3.example.com", "bucket": "dev-bucket-1", "prefix": "folder", "region": "eu-west-1", "access_key": "my_access_key", "secret_key": "my_secret_key"}`,
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
