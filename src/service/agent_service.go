package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	utils "github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type AgentServiceInterface interface {
	CreateAgent(cmd *cobra.Command, args []string) error
	CreateAgentBatch(cmd *cobra.Command, args []string) error
	DescribeAgent(cmd *cobra.Command, args []string) error
	EditAgent(cmd *cobra.Command, args []string) error
	ListAgents(cmd *cobra.Command, args []string) error
	RemoveAgent(cmd *cobra.Command, args []string) error
	CheckAgentStatus(cmd *cobra.Command, args []string) error
}

type AgentService struct {
	configuration *configuration.Config
}

func NewAgentService(
	configuration *configuration.Config,
) AgentService {
	return AgentService{
		configuration: configuration,
	}
}

func (s AgentService) CreateAgent(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID, agentFeaturesStr, agentDisk, agentMountPoint, nodeConfigStr string
	var agentPort int
	var agentFeatures map[string]interface{}
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var agent *api.NewAgentsResponse

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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
			return fmt.Errorf("%s: %w", constants.ErrorParsingJSONConfiguration, err)
		}
	}

	if resolvedProfile, urls, err = s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	agentsBodyRequest := api.BulkInsertNewAgentRequestBody{
		Agents: []api.CreateNewAgentRequestBody{
			{
				Port: agentPort,
				Volume: api.AgentVolume{
					MountPoint: agentMountPoint,
					Disk:       agentDisk,
				},
				Features: agentFeatures,
			},
		},
	}

	if agent, err = api.CreateAgent(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, agentsBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		agent.Agents,
		func(a *api.NewAgentResponse) []string {
			return []string{
				a.ID,
			}
		},
		&utils.SmartOutputConfig[*api.NewAgentResponse]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func (s AgentService) CreateAgentBatch(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID, filePath string
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if resolvedProfile, urls, err = s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
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

		return fmt.Errorf("error parsing batch file: %w\nExpected file schema:\n%s", err, utils.GetStructSchema(schemaTemplate))
	}

	if len(batchAgents.Agents) == 0 {
		return fmt.Errorf("no agents found in batch file")
	}

	agents, err := api.CreateAgent(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, batchAgents)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		agents.Agents,
		func(a *api.NewAgentResponse) []string {
			return []string{
				a.ID,
			}
		},
		&utils.SmartOutputConfig[*api.NewAgentResponse]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func (s AgentService) DescribeAgent(cmd *cobra.Command, args []string) error {
	var err error
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var agents *api.GenericPaginatedResponse[*api.NewAgent]
	var swarmID, nexusID, nodeID, agentID string

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if agents, err = api.GetAgent(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, agentID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
	}

	if len(agents.Data) == 0 {
		return fmt.Errorf("no agent found with ID %s", agentID)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.NewAgent{agents.Data[0]},
		func(a *api.NewAgent) []string {
			return []string{
				a.ID,
				fmt.Sprintf("%d", a.Port),
				a.Volume.MountPoint,
				a.Volume.Disk,
			}
		},
		&utils.SmartOutputConfig[*api.NewAgent]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func (s AgentService) EditAgent(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID, agentID, agentDisk, agentMountPoint, agentFeaturesStr string
	var agentPort int
	var agentFeatures map[string]interface{}
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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
			return fmt.Errorf("%s: %w", constants.ErrorParsingJSONConfiguration, err)
		}
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
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

	if err = api.UpdateAgent(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, agentID, agentBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.UpdateNewAgentRequestBody{
			&agentBodyRequest},
		func(a *api.UpdateNewAgentRequestBody) []string {
			return []string{
				agentID,
			}
		},
		&utils.SmartOutputConfig[*api.UpdateNewAgentRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func (s AgentService) ListAgents(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID, rcID, sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var agents *api.GenericPaginatedResponse[*api.NewAgent]

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if rcID != "" {
		if agents, err = api.ListAgentsForRC(*urls, resolvedProfile.APIKey, swarmID, rcID, sort, filter); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingAgentsRequest, err)
		}
	} else if nexusID != "" && nodeID != "" {
		if agents, err = api.ListAgents(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, sort, filter); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
		}
	}

	return utils.PrintSmartOutput(
		cmd,
		agents.Data,
		func(a *api.NewAgent) []string {
			return []string{
				a.ID,
				fmt.Sprintf("%d", a.Port),
				a.Volume.MountPoint,
				a.Volume.Disk,
			}
		},
		&utils.SmartOutputConfig[*api.NewAgent]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func (s AgentService) RemoveAgent(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nodeID, agentID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.DeleteAgent(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, agentID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNodeRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{agentID},
		func(s string) []string { return []string{s} },
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func (s AgentService) CheckAgentStatus(cmd *cobra.Command, args []string) error {
	var err error

	var swarmID, nexusID, nodeID, agentID string

	var conf *configuration.Config

	var resolvedProfile *configuration.ResolvedProfile

	var urls *configuration.URLs

	var status *api.GetAgentEvaluatedStatusResponse

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {

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

	if conf, err = configuration.LoadConfig(); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)

	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)

	}

	if status, err = api.GetAgentStatus(*urls, resolvedProfile.APIKey, swarmID, nexusID, nodeID, agentID); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRetrievingAgentStatusRequest, err)

	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.GetAgentEvaluatedStatusResponse{status},
		func(s *api.GetAgentEvaluatedStatusResponse) []string {
			return []string{
				string(s.Status),
			}
		},
		&utils.SmartOutputConfig[*api.GetAgentEvaluatedStatusResponse]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
		},
	)
}
