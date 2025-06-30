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

func CreateSwarmRedundancyClassInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, redundancyClassName, redundancyClassDescription, configPath string
	var innerK, innerN, outerK, outerN, antiAffinityGroup string
	var innerKInt, innerNInt, outerKInt, outerNInt, antiAffinityGroupInt int
	var conf *configuration.Config
	var redundancyClass *api.RedundancyClass
	var operator *api.Operator
	var choice string
	var choices []string
	var swarms []*api.Swarm
	var nexuses *api.NexusList
	var nexusIDs []string

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
		if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
		}

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, utils.StringOrEmpty(swarm.Description)))
		}

		if choice, err = tui.ChooseOne("Which swarm would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if _, err = tui.TextInputs(
		"Fill in the form below",
		false,
		tui.Input{Placeholder: "Name*", IsPassword: false, Value: &redundancyClassName},
		tui.Input{Placeholder: "Description", IsPassword: false, Value: &redundancyClassDescription},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	type nexusInfo struct {
		ID                  string
		NodeCount           int
		NodesAgentsCountMap map[string]int
	}

	nexusInfoMap := make(map[string]*nexusInfo)

	for _, nx := range nexuses.Nexuses {
		var nodes *api.GenericPaginatedResponse[*api.NewNode]

		if nodes, err = api.ListNodesV4(conf.Urls, *accessToken, id, nx.ID, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
		}

		nodeCount := len(nodes.Data)
		nodesAgentsCountMap := make(map[string]int)

		for _, node := range nodes.Data {
			var agents *api.GenericPaginatedResponse[*api.NewAgent]
			if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nx.ID, node.ID, "", ""); err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
			}

			nodesAgentsCountMap[node.ID] = len(agents.Data)

		}

		info := &nexusInfo{
			ID:                  nx.ID,
			NodeCount:           nodeCount,
			NodesAgentsCountMap: nodesAgentsCountMap,
		}

		nexusInfoMap[nx.ID] = info

	}

	choices = []string{}

	minDisks := 0
	for _, nx := range nexuses.Nexuses {
		for _, count := range nexusInfoMap[nx.ID].NodesAgentsCountMap {
			minDisks = count
			if count < minDisks {
				minDisks = count
			}
		}
	}

	for _, nx := range nexuses.Nexuses {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s (nodes: %d, agents: %d)",
			nx.ID, nx.Name, nx.Description, nexusInfoMap[nx.ID].NodeCount, minDisks))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	if nexusIDs, err = tui.ChooseMany("Which nexuses would you like to choose?", false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNexusRequest, err)
	}

	for i, nexus := range nexusIDs {
		_, withoutPrefix, _ := strings.Cut(nexus, " ")
		nexusID, _, _ := strings.Cut(withoutPrefix, ",")
		nexusIDs[i] = nexusID
	}

	for _, nexusID := range nexusIDs {
		if info, exists := nexusInfoMap[nexusID]; exists {
			if info.NodeCount == 0 {
				utils.PrintWarn(fmt.Sprintf("Nexus %s has no nodes available", nexusID))
				return nil
			}

			totalAgents := 0
			for _, agentCount := range info.NodesAgentsCountMap {
				totalAgents += agentCount
			}

			if totalAgents == 0 {
				utils.PrintWarn(fmt.Sprintf("Nexus %s has no agents available", nexusID))
				return nil
			}
		}
	}

	if _, err = tui.TextInputs(
		"Fill in the redundancy parameters",
		true,
		tui.Input{Placeholder: "Inner K*", IsPassword: false, Value: &innerK},
		tui.Input{Placeholder: "Inner N*", IsPassword: false, Value: &innerN},
		tui.Input{Placeholder: "Outer K*", IsPassword: false, Value: &outerK},
		tui.Input{Placeholder: "Outer N*", IsPassword: false, Value: &outerN},
		tui.Input{Placeholder: "Anti Affinity Group*", IsPassword: false, Value: &antiAffinityGroup}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if innerKInt, err = strconv.Atoi(innerK); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if innerNInt, err = strconv.Atoi(innerN); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if outerKInt, err = strconv.Atoi(outerK); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if outerNInt, err = strconv.Atoi(outerN); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if antiAffinityGroupInt, err = strconv.Atoi(antiAffinityGroup); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if len(nexusIDs) < outerKInt+outerNInt {
		utils.PrintWarn(fmt.Sprintf("not enough nexuses selected to create redundancy class: need at least %d, selected %d", outerKInt+outerNInt, len(nexusIDs)))
		return nil
	}

	for _, nexusID := range nexusIDs {
		totalNodes := 0
		if info, exists := nexusInfoMap[nexusID]; exists {
			totalNodes += info.NodeCount

			totalAgents := 0
			for nodeID, agentCount := range info.NodesAgentsCountMap {
				totalAgents += agentCount

				if agentCount < innerKInt+innerNInt {
					utils.PrintWarn(fmt.Sprintf("Node %s in nexus %s has insufficient agents: %d (need at least %d)",
						nodeID, nexusID, agentCount, innerKInt+innerNInt))
				}

				if agentCount < antiAffinityGroupInt {
					utils.PrintWarn("The anti-affinity parameter is greater than the number of agents available per single machine")
				}
			}
		}
	}

	if outerKInt == 0 {
		utils.PrintWarn("The configuration chosen will result in the use of a single availability zone. This configuration is allowed but strongly discouraged")
		return nil
	}

	if innerKInt == 0 {
		utils.PrintWarn("The current configuration does not add redundancy. Any offline agent will compromise access to the data distributed on it")
		return nil
	}

	bodyRequest := api.CreateRedundancyClassRequestBody{
		Name:              redundancyClassName,
		Description:       redundancyClassDescription,
		InnerK:            innerKInt,
		InnerN:            innerNInt,
		OuterK:            outerKInt,
		OuterN:            outerNInt,
		AntiAffinityGroup: antiAffinityGroupInt,
		Nexuses:           nexusIDs,
	}

	if redundancyClass, err = api.CreateRedundancyClassV4(conf.Urls, *accessToken, id, bodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingRedundancyClassRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("redundancy class %s created successfully", redundancyClass.ID))
	return nil
}

func ListSwarmRedundancyClassesInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath string
	var conf *configuration.Config
	var RedundancyClassList *api.GenericPaginatedResponse[*api.RedundancyClass]
	var operator *api.Operator
	var choice string
	var choices []string
	var swarms []*api.Swarm

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
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	utils.PrintList("Your Redundancy Classes")

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string

	for _, rc := range RedundancyClassList.Data {
		list = append(list, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	tui.List(list)

	return nil
}

func DescribeSwarmRedundancyClassInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath, format string
	var conf *configuration.Config
	var RedundancyClassList *api.GenericPaginatedResponse[*api.RedundancyClass]
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var swarms []*api.Swarm

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
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintNotFound("redundancy classes not found")
		return nil
	}

	choices = []string{}
	for _, rc := range RedundancyClassList.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	if choice, err = tui.ChooseOne("Which redundancy class would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, rc := range RedundancyClassList.Data {
		if rc.ID == redundancyClassID {
			utils.PrintFormattedData(rc, format)
			return nil
		}
	}

	return nil
}

func CheckSwarmRedundancyClassStatusInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath, format string
	var conf *configuration.Config
	var RedundancyClassList *api.GenericPaginatedResponse[*api.RedundancyClass]
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var swarms []*api.Swarm
	var redundancyClassStatus *api.SummaryDetailsWithStatusNullable

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
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintNotFound("redundancy classes not found")
		return nil
	}

	choices = []string{}
	for _, rc := range RedundancyClassList.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	if choice, err = tui.ChooseOne("Which redundancy class would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if redundancyClassStatus, err = api.CheckRedundancyClassStatus(conf.Urls, *accessToken, id, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCheckingRedundancyClassStatusRequest, err)
	}

	utils.PrintFormattedData(redundancyClassStatus, format)
	return nil
}

func CheckSwarmRedundancyClassRecoveryStatusInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath, format string
	var conf *configuration.Config
	var RedundancyClassList *api.GenericPaginatedResponse[*api.RedundancyClass]
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var swarms []*api.Swarm
	var redundancyClassStatus *api.RedundancyClassRecoveryStatus

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
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintNotFound("redundancy classes not found")
		return nil
	}

	choices = []string{}
	for _, rc := range RedundancyClassList.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	if choice, err = tui.ChooseOne("Which redundancy class would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if redundancyClassStatus, err = api.CheckRedundancyClassRecoveryStatus(conf.Urls, *accessToken, id, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCheckingRedundancyClassRecoveryStatusRequest, err)
	}

	utils.PrintFormattedData(redundancyClassStatus, format)
	return nil
}

func ExpandSwarmRedundancyClassInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath, format, dryRunStr string
	var conf *configuration.Config
	var RedundancyClassList *api.GenericPaginatedResponse[*api.RedundancyClass]
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var swarms []*api.Swarm
	var redundancyClassExpanded *api.RedundancyClassExpanded
	var dryRun bool

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if dryRunStr, err = tui.ChooseOne("Would you like to run this command in dry run mode?", false, false, []string{"true", "false"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if dryRun, err = strconv.ParseBool(dryRunStr); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
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
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintNotFound("redundancy classes not found")
		return nil
	}

	choices = []string{}
	for _, rc := range RedundancyClassList.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	if choice, err = tui.ChooseOne("Which redundancy class would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if redundancyClassExpanded, err = api.ExpandRedundancyClass(conf.Urls, *accessToken, id, redundancyClassID, dryRun); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorExpandingRedundancyClassRequest, err)
	}

	utils.PrintFormattedData(redundancyClassExpanded, format)
	return nil
}

func RecoverSwarmRedundancyClassInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath, format, dryRunStr string
	var conf *configuration.Config
	var RedundancyClassList *api.GenericPaginatedResponse[*api.RedundancyClass]
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var swarms []*api.Swarm
	var redundancyClassRecovery *api.RedundancyClassRecovery
	var dryRun bool

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if dryRunStr, err = tui.ChooseOne("Would you like to run this command in dry run mode?", false, false, []string{"true", "false"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if dryRun, err = strconv.ParseBool(dryRunStr); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
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
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintNotFound("redundancy classes not found")
		return nil
	}

	choices = []string{}
	for _, rc := range RedundancyClassList.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	if choice, err = tui.ChooseOne("Which redundancy class would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if redundancyClassRecovery, err = api.RecoverRedundancyClass(conf.Urls, *accessToken, id, redundancyClassID, dryRun); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRecoveringRedundancyClassRequest, err)
	}

	utils.PrintFormattedData(redundancyClassRecovery, format)
	return nil
}
