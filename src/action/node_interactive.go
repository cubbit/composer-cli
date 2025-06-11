package action

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func DescribeSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nexusID, format string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
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

	if choice, err = tui.ChooseOne("Which node would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ := strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, node := range nodes.Data {
		if node.ID == nodeID {
			utils.PrintFormattedData(node, format)
		}
	}

	return nil
}

func EditSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nexusID, nodeName, label, publicIP, privateIP string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
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
		return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
	}

	if len(nodes.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	choices = []string{}

	for _, node := range nodes.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s", node.ID, node.Name))
	}

	if choice, err = tui.ChooseOne("Which node would you like to edit?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ := strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Fill in the form below", true,
		tui.Input{Placeholder: "Name", IsPassword: false, Value: &nodeName},
		tui.Input{Placeholder: "Label", IsPassword: false, Value: &label},
		tui.Input{Placeholder: "Public IP", IsPassword: false, Value: &publicIP},
		tui.Input{Placeholder: "Private IP", IsPassword: false, Value: &privateIP},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	var nodeBody api.UpdateNewNodeRequestBody
	if nodeName != "" {
		nodeBody.Name = &nodeName
	}

	if label != "" {
		nodeBody.Label = &label
	}

	if publicIP != "" {
		nodeBody.PublicIP = &publicIP
	}

	if privateIP != "" {
		nodeBody.PrivateIP = &privateIP
	}

	if err = api.UpdateNodeV4(conf.Urls, *accessToken, id, nexusID, nodeID, nodeBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Node %s updated successfully", nodeID))
	return nil
}

func RemoveSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nexusID, email, password, code, deleteSwarmNodeToken string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
	var choice string
	var choices []string
	var swarms []*api.Swarm
	var challenge *api.ChallengeResponseModel

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

	if choice, err = tui.ChooseOne("Which node would you like to remove?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ := strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs(fmt.Sprintf("Confirm your login to remove node %s 🚮", utils.RedBg.Render(nodeID)), true, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if deleteSwarmNodeToken, err = api.ForgeOperatorSwarmNodeToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, nodeID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteTokenRequest, err)
	}

	if err = api.DeleteNodeV4(conf.Urls, *accessToken, id, nexusID, nodeID, deleteSwarmNodeToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNodeRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("Node %s removed successfully", nodeID))

	return nil
}

func ListSwarmNodesInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, nexusID string
	var conf *configuration.Config
	var operator *api.Operator
	var nexuses *api.NexusList
	var nodes *api.GenericPaginatedResponse[*api.NewNode]
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

	if choice, err = tui.ChooseOne("Which nexus would you like to choose?", true, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	nexusID, _, _ = strings.Cut(withoutPrefix, ",")

	if nodes, err = api.ListNodesV4(conf.Urls, *accessToken, id, nexusID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
	}

	utils.PrintList("Your Swarm Nodes List")

	if len(nodes.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string
	for _, node := range nodes.Data {
		list = append(list, fmt.Sprintf("• %s, %s", node.ID, node.Name))
	}

	tui.List(list)

	return nil
}

func CreateSwarmNodeInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, nexusID, configPath string
	var conf *configuration.Config
	var operator *api.Operator
	var swarms []*api.Swarm
	var nexuses *api.NexusList
	var choice string
	var choices []string
	var createdNodes *api.NewNodesResponse

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

	nodeConfigs, err := collectNodeConfiguration()
	if err != nil {
		return fmt.Errorf("error collecting node configuration: %w", err)
	}

	if err = collectAgentConfiguration(&nodeConfigs); err != nil {
		return fmt.Errorf("error collecting agent configuration: %w", err)
	}

	if createdNodes, err = createNodeWithAgents(conf, *accessToken, id, nexusID, nodeConfigs); err != nil {
		return fmt.Errorf("error creating nodes with agents: %w %v", err, nodeConfigs)
	}

	utils.PrintSuccess(fmt.Sprintf("Nodes and agents created successfully in swarm %s", id))

	var generateFilesChoice string
	if generateFilesChoice, err = tui.ChooseOne("Would you like to generate Deployment files (Ansible, YAML)?", false, true, []string{"Yaml", "Ansible", "Both", "Skip"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	generateAnsible := generateFilesChoice == "Ansible"
	generateYAML := generateFilesChoice == "Yaml"
	generateBoth := generateFilesChoice == "Both"

	var nodeConfigsCreated []api.NodeConfig
	for _, node := range createdNodes.Nodes {
		nodeConfig := api.NodeConfig{
			ID:        node.ID,
			Name:      node.Name,
			PublicIP:  node.PublicIP,
			PrivateIP: node.PrivateIP,
			Agents:    make([]api.AgentConfig, len(node.Agents)),
		}

		for i, agent := range node.Agents {
			nodeConfig.Agents[i] = api.AgentConfig{
				ID:         agent.ID,
				MountPoint: agent.Volume.MountPoint,
				Disk:       agent.Volume.Disk,
				Port:       agent.Port,
				Secret:     agent.Secret,
			}
		}
		nodeConfigsCreated = append(nodeConfigsCreated, nodeConfig)
	}

	basePath := "./cubbit-agent-playbook.tar"
	outputPath := "./cubbit-agent-yaml.tar"

	if generateAnsible {
		if err = downloadAndGenerateAnsibleTar(constants.AnsibleTarUrl, nodeConfigsCreated, basePath); err != nil {
			return fmt.Errorf("error downloading and generating Ansible tar: %w", err)
		}

		utils.PrintSuccess("Ansible files generated successfully")
	}

	if generateYAML {
		if err := generateYAMLFiles(conf, nodeConfigsCreated, outputPath); err != nil {
			return fmt.Errorf("error generating YAML file: %w", err)
		}

		utils.PrintSuccess("YAML file generated successfully")
	}

	if generateBoth {
		if err = downloadAndGenerateAnsibleTar(constants.AnsibleTarUrl, nodeConfigsCreated, basePath); err != nil {
			return fmt.Errorf("error downloading and generating Ansible tar: %w", err)
		}

		if err = generateYAMLFiles(conf, nodeConfigsCreated, outputPath); err != nil {
			return fmt.Errorf("error generating YAML file: %w", err)
		}

		utils.PrintSuccess("Ansible and YAML files generated successfully")
	}

	return nil
}

func GenerateSwarmNodeDeployFilesInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, nexusID, nodeID, configPath string
	var conf *configuration.Config
	var operator *api.Operator
	var swarms []*api.Swarm
	var nexuses *api.NexusList
	var choice string
	var choices []string

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

	var nodes *api.GenericPaginatedResponse[*api.NewNode]
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

	if choice, err = tui.ChooseOne("Which node would you like to choose?", true, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNodeRequest, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	nodeID, _, _ = strings.Cut(withoutPrefix, ",")

	var generateFilesChoice string
	if generateFilesChoice, err = tui.ChooseOne("Would you like to generate Deployment files (Ansible, YAML)?", false, true, []string{"yaml", "ansible", "both", "skip"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	var nodeConfigs []api.NodeConfig
	if nodeID != "" {
		var node *api.NewNode
		for _, n := range nodes.Data {
			if n.ID == nodeID {
				node = n
				break
			}
		}

		if node == nil {
			utils.PrintNotFound("Node not found")
			return nil
		}

		var agents *api.GenericPaginatedResponse[*api.NewAgent]
		if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, nodeID, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
		}

		if len(agents.Data) == 0 {
			utils.PrintNotFound("No agents found for the selected node")
			return nil
		}

		nodeConfigs = make([]api.NodeConfig, 0, len(agents.Data))
		for _, agent := range agents.Data {
			nodeConfig := api.NodeConfig{
				ID:        node.ID,
				Name:      node.Name,
				PublicIP:  node.PublicIP,
				PrivateIP: node.PrivateIP,
				Agents: []api.AgentConfig{
					{
						ID:         agent.ID,
						MountPoint: agent.Volume.MountPoint,
						Disk:       agent.Volume.Disk,
						Port:       agent.Port,
						Secret:     agent.Secret,
					},
				},
			}
			nodeConfigs = append(nodeConfigs, nodeConfig)
		}
	} else {

		nodeConfigs = make([]api.NodeConfig, 0, len(nodes.Data))
		for _, node := range nodes.Data {
			nodeConfig := api.NodeConfig{
				ID:        node.ID,
				Name:      node.Name,
				PublicIP:  node.PublicIP,
				PrivateIP: node.PrivateIP,
				Agents:    make([]api.AgentConfig, 0),
			}

			var agents *api.GenericPaginatedResponse[*api.NewAgent]
			if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, node.ID, "", ""); err != nil {
				return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
			}

			if len(agents.Data) == 0 {
				utils.PrintNotFound(fmt.Sprintf("No agents found for node %s", node.ID))
				continue
			}

			for _, agent := range agents.Data {
				nodeConfig.Agents = append(nodeConfig.Agents, api.AgentConfig{
					ID:         agent.ID,
					MountPoint: agent.Volume.MountPoint,
					Disk:       agent.Volume.Disk,
					Port:       agent.Port,
					Secret:     agent.Secret,
				})
			}

			nodeConfigs = append(nodeConfigs, nodeConfig)

		}
	}

	if err = generateDeployFiles(conf, generateFilesChoice, nodeConfigs); err != nil {
		return fmt.Errorf("error generating deployment files: %w", err)
	}

	return nil
}

func collectNodeConfiguration() ([]api.NodeConfig, error) {
	var numNodesStr, publicIPs, privateIPs, nodeNames, label string
	var err error

	if _, err = tui.TextInputs(
		"Node Configuration",
		false,
		tui.Input{Placeholder: "Number of nodes", IsPassword: false, Value: &numNodesStr},
		tui.Input{Placeholder: "Public IPs (comma-separated)", IsPassword: false, Value: &publicIPs},
		tui.Input{Placeholder: "Private IPs (optional, comma-separated)", IsPassword: false, Value: &privateIPs},
		tui.Input{Placeholder: "Node names (optional, comma-separated)", IsPassword: false, Value: &nodeNames},
		tui.Input{Placeholder: "Label for all nodes (optional)", IsPassword: false, Value: &label},
	); err != nil {
		return nil, err
	}

	if privateIPs == "" {
		privateIPs = publicIPs
	}

	numNodes, err := strconv.Atoi(numNodesStr)
	if err != nil || numNodes <= 0 {
		return nil, fmt.Errorf("invalid number of nodes: %s", numNodesStr)
	}

	if err = utils.ValidateIPsInput(publicIPs, numNodes); err != nil {
		return nil, fmt.Errorf("error validating public IPs: %w", err)
	}

	if err = utils.ValidateIPsInput(privateIPs, numNodes); err != nil {
		return nil, fmt.Errorf("error validating private IPs: %w", err)
	}

	if err = utils.ValidateNamesInput(nodeNames, numNodes); err != nil {
		return nil, fmt.Errorf("error validating node names: %w", err)
	}

	computedPublicIPs, err := utils.ComputeIPsArray(publicIPs, numNodes)
	if err != nil {
		return nil, fmt.Errorf("error computing public IPs: %w", err)
	}

	computedPrivateIPs, err := utils.ComputeIPsArray(privateIPs, numNodes)
	if err != nil {
		return nil, fmt.Errorf("error computing private IPs: %w", err)
	}

	computedNames := utils.ComputeNamesArray(nodeNames, numNodes)

	nodeConfigs := make([]api.NodeConfig, numNodes)
	for i := 0; i < numNodes; i++ {
		nodeConfigs[i] = api.NodeConfig{
			Name:      computedNames[i],
			PublicIP:  computedPublicIPs[i],
			PrivateIP: computedPrivateIPs[i],
			Label:     strings.TrimSpace(label),
		}
	}

	return nodeConfigs, nil
}

func collectAgentConfiguration(nodeConfigs *[]api.NodeConfig) error {
	var numDisksStr, mountPoints, diskIdentifiers, ports string
	var err error

	if _, err = tui.TextInputs(
		"Agent Configuration",
		false,
		tui.Input{Placeholder: "Number of disks per node", IsPassword: false, Value: &numDisksStr},
		tui.Input{Placeholder: "Disk identifiers (comma-separated)", IsPassword: false, Value: &diskIdentifiers},
		tui.Input{Placeholder: "Ports (comma-separated)", IsPassword: false, Value: &ports},
		tui.Input{Placeholder: "Mountpoints (optional, comma-separated)", IsPassword: false, Value: &mountPoints},
	); err != nil {
		return err
	}

	numDisks, err := strconv.Atoi(numDisksStr)
	if err != nil || numDisks <= 0 {
		return fmt.Errorf("invalid number of disks: %s", numDisksStr)
	}

	computedMountPoints, err := utils.ComputeMountPointsArray(mountPoints, numDisks)
	if err != nil {
		return fmt.Errorf("error computing mount points: %w", err)
	}

	computedDisks, err := utils.ComputeDisksArray(diskIdentifiers, numDisks)
	if err != nil {
		return fmt.Errorf("error computing disk identifiers: %w", err)
	}

	computedPorts, err := utils.ComputePortsArray(ports, numDisks)
	if err != nil {
		return fmt.Errorf("error computing ports: %w", err)
	}

	for i := range *nodeConfigs {
		(*nodeConfigs)[i].Agents = make([]api.AgentConfig, numDisks)
		for j := 0; j < numDisks; j++ {
			(*nodeConfigs)[i].Agents[j] = api.AgentConfig{
				MountPoint: computedMountPoints[j],
				Disk:       computedDisks[j],
				Port:       computedPorts[j],
			}
		}
	}

	return nil
}

func createNodeWithAgents(conf *configuration.Config, accessToken, swarmID string, nexusID string, nodeConfigs []api.NodeConfig) (*api.NewNodesResponse, error) {
	var err error
	req := api.BulkInsertNewNodeRequestBody{
		Nodes: make([]api.CreateNewNodeRequestBody, len(nodeConfigs)),
	}
	for i, nodeConfig := range nodeConfigs {
		nodeBody := api.CreateNewNodeRequestBody{
			Name:      nodeConfig.Name,
			Label:     &nodeConfig.Label,
			PublicIP:  nodeConfig.PublicIP,
			PrivateIP: nodeConfig.PrivateIP,
		}

		agents := make([]api.CreateNewAgentRequestBody, len(nodeConfig.Agents))
		for j, agentConfig := range nodeConfig.Agents {
			agents[j] = api.CreateNewAgentRequestBody{
				Port: agentConfig.Port,
				Volume: api.AgentVolume{
					MountPoint: agentConfig.MountPoint,
					Disk:       agentConfig.Disk,
				},
			}
		}

		nodeBody.Agents = agents
		req.Nodes[i] = nodeBody
	}

	var createdNodes *api.NewNodesResponse
	if createdNodes, err = api.CreateNodeV4(conf.Urls, accessToken, swarmID, nexusID, req); err != nil {
		return nil, fmt.Errorf("error creating nodes and agents: %w", err)
	}

	return createdNodes, nil
}

func generateInventoryContent(nodeConfigs []api.NodeConfig) string {
	var content strings.Builder
	for _, nodeConfig := range nodeConfigs {
		content.WriteString(fmt.Sprintf("[%s]\n%s\n", nodeConfig.Name, nodeConfig.PublicIP))
	}
	return content.String()
}

func generateHostNamesContent(nodeConfigs []api.NodeConfig) string {
	var content strings.Builder
	for _, nodeConfig := range nodeConfigs {
		content.WriteString(fmt.Sprintf("%s\n", nodeConfig.Name))
	}
	return content.String()
}

func generateAgentSecretsContent(nodeConfig api.NodeConfig) map[string]interface{} {
	agentsMap := make(map[string]interface{})

	for i, agent := range nodeConfig.Agents {
		agentsMap[strconv.Itoa(i)] = map[string]interface{}{
			"agent_secret":         agent.Secret,
			"mount_point":          agent.MountPoint,
			"disk":                 agent.Disk,
			"cccp_server_port":     agent.Port,
			"cccp_server_local_ip": nodeConfig.PrivateIP,
			"machine_id":           nodeConfig.ID,
		}
	}

	return agentsMap
}

func generateYAMLFiles(conf *configuration.Config, nodeConfigs []api.NodeConfig, outputPath string) error {
	envs := api.YAMLGenerationEnvs{
		HiveURL:                  conf.Urls.ChUrl,
		MetricsURL:               conf.Urls.MetricsUrl,
		MetricsRoutesSend:        constants.MetricsSender,
		CCCPSwarmGatewayEndpoint: conf.Urls.SwarmGatewayUrl,
		CCCPSwarmGatewayPort:     "",
		CCCPSwarmGatewaySecure:   "false",
	}

	tmpDir, err := os.MkdirTemp("", "ansible-tar")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	for _, nodeConfig := range nodeConfigs {
		yamlContent, err := generateClusterYAMLFiles(nodeConfig, envs)
		if err != nil {
			return fmt.Errorf("error generating YAML for node %s: %w", nodeConfig.Name, err)
		}

		filename := fmt.Sprintf("cluster-%s.yaml", nodeConfig.Name)
		if err := utils.WriteFile(filepath.Join(tmpDir, filename), []byte(yamlContent)); err != nil {
			return fmt.Errorf("error writing YAML file for node %s: %w", nodeConfig.Name, err)
		}
	}

	if err := utils.CreateTar(outputPath, tmpDir); err != nil {
		return fmt.Errorf("failed to repack tarball: %w", err)
	}

	return nil
}

func generateClusterYAMLFiles(nodeConfig api.NodeConfig, envs api.YAMLGenerationEnvs) (string, error) {
	secretData := make(map[string]string)
	agentsDetail := make(map[string]api.AgentDetail)

	for i, agent := range nodeConfig.Agents {
		agentName := fmt.Sprintf("%s-agent%02d", nodeConfig.Name, i)

		agentSecret := api.AgentSecret{
			AgentSecret: agent.Secret,
			AgentUUID:   agent.ID,
		}

		secretJSON, err := json.Marshal(agentSecret)
		if err != nil {
			return "", fmt.Errorf("error marshaling agent secret: %w", err)
		}
		secretData[agentName] = base64.StdEncoding.EncodeToString(secretJSON)

		agentsDetail[agentName] = api.AgentDetail{
			LocalPath:        agent.MountPoint,
			NodeNameSelector: nodeConfig.PublicIP,
		}
	}

	secretYAML := api.SecretYAML{
		APIVersion: "v1",
		Kind:       "Secret",
		Metadata: api.SecretMetadata{
			Name: fmt.Sprintf("%s-secret", nodeConfig.Name),
		},
		Type: "Opaque",
		Data: secretData,
	}

	clusterAgentYAML := api.ClusterAgentYAML{
		APIVersion: "agent.cubbit.io/v1alpha1",
		Kind:       "ClusterAgent",
		Metadata: api.ClusterAgentMetadata{
			Name: fmt.Sprintf("cluster-%s", nodeConfig.Name),
		},
		Spec: api.ClusterAgentSpec{
			InstancesCounter: len(nodeConfig.Agents),
			BaseName:         fmt.Sprintf("%s-agent", nodeConfig.Name),
			SecretName:       fmt.Sprintf("%s-secret", nodeConfig.Name),
			AgentImage:       "cubbit/agent:latest",
			AdditionalEnvVars: []api.EnvVar{
				{Name: "HIVE_URL", Value: envs.HiveURL},
				{Name: "METRICS_URL", Value: envs.MetricsURL},
				{Name: "METRICS_ROUTES_SEND", Value: envs.MetricsRoutesSend},
				{Name: "CCCP_SWARM_GATEWAY_ENDPOINT", Value: envs.CCCPSwarmGatewayEndpoint},
				{Name: "CCCP_SWARM_GATEWAY_PORT", Value: envs.CCCPSwarmGatewayPort},
				{Name: "CCCP_SWARM_GATEWAY_SECURE", Value: envs.CCCPSwarmGatewaySecure},
			},
			Volume: api.VolumeSpec{
				Type:    "local-storage",
				PVCSize: "1Ti",
			},
			AgentsDetail: agentsDetail,
		},
	}

	secretYAMLBytes, err := yaml.Marshal(&secretYAML)
	if err != nil {
		return "", fmt.Errorf("error marshaling secret YAML: %w", err)
	}

	clusterAgentYAMLBytes, err := yaml.Marshal(&clusterAgentYAML)
	if err != nil {
		return "", fmt.Errorf("error marshaling cluster agent YAML: %w", err)
	}

	return string(secretYAMLBytes) + "\n---\n" + string(clusterAgentYAMLBytes), nil
}

func downloadAndGenerateAnsibleTar(tarURL string, nodeConfigs []api.NodeConfig, outputPath string) error {
	resp, err := http.Get(tarURL)
	if err != nil {
		return fmt.Errorf("failed to download tarball: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download tarball: %s", resp.Status)
	}

	tmpDir, err := os.MkdirTemp("", "ansible-tar")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	if err := utils.ExtractTar(resp.Body, tmpDir); err != nil {
		return fmt.Errorf("failed to extract tarball: %w", err)
	}

	if err := injectAnsibleFiles(tmpDir+"/cubbit-agent-playbook", nodeConfigs); err != nil {
		return fmt.Errorf("failed to inject generated files: %w", err)
	}

	if err := utils.CreateTar(outputPath, tmpDir); err != nil {
		return fmt.Errorf("failed to repack tarball: %w", err)
	}

	return nil
}

func injectAnsibleFiles(root string, nodes []api.NodeConfig) error {
	filesDir := filepath.Join(root, "files")
	invDir := filepath.Join(root, "inventory")

	if err := os.MkdirAll(filesDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(invDir, 0755); err != nil {
		return err
	}

	invPath := filepath.Join(invDir, "hosts.ini")
	if err := os.WriteFile(invPath, []byte(generateInventoryContent(nodes)), 0644); err != nil {
		return err
	}

	hostNamesPath := filepath.Join(filesDir, "host-names")
	if err := os.WriteFile(hostNamesPath, []byte(generateHostNamesContent(nodes)), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(filesDir, "ssh-public-keys"), []byte(""), 0644); err != nil {
		return err
	}

	for _, node := range nodes {
		path := filepath.Join(filesDir, fmt.Sprintf("%s-agent-secrets.json", node.Name))
		data, _ := json.MarshalIndent(generateAgentSecretsContent(node), "", "  ")
		if err := os.WriteFile(path, data, 0644); err != nil {
			return err
		}
	}

	return nil
}

func generateDeployFiles(conf *configuration.Config, deployOption string, nodeConfigs []api.NodeConfig) error {
	basePath := "./cubbit-agent-playbook.tar"
	outputPath := "./cubbit-agent-yaml.tar"

	switch deployOption {
	case "ansible":
		return downloadAndGenerateAnsibleTar(constants.AnsibleTarUrl, nodeConfigs, basePath)
	case "yaml":
		return generateYAMLFiles(conf, nodeConfigs, outputPath)
	case "both":
		if err := downloadAndGenerateAnsibleTar(constants.AnsibleTarUrl, nodeConfigs, basePath); err != nil {
			return fmt.Errorf("error generating Ansible files: %w", err)
		}

		if err := generateYAMLFiles(conf, nodeConfigs, outputPath); err != nil {
			return fmt.Errorf("error generating YAML files: %w", err)
		}
	case "skip":
		return nil
	default:
		return fmt.Errorf("invalid deployment option: %s", deployOption)
	}

	return nil
}
