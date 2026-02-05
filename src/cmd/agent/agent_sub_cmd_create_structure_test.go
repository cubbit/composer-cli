package cmd_agent

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAgentSubCmd_Structure_CreateSingleAgentWithAllRequiredFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.CreateAgentFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent created successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"create",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--agent-port", "8080",
		"--agent-disk", "/dev/sda1",
		"--agent-mount-point", "/mnt/agent",
	})
	err := agentCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Agent created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAgentSubCmd_Structure_CreateSingleAgentWithMissingFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.CreateAgentFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Agent created successfully")
		return nil
	}

	flags := []string{
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--agent-port", "8080",
		"--agent-disk", "/dev/sda1",
		"--agent-mount-point", "/mnt/agent",
	}

	for i := 0; i < len(flags); i += 2 {
		agentCmd := NewAgentCmd(mockService)

		commandOutput := new(bytes.Buffer)
		agentCmd.SetOut(commandOutput)
		agentCmd.SetErr(commandOutput)

		args := []string{"create"}
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

func TestAgentSubCmd_Structure_CreateAgentBatchWithAllRequiredFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.CreateAgentBatchFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Batch agent created successfully")
		return nil
	}

	// Create a temporary directory, this is autocleaned up after the test
	tmpDir := t.TempDir()
	batchFilePath := filepath.Join(tmpDir, "agents_batch.json")
	content := []byte(`{
		"agents": [
			{"id": "agent-1"},
			{"id": "agent-2"}
		]
	}`)
	err := os.WriteFile(batchFilePath, content, 0644)
	if err != nil {
		t.Fatalf("failed to write temp batch file: %v", err)
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"create",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--batch",
		"--file", batchFilePath,
	})
	err = agentCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Batch agent created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAgentSubCmd_Structure_CreateAgentBatchWithMissingFlags(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.CreateAgentBatchFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Batch agent created successfully")
		return nil
	}

	// Create a temporary directory, this is autocleaned up after the test
	tmpDir := t.TempDir()
	batchFilePath := filepath.Join(tmpDir, "agents_batch.json")
	content := []byte(`{
		"agents": [
			{"id": "agent-1"},
			{"id": "agent-2"}
		]
	}`)
	err := os.WriteFile(batchFilePath, content, 0644)
	if err != nil {
		t.Fatalf("failed to write temp batch file: %v", err)
	}

	flags := []string{
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
	}

	for i := 0; i < len(flags); i += 2 {
		agentCmd := NewAgentCmd(mockService)

		commandOutput := new(bytes.Buffer)
		agentCmd.SetOut(commandOutput)
		agentCmd.SetErr(commandOutput)

		args := []string{"create", "--batch", "--file", batchFilePath}
		flagsCopy := append([]string(nil), flags...)
		missingFlag := flagsCopy[i][2:]
		filteredFlags := append(flagsCopy[:i], flagsCopy[i+2:]...)
		args = append(args, filteredFlags...)

		agentCmd.SetArgs(args)
		err = agentCmd.Execute()
		if err == nil {
			t.Fatalf("Expected error due to missing flag %s, got none", missingFlag)
		}

		if err != nil && err.Error() != fmt.Sprintf("required flag(s) \"%s\" not set", missingFlag) {
			t.Fatalf("error %s", err)
		}
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)

	args := []string{"create", "--batch"}
	flagsCopy := append([]string(nil), flags...)

	args = append(args, flagsCopy...)

	agentCmd.SetArgs(args)
	err = agentCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error due to missing flag --file, got none")
	}

	if err != nil && !strings.Contains(err.Error(), "--file flag is required when using --batch mode") {
		t.Fatalf("error %s", err)
	}
}

func TestAgentSubCmd_Structure_CreateAgentBatchWithMissingFile(t *testing.T) {
	mockService := service.NewAgentServiceMock()

	mockService.CreateAgentBatchFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Batch agent created successfully")
		return nil
	}

	agentCmd := NewAgentCmd(mockService)

	commandOutput := new(bytes.Buffer)
	agentCmd.SetOut(commandOutput)
	agentCmd.SetErr(commandOutput)
	agentCmd.SetArgs([]string{
		"create",
		"--nexus-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--node-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--swarm-id", "b3e1a8f2-4c2d-4e6a-9c1e-2f7e6b8a9d3c",
		"--batch",
		"--file", "/non/existent/path/agents_batch.json",
	})
	err := agentCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error due to missing file, got none")
	}

	if err != nil && !strings.Contains(err.Error(), "file does not exist") {
		t.Fatalf("error %s", err)
	}
}
