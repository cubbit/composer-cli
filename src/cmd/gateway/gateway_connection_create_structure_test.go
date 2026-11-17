package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewayConnectionCreateCmd_Structure_FlagWiring(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.CreateConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			t.Fatalf("Expected no positional args, got %v", args)
		}

		gatewayID, _ := cmd.Flags().GetString("gateway-id")
		if gatewayID != "gateway-123" {
			t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
		}

		connections, _ := cmd.Flags().GetStringArray("connection")
		expected := []string{
			testConnectionSwarmID + ":" + testConnectionRCID + "," + testConnectionRCID2,
		}
		if strings.Join(connections, "|") != strings.Join(expected, "|") {
			t.Fatalf("Expected connections %v, got %v", expected, connections)
		}

		cmd.Println("Mock: Connections created successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"connection", "create",
		"--gateway-id", "gateway-123",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID + "," + testConnectionRCID2,
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Connections created successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewayConnectionCreateCmd_Structure_RepeatedFlags(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.CreateConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		connections, _ := cmd.Flags().GetStringArray("connection")
		if len(connections) != 2 {
			t.Fatalf("Expected two --connection flags, got %v", connections)
		}
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{
		"connection", "create", "--gateway-id", "gateway-123",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID,
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID2,
	})

	if err := gatewayCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGatewayConnectionCreateCmd_Structure_RequiresGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "create"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --gateway-id flag")
	}
	if !strings.Contains(err.Error(), `"gateway-id"`) {
		t.Fatalf("Expected required gateway-id flag error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateConnections should not be called without --gateway-id")
	}
}

func TestGatewayConnectionCreateCmd_Structure_RequiresConnection(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{"connection", "create", "--gateway-id", "gateway-123"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for missing --connection flag")
	}
	if !strings.Contains(err.Error(), `required flag(s) "connection" not set`) {
		t.Fatalf("Expected required connection flag error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateConnections should not be called without --connection")
	}
}

func TestGatewayConnectionCreateCmd_Structure_RejectsPositionalGatewayID(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	createCalled := false
	mockService.CreateConnectionsFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)
	gatewayCmd.SetOut(new(bytes.Buffer))
	gatewayCmd.SetErr(new(bytes.Buffer))
	gatewayCmd.SetArgs([]string{
		"connection", "create", "gateway-123",
		"--connection", testConnectionSwarmID + ":" + testConnectionRCID,
	})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for positional gateway ID argument")
	}
	if createCalled {
		t.Fatal("CreateConnections should not be called with positional arguments")
	}
}

func TestGatewayConnectionCreateCmd_Structure_Help(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	connectionCmd, _, err := gatewayCmd.Find([]string{"connection", "create"})
	if err != nil {
		t.Fatalf("Expected connection create command, got %v", err)
	}

	if !strings.Contains(connectionCmd.Long, "<swarm-id>:<redundancy-class-id>") ||
		!strings.Contains(connectionCmd.Long, "can be repeated") {
		t.Fatalf("Expected create help to document the multi-pair syntax, got %q", connectionCmd.Long)
	}
	if !strings.Contains(connectionCmd.Long, "--gateway-id <gateway-id>") {
		t.Fatalf("Expected create help to mention the --gateway-id flag, got %q", connectionCmd.Long)
	}
}
