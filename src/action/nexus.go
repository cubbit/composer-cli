// Package action provides CLI actions for managing nexuses.
package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateNexus(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, name, description, location, providerID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nexus *api.Nexus

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if location, err = cmd.Flags().GetString("location"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if providerID, err = cmd.Flags().GetString("provider-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	nexusBodyRequest := api.CreateNexusRequestBody{
		Name:        name,
		Description: description,
		Location:    location,
		ProviderID:  providerID,
	}

	if nexus, err = api.CreateNexus(*urls, resolvedProfile.APIKey, swarmID, nexusBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNexusRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.CreateNexusRequestBody{nexusBodyRequest},
		func(n api.CreateNexusRequestBody) []string {
			return []string{nexus.ID}
		},
		&utils.SmartOutputConfig[api.CreateNexusRequestBody]{
			SingleResourceCompactOutput: true,
			SingleResource:              true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func DescribeNexus(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nexus *api.Nexus

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if nexus, err = api.GetNexus(*urls, resolvedProfile.APIKey, swarmID, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNexusRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.Nexus{nexus},
		func(n *api.Nexus) []string {
			return []string{n.ID, n.Name, n.Description, n.Location}
		},
		&utils.SmartOutputConfig[*api.Nexus]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)

}

func EditNexus(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, nexusName, description, location string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusName, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if location, err = cmd.Flags().GetString("location"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	nexusBodyRequest := api.UpdateNexusRequestBody{
		Name:        nexusName,
		Description: description,
		Location:    location,
	}

	if err = api.UpdateNexus(*urls, resolvedProfile.APIKey, swarmID, nexusID, nexusBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingNexusRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.UpdateNexusRequestBody{&nexusBodyRequest},
		func(n *api.UpdateNexusRequestBody) []string {
			return []string{nexusID}
		},
		&utils.SmartOutputConfig[*api.UpdateNexusRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func ListNexuses(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nexuses *api.GenericPaginatedResponse[*api.Nexus]

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if nexuses, err = api.ListNexuses(*urls, resolvedProfile.APIKey, swarmID, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		nexuses.Data,
		func(n *api.Nexus) []string {
			return []string{n.ID}
		},
		&utils.SmartOutputConfig[*api.Nexus]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func RemoveNexus(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.DeleteNexus(*urls, resolvedProfile.APIKey, swarmID, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNexusRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{nexusID},
		func(s string) []string { return []string{s} },
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func GenerateNexusDeployFiles(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, nexusID, outputDir string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nodes *api.GenericPaginatedResponse[*api.NewNode]

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusID, err = cmd.Flags().GetString("nexus-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if outputDir, err = cmd.Flags().GetString("output-dir"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if nodes, err = api.ListNodes(*urls, resolvedProfile.APIKey, swarmID, nexusID, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nodes.Data) == 0 {
		utils.PrintNotFound("No nodes found")
		return nil
	}

	var nodeConfigs []api.NodeConfig

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
		if agents, err = api.ListAgents(*urls, resolvedProfile.APIKey, swarmID, nexusID, node.ID, "", ""); err != nil {
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

	if err = generateDeployFiles(urls, "ansible", nodeConfigs, outputDir); err != nil {
		return fmt.Errorf("error generating deployment files: %w", err)
	}

	utils.PrintSuccess("Deployment files generated successfully in directory: " + outputDir)
	return nil
}
