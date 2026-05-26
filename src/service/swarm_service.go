package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/tui"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type SwarmServiceInterface interface {
	Create(cmd *cobra.Command, args []string) error
	CreateInteractive(cmd *cobra.Command, args []string) error
	Describe(cmd *cobra.Command, args []string) error
	List(cmd *cobra.Command, args []string) error
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

func (s SwarmService) CreateInteractive(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	return createInteractive(resolvedProfile, urls, s.swarmAPI, s.locationAPI, s.processAPI, s.redundancyClassValidator)
}

type infrastructureLimits struct {
	minDisksPerNode     int
	minNodesPerLocation int
	minLocationDisks    int
}

func createInteractive(
	profile *configuration.ResolvedProfile,
	urls *configuration.URLs,
	swarmAPI api.SwarmAPIInterface,
	locationAPI api.LocationAPIInterface,
	processAPI api.ProcessAPIInterface,
	rcValidator RedundancyClassValidatorInterface,
) error {
	existing, err := checkExistingProcess(processAPI, profile, urls)
	if err != nil {
		return err
	}
	if existing {
		return nil
	}

	name, description, err := collectSwarmInfo()
	if err != nil {
		return err
	}
	if name == "" {
		return nil
	}

	clusters, nodeIDsByCluster, totalNodes, err := selectInfrastructure(locationAPI, profile, urls)
	if err != nil {
		return err
	}
	if clusters == nil {
		return nil
	}

	ok, err := confirmInfrastructure(clusters, nodeIDsByCluster, totalNodes)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	limits := computeInfrastructureLimits(clusters, nodeIDsByCluster)

	redundancyClasses, ok, err := collectRedundancyClasses(clusters, nodeIDsByCluster, limits)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	ok, err = confirmSubmission(name, len(clusters), totalNodes, len(redundancyClasses))
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	return submitSwarm(swarmAPI, processAPI, rcValidator, profile, urls, name, description, clusters, nodeIDsByCluster, redundancyClasses)
}

func checkExistingProcess(processAPI api.ProcessAPIInterface, profile *configuration.ResolvedProfile, urls *configuration.URLs) (bool, error) {
	processes, err := processAPI.ListProcesses(*urls, profile.APIKey, profile.OrganizationID, api.WithProcessType(api.ProcessTypeSwarmCreation))
	if err != nil {
		return false, nil
	}

	for _, proc := range processes {
		if proc.Status == api.ProcessStatusRunning {
			choice, err := tui.ConfirmDialog(
				"Swarm creation already in progress",
				fmt.Sprintf("An ongoing swarm creation process was found (ID: %s).", proc.ID),
				[]string{"Monitor existing process", "Exit"},
			)
			if err != nil {
				return false, err
			}
			if choice == "Monitor existing process" {
				monitorSwarmCreation(processAPI, *urls, profile.APIKey, profile.OrganizationID, proc.ID)
			}
			return true, nil
		}
	}
	return false, nil
}

func collectSwarmInfo() (name string, description string, err error) {
	swarmName := ""
	swarmDesc := ""

	inputs, err := tui.TextInputs("Phase 1 - Swarm Information\nStep 1 of 9: Enter swarm details", false,
		tui.Input{Placeholder: "Swarm name (3-63 characters)", Value: &swarmName},
		tui.Input{Placeholder: "Description (optional)", Value: &swarmDesc},
	)
	if err != nil {
		return "", "", err
	}

	if len(inputs) < 2 || len(*inputs[0].Value) < 3 || len(*inputs[0].Value) > 63 {
		tui.ShowError("Swarm name must be between 3 and 63 characters")
		return "", "", nil
	}

	return *inputs[0].Value, *inputs[1].Value, nil
}

