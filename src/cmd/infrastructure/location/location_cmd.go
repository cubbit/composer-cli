package cmd_location

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewLocationCmd(
	locationServiceFn service.LocationServiceInterface,
) *cobra.Command {
	var locationCmd = &cobra.Command{
		Use:   "location",
		Short: "Execute commands in location sections",
	}

	locationListSubCmd := NewLocationSubCmdList(locationServiceFn)
	locationDescribeSubCmd := NewLocationSubCmdDescribe(locationServiceFn)
	locationCmd.AddCommand(locationListSubCmd, locationDescribeSubCmd)

	createVirtualCMD := NewLocationSubCmdCreateVirtual(locationServiceFn)
	locationCmd.AddCommand(createVirtualCMD)

	createVirtualNodeCMD := NewLocationSubCmdCreateVirtualNode(locationServiceFn)
	locationCmd.AddCommand(createVirtualNodeCMD)

	return locationCmd
}
