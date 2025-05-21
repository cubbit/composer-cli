package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func DescribeSwarmAgentInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nexusID, nodeID, agentID, format string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
	var agents *api.GenericPaginatedResponse[*api.NewAgent]
	var choice string
	var choices []string
	var swarms []*api.Swarm

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

	choices = []string{}

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

	if nodes, err = api.ListNodesV4(conf.Urls, *accessToken, id, nexusID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nodes.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, node := range nodes.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s", node.ID, node.Name))
	}

	if choice, err = tui.ChooseOne("Which node would you like to choose?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ = strings.Cut(withoutPrefix, ",")

	if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, nodeID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(agents.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, agent := range agents.Data {
		choices = append(choices, fmt.Sprintf("• %s", agent.ID))
	}

	if choice, err = tui.ChooseOne("Which agent would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	agentID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, agent := range agents.Data {
		if agent.ID == agentID {
			utils.PrintFormattedData(agent, format)
		}
	}

	return nil
}

func EditSwarmAgentInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, nexusID, nodeID, agentID, agentDisk, agentMountPoint, agentPort, configPath string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
	var agents *api.GenericPaginatedResponse[*api.NewAgent]
	var choice string
	var choices []string
	var swarms []*api.Swarm

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

	choices = []string{}

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

	if nodes, err = api.ListNodesV4(conf.Urls, *accessToken, id, nexusID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nodes.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, node := range nodes.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s", node.ID, node.Name))
	}

	if choice, err = tui.ChooseOne("Which node would you like to choose ?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ = strings.Cut(withoutPrefix, ",")

	if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, nodeID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(agents.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, agent := range agents.Data {
		choices = append(choices, fmt.Sprintf("• %s", agent.ID))
	}

	if choice, err = tui.ChooseOne("Which agent would you like to edit?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	agentID, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Fill in the form below", true,
		tui.Input{Placeholder: "Port", IsPassword: false, Value: &agentPort},
		tui.Input{Placeholder: "Disk", IsPassword: false, Value: &agentDisk},
		tui.Input{Placeholder: "Mount Point", IsPassword: false, Value: &agentMountPoint},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	var agentBodyRequest api.UpdateNewAgentRequestBody
	if agentPort != "" {
		agentPortInt, err := strconv.Atoi(agentPort)
		if err != nil {
			return fmt.Errorf("invalid port value: %w", err)
		}
		agentBodyRequest.Port = &agentPortInt
	}

	if agentDisk != "" {
		agentBodyRequest.Volume = &api.UpdateAgentVolumeRequest{
			Disk: &agentDisk,
		}
	}

	if agentMountPoint != "" {
		if agentBodyRequest.Volume == nil {
			agentBodyRequest.Volume = &api.UpdateAgentVolumeRequest{}
		}
		agentBodyRequest.Volume.MountPoint = &agentMountPoint
	}

	if err = api.UpdateAgentV4(conf.Urls, *accessToken, id, nexusID, nodeID, agentID, agentBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Agent %s updated successfully", nodeID))
	return nil
}

func RemoveSwarmAgentInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nexusID, nodeID, agentID string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
	var agents *api.GenericPaginatedResponse[*api.NewAgent]
	var choice string
	var choices []string
	var swarms []*api.Swarm

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

	choices = []string{}

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

	if nodes, err = api.ListNodesV4(conf.Urls, *accessToken, id, nexusID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nodes.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, node := range nodes.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s", node.ID, node.Name))
	}

	if choice, err = tui.ChooseOne("Which node would you like to choose?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ = strings.Cut(withoutPrefix, ",")

	if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, nodeID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(agents.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, agent := range agents.Data {
		choices = append(choices, fmt.Sprintf("• %s", agent.ID))
	}

	if choice, err = tui.ChooseOne("Which agent would you like to delete?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	agentID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.DeleteAgentV4(conf.Urls, *accessToken, id, nexusID, nodeID, agentID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNodeRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("Agent %s deleted successfully", agentID))
	return nil
}

func ListSwarmAgentsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nexusID, nodeID string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
	var agents *api.GenericPaginatedResponse[*api.NewAgent]
	var choice string
	var choices []string
	var swarms []*api.Swarm
	var RedundancyClassList *api.GenericPaginatedResponse[*api.RedundancyClass]
	var listType string

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

	listTypeChoices := []string{
		"• List agents by node",
		"• List agents by redundancy class",
	}

	if listType, err = tui.ChooseOne("How would you like to list agents?", false, false, listTypeChoices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if strings.Contains(listType, "by node") {
		if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
		}

		if len(nexuses.Nexuses) == 0 {
			utils.PrintNotFound("No nexuses found")
			return nil
		}

		choices = []string{}
		for _, nx := range nexuses.Nexuses {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", nx.ID, nx.Name, nx.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No nexuses found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which nexus would you like to choose?", true, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		nexusID, _, _ = strings.Cut(withoutPrefix, ",")

		if nodes, err = api.ListNodesV4(conf.Urls, *accessToken, id, nexusID, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
		}

		if len(nodes.Data) == 0 {
			utils.PrintNotFound("No nodes found")
			return nil
		}

		choices = []string{}
		for _, node := range nodes.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s", node.ID, node.Name))
		}

		if choice, err = tui.ChooseOne("Which node would you like to choose?", true, true, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
		}

		_, withoutPrefix, _ = strings.Cut(choice, " ")
		nodeID, _, _ = strings.Cut(withoutPrefix, ",")

		if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, nodeID, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
		}

		utils.PrintList("Your Agents List")
		if len(agents.Data) == 0 {
			utils.PrintEmptyList()
			return nil
		}

		var list []string
		for _, agent := range agents.Data {
			list = append(list, fmt.Sprintf("• %s", agent.ID))
		}

		tui.List(list)
	} else {
		if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
		}

		if len(RedundancyClassList.Data) == 0 {
			utils.PrintNotFound("No redundancy classes found")
			return nil
		}

		choices = []string{}
		for _, rc := range RedundancyClassList.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s", rc.ID, rc.Name))
		}

		if choice, err = tui.ChooseOne("Which redundancy class would you like to choose?", true, true, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		rcID, _, _ := strings.Cut(withoutPrefix, ",")

		if agents, err = api.ListAgentsForRCV4(conf.Urls, *accessToken, id, rcID, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
		}

		utils.PrintList("Your Agents List for Redundancy Class")
		if len(agents.Data) == 0 {
			utils.PrintEmptyList()
			return nil
		}

		var list []string
		for _, agent := range agents.Data {
			list = append(list, fmt.Sprintf("• %s", agent.ID))
		}

		tui.List(list)
	}

	return nil
}
