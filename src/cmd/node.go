// Package cmd provides CLI commands for managing nodes.
package cmd

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Execute commands in node sections",
}

var createNodeSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a new node or a batch of nodes",
	Example: "cubbit node create --swarm-id <swarm-id> --nexus-id <nexus-id> --batch --file ./batch.json\n cubbit node create --swarm-id <swarm-id> --nexus-id <nexus-id> --name <name> --private-ip <private-ip> --public-ip <public-ip> --label <label>",
	PreRun: func(cmd *cobra.Command, args []string) {

		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")

		cmd.MarkFlagsOneRequired("batch", "name")

		cmd.MarkFlagsRequiredTogether("batch", "file")
		cmd.MarkFlagsRequiredTogether("name", "private-ip", "public-ip")
		cmd.MarkFlagsMutuallyExclusive("batch", "name")
		cmd.MarkFlagsMutuallyExclusive("batch", "private-ip")
		cmd.MarkFlagsMutuallyExclusive("batch", "public-ip")
		cmd.MarkFlagsMutuallyExclusive("batch", "label")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		batch, err := cmd.Flags().GetBool("batch")
		if err != nil {
			err = fmt.Errorf("error retrieving batch flag: %w", err)
			utils.PrintError(err)
			return err
		}

		if batch {
			if err = action.CreateNodeBatch(cmd, args); err != nil {
				utils.PrintError(err)
				return err
			}
			return nil
		}

		nodePrivateIP, _ := cmd.Flags().GetString("private-ip")
		if nodePrivateIP != "" && !utils.IsValidIP(nodePrivateIP) {
			err = fmt.Errorf("invalid node private IP address")
			utils.PrintError(err)
			return err
		}

		nodePublicIP, _ := cmd.Flags().GetString("public-ip")
		if nodePublicIP != "" && !utils.IsValidIP(nodePublicIP) {
			err = fmt.Errorf("invalid node public IP address")
			utils.PrintError(err)
			return err
		}

		if err = action.CreateNode(cmd, args); err != nil {
			utils.PrintError(err)
			return err
		}

		return nil
	},
}

var describeNodeSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a node",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("node-id")
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeNode(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var editNodeSubCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a node",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("node-id")
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.EditNode(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var removeNodeSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a node",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("node-id")
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveNode(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listNodesSubCmd = &cobra.Command{
	Use:   "list",
	Short: "list nodes",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		allowedSortingKeys := []string{"id", "name", "label", "created_at", "deleted_at", "nexus_id"}
		sort, _ := cmd.Flags().GetString("sort")
		if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
			return fmt.Errorf("Error: invalid sort key provided, allowed keys are: \"id\", \"name\", \"label\", \"created_at\", \"deleted_at\", \"nexus_id\"")
		}

		filter, _ := cmd.Flags().GetString("filter")
		if filter != "" && !utils.IsValidFilter(filter) {
			return fmt.Errorf("Error: invalid filter provided, allowed format is: key:value key:value ...")
		}

		if err := action.ListNodes(cmd, args); err != nil {
			return err
		}

		return nil
	},
}

var deployNodeSubCmd = &cobra.Command{
	Use:   "deploy",
	Short: "generate deploy files for a node",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("swarm-id")
		cmd.MarkFlagRequired("nexus-id")
		cmd.MarkFlagRequired("node-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.GenerateNodeDeployFiles(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	nodeCmd.AddCommand(createNodeSubCmd)
	createNodeSubCmd.Flags().String("name", "", "Name of the node")
	createNodeSubCmd.Flags().String("private-ip", "", "Private IP of the node")
	createNodeSubCmd.Flags().String("public-ip", "", "Public IP of the node")
	createNodeSubCmd.Flags().String("label", "", "Label of the node")
	createNodeSubCmd.Flags().Bool("batch", false, "Create multiple nodes from a batch file")
	createNodeSubCmd.Flags().String("file", "", "Path to the JSON file containing node definitions")

	nodeCmd.AddCommand(describeNodeSubCmd)
	describeNodeSubCmd.Flags().String("node-id", "", "ID of the node")

	nodeCmd.AddCommand(editNodeSubCmd)
	editNodeSubCmd.Flags().String("node-id", "", "ID of the node")
	editNodeSubCmd.Flags().String("name", "", "Name of the node")
	editNodeSubCmd.Flags().String("private-ip", "", "Private IP of the node")
	editNodeSubCmd.Flags().String("public-ip", "", "Public IP of the node")
	editNodeSubCmd.Flags().String("label", "", "Label of the node")

	nodeCmd.AddCommand(removeNodeSubCmd)
	removeNodeSubCmd.Flags().String("node-id", "", "ID of the node")

	nodeCmd.AddCommand(listNodesSubCmd)
	listNodesSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listNodesSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")

	nodeCmd.AddCommand(deployNodeSubCmd)
	deployNodeSubCmd.Flags().String("node-id", "", "ID of the node")
	deployNodeSubCmd.Flags().String("output-dir", ".", "Directory to save the deployment files (default: current directory)")

	rootCmd.AddCommand(nodeCmd)
	nodeCmd.PersistentFlags().String("swarm-id", "", "ID of the swarm")
	nodeCmd.PersistentFlags().String("nexus-id", "", "ID of the nexus")
}
