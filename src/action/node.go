package action

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarmNode(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, nodeName, label, nexusID, privateIP, publicIP, configPath, nodeConfigStr string
	var nodeConfig map[string]interface{}
	var conf *configuration.Config
	var node *api.NewNodesResponse

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeName, err = cmd.Flags().GetString("node-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if privateIP, err = cmd.Flags().GetString("node-private-ip"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if publicIP, err = cmd.Flags().GetString("node-public-ip"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if label, err = cmd.Flags().GetString("node-label"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeConfigStr, err = cmd.Flags().GetString("node-config"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeConfigStr != "" {
		if err = json.Unmarshal([]byte(nodeConfigStr), &nodeConfig); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorParsingJsonConfiguration, err)
		}
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	nodesBodyRequest := api.BulkInsertNewNodeRequestBody{
		Nodes: []api.CreateNewNodeRequestBody{
			{
				Name:          nodeName,
				PrivateIP:     privateIP,
				PublicIP:      publicIP,
				Label:         &label,
				Configuration: nodeConfig,
			},
		},
	}

	if node, err = api.CreateNodeV4(conf.Urls, *accessToken, id, nexusID, nodesBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Node %s created successfully\n", node.Nodes[0].ID))

	return nil
}

func CreateSwarmNodeBatch(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, nexusID, configPath, filePath string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filePath, err = cmd.Flags().GetString("file"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading batch file: %w", err)
	}

	var batchNodes api.BulkInsertNewNodeRequestBody
	if err = json.Unmarshal(fileData, &batchNodes); err != nil {
		schemaTemplate := map[string]interface{}{
			"nodes": []map[string]interface{}{
				{
					"name":       "",
					"label":      "",
					"config":     map[string]interface{}{},
					"private_ip": "",
					"public_ip":  "",

					"agents": []map[string]interface{}{
						{
							"port":     0,
							"features": map[string]interface{}{},
							"volume": map[string]interface{}{
								"mount_point": "",
								"disk":        "",
							},
						},
					},
				},
			},
		}

		return fmt.Errorf("error parsing batch file: %w\nExpected file schema:\n%s", err, getStructSchema(schemaTemplate))
	}

	if len(batchNodes.Nodes) == 0 {
		return fmt.Errorf("no nodes found in batch file")
	}

	for i, node := range batchNodes.Nodes {
		if node.Name == "" {
			return fmt.Errorf("node at index %d is missing a name", i)
		}
		if node.PrivateIP == "" || !utils.IsValidIP(node.PrivateIP) {
			return fmt.Errorf("node '%s' has an invalid private IP address", node.Name)
		}
		if node.PublicIP == "" || !utils.IsValidIP(node.PublicIP) {
			return fmt.Errorf("node '%s' has an invalid public IP address", node.Name)
		}
	}

	nodes, err := api.CreateNodeV4(conf.Urls, *accessToken, id, nexusID, batchNodes)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Successfully created %d nodes\n", len(nodes.Nodes)))
	for _, node := range nodes.Nodes {
		fmt.Printf("  • %s  %s\n", node.ID, node.Name)
	}

	return nil
}

func DescribeSwarmNode(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath, format string
	var conf *configuration.Config
	var node *api.Node

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	nodeID := args[0]

	if node, err = api.GetNode(conf.Urls, *accessToken, nodeID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	utils.PrintFormattedData(*node, format)

	return nil
}

func EditSwarmNode(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, description, configPath string
	var conf *configuration.Config

	if name, err = cmd.Flags().GetString("node-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	nodeID := args[0]

	nodeBodyRequest := api.UpdateNodeBodyRequest{
		Name:        name,
		Description: description,
	}

	if err = api.UpdateNode(conf.Urls, *accessToken, nodeID, nodeBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Node %s edited successfully", nodeID))
	return nil
}

func RemoveSwarmNode(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	nodeID := args[0]

	if err = api.DeleteNode(conf.Urls, *accessToken, nodeID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNodeRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("Node %s removed successfully", nodeID))
	return nil
}

func ListSwarmNodes(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, nexusID, configPath string
	var conf *configuration.Config
	var nodes *api.NodeList
	var verbose, l bool

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if nodes, err = api.ListNodes(conf.Urls, *accessToken, id, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
	}

	if len(nodes.Nodes) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if verbose {
		utils.PrintVerbose(nodes.Nodes, l)
		return nil
	}

	for _, node := range nodes.Nodes {

		fmt.Printf(" • %s\n", node.Name)

		if l {
			fmt.Println()
		}
	}

	return nil
}

func getStructSchema(schemaTemplate map[string]interface{}) string {
	bytes, err := json.MarshalIndent(schemaTemplate, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error generating schema: %v", err)
	}

	return string(bytes)
}
