package cmd_gateway

import (
	"bytes"
	"strings"
	"testing"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

func TestGatewayRoutingRulesCmd_Structure_RegistersSubcommands(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	routingRulesCmd, _, err := gatewayCmd.Find([]string{"routing-rules"})
	if err != nil {
		t.Fatalf("Expected routing-rules command, got %v", err)
	}

	for _, name := range []string{"list", "get", "create", "update", "delete"} {
		subCmd, _, err := routingRulesCmd.Find([]string{name})
		if err != nil {
			t.Fatalf("Expected routing-rules %s command, got %v", name, err)
		}
		if subCmd.Name() != name {
			t.Fatalf("Expected command %q, got %q", name, subCmd.Name())
		}
	}

	for _, removed := range []string{"set", "export", "import"} {
		for _, subCmd := range routingRulesCmd.Commands() {
			if subCmd.Name() == removed {
				t.Fatalf("Expected no routing-rules %s command", removed)
			}
		}
	}
}

func TestGatewayRoutingRulesCmd_Structure_GroupAliasAndHelp(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	aliasedCmd, _, err := gatewayCmd.Find([]string{"rr"})
	if err != nil {
		t.Fatalf("Expected rr alias, got %v", err)
	}
	if aliasedCmd.Name() != "routing-rules" {
		t.Fatalf("Expected rr alias to resolve to routing-rules, got %q", aliasedCmd.Name())
	}

	routingRulesCmd, _, err := gatewayCmd.Find([]string{"routing-rules"})
	if err != nil {
		t.Fatalf("Expected routing-rules command, got %v", err)
	}

	for _, expected := range []string{
		"zero or more",
		"evaluated in order",
		"list",
		"get --output json",
		"update --file-path",
	} {
		if !strings.Contains(routingRulesCmd.Long, expected) {
			t.Fatalf("Expected group help to document %q, got %q", expected, routingRulesCmd.Long)
		}
	}
	for _, removed := range []string{"single SDP policy", "needs --id when more than one", "export", "import"} {
		if strings.Contains(routingRulesCmd.Long, removed) {
			t.Fatalf("Expected group help not to contain %q, got %q", removed, routingRulesCmd.Long)
		}
	}
}

func TestGatewayRoutingRulesCmd_Structure_SubcommandAliasesAndFlags(t *testing.T) {
	gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())

	tests := []struct {
		name        string
		alias       string
		required    []string
		notRequired []string
	}{
		{
			name:     "list",
			alias:    "ls",
			required: []string{"gateway-id"},
		},
		{
			name:     "get",
			required: []string{"gateway-id", "id"},
		},
		{
			name:     "create",
			required: []string{"gateway-id", "file-path"},
		},
		{
			name:     "update",
			required: []string{"gateway-id", "id", "file-path"},
		},
		{
			name:     "delete",
			alias:    "rm",
			required: []string{"gateway-id", "id"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subCmd, _, err := gatewayCmd.Find([]string{"routing-rules", test.name})
			if err != nil {
				t.Fatalf("Expected routing-rules %s command, got %v", test.name, err)
			}

			if test.alias != "" {
				aliasedCmd, _, err := gatewayCmd.Find([]string{"routing-rules", test.alias})
				if err != nil {
					t.Fatalf("Expected alias %s, got %v", test.alias, err)
				}
				if aliasedCmd.Name() != test.name {
					t.Fatalf("Expected alias %s to resolve to %s, got %q", test.alias, test.name, aliasedCmd.Name())
				}
			}

			for _, flag := range test.required {
				if subCmd.Flags().Lookup(flag) == nil {
					t.Fatalf("Expected %s to have --%s flag", test.name, flag)
				}
			}

			if subCmd.Flags().Lookup("file") != nil {
				t.Fatalf("Expected %s to use --file-path only, found --file", test.name)
			}
			if test.name != "create" && test.name != "update" && subCmd.Flags().Lookup("file-path") != nil {
				t.Fatalf("Expected %s not to have --file-path", test.name)
			}
		})
	}
}

func TestGatewayRoutingRulesCmd_Structure_RequiredFlagsEnforced(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "list gateway-id",
			args:     []string{"routing-rules", "list"},
			expected: `required flag(s) "gateway-id" not set`,
		},
		{
			name:     "get id",
			args:     []string{"routing-rules", "get", "--gateway-id", testRoutingRuleGatewayID},
			expected: `required flag(s) "id" not set`,
		},
		{
			name:     "create file-path",
			args:     []string{"routing-rules", "create", "--gateway-id", testRoutingRuleGatewayID},
			expected: `required flag(s) "file-path" not set`,
		},
		{
			name:     "update file-path",
			args:     []string{"routing-rules", "update", "--gateway-id", testRoutingRuleGatewayID, "--id", testRoutingRuleID},
			expected: `required flag(s) "file-path" not set`,
		},
		{
			name:     "delete id",
			args:     []string{"routing-rules", "delete", "--gateway-id", testRoutingRuleGatewayID},
			expected: `required flag(s) "id" not set`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gatewayCmd := NewGatewayCmd(servicegateway.NewGatewayServiceMock())
			gatewayCmd.SetOut(new(bytes.Buffer))
			gatewayCmd.SetErr(new(bytes.Buffer))
			gatewayCmd.SetArgs(test.args)

			err := gatewayCmd.Execute()
			if err == nil {
				t.Fatal("Expected required flag error, got nil")
			}
			if !strings.Contains(err.Error(), test.expected) {
				t.Fatalf("Expected error %q, got %q", test.expected, err.Error())
			}
		})
	}
}
