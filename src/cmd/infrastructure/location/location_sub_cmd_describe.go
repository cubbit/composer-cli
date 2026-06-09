package cmd_location

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewLocationSubCmdDescribe(
	locationService service.LocationServiceInterface,
) *cobra.Command {
	var clusterName string
	var clusterID string

	var locationDescribeCmd = &cobra.Command{
		Use:   "describe",
		Short: "Describe aggregated locations",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if clusterName != "" && clusterID != "" {
				return fmt.Errorf("--cluster-name and --cluster-id are mutually exclusive")
			}

			if clusterName == "" && clusterID == "" {
				return fmt.Errorf("either --cluster-name or --cluster-id must be provided")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := locationService.ListAggregated(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	locationDescribeCmd.Flags().StringVar(&clusterName, "cluster-name", "", "Cluster name to filter and describe in detail")
	locationDescribeCmd.Flags().StringVar(&clusterID, "cluster-id", "", "Cluster ID to filter and describe in detail")

	return locationDescribeCmd
}
