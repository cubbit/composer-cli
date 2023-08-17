package cmd

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/spf13/cobra"
)

var swarmCmd = &cobra.Command{
	Use:   "swarm",
	Short: "Execute commands in swarm sections",
}

var createSwarmSubCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new swarm",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.CreateSwarm); err != nil {
				fmt.Println(err)
			}
		}
	},
}

var listSwarmSubCmd = &cobra.Command{
	Use:   "list",
	Short: "list swarms",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.ListSwarms); err != nil {
				fmt.Println(err)
			}
		}
	},
}

var describeSwarmSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a swarm",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.DescribeSwarm); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	swarmCmd.AddCommand(createSwarmSubCmd)
	createSwarmSubCmd.Flags().String("name", "", "Name of the swarm")
	createSwarmSubCmd.Flags().String("description", "", "Description of the swarm")
	createSwarmSubCmd.Flags().String("configuration", "", "A Json object containing the swarm configuration")

	swarmCmd.AddCommand(listSwarmSubCmd)
	listSwarmSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for swarms")
	listSwarmSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different swarms")

	swarmCmd.AddCommand(describeSwarmSubCmd)
	describeSwarmSubCmd.Flags().String("id", "", "ID of the swarm")
	describeSwarmSubCmd.Flags().String("name", "", "Name of the swarm")
	describeSwarmSubCmd.Flags().String("format", "default", "Format of the output")

	rootCmd.AddCommand(swarmCmd)
}
