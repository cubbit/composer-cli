package create

import (
	"errors"
	"fmt"
	"strings"
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/interactive"
	"github.com/cubbit/composer-cli/utils/interactive/tui/input"
	"github.com/spf13/cobra"
)

func createInteractive(deps Dependencies, cmd *cobra.Command, resolvedProfile configuration.ResolvedProfile, urls configuration.URLs) error {
	createRequest, err := promptForGatewayCreate(deps, cmd, resolvedProfile, urls)
	if err != nil {
		if errors.Is(err, interactive.ErrCancelled) {
			utils.PrintErrorWithWriter(cmd.ErrOrStderr(), fmt.Errorf("command was cancelled"))
			return nil
		}
		return err
	}

	response, err := deps.GatewayAPI.CreateGatewayV5(urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, createRequest)
	if err != nil {
		return fmt.Errorf("failed to create gateway: %w", err)
	}

	return waitForGatewayDeployment(deps, cmd, urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, response.ID)
}

var gatewayStepOrder = []api.ProcessStep{
	api.ProcessStepInitializing,
	api.ProcessStepGatewayProfileDeployment,
	api.ProcessStepGatewayInstallation,
	api.ProcessStepCompleted,
}

var gatewayStepMessages = map[api.ProcessStep]string{
	api.ProcessStepInitializing:             "Initializing gateway deployment...",
	api.ProcessStepGatewayProfileDeployment: "Deploying gateway profile...",
	api.ProcessStepGatewayInstallation:      "Installing gateway...",
	api.ProcessStepCompleted:                "Gateway deployed successfully",
}

func stepProgress(step api.ProcessStep) float64 {
	for i, s := range gatewayStepOrder {
		if s == step {
			return float64(i) / float64(len(gatewayStepOrder))
		}
	}
	return 0
}

func waitForGatewayDeployment(deps Dependencies, cmd *cobra.Command, urls configuration.URLs, apiKey string, organizationID string, processID string) error {
	ic := interactive.New(interactive.Config{
		Stdout: cmd.OutOrStdout(),
		Stdin:  cmd.InOrStdin(),
	})

	h := ic.StartProgress("Deploying gateway")

	pollTicker := time.NewTicker(3 * time.Second)
	defer pollTicker.Stop()

	timeout := time.After(1 * time.Hour)

	for {
		select {
		case <-pollTicker.C:
			process, err := deps.ProcessAPI.GetProcess(urls, apiKey, organizationID, processID)
			if err != nil {
				h.Stop()
				return fmt.Errorf("failed to poll gateway deployment status: %w", err)
			}

			processIdentification := ""
			gcProcess, ok := process.CastToGatewayCreationProcess()
			if ok {
				processIdentification = fmt.Sprintf("Gateway ID: %s", gcProcess.Data.ID)
				if gcProcess.Data.Error != nil {
					processIdentification += fmt.Sprintf(", Error %s: %s", gcProcess.Data.Error.Code, gcProcess.Data.Error.Message)
				}
			}

			switch process.Status {
			case api.ProcessStatusRunning:
				msg, ok := gatewayStepMessages[process.Step]
				if !ok {
					msg = string(process.Step)
				}
				if msg == "" {
					msg = "Deploying gateway..."
				}
				msg = fmt.Sprintf("%s %s", msg, processIdentification)
				h.Set(stepProgress(process.Step), msg)

			case api.ProcessStatusSuccess:
				h.Set(1.0, "Gateway deployed successfully")
				h.Stop()
				ic.Success(fmt.Sprintf("Gateway deployed — %s", processIdentification))
				return nil

			case api.ProcessStatusFailed:
				h.Stop()
				return fmt.Errorf("gateway deployment failed — %s", processIdentification)

			default:
				h.Set(stepProgress(process.Step), string(process.Status))
			}

		case <-timeout:
			h.Stop()
			return fmt.Errorf("gateway deployment timed out after 1 hour")
		}
	}
}

