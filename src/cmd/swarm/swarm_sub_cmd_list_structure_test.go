package cmd_swarm

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestSwarmSubCmd_Structure_List(t *testing.T) {
	mockService := service.NewSwarmServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarms listed successfully")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{"list"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Swarms listed successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestSwarmSubCmd_Structure_List_Alias(t *testing.T) {
	mockService := service.NewSwarmServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarms listed via alias")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{"ls"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Swarms listed via alias\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
