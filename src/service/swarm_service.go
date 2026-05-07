package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type SwarmServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
	Describe(cmd *cobra.Command, args []string) error
}

type SwarmService struct {
	configuration            configuration.ConfigInterface
	swarmAPI                 api.SwarmAPIInterface
	locationAPI              api.LocationAPIInterface
	processAPI               api.ProcessAPIInterface
	redundancyClassValidator RedundancyClassValidatorInterface
}

func NewSwarmService(
	configuration configuration.ConfigInterface,
	swarmAPI api.SwarmAPIInterface,
	locationAPI api.LocationAPIInterface,
	processAPI api.ProcessAPIInterface,
	redundancyClassValidator RedundancyClassValidatorInterface,
) SwarmService {
	return SwarmService{
		configuration:            configuration,
		swarmAPI:                 swarmAPI,
		locationAPI:              locationAPI,
		processAPI:               processAPI,
		redundancyClassValidator: redundancyClassValidator,
	}
}

func (s SwarmService) Create(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
	}

	description, err := utils.GetOptionalStringFlag(cmd, "description")
	if err != nil {
		return fmt.Errorf("%s description: %w", constants.ErrorRetrievingField, err)
	}

	nexusFlags, err := cmd.Flags().GetStringArray("nexus")
	if err != nil {
		return fmt.Errorf("%s nexus: %w", constants.ErrorRetrievingField, err)
	}

	rcFlags, err := cmd.Flags().GetStringArray("redundancy-class")
	if err != nil {
		return fmt.Errorf("%s redundancy-class: %w", constants.ErrorRetrievingField, err)
	}

	rcFile, err := cmd.Flags().GetString("redundancy-class-file")
	if err != nil {
		return fmt.Errorf("%s redundancy-class-file: %w", constants.ErrorRetrievingField, err)
	}

	nexuses, err := s.parseNexuses(cmd, nexusFlags, *resolvedProfile, *urls)
	if err != nil {
		return err
	}

	redundancyClasses, err := s.parseRedundancyClasses(rcFlags, rcFile)
	if err != nil {
		return err
	}

	if err := s.redundancyClassValidator.ValidateRedundancyClasses(redundancyClasses, nexuses); err != nil {
		return fmt.Errorf("redundancy class validation failed: %w", err)
	}

	processes, err := s.processAPI.ListProcesses(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, api.ProcessTypeSwarmCreation)
	if err != nil {
		return fmt.Errorf("failed to check for ongoing swarm creation processes: %w", err)
	}

	for _, process := range processes {
		if process.Status == api.ProcessStatusRunning {
			return fmt.Errorf("swarm creation already in progress")
		}
	}

	createRequest := &api.CreateSwarmV5Request{
		Name:              name,
		Description:       description,
		Configuration:     make(map[string]interface{}),
		Nexuses:           nexuses,
		RedundancyClasses: redundancyClasses,
	}

	response, err := s.swarmAPI.CreateSwarmV5(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, createRequest)
	if err != nil {
		return fmt.Errorf("failed to create swarm: %w", err)
	}

	return printer.PrintText(cmd, fmt.Sprintf("Swarm creation started with Process ID: %s\n", response.ID))
}

func (s SwarmService) Describe(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	swarmID, err := s.resolveSwarmID(cmd, args, *resolvedProfile, *urls)
	if err != nil {
		return err
	}

	swarm, err := s.swarmAPI.GetSwarmV5(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, swarmID)
	if err != nil {
		return fmt.Errorf("failed to describe swarm: %w", err)
	}

	output, err := resolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		return PrintSwarmDetails(cmd, *swarm)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), swarm, output)
	return nil
}

func (s SwarmService) resolveSwarmID(cmd *cobra.Command, args []string, resolvedProfile configuration.ResolvedProfile, urls configuration.URLs) (string, error) {
	identifiers := 0

	swarmIDFlag, err := cmd.Flags().GetString("swarm-id")
	if err != nil {
		return "", fmt.Errorf("%s swarm-id: %w", constants.ErrorRetrievingField, err)
	}
	if swarmIDFlag != "" {
		identifiers++
	}

	swarmNameFlag, err := cmd.Flags().GetString("swarm-name")
	if err != nil {
		return "", fmt.Errorf("%s swarm-name: %w", constants.ErrorRetrievingField, err)
	}
	if swarmNameFlag != "" {
		identifiers++
	}

	swarmIDPositional := ""
	if len(args) > 0 {
		swarmIDPositional = strings.TrimSpace(args[0])
		if swarmIDPositional != "" {
			identifiers++
		}
	}

	if identifiers != 1 {
		return "", fmt.Errorf("specify exactly one of SWARM_ID, --swarm-id or --swarm-name")
	}

	if swarmIDPositional != "" {
		return swarmIDPositional, nil
	}

	if swarmIDFlag != "" {
		return swarmIDFlag, nil
	}

	swarmID, err := s.resolveSwarmIDByName(urls, resolvedProfile, swarmNameFlag)
	if err != nil {
		return "", err
	}
	return swarmID, nil
}

func (s SwarmService) resolveSwarmIDByName(urls configuration.URLs, resolvedProfile configuration.ResolvedProfile, swarmName string) (string, error) {
	page := 1
	const itemsPerPage = 1000

	for {
		response, err := s.swarmAPI.ListSwarmsV5(urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, page, itemsPerPage)
		if err != nil {
			return "", fmt.Errorf("failed to resolve swarm name '%s': %w", swarmName, err)
		}

		for _, swarm := range response.Data {
			if swarm.Name == swarmName {
				return swarm.ID, nil
			}
		}

		if len(response.Data) == 0 {
			break
		}

		if response.NextPage != nil {
			page = *response.NextPage
			continue
		}

		page++
	}

	return "", fmt.Errorf("swarm with name '%s' not found", swarmName)
}