func promptForGatewayCreate(
	deps Dependencies,
	cmd *cobra.Command,
	resolvedProfile configuration.ResolvedProfile,
	urls configuration.URLs,
) (*api.CreateGatewayV5Request, error) {
	ic := interactive.New(interactive.Config{
		Stdout: cmd.OutOrStdout(),
		Stdin:  cmd.InOrStdin(),
	})

	name, err := ic.Input("Gateway name",
		input.WithRequired(),
	)
	if err != nil {
		return nil, fmt.Errorf("cancelled: %w", err)
	}

	slug, err := ic.Input("Gateway slug",
		input.WithRequired(),
	)
	if err != nil {
		return nil, fmt.Errorf("cancelled: %w", err)
	}

	description, err := ic.Input("Description (optional)")
	if err != nil {
		return nil, fmt.Errorf("cancelled: %w", err)
	}

	var clusters []api.InfrastructureCluster
	{
		h := ic.StartSpinner("Fetching available clusters...")
		clusters, err = deps.LocationAPI.List(
			urls,
			resolvedProfile.APIKey,
			resolvedProfile.OrganizationID,
			api.WithProfileType("not-it", api.LocationProfileGateway),
		)
		h.Stop()
		if err != nil {
			return nil, fmt.Errorf("failed to list clusters: %w", err)
		}
	}

	if len(clusters) == 0 {
		return nil, fmt.Errorf("no clusters available")
	}

	clusterNames := make([]string, len(clusters))
	for i, cl := range clusters {
		clusterNames[i] = fmt.Sprintf("%s (%s)", cl.Name, cl.ClusterID)
	}

	selectedCluster, err := ic.Select("Select cluster", clusterNames)
	if err != nil {
		return nil, fmt.Errorf("cancelled: %w", err)
	}

	clusterID := extractID(selectedCluster)

	ingressType, err := ic.Select("Ingress type", []string{
		string(api.IngressTypeManual),
	})
	if err != nil {
		return nil, fmt.Errorf("cancelled: %w", err)
	}

	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}

	var allSwarms []api.ListSwarmV5ItemPresentation
	{
		h := ic.StartSpinner("Fetching available swarms...")
		allSwarms, err = fetchAllSwarms(deps, urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID)
		h.Stop()
		if err != nil {
			return nil, fmt.Errorf("failed to list swarms: %w", err)
		}
	}

	if len(allSwarms) == 0 {
		return nil, fmt.Errorf("no swarms available")
	}

	swarmNames := make([]string, len(allSwarms))
	for i, sw := range allSwarms {
		swarmNames[i] = fmt.Sprintf("%s (%s)", sw.Name, sw.ID)
	}

	selectedSwarms, err := ic.MultiSelect("Select swarms for the gateway", swarmNames)
	if err != nil {
		return nil, fmt.Errorf("cancelled: %w", err)
	}

	selectedIDs := make([]string, 0, len(selectedSwarms))
	for _, s := range selectedSwarms {
		id := extractID(s)
		selectedIDs = append(selectedIDs, id)
	}

	swarmNamesByID := make(map[string]string, len(allSwarms))
	for _, sw := range allSwarms {
		swarmNamesByID[sw.ID] = sw.Name
	}

	swarmsAndRC, err := promptForSwarmsAndRedundancyClasses(deps, ic, urls, resolvedProfile, selectedIDs, swarmNamesByID)
	if err != nil {
		return nil, err
	}

	return &api.CreateGatewayV5Request{
		ClusterID:                clusterID,
		Name:                     name,
		Slug:                     slug,
		Description:              descriptionPtr,
		SwarmsAndRedundancyClass: swarmsAndRC,
		CubbitIngress: api.CubbitIngress{
			Type: api.CubbitIngressType(ingressType),
		},
	}, nil
}

type swarmRCSelection struct {
	SwarmID           string
	RedundancyClassID string
	SwarmName         string
	RCName            string
	isDefault         bool
}

func promptForSwarmsAndRedundancyClasses(
	deps Dependencies,
	ic *interactive.Interactive,
	urls configuration.URLs,
	resolvedProfile configuration.ResolvedProfile,
	swarmIDs []string,
	swarmNamesByID map[string]string,
) ([]api.SwarmAndRedundancyClassV5, error) {
	var selections []swarmRCSelection

	for _, swarmID := range swarmIDs {
		sh := ic.StartSpinner(fmt.Sprintf("Fetching redundancy classes for swarm %s...", swarmID))
		rcItems, err := deps.RedundancyClassAPI.ListRedundancyClassesBySwarm(urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, swarmID)
		sh.Stop()
		if err != nil {
			return nil, fmt.Errorf("failed to list redundancy classes for swarm %s: %w", swarmID, err)
		}

		if len(rcItems) == 0 {
			return nil, fmt.Errorf("no redundancy classes available for swarm %s", swarmID)
		}

		swarmName := swarmNamesByID[swarmID]
		if swarmName == "" {
			swarmName = swarmID
		}

		rcNames := make([]string, len(rcItems))
		rcNameByID := make(map[string]string, len(rcItems))
		for j, rc := range rcItems {
			label := fmt.Sprintf("%s (%s)", rc.Name, rc.ID)
			rcNames[j] = label
			rcNameByID[rc.ID] = rc.Name
		}

		selectedRCs, err := ic.MultiSelect(
			fmt.Sprintf("Select redundancy class(es) for swarm %s", swarmName),
			rcNames,
		)
		if err != nil {
			return nil, fmt.Errorf("cancelled: %w", err)
		}

		for _, s := range selectedRCs {
			rcID := extractID(s)
			selections = append(selections, swarmRCSelection{
				SwarmID:           swarmID,
				RedundancyClassID: rcID,
				SwarmName:         swarmName,
				RCName:            rcNameByID[rcID],
			})
		}
	}

	if len(selections) == 0 {
		return nil, fmt.Errorf("at least one swarm-RC pair is required")
	}

	if len(selections) == 1 {
		selections[0].isDefault = true
	} else {
		defaultNames := make([]string, len(selections))
		for i, s := range selections {
			defaultNames[i] = fmt.Sprintf("Swarm: %s, RC: %s", s.SwarmName, s.RCName)
		}

		chosen, err := ic.Select("Select the default swarm-RC pair", defaultNames)
		if err != nil {
			return nil, fmt.Errorf("cancelled: %w", err)
		}

		for i, n := range defaultNames {
			if n == chosen {
				selections[i].isDefault = true
				break
			}
		}
	}

	swarmsAndRC := make([]api.SwarmAndRedundancyClassV5, len(selections))
	for i, s := range selections {
		swarmsAndRC[i] = api.SwarmAndRedundancyClassV5{
			SwarmID:           s.SwarmID,
			RedundancyClassID: s.RedundancyClassID,
			IsDefault:         s.isDefault,
		}
	}

	return swarmsAndRC, nil
}

func fetchAllSwarms(deps Dependencies, urls configuration.URLs, apiKey string, organizationID string) ([]api.ListSwarmV5ItemPresentation, error) {
	page := 1
	itemsPerPage := 100
	var all []api.ListSwarmV5ItemPresentation

	for {
		response, err := deps.SwarmAPI.ListSwarmsV5(urls, apiKey, organizationID, page, itemsPerPage)
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

func extractID(s string) string {
	if idx := strings.LastIndex(s, " ("); idx != -1 {
		return s[idx+2 : len(s)-1]
	}
	return s
}
