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

func CreateSwarmAgent(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, nexusID, nodeID, agentFeaturesStr, agentDisk, agentMountPoint, configPath, nodeConfigStr string
	var agentPort int
	var agentFeatures map[string]interface{}
	var conf *configuration.Config
	var agent *api.NewAgentsResponse

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentPort, err = cmd.Flags().GetInt("agent-port"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentDisk, err = cmd.Flags().GetString("agent-disk"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentMountPoint, err = cmd.Flags().GetString("agent-mount-point"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentFeaturesStr, err = cmd.Flags().GetString("agent-features"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentFeaturesStr != "" {
		if err = json.Unmarshal([]byte(nodeConfigStr), &agentFeatures); err != nil {
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

	agentsBodyRequest := api.BulkInsertNewAgentRequestBody{
		Agents: []api.CreateNewAgentRequestBody{
			{
				Port: agentPort,
				Volume: api.AgentVolume{
					MountPoint: agentMountPoint,
					Disk:       agentDisk},
				Features: agentFeatures,
			},
		},
	}

	if agent, err = api.CreateAgentV4(conf.Urls, *accessToken, id, nexusID, nodeID, agentsBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Agent %s created successfully\n", agent.Agents[0].ID))

	return nil
}

func CreateSwarmAgentBatch(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, nexusID, nodeID, configPath, filePath string
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

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
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

	var batchAgents api.BulkInsertNewAgentRequestBody
	if err = json.Unmarshal(fileData, &batchAgents); err != nil {
		schemaTemplate := map[string]interface{}{
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
		}

		return fmt.Errorf("error parsing batch file: %w\nExpected file schema:\n%s", err, getStructSchema(schemaTemplate))
	}

	if len(batchAgents.Agents) == 0 {
		return fmt.Errorf("no agents found in batch file")
	}

	agents, err := api.CreateAgentV4(conf.Urls, *accessToken, id, nexusID, nodeID, batchAgents)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Successfully created %d agents\n", len(agents.Agents)))
	for _, agent := range agents.Agents {
		fmt.Printf("  • %s\n", agent.ID)
	}

	return nil
}

func DescribeSwarmAgent(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath, format string
	var conf *configuration.Config
	var agents *api.GenericPaginatedResponse[*api.NewAgent]
	var id, name, nexusID, nodeID, agentID string

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentID, err = cmd.Flags().GetString("agent-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
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

	if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, nodeID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(agents.Data) == 0 {
		utils.PrintError(fmt.Errorf("agent not found"))
		return nil
	}

	for _, agent := range agents.Data {
		if agent.ID == agentID {
			utils.PrintFormattedData(agent, format)
		}
	}

	return nil
}

func ListSwarmAgents(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, nexusID, nodeID, rcID, configPath, sort, filter string
	var conf *configuration.Config
	var agents *api.GenericPaginatedResponse[*api.NewAgent]
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

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if rcID, err = cmd.Flags().GetString("redundancy-class-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if sort, err = cmd.Flags().GetString("sort"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter, err = cmd.Flags().GetString("filter"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter != "" {
		filter = utils.BuildFilterQuery(filter)
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

	if rcID != "" {
		if agents, err = api.ListAgentsForRCV4(conf.Urls, *accessToken, id, rcID, sort, filter); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
		}
	} else if nexusID != "" && nodeID != "" {
		if agents, err = api.ListAgentsV4(conf.Urls, *accessToken, id, nexusID, nodeID, sort, filter); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
		}
	}

	utils.PrintList("Your Agents List")

	if len(agents.Data) == 0 {
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
		utils.PrintVerbose(agents.Data, l)
		return nil
	}

	for _, agent := range agents.Data {
		fmt.Printf(" • %s\n", agent.ID)
		if l {
			fmt.Println()
		}
	}

	return nil
}

func EditSwarmAgent(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, nexusID, nodeID, agentID, agentDisk, agentMountPoint, configPath, agentFeaturesStr string
	var agentPort int
	var agentFeatures map[string]interface{}
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

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentID, err = cmd.Flags().GetString("agent-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentPort, err = cmd.Flags().GetInt("agent-port"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentDisk, err = cmd.Flags().GetString("agent-disk"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentMountPoint, err = cmd.Flags().GetString("agent-mount-point"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentFeaturesStr, err = cmd.Flags().GetString("agent-features"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentFeaturesStr != "" {
		if err = json.Unmarshal([]byte(agentFeaturesStr), &agentFeatures); err != nil {
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

	var agentBodyRequest api.UpdateNewAgentRequestBody

	if agentPort != 0 {
		agentBodyRequest.Port = &agentPort
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

	if agentFeaturesStr != "" {
		agentBodyRequest.Features = agentFeatures
	}

	if err = api.UpdateAgentV4(conf.Urls, *accessToken, id, nexusID, nodeID, agentID, agentBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Agent %s updated successfully\n", agentID))

	return nil
}

func RemoveSwarmAgent(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var id, name, nexusID, nodeID, agentID string

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nodeID, err = cmd.Flags().GetString("node-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if agentID, err = cmd.Flags().GetString("agent-id"); err != nil {
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

	if err = api.DeleteAgentV4(conf.Urls, *accessToken, id, nexusID, nodeID, agentID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNodeRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("Agent %s deleted successfully\n", agentID))

	return nil
}