func selectInfrastructure(
	locationAPI api.LocationAPIInterface,
	profile *configuration.ResolvedProfile,
	urls *configuration.URLs,
) ([]api.InfraAggregateCluster, map[string][]string, int, error) {
	clusters, err := locationAPI.ListAggregated(*urls, profile.APIKey, profile.OrganizationID)
	if err != nil {
		tui.ShowError(fmt.Sprintf("Failed to fetch locations: %v", err))
		return nil, nil, 0, nil
	}

	if len(clusters) == 0 {
		tui.ShowError("No locations available")
		return nil, nil, 0, nil
	}

	clusterOptions := make([]string, len(clusters))
	for i, c := range clusters {
		clusterOptions[i] = fmt.Sprintf("%s (%s) - %s", c.Name, c.ClusterID, c.Type)
	}

	selected, err := tui.ChooseMany("Phase 1 - Location Selection\nStep 2 of 9: Select one or more locations (space to select, enter to confirm)", false, clusterOptions)
	if err != nil {
		return nil, nil, 0, err
	}

	selectedClusters := make([]api.InfraAggregateCluster, 0, len(selected))
	for _, sel := range selected {
		for i, opt := range clusterOptions {
			if sel == opt {
				selectedClusters = append(selectedClusters, clusters[i])
				break
			}
		}
	}

	if len(selectedClusters) == 0 {
		tui.ShowError("At least one location must be selected")
		return nil, nil, 0, nil
	}

	nodeIDsByCluster := make(map[string][]string)
	totalNodes := 0

	for _, cluster := range selectedClusters {
		nodeOptions, nodeIDMap := nodeOptionsForCluster(cluster)

		if len(nodeOptions) == 0 {
			tui.ShowError(fmt.Sprintf("No available servers in location '%s'", cluster.Name))
			return nil, nil, 0, nil
		}

		selectedNodes, err := tui.ChooseMany(
			fmt.Sprintf("Phase 1 - Node Selection\nStep 3 of 9: Select servers for location '%s' (space to select, enter to confirm)", cluster.Name),
			false,
			nodeOptions,
		)
		if err != nil {
			return nil, nil, 0, err
		}

		nodeIDs := make([]string, 0, len(selectedNodes))
		for _, sel := range selectedNodes {
			if id, ok := nodeIDMap[sel]; ok {
				nodeIDs = append(nodeIDs, id)
			}
		}
		nodeIDsByCluster[cluster.ClusterID] = nodeIDs
		totalNodes += len(nodeIDs)
	}

	return selectedClusters, nodeIDsByCluster, totalNodes, nil
}

func nodeOptionsForCluster(cluster api.InfraAggregateCluster) ([]string, map[string]string) {
	nodeOptions := make([]string, 0)
	nodeIDMap := make(map[string]string)

	if cluster.Type == api.ClusterTypePhysical {
		for _, node := range cluster.Details.Nodes {
			if node.Status.Code == string(api.StatusCodeOk) {
				label := fmt.Sprintf("%s (%s)", node.NodeName, node.NodeID)
				if node.InternalIP != nil {
					label = fmt.Sprintf("%s (%s) - IP: %s", node.NodeName, node.NodeID, *node.InternalIP)
				}
				nodeOptions = append(nodeOptions, label)
				nodeIDMap[label] = node.NodeID
			}
		}
	} else {
		for _, vnode := range cluster.Details.VirtualNodes {
			if vnode.Status.Code == string(api.StatusCodeOk) {
				label := fmt.Sprintf("%s (%s) [virtual]", vnode.NodeName, vnode.NodeID)
				nodeOptions = append(nodeOptions, label)
				nodeIDMap[label] = vnode.NodeID
			}
		}
	}

	return nodeOptions, nodeIDMap
}

