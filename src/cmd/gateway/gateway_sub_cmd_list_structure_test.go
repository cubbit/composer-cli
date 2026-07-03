package cmd_gateway

import (
	"bytes"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewaySubCmd_Structure_List(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateways listed successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"list"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateways listed successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_List_Alias(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateways listed via alias")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"ls"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateways listed via alias\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_List_WithQueryFlag(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		actualQuery, _ := cmd.Flags().GetStringArray("query")
		if len(actualQuery) != 1 || actualQuery[0] != "name:eq(test)" {
			t.Fatalf("Expected query 'name:eq(test)', got %v", actualQuery)
		}
		cmd.Println("Mock: Gateways listed with query")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"list", "--query", "name:eq(test)"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateways listed with query\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_List_WithMultipleQueryFlags(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		actualQuery, _ := cmd.Flags().GetStringArray("query")
		if len(actualQuery) != 2 {
			t.Fatalf("Expected 2 query values, got %d: %v", len(actualQuery), actualQuery)
		}
		if actualQuery[0] != "name:eq(test)" {
			t.Fatalf("Expected first query 'name:eq(test)', got %q", actualQuery[0])
		}
		if actualQuery[1] != "type:eq(manual)" {
			t.Fatalf("Expected second query 'type:eq(manual)', got %q", actualQuery[1])
		}
		cmd.Println("Mock: Gateways listed with multiple queries")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"list", "--query", "name:eq(test)", "--query", "type:eq(manual)"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateways listed with multiple queries\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_List_WithPageAndItems(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		page, _ := cmd.Flags().GetInt("page")
		items, _ := cmd.Flags().GetInt("items")

		if page != 2 {
			t.Fatalf("Expected page 2, got %d", page)
		}
		if items != 50 {
			t.Fatalf("Expected items 50, got %d", items)
		}

		cmd.Println("Mock: Gateways listed with pagination")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"list", "--page", "2", "--items", "50"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateways listed with pagination\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
