package cmd_agent

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAgentSubCmd_Structure_StatusWithAllRequiredFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.CheckAgentStatusFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent status retrieved successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"status",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--agent-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	})
	err := agentCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Agent status retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAgentSubCmd_Structure_StatusWithMissingRequiredFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.CheckAgentStatusFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent status retrieved successfully")
		return nil
	}

	flags := []string{
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--agent-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	}

	for i := 0; i < len(flags); i += 2 {
		agentCmd := NewAgentCmd(mockService)

		commandOutput := new(bytes.Buffer)
		agentCmd.SetOut(commandOutput)
		agentCmd.SetErr(commandOutput)

		args := []string{"status"}
		flagsCopy := append([]string(nil), flags...)
		missingFlag := flagsCopy[i][2:]
		filteredFlags := append(flagsCopy[:i], flagsCopy[i+2:]...)
		args = append(args, filteredFlags...)

		agentCmd.SetArgs(args)
		err := agentCmd.Execute()
		if err == nil {
			t.Fatalf("Expected error due to missing flag %s, got none", missingFlag)
		}

		if err != nil && err.Error() != fmt.Sprintf("required flag(s) \"%s\" not set", missingFlag) {
			t.Fatalf("error %s", err)
		}
	}
}
