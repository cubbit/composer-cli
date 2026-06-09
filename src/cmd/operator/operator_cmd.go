package cmd_operator

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewOperatorCmd(
	operatorService service.OperatorServiceInterface,
) *cobra.Command {
	var authCmd = &cobra.Command{
		Use:   "operator",
		Short: "Execute commands in k8s operator sections",
	}

	operatorConnectSubCmd := NewOperatorSubCmdConnect(operatorService)
	authCmd.AddCommand(operatorConnectSubCmd)

	return authCmd
}