func confirmInfrastructure(clusters []api.InfraAggregateCluster, nodeIDsByCluster map[string][]string, totalNodes int) (bool, error) {
	recapBody := fmt.Sprintf("Locations: %d\nTotal nodes: %d\n\n", len(clusters), totalNodes)
	for _, c := range clusters {
		recapBody += fmt.Sprintf("  %s (%s): %d nodes\n", c.Name, c.ClusterID, len(nodeIDsByCluster[c.ClusterID]))
	}

	choice, err := tui.ConfirmDialog(
		"Phase 1 - Infrastructure Recap\nStep 4 of 9: Review your selections",
		recapBody,
		[]string{"Yes, continue", "No, go back"},
	)
	if err != nil {
		return false, err
	}
	if choice != "Yes, continue" {
		tui.ShowError("Cancelled by user")
		return false, nil
	}
	return true, nil
}

func computeInfrastructureLimits(clusters []api.InfraAggregateCluster, nodeIDsByCluster map[string][]string) infrastructureLimits {
	limits := infrastructureLimits{}
	selectedNodeIDs := make(map[string]map[string]bool)
	for _, c := range clusters {
		ids := make(map[string]bool)
		for _, nid := range nodeIDsByCluster[c.ClusterID] {
			ids[nid] = true
		}
		selectedNodeIDs[c.ClusterID] = ids
	}

	perNodeDisks := make([]int, 0)
	perLocationDisks := make([]int, 0)

	for _, c := range clusters {
		nodeCount := len(nodeIDsByCluster[c.ClusterID])
		if limits.minNodesPerLocation == 0 || nodeCount < limits.minNodesPerLocation {
			limits.minNodesPerLocation = nodeCount
		}

		selected := selectedNodeIDs[c.ClusterID]

		locationTotal := 0
		for _, node := range c.Details.Nodes {
			if !selected[node.NodeID] {
				continue
			}
			healthy := 0
			for _, d := range node.Disks {
				if d.Status.Code == string(api.StatusCodeOk) {
					healthy++
				}
			}
			if healthy > 0 {
				perNodeDisks = append(perNodeDisks, healthy)
			}
			locationTotal += healthy
		}
		for _, vnode := range c.Details.VirtualNodes {
			if !selected[vnode.NodeID] {
				continue
			}
			perNodeDisks = append(perNodeDisks, 1)
			locationTotal++
		}
		if locationTotal > 0 {
			perLocationDisks = append(perLocationDisks, locationTotal)
		}
	}

	if len(perNodeDisks) > 0 {
		limits.minDisksPerNode = minIntSlice(perNodeDisks)
	}
	if len(perLocationDisks) > 0 {
		limits.minLocationDisks = minIntSlice(perLocationDisks)
	}
	if len(perLocationDisks) > 0 {
		limits.minLocationDisks = minIntSlice(perLocationDisks)
	}

	return limits
}

