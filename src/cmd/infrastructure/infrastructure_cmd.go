package cmd_infrastructure

import (
	cmd_location "github.com/cubbit/composer-cli/src/cmd/infrastructure/location"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewInfrastructureCmd(
	locationService service.LocationServiceInterface,
) *cobra.Command {
	var infrastructureCmd = &cobra.Command{
		Use:     "infrastructure",
		Aliases: []string{"infra"},
		Short:   "Execute commands in infrastructure sections",
	}

	locationCmd := cmd_location.NewLocationCmd(locationService)
	infrastructureCmd.AddCommand(locationCmd)

	return infrastructureCmd
}
