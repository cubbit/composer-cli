package cmd_swarm

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestSwarmSubCmd_Structure_Create_WithAllFlags(t *testing.T) {
	mockService := service.NewSwarmServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm created successfully with all flags")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"create",
		"--name", "test-swarm",
		"--description", "Test swarm description",
		"--nexus", "test-cluster:test-node",
		"--redundancy-class", `{"name":"test-rc","outer_n":1,"outer_k":0,"inner_n":1,"inner_k":0,"cluster_ids":["test-cluster"]}`,
	})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Swarm created successfully with all flags\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestSwarmSubCmd_Structure_Create_WithMultipleNexusFlags(t *testing.T) {
	mockService := service.NewSwarmServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm created successfully with multiple nexus flags")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"create",
		"--name", "test-swarm",
		"--nexus", "test-cluster-1:test-node-1",
		"--nexus", "test-cluster-2:test-node-2",
		"--redundancy-class", `{"name":"test-rc","outer_n":1,"outer_k":0,"inner_n":1,"inner_k":0,"cluster_ids":["test-cluster-1","test-cluster-2"]}`,
	})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Swarm created successfully with multiple nexus flags\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestSwarmSubCmd_Structure_Create_WithMultipleRedundancyClassFlags(t *testing.T) {
	mockService := service.NewSwarmServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm created successfully with multiple RC flags")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"create",
		"--name", "test-swarm",
		"--nexus", "test-cluster:test-node",
		"--redundancy-class", `{"name":"rc-1","outer_n":1,"outer_k":0,"inner_n":1,"inner_k":0,"cluster_ids":["test-cluster"]}`,
		"--redundancy-class", `{"name":"rc-2","outer_n":2,"outer_k":1,"inner_n":2,"inner_k":1,"cluster_ids":["test-cluster"]}`,
	})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Swarm created successfully with multiple RC flags\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestSwarmSubCmd_Structure_Create_WithOptionalDescription(t *testing.T) {
	mockService := service.NewSwarmServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm created successfully with description")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"create",
		"--name", "test-swarm",
		"--description", "This is a test swarm for unit testing",
		"--nexus", "test-cluster:test-node",
		"--redundancy-class", `{"name":"test-rc","outer_n":1,"outer_k":0,"inner_n":1,"inner_k":0,"cluster_ids":["test-cluster"]}`,
	})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Swarm created successfully with description\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestSwarmSubCmd_Structure_Create_MissingName(t *testing.T) {
	mockService := service.NewSwarmServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"create",
		"--nexus", "test-cluster:test-node",
	})

	err := swarmCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --name is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"name\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}

func TestSwarmSubCmd_Structure_Create_MissingNexus(t *testing.T) {
	mockService := service.NewSwarmServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"create",
		"--name", "test-swarm",
	})

	err := swarmCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --nexus is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"nexus\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}