func minIntSlice(vals []int) int {
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func collectRedundancyClasses(
	clusters []api.InfraAggregateCluster,
	nodeIDsByCluster map[string][]string,
	limits infrastructureLimits,
) ([]api.RedundancyClassRequest, bool, error) {
	redundancyClasses := make([]api.RedundancyClassRequest, 0)
	allClusterIDs := make([]string, len(clusters))
	for i, c := range clusters {
		allClusterIDs[i] = c.ClusterID
	}

	for {
		rcName, rcDesc, err := collectRCBasicInfo()
		if err != nil {
			return nil, false, err
		}
		if rcName == "" {
			return nil, false, nil
		}

		aagValue := tui.CollectIntParameter(
			"Phase 2 - Redundancy Class Definition\nStep 6.1 of 9: Anti-Affinity Group (disks per node)",
			fmt.Sprintf("How many disks per node would you like to use?\n(Min: 1, Max: %d)", limits.minDisksPerNode),
			"AAG defines the max disks we can use per node to satisfy Local N+K",
			func(v string) string {
				return ValidateAAG(parseInt(v), limits.minDisksPerNode)
			},
			func(v string) (string, string) {
				val := parseInt(v)
				if val <= 0 || val > limits.minDisksPerNode {
					return "", ValidateAAG(val, limits.minDisksPerNode)
				}
				nodesRequired := (val + limits.minNodesPerLocation - 1) / limits.minNodesPerLocation
				if limits.minLocationDisks > 0 {
					nodesRequired = (limits.minLocationDisks + val - 1) / val
				}
				return fmt.Sprintf("This will require ~%d nodes per location", nodesRequired), ""
			},
		)
		if aagValue < 0 {
			return nil, false, nil
		}

		outerKValue := tui.CollectIntParameter(
			"Phase 2 - Redundancy Class Definition\nStep 6.2 of 9: Geographical K",
			fmt.Sprintf("How many locations can you afford to lose?\n(Min: 0, Max: %d)", len(allClusterIDs)-1),
			"Data will survive loss of K locations",
			func(v string) string {
				return ValidateOuterK(parseInt(v), len(allClusterIDs))
			},
			func(v string) (string, string) {
				val := parseInt(v)
				if err := ValidateOuterK(val, len(allClusterIDs)); err != "" {
					return "", err
				}
				outerN := len(allClusterIDs) - val
				return fmt.Sprintf("Geographical N (locations needed): %d", outerN), ""
			},
		)
		if outerKValue < 0 {
			return nil, false, nil
		}

		localNKValue := tui.CollectIntParameter(
			"Phase 2 - Redundancy Class Definition\nStep 6.3 of 9: Local N+K (total disks per location)",
			fmt.Sprintf("How many disks to use on each location?\n(Min: 1, Max: %d)", limits.minLocationDisks),
			"Total number of disks used per location across all nodes",
			func(v string) string {
				return ValidateLocalNK(parseInt(v), aagValue, limits.minNodesPerLocation, limits.minLocationDisks)
			},
			func(v string) (string, string) {
				val := parseInt(v)
				if err := ValidateLocalNK(val, aagValue, limits.minNodesPerLocation, limits.minLocationDisks); err != "" {
					return "", err
				}
				nodesRequired := val / aagValue
				return fmt.Sprintf("Nodes required per location: %d / %d = %d", val, aagValue, nodesRequired), ""
			},
		)
		if localNKValue < 0 {
			return nil, false, nil
		}

		localKValue := tui.CollectIntParameter(
			"Phase 2 - Redundancy Class Definition\nStep 6.4 of 9: Local K",
			fmt.Sprintf("How many disks can you afford to lose per location?\n(Min: 0, Max: %d)", localNKValue-1),
			"Data will survive loss of K disks per location",
			func(v string) string {
				return ValidateLocalK(parseInt(v), localNKValue)
			},
			func(v string) (string, string) {
				val := parseInt(v)
				if err := ValidateLocalK(val, localNKValue); err != "" {
					return "", err
				}
				localN := localNKValue - val
				return fmt.Sprintf("Local N (data disks): %d", localN), ""
			},
		)
		if localKValue < 0 {
			return nil, false, nil
		}

		outerN := len(allClusterIDs) - outerKValue
		localN := localNKValue - localKValue

		var rcDescPtr *string
		if rcDesc != "" {
			rcDescPtr = &rcDesc
		}

		rc := api.RedundancyClassRequest{
			Name:              rcName,
			Description:       rcDescPtr,
			OuterN:            outerN,
			OuterK:            outerKValue,
			InnerN:            localN,
			InnerK:            localKValue,
			AntiAffinityGroup: aagValue,
			ClusterIDs:        allClusterIDs,
		}

		redundancyClasses = append(redundancyClasses, rc)

		rcPreview := fmt.Sprintf(
			"Name: %s\nGeographical N: %d\nGeographical K: %d\nLocal N: %d\nLocal K: %d\nAAG: %d\nLocations: %d",
			rcName, outerN, outerKValue, localN, localKValue, aagValue, len(allClusterIDs),
		)

		rcChoice, err := tui.ConfirmDialog(
			"Phase 2 - RC Preview\nStep 7 of 9: Redundancy class created",
			rcPreview+"\n\nCreate another RC or continue?",
			[]string{"Create another RC", "Continue", "Go back"},
		)
		if err != nil {
			return nil, false, err
		}
		if rcChoice == "Create another RC" {
			continue
		}
		if rcChoice == "Go back" {
			redundancyClasses = redundancyClasses[:len(redundancyClasses)-1]
			continue
		}
		break
	}

	return redundancyClasses, true, nil
}

func collectRCBasicInfo() (name string, description string, err error) {
	rcName := ""
	rcDesc := ""

	inputs, err := tui.TextInputs("Phase 2 - Redundancy Class\nStep 5 of 9: Enter redundancy class details", false,
		tui.Input{Placeholder: "RC name (required)", Value: &rcName},
		tui.Input{Placeholder: "RC description (optional)", Value: &rcDesc},
	)
	if err != nil {
		return "", "", err
	}

	if len(inputs) < 1 || len(*inputs[0].Value) == 0 {
		tui.ShowError("RC name is required")
		return "", "", nil
	}

	return *inputs[0].Value, *inputs[1].Value, nil
}

func confirmSubmission(name string, numLocations, totalNodes, numRCs int) (bool, error) {
	summary := fmt.Sprintf("Swarm: %s\nLocations: %d\nTotal Nodes: %d\nRedundancy Classes: %d\n",
		name, numLocations, totalNodes, numRCs)

	choice, err := tui.ConfirmDialog(
		"Phase 3 - Final Confirmation\nStep 8 of 9: Confirm swarm creation",
		summary+"\nProceed with swarm creation?",
		[]string{"Yes, create swarm", "No, cancel"},
	)
	if err != nil {
		return false, err
	}
	if choice != "Yes, create swarm" {
		return false, nil
	}
	return true, nil
}

func submitSwarm(
	swarmAPI api.SwarmAPIInterface,
	processAPI api.ProcessAPIInterface,
	rcValidator RedundancyClassValidatorInterface,
	profile *configuration.ResolvedProfile,
	urls *configuration.URLs,
	name string,
	description string,
	clusters []api.InfraAggregateCluster,
	nodeIDsByCluster map[string][]string,
	redundancyClasses []api.RedundancyClassRequest,
) error {
	nexuses := buildNexuses(clusters, nodeIDsByCluster)

	if err := rcValidator.ValidateRedundancyClasses(redundancyClasses, nexuses); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	var createDesc *string
	if description != "" {
		createDesc = &description
	}

	createRequest := &api.CreateSwarmV5Request{
		Name:              name,
		Description:       createDesc,
		Configuration:     make(map[string]interface{}),
		Nexuses:           nexuses,
		RedundancyClasses: redundancyClasses,
	}

	response, err := swarmAPI.CreateSwarmV5(*urls, profile.APIKey, profile.OrganizationID, createRequest)
	if err != nil {
		return fmt.Errorf("failed to create swarm: %w", err)
	}

	monitorSwarmCreation(processAPI, *urls, profile.APIKey, profile.OrganizationID, response.ID)

	return nil
}

func parseInt(s string) int {
	if s == "" {
		return 0
	}
	val := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			val = val*10 + int(c-'0')
		} else {
			return -1
		}
	}
	return val
}

