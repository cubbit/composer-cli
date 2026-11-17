package cmd_gateway

import (
	"bytes"

	"github.com/cubbit/composer-cli/src/api"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
)

const (
	testConnectionSwarmID = "11111111-1111-4111-8111-111111111111"
	testConnectionRCID    = "33333333-3333-4333-8333-333333333333"
	testConnectionRCID2   = "44444444-4444-4444-8444-444444444444"
	testConnectionID      = "66666666-6666-4666-8666-666666666666"
)

func setupGatewayConnectionIntegrationCommand(gatewayService servicegateway.GatewayServiceInterface) (*bytes.Buffer, func([]string) error) {
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

func newGatewayConnectionService(gatewayAPI api.GatewayAPIInterface) servicegateway.GatewayService {
	return servicegateway.NewGatewayService(
		newTestGatewayConfig(),
		gatewayAPI,
		&api.MockSwarmAPI{},
		&api.MockRedundancyClassAPI{},
		&api.MockProcessAPI{},
		&api.MockLocationAPI{},
	)
}
