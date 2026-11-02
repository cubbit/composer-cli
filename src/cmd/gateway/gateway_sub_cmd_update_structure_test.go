package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func TestGatewaySubCmd_Structure_Update(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.UpdateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Gateway updated successfully")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"update", "gateway-123", "--name", "renamed-gateway"})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateway updated successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_Update_Flags(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	mockService.UpdateFunc = func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || args[0] != "test-gateway" {
			t.Fatalf("Expected positional gateway ID test-gateway, got %v", args)
		}

		name, _ := cmd.Flags().GetString("name")
		if name != "renamed-gateway" {
			t.Fatalf("Expected name renamed-gateway, got %q", name)
		}

		description, _ := cmd.Flags().GetString("description")
		if description != "Updated description" {
			t.Fatalf("Expected description %q, got %q", "Updated description", description)
		}

		location, _ := cmd.Flags().GetString("location")
		if location != "eu-west-1" {
			t.Fatalf("Expected location eu-west-1, got %q", location)
		}

		cmd.Println("Mock: Gateway updated with flags")
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{
		"update", "test-gateway",
		"--name", "renamed-gateway",
		"--description", "Updated description",
		"--location", "eu-west-1",
	})

	err := gatewayCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Gateway updated with flags\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestGatewaySubCmd_Structure_Update_DoesNotExposeGatewayNameFlag(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	updateCmd, _, err := gatewayCmd.Find([]string{"update"})
	if err != nil {
		t.Fatalf("Expected update command, got %v", err)
	}

	if updateCmd.Flags().Lookup("gateway-name") != nil {
		t.Fatal("Expected gateway update command not to expose --gateway-name")
	}
}

func TestGatewaySubCmd_Structure_Update_DoesNotExposeGatewayIDFlag(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	updateCmd, _, err := gatewayCmd.Find([]string{"update"})
	if err != nil {
		t.Fatalf("Expected update command, got %v", err)
	}

	if updateCmd.Flags().Lookup("gateway-id") != nil {
		t.Fatal("Expected gateway update command not to expose --gateway-id")
	}
}

func TestGatewaySubCmd_Structure_Update_ValidatesFlags(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedError string
	}{
		{
			name:          "no update flag",
			args:          []string{"update", "gateway-123"},
			expectedError: "specify at least one of --name, --description or --location",
		},
		{
			name:          "empty name",
			args:          []string{"update", "gateway-123", "--name", ""},
			expectedError: "--name must be between 3 and 63 characters",
		},
		{
			name:          "name too short",
			args:          []string{"update", "gateway-123", "--name", "ab"},
			expectedError: "--name must be between 3 and 63 characters",
		},
		{
			name:          "name too long",
			args:          []string{"update", "gateway-123", "--name", strings.Repeat("a", 64)},
			expectedError: "--name must be between 3 and 63 characters",
		},
		{
			name:          "description too long",
			args:          []string{"update", "gateway-123", "--description", strings.Repeat("a", 121)},
			expectedError: "--description must be at most 120 characters",
		},
		{
			name:          "location too long",
			args:          []string{"update", "gateway-123", "--location", strings.Repeat("a", 241)},
			expectedError: "--location must be at most 240 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateCalled := false
			mockService := servicegateway.NewGatewayServiceMock()
			mockService.UpdateFunc = func(cmd *cobra.Command, args []string) error {
				updateCalled = true
				return nil
			}

			gatewayCmd := NewGatewayCmd(mockService)
			gatewayCmd.SetOut(new(bytes.Buffer))
			gatewayCmd.SetErr(new(bytes.Buffer))
			gatewayCmd.SetArgs(tt.args)

			err := gatewayCmd.Execute()
			if err == nil {
				t.Fatalf("Expected error %q, got nil", tt.expectedError)
			}
			if !strings.Contains(err.Error(), tt.expectedError) {
				t.Fatalf("Expected error %q, got %q", tt.expectedError, err.Error())
			}
			if updateCalled {
				t.Fatal("Update should not be called when update flags are invalid")
			}
		})
	}
}

func TestGatewaySubCmd_Structure_Update_CountsUnicodeCharacters(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedError string
	}{
		{
			name: "three-rune name is accepted",
			args: []string{"update", "gateway-123", "--name", "東京!"},
		},
		{
			name: "forty-rune name is accepted",
			args: []string{"update", "gateway-123", "--name", strings.Repeat("é", 40)},
		},
		{
			name: "120-rune description is accepted",
			args: []string{"update", "gateway-123", "--description", strings.Repeat("é", 120)},
		},
		{
			name: "240-rune location is accepted",
			args: []string{"update", "gateway-123", "--location", strings.Repeat("界", 240)},
		},
		{
			name:          "two-rune name is rejected",
			args:          []string{"update", "gateway-123", "--name", "東京"},
			expectedError: "--name must be between 3 and 63 characters",
		},
		{
			name:          "121-rune description is rejected",
			args:          []string{"update", "gateway-123", "--description", strings.Repeat("é", 121)},
			expectedError: "--description must be at most 120 characters",
		},
		{
			name:          "241-rune location is rejected",
			args:          []string{"update", "gateway-123", "--location", strings.Repeat("界", 241)},
			expectedError: "--location must be at most 240 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateCalled := false
			mockService := servicegateway.NewGatewayServiceMock()
			mockService.UpdateFunc = func(cmd *cobra.Command, args []string) error {
				updateCalled = true
				return nil
			}

			gatewayCmd := NewGatewayCmd(mockService)
			gatewayCmd.SetOut(new(bytes.Buffer))
			gatewayCmd.SetErr(new(bytes.Buffer))
			gatewayCmd.SetArgs(tt.args)

			err := gatewayCmd.Execute()

			if tt.expectedError != "" {
				if err == nil {
					t.Fatalf("Expected error %q, got nil", tt.expectedError)
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Fatalf("Expected error %q, got %q", tt.expectedError, err.Error())
				}
				if updateCalled {
					t.Fatal("Update should not be called when unicode limits are exceeded")
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if !updateCalled {
				t.Fatal("Update should be called when unicode limits are respected")
			}
		})
	}
}

func TestGatewaySubCmd_Structure_Update_TooManyArgs(t *testing.T) {
	mockService := servicegateway.NewGatewayServiceMock()
	updateCalled := false
	mockService.UpdateFunc = func(cmd *cobra.Command, args []string) error {
		updateCalled = true
		return nil
	}

	gatewayCmd := NewGatewayCmd(mockService)

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)
	gatewayCmd.SetArgs([]string{"update", "gateway-001", "gateway-002", "--name", "renamed-gateway"})

	err := gatewayCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for too many positional arguments")
	}
	if updateCalled {
		t.Fatal("Update should not be called with too many positional arguments")
	}
}
