package cmd_agent

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAgentSubCmd_Structure_ListWithAllRequiredFlags_ForNexusAndNode(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.ListAgentsFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent listed successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"list",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--sort", "id",
		"--filter", "node_id:b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	})
	err := agentCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Agent listed successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAgentSubCmd_Structure_ListWithAllRequiredFlags_ForRedundancyClass(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.ListAgentsFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent listed successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"list",
		"--redundancy-class-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--sort", "id",
		"--filter", "node_id:b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	})
	err := agentCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Agent listed successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAgentSubCmd_Structure_ListWithAllRequiredFlags_WithMutuallyExclusiveFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.ListAgentsFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent listed successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"list",
		"--redundancy-class-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--sort", "id",
		"--filter", "node_id:b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	})
	err := agentCmd.Execute()
	if err == nil || !strings.Contains(err.Error(), "if any flags in the group [nexus-id redundancy-class-id] are set none of the others can be; [nexus-id redundancy-class-id] were all set") {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestAgentSubCmd_Structure_ListWithAllRequiredFlags_WithInvalidSort(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.ListAgentsFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent listed successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"list",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--sort", "invalid",
		"--filter", "node_id:b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	})
	err := agentCmd.Execute()
	if err == nil || !strings.Contains(err.Error(), "error: invalid sort key provided, allowed keys are: id, node_id, port, created_at") {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestAgentSubCmd_Structure_ListWithAllRequiredFlags_WithInvalidFilter(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.ListAgentsFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent listed successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"list",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--sort", "node_id",
		"--filter", "invalid_filter_format",
	})
	err := agentCmd.Execute()
	if err == nil || !strings.Contains(err.Error(), "error: invalid filter provided, allowed format is: key:value key:value") {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestAgentSubCmd_Structure_ListWithMissingFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.ListAgentsFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent listed successfully")
		return nil
	}

	flags := []string{
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	}

	for i := 0; i < len(flags); i += 2 {
		agentCmd := NewAgentCmd(mockService)

		commandOutput := new(bytes.Buffer)
		agentCmd.SetOut(commandOutput)
		agentCmd.SetErr(commandOutput)

		args := []string{"list"}
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
