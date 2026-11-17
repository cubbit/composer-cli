package update

import (
	"fmt"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/service/gateway/describe"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

func PrintGatewayUpdate(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, gateway api.GatewayV5GetResponse) error {
	if err := printer.PrintText(cmd, handler, fmt.Sprintf("Gateway %s updated successfully.\n\n", gateway.Name)); err != nil {
		return err
	}

	return describe.PrintGatewayDetails(cmd, handler, gateway)
}