func (s SwarmService) parseNexuses(cmd *cobra.Command, nexusFlags []string, resolvedProfile configuration.ResolvedProfile, urls configuration.URLs) ([]api.NexusV5Request, error) {
	clusters, err := s.locationAPI.ListAggregated(urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cluster information: %w", err)
	}

	clusterMap := make(map[string]api.InfraAggregateCluster)
	for _, cluster := range clusters {
		clusterMap[cluster.ClusterID] = cluster
	}

	nexuses := make([]api.NexusV5Request, 0)

	for _, nexusFlag := range nexusFlags {
		nexus, err := s.parseNexus(nexusFlag, clusterMap)
		if err != nil {
			return nil, err
		}
		nexuses = append(nexuses, nexus)
	}

	return nexuses, nil
}

func (s SwarmService) parseNexus(nexusFlag string, clusterMap map[string]api.InfraAggregateCluster) (api.NexusV5Request, error) {
	parts := strings.SplitN(nexusFlag, ":", 2)
	if len(parts) != 2 {
		return api.NexusV5Request{}, fmt.Errorf("invalid nexus format: expected 'cluster-id:node-id1,node-id2', got '%s'", nexusFlag)
	}

	clusterID := strings.TrimSpace(parts[0])
	if clusterID == "" {
		return api.NexusV5Request{}, fmt.Errorf("invalid nexus format: cluster ID cannot be empty in '%s'", nexusFlag)
	}
	rawNodeIDs := strings.Split(parts[1], ",")
	nodeIDs := make([]string, 0, len(rawNodeIDs))
	for _, rawNodeID := range rawNodeIDs {
		nodeID := strings.TrimSpace(rawNodeID)
		if nodeID == "" {
			return api.NexusV5Request{}, fmt.Errorf("invalid nexus format: node ID cannot be empty in '%s'", nexusFlag)
		}
		nodeIDs = append(nodeIDs, nodeID)
	}

	cluster, found := clusterMap[clusterID]
	if !found {
		return api.NexusV5Request{}, fmt.Errorf("cluster with ID '%s' not found", clusterID)
	}

	nexus := api.NexusV5Request{
		ClusterID:   clusterID,
		ClusterType: cluster.Type,
	}

	if cluster.Type == api.ClusterTypeVirtual {
		virtualNodes := make([]api.VirtualNodeRequest, 0)
		for _, nodeID := range nodeIDs {
			found := false
			for _, virtualNode := range cluster.Details.VirtualNodes {
				if virtualNode.NodeID == nodeID {
					found = true
					break
				}
			}
			if !found {
				return api.NexusV5Request{}, fmt.Errorf("virtual node with ID '%s' not found in cluster '%s'", nodeID, clusterID)
			}
			virtualNodes = append(virtualNodes, api.VirtualNodeRequest{
				ServerID: nodeID,
			})
		}
		nexus.VirtualNodes = virtualNodes
	} else {
		nodes := make([]api.NodeRequest, 0)
		for _, nodeID := range nodeIDs {
			var clusterNode *api.InfraAggregateNodeDetail
			for i := range cluster.Details.Nodes {
				if cluster.Details.Nodes[i].NodeID == nodeID {
					clusterNode = &cluster.Details.Nodes[i]
					break
				}
			}
			if clusterNode == nil {
				return api.NexusV5Request{}, fmt.Errorf("node with ID '%s' not found in cluster '%s'", nodeID, clusterID)
			}

			volumes := make([]api.VolumeRequest, 0)
			for _, disk := range clusterNode.Disks {
				if disk.Status.Code == string(api.StatusCodeOk) {
					volumes = append(volumes, api.VolumeRequest{
						VolumeID: disk.DiskUUID,
					})
				}
			}

			if len(volumes) == 0 {
				return api.NexusV5Request{}, fmt.Errorf("node '%s' in cluster '%s' has no usable volumes", nodeID, clusterID)
			}

			nodes = append(nodes, api.NodeRequest{
				ServerID: nodeID,
				Volumes:  volumes,
			})
		}
		nexus.Nodes = nodes
	}

	return nexus, nil
}

func (s SwarmService) parseRedundancyClasses(rcFlags []string, rcFile string) ([]api.RedundancyClassRequest, error) {
	redundancyClasses := make([]api.RedundancyClassRequest, 0)

	for _, rcFlag := range rcFlags {
		var rc api.RedundancyClassRequest
		err := json.Unmarshal([]byte(rcFlag), &rc)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", constants.ErrorParsingJSONConfiguration, err)
		}
		redundancyClasses = append(redundancyClasses, rc)
	}

	if rcFile != "" {
		fileContent, err := os.ReadFile(rcFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read redundancy class file: %w", err)
		}

		var fileRCs []api.RedundancyClassRequest
		if err := json.Unmarshal(fileContent, &fileRCs); err != nil {
			return nil, fmt.Errorf("%s: %w", constants.ErrorParsingJSONConfiguration, err)
		}

		redundancyClasses = append(redundancyClasses, fileRCs...)
	}

	return redundancyClasses, nil
}

func resolveCommandOutput(cmd *cobra.Command, defaultOutput configuration.OutputFormat) (string, error) {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return "", fmt.Errorf("%s output: %w", constants.ErrorRetrievingField, err)
	}

	if defaultOutput != "" &&
		!cmd.Flags().Changed("output") &&
		!cmd.Flags().Changed("quiet") {
		output = string(defaultOutput)
	}

	return output, nil
}
