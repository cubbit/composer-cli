package cmd_location

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewLocationCmd(
	locationService service.LocationServiceInterface,
) *cobra.Command {
	var locationCmd = &cobra.Command{
		Use:   "location",
		Short: "Execute commands in location sections",
	}

	locationListSubCmd := NewLocationSubCmdList(locationService)
	locationDescribeSubCmd := NewLocationSubCmdDescribe(locationService)
	locationCmd.AddCommand(locationListSubCmd, locationDescribeSubCmd)

	return locationCmd
}