func buildNexuses(clusters []api.InfraAggregateCluster, nodeIDsByCluster map[string][]string) []api.NexusV5Request {
	nexuses := make([]api.NexusV5Request, 0)
	for _, cluster := range clusters {
		nodeIDs := nodeIDsByCluster[cluster.ClusterID]

		if cluster.Type == api.ClusterTypeVirtual {
			virtualNodes := make([]api.VirtualNodeRequest, len(nodeIDs))
			for i, nid := range nodeIDs {
				virtualNodes[i] = api.VirtualNodeRequest{ServerID: nid}
			}
			nexuses = append(nexuses, api.NexusV5Request{
				ClusterID:    cluster.ClusterID,
				ClusterType:  cluster.Type,
				VirtualNodes: virtualNodes,
			})
		} else {
			nodes := make([]api.NodeRequest, 0)
			for _, node := range cluster.Details.Nodes {
				found := false
				for _, nid := range nodeIDs {
					if node.NodeID == nid {
						found = true
						break
					}
				}
				if !found {
					continue
				}

				volumes := make([]api.VolumeRequest, 0)
				for _, disk := range node.Disks {
					if disk.Status.Code == string(api.StatusCodeOk) {
						volumes = append(volumes, api.VolumeRequest{
							VolumeID: disk.DiskUUID,
						})
					}
				}
				nodes = append(nodes, api.NodeRequest{
					ServerID: node.NodeID,
					Volumes:  volumes,
				})
			}
			nexuses = append(nexuses, api.NexusV5Request{
				ClusterID:   cluster.ClusterID,
				ClusterType: cluster.Type,
				Nodes:       nodes,
			})
		}
	}
	return nexuses
}

