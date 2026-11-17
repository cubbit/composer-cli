package cmd_gateway

import (
	"bytes"
	"strings"

	"github.com/cubbit/composer-cli/src/api"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

// containsAll reports whether s contains every expected substring.
func containsAll(s string, expected ...string) bool {
	for _, e := range expected {
		if !strings.Contains(s, e) {
			return false
		}
	}

	return true
}

func setupGatewayStorageClassIntegrationCommand(gatewayService servicegateway.GatewayServiceInterface) (*bytes.Buffer, func([]string) error) {
	gatewayCmd := NewGatewayCmd(gatewayService)
	gatewayCmd.PersistentFlags().String("profile", "", "Profile")
	gatewayCmd.PersistentFlags().String("output", "human", "Output format")
	gatewayCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	gatewayCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	gatewayCmd.SetOut(commandOutput)
	gatewayCmd.SetErr(commandOutput)

	return commandOutput, func(args []string) error {
		gatewayCmd.SetArgs(args)
		return gatewayCmd.Execute()
	}
}

func newGatewayStorageClassService(gatewayAPI api.GatewayAPIInterface) servicegateway.GatewayService {
	return servicegateway.NewGatewayService(
		newTestGatewayConfig(),
		gatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)
}
