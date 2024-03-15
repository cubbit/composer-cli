package action

import (
	"fmt"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nodeName, description, nexusID, choice, providerID string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var providers *api.ProviderList
	var secret *api.GenericIDResponseModel
	var node *api.Node

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var swarms []api.Swarm

		if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
		}

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if choice, err = tui.ChooseOne("Which swarm would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	var choices []string

	for _, nx := range nexuses.Nexuses {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", nx.ID, nx.Name, nx.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which nexus would you like to create your node?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNexusRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	nexusID, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Fill in the form below", false, tui.Input{Placeholder: "Name", IsPassword: false, Value: &nodeName}, tui.Input{Placeholder: "Description", IsPassword: false, Value: &description}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if providers, err = api.ListSwarmProviders(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingProvidersRequest, err)
	}

	if len(providers.Providers) == 0 {
		utils.PrintNotFound("No providers found")
		return nil
	}

	choices = []string{}

	for _, provider := range providers.Providers {
		choices = append(choices, fmt.Sprintf("• %s, %s", provider.ID, provider.Name))
	}

	if choice, err = tui.ChooseOne("Which provider would you like to use?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingProvidersRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	providerID, _, _ = strings.Cut(withoutPrefix, ",")

	if secret, err = api.CreateSwarmSecret(conf.Urls, *accessToken, id, providerID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingSwarmSecret, err)
	}

	nodeBody := api.CreateNodeBodyRequest{
		Name:        nodeName,
		Description: description,
		NexusID:     nexusID,
		SecretID:    secret.ID,
	}

	if node, err = api.CreateNode(conf.Urls, *accessToken, nodeBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Node %s created successfully", node.ID))

	return nil
}

func DescribeSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, nexusID, format string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.NodeList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var swarms []api.Swarm

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No swarms found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your swarm", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	var choices []string

	for _, nx := range nexuses.Nexuses {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", nx.ID, nx.Name, nx.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which nexus would you like to choose?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNexusRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	nexusID, _, _ = strings.Cut(withoutPrefix, ",")

	if nodes, err = api.ListNodes(conf.Urls, *accessToken, id, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nodes.Nodes) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, node := range nodes.Nodes {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", node.ID, node.Name, node.Description))
	}

	if choice, err = tui.ChooseOne("Which node would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ := strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, node := range nodes.Nodes {
		if node.ID == nodeID {
			utils.PrintFormattedData(node, format)
		}
	}

	return nil
}

func EditSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, nexusID, nodeName, description string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.NodeList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var swarms []api.Swarm

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No swarms found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your swarm", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	var choices []string

	for _, nx := range nexuses.Nexuses {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", nx.ID, nx.Name, nx.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which nexus would you like to choose?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNexusRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	nexusID, _, _ = strings.Cut(withoutPrefix, ",")

	if nodes, err = api.ListNodes(conf.Urls, *accessToken, id, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
	}

	if len(nodes.Nodes) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, node := range nodes.Nodes {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", node.ID, node.Name, node.Description))
	}

	if choice, err = tui.ChooseOne("Which node would you like to edit?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ := strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Fill in the form below", true, tui.Input{Placeholder: "Name", IsPassword: false, Value: &nodeName}, tui.Input{Placeholder: "Description", IsPassword: false, Value: &description}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	nodeBody := api.UpdateNodeBodyRequest{
		Name:        nodeName,
		Description: description,
	}

	if err = api.UpdateNode(conf.Urls, *accessToken, nodeID, nodeBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Node %s updated successfully", nodeID))
	return nil
}

func DeleteSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, nexusID string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.NodeList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var swarms []api.Swarm

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No swarms found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your swarm", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	var choices []string

	for _, nx := range nexuses.Nexuses {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", nx.ID, nx.Name, nx.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which nexus would you like to choose?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNexusRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	nexusID, _, _ = strings.Cut(withoutPrefix, ",")

	if nodes, err = api.ListNodes(conf.Urls, *accessToken, id, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nodes.Nodes) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, node := range nodes.Nodes {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", node.ID, node.Name, node.Description))
	}

	if choice, err = tui.ChooseOne("Which node would you like to delete?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ := strings.Cut(withoutPrefix, ",")

	if err = api.DeleteNode(conf.Urls, *accessToken, nodeID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNodeRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("Node %s deleted successfully", nodeID))

	return nil
}

func ListSwarmNodesInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, nexusID string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.NodeList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var swarms []api.Swarm

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No swarms found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your swarm", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	var choices []string

	for _, nx := range nexuses.Nexuses {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", nx.ID, nx.Name, nx.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which nexus would you like to choose?", true, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	nexusID, _, _ = strings.Cut(withoutPrefix, ",")

	if nodes, err = api.ListNodes(conf.Urls, *accessToken, id, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
	}

	utils.PrintList("Your Swarm Nodes List")

	if len(nodes.Nodes) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string
	for _, node := range nodes.Nodes {
		list = append(list, fmt.Sprintf("• %s, %s, %s", node.ID, node.Name, node.Description))
	}

	tui.List(list)

	return nil
}