var swarmCreationSteps = []api.ProcessStep{
	api.ProcessStepInitializing,
	api.ProcessStepCreatingNexuses,
	api.ProcessStepCreatingAgents,
	api.ProcessStepCreatingRedundancyCls,
	api.ProcessStepCompleted,
}

func stepProgress(current api.ProcessStep) int {
	for i, s := range swarmCreationSteps {
		if s == current {
			return (i * 100) / (len(swarmCreationSteps) - 1)
		}
	}
	return 0
}

func stepLabel(s api.ProcessStep) string {
	switch s {
	case api.ProcessStepInitializing:
		return "Initializing"
	case api.ProcessStepCreatingNexuses:
		return "Creating nexuses"
	case api.ProcessStepCreatingAgents:
		return "Creating agents"
	case api.ProcessStepCreatingRedundancyCls:
		return "Creating redundancy classes"
	case api.ProcessStepCompleted:
		return "Completed"
	default:
		return string(s)
	}
}

func monitorSwarmCreation(
	processAPI api.ProcessAPIInterface,
	urls configuration.URLs,
	apiKey string,
	organizationID string,
	processID string,
) {
	pollFn := func() (string, string, int, error) {
		detail, err := processAPI.GetProcess(urls, apiKey, organizationID, processID)
		if err != nil {
			return "", "", 0, err
		}

		statusStr := string(detail.Status)
		stepStr := stepLabel(detail.Step)
		progress := stepProgress(detail.Step)

		return statusStr, stepStr, progress, nil
	}

	_, err := tui.NewMonitorSession(processID, pollFn).Run()
	if err != nil {
		fmt.Println("Monitoring error:", err)
		return
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

	processes, err := s.processAPI.ListProcesses(
		*urls,
		resolvedProfile.APIKey,
		resolvedProfile.OrganizationID,
		api.WithProcessType(api.ProcessTypeSwarmCreation),
		api.WithProcessStatus(api.ProcessStatusRunning),
	)
	if err != nil {
		return fmt.Errorf("failed to check for ongoing swarm creation processes: %w", err)
	}

	if len(processes) > 0 {
		return fmt.Errorf("swarm creation already in progress")
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

func (s SwarmService) List(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	allSwarms, err := s.fetchAllSwarms(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID)
	if err != nil {
		return fmt.Errorf("failed to list swarms: %w", err)
	}

	output, err := resolveCommandOutput(cmd, resolvedProfile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration.OutputHuman) {
		return PrintSwarmList(cmd, allSwarms)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), allSwarms, output)
	return nil
}

func (s SwarmService) fetchAllSwarms(urls configuration.URLs, apiKey string, organizationID string) ([]api.ListSwarmV5ItemPresentation, error) {
	page := 1
	itemsPerPage := 100
	var all []api.ListSwarmV5ItemPresentation

	for {
		response, err := s.swarmAPI.ListSwarmsV5(urls, apiKey, organizationID, page, itemsPerPage)
		if err != nil {
			return nil, err
		}

		all = append(all, response.Data...)

		if response.NextPage == nil {
			break
		}

		page = *response.NextPage
	}

	return all, nil
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
