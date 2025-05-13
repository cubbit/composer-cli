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
