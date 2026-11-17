package cmd_location

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestLocationSubCmd_Structure_UpdateWithAllFlags(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	mockService.UpdateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Location updated successfully")
		return nil
	}

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"update",
		"--cluster-id", "cluster-1",
		"--name", "Milan datacenter",
		"--description", "Primary datacenter",
		"--position", "Milan, Italy",
		"--contact-user", "jdoe",
	})

	if err := locationCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if commandOutput.String() != "Mock: Location updated successfully\n" {
		t.Fatalf("Unexpected output %q", commandOutput.String())
	}
}

func TestLocationSubCmd_Structure_UpdateRequiresClusterID(t *testing.T) {
	mockService := service.NewLocationServiceMock()

	locationCmd := NewLocationCmd(mockService)

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"update",
		"--name", "Milan datacenter",
	})

	err := locationCmd.Execute()
	if err == nil {
		t.Fatal("Expected an error when --cluster-id is missing")
	}
	if !strings.Contains(err.Error(), "cluster-id") {
		t.Fatalf("Expected an error about the required cluster-id flag, got %v", err)
	}
}
