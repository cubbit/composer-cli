package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarmNode(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, nodeName, description, nexusID, providerID, configPath string
	var secret *api.GenericIDResponseModel
	var conf *configuration.Config
	var node *api.Node

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeName, err = cmd.Flags().GetString("node-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if providerID, err = cmd.Flags().GetString("provider-id"); err != nil {
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

	if secret, err = api.CreateSwarmSecret(conf.Urls, *accessToken, id, providerID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingSwarmSecret, err)
	}

	nodeBodyRequest := api.CreateNodeBodyRequest{
		Name:        nodeName,
		Description: description,
		NexusID:     nexusID,
		SecretID:    secret.ID,
	}

	if node, err = api.CreateNode(conf.Urls, *accessToken, nodeBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Node %s created successfully\n", node.ID))

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

func DeleteSwarmNode(cmd *cobra.Command, args []string) error {
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

	utils.PrintDelete(fmt.Sprintf("Node %s deleted successfully", nodeID))
	return nil
}

func ListSwarmNodes(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, nexusID, configPath string
	var conf *configuration.Config
	var nodes *api.NodeList

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

	var verbose, l bool

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	for _, node := range nodes.Nodes {
		if verbose {
			fmt.Printf(" • %s, %s, %s, %s\n", node.ID, node.Name, node.Description, node.NexusID)
		} else {
			fmt.Printf(" • %s\n", node.Name)
		}
		if l {
			fmt.Println()
		}
	}

	return nil
}
