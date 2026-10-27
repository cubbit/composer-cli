package create

import (
	"errors"
	"fmt"
	"strings"
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/interactive"
	"github.com/cubbit/composer-cli/utils/interactive/tui/input"
	"github.com/spf13/cobra"
)

func createInteractive(deps Dependencies, cmd *cobra.Command, endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) error {
	createRequest, err := promptForTenantCreate(deps, cmd, endpoints, apiKey, organizationID)
	if err != nil {
		if errors.Is(err, interactive.ErrCancelled) {
			utils.PrintErrorWithWriter(cmd.ErrOrStderr(), fmt.Errorf("command was cancelled"))
			return nil
		}
		return err
	}

	response, err := deps.TenantAPI.CreateTenantV5(endpoints, apiKey, organizationID, createRequest)
	if err != nil {
		return fmt.Errorf("failed to create tenant: %w", err)
	}

	return waitForTenantDeployment(deps, cmd, endpoints, apiKey, organizationID, response.ID)
}

var tenantStepOrder = []api.ProcessStep{
	api.ProcessStepInitializing,
	api.ProcessStepCreatingTenantGateway,
	api.ProcessStepWaitingForTenantGatewayToBeReady,
	api.ProcessStepCompleted,
}

var tenantStepMessages = map[api.ProcessStep]string{
	api.ProcessStepInitializing:                     "Initializing tenant creation...",
	api.ProcessStepCreatingTenantGateway:            "Creating tenant gateways...",
	api.ProcessStepWaitingForTenantGatewayToBeReady: "Waiting for tenant gateways to be ready...",
	api.ProcessStepCompleted:                        "Tenant created successfully",
}

func tenantStepProgress(step api.ProcessStep) float64 {
	denom := len(tenantStepOrder) - 1
	if denom <= 0 {
		denom = 1
	}
	for i, s := range tenantStepOrder {
		if s == step {
			return float64(i) / float64(denom)
		}
	}
	return 0
}

func waitForTenantDeployment(deps Dependencies, cmd *cobra.Command, endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, processID string) error {
	ic := interactive.New(interactive.Config{
		Stdout: cmd.OutOrStdout(),
		Stdin:  cmd.InOrStdin(),
	})

	h := ic.StartProgress("Creating tenant")

	pollTicker := time.NewTicker(3 * time.Second)
	defer pollTicker.Stop()

	timeout := time.After(1 * time.Hour)

	for {
		select {
		case <-pollTicker.C:
			process, err := deps.ProcessAPI.GetProcess(endpoints, apiKey, organizationID, processID)
			if err != nil {
				h.Stop()
				return fmt.Errorf("failed to poll tenant creation status: %w", err)
			}

			processIdentification := ""
			tcProcess, ok := process.CastToTenantCreationProcess()
			if ok {
				if tcProcess.Data.Error != nil {
					processIdentification = fmt.Sprintf(", Error %s: %s", tcProcess.Data.Error.Code, tcProcess.Data.Error.Message)
				}
			}

			switch process.Status {
			case api.ProcessStatusRunning:
				msg, ok := tenantStepMessages[process.Step]
				if !ok {
					msg = string(process.Step)
				}
				if msg == "" {
					msg = "Creating tenant..."
				}
				msg = fmt.Sprintf("%s%s", msg, processIdentification)
				h.Set(tenantStepProgress(process.Step), msg)

			case api.ProcessStatusSuccess:
				h.Set(1.0, "Tenant created successfully")
				h.Stop()

				tenantID := ""
				if ok && tcProcess.Data.TenantID != "" {
					tenantID = tcProcess.Data.TenantID
				}

				ic.Success(fmt.Sprintf("Tenant created successfully — Tenant ID: %s", tenantID))
				return nil

			case api.ProcessStatusFailed:
				h.Stop()
				if processIdentification != "" {
					return fmt.Errorf("tenant creation failed%s", processIdentification)
				}
				return fmt.Errorf("tenant creation failed — step: %s", process.Step)

			default:
				h.Set(tenantStepProgress(process.Step), string(process.Status))
			}

		case <-timeout:
			h.Stop()
			return fmt.Errorf("tenant creation timed out after 1 hour")
		}
	}
}

func promptForTenantCreate(
	deps Dependencies,
	cmd *cobra.Command,
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
) (*api.CreateTenantV5Request, error) {
	ic := interactive.New(interactive.Config{
		Stdout: cmd.OutOrStdout(),
		Stdin:  cmd.InOrStdin(),
	})

	name, slug, description, err := promptBasicInfo(ic)
	if err != nil {
		return nil, err
	}

	domains, err := fetchVerifiedDomains(ic, deps, endpoints, apiKey, organizationID)
	if err != nil {
		return nil, err
	}

	selectedDomainIDs, domainByID, err := selectDomains(ic, domains)
	if err != nil {
		return nil, err
	}

	gatewayNames, err := fetchGatewayNames(ic, deps, endpoints, apiKey, organizationID)
	if err != nil {
		return nil, err
	}

	connections, err := promptConnections(ic, selectedDomainIDs, domainByID, gatewayNames)
	if err != nil {
		return nil, err
	}

	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}

	return &api.CreateTenantV5Request{
		Name:        name,
		Slug:        slug,
		Description: descriptionPtr,
		Connections: connections,
	}, nil
}

func promptBasicInfo(ic *interactive.Interactive) (string, string, string, error) {
	name, err := ic.Input("Tenant name",
		input.WithRequired(),
	)
	if err != nil {
		return "", "", "", fmt.Errorf("cancelled: %w", err)
	}

	slug, err := ic.Input("Tenant slug",
		input.WithRequired(),
	)
	if err != nil {
		return "", "", "", fmt.Errorf("cancelled: %w", err)
	}

	description, err := ic.Input("Description (optional)")
	if err != nil {
		return "", "", "", fmt.Errorf("cancelled: %w", err)
	}

	return name, slug, description, nil
}

func fetchVerifiedDomains(
	ic *interactive.Interactive,
	deps Dependencies,
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
) ([]api.DomainDTO, error) {
	h := ic.StartSpinner("Fetching verified domains...")

	var domains []api.DomainDTO
	page := 1
	itemsPerPage := 1000

	for {
		response, err := deps.DomainAPI.List(endpoints, apiKey, organizationID, page, itemsPerPage)
		if err != nil {
			h.Stop()
			return nil, fmt.Errorf("failed to list domains: %w", err)
		}

		for _, d := range response.Data {
			if d.VerifiedAt != nil {
				domains = append(domains, d)
			}
		}

		if response.NextPage == nil {
			break
		}

		page = *response.NextPage
	}

	h.Stop()

	if len(domains) == 0 {
		return nil, fmt.Errorf("no verified domains available for this organization")
	}

	return domains, nil
}

func selectDomains(ic *interactive.Interactive, domains []api.DomainDTO) ([]string, map[string]api.DomainDTO, error) {
	domainNames := make([]string, len(domains))
	domainByID := make(map[string]api.DomainDTO, len(domains))

	for i, d := range domains {
		domainNames[i] = fmt.Sprintf("%s (%s)", d.DomainName, d.ID)
		domainByID[d.ID] = d
	}

	selected, err := ic.MultiSelect("Select domains", domainNames)
	if err != nil {
		return nil, nil, fmt.Errorf("cancelled: %w", err)
	}

	if len(selected) == 0 {
		return nil, nil, fmt.Errorf("at least one domain must be selected")
	}

	return selected, domainByID, nil
}

func fetchGatewayNames(
	ic *interactive.Interactive,
	deps Dependencies,
	endpoints configuration_models.EndpointsV2,
	apiKey string,
	organizationID string,
) ([]string, error) {
	h := ic.StartSpinner("Fetching available gateways...")

	response, err := deps.GatewayAPI.ListGatewaysV5(endpoints, apiKey, organizationID, api.WithListAllPages(true))
	h.Stop()
	if err != nil {
		return nil, fmt.Errorf("failed to list gateways: %w", err)
	}

	if len(response.Data) == 0 {
		return nil, fmt.Errorf("no gateways available for the organization")
	}

	gatewayNames := make([]string, len(response.Data))
	for i, g := range response.Data {
		gatewayNames[i] = fmt.Sprintf("%s (%s)", g.Name, g.ID)
	}

	return gatewayNames, nil
}

func promptConnections(
	ic *interactive.Interactive,
	selectedDomains []string,
	domainByID map[string]api.DomainDTO,
	gatewayNames []string,
) ([]api.CreateTenantV5Connection, error) {
	connections := make([]api.CreateTenantV5Connection, 0, len(selectedDomains))

	for _, sel := range selectedDomains {
		domainID := extractID(sel)
		domain := domainByID[domainID]

		selectedGateways, err := ic.MultiSelect(
			fmt.Sprintf("Select gateways for domain %s", domain.DomainName),
			gatewayNames,
		)
		if err != nil {
			return nil, fmt.Errorf("cancelled: %w", err)
		}

		if len(selectedGateways) == 0 {
			return nil, fmt.Errorf("at least one gateway must be selected for domain %s", domain.DomainName)
		}

		gatewayIDs := make([]string, 0, len(selectedGateways))
		for _, s := range selectedGateways {
			gatewayIDs = append(gatewayIDs, extractID(s))
		}

		subdomain, err := ic.Input(fmt.Sprintf("Subdomain for domain %s (optional)", domain.DomainName))
		if err != nil {
			return nil, fmt.Errorf("cancelled: %w", err)
		}

		conn := api.CreateTenantV5Connection{
			DomainID:   domainID,
			GatewayIDs: gatewayIDs,
		}

		if subdomain != "" {
			conn.Subdomain = &subdomain
		}

		connections = append(connections, conn)
	}

	if len(connections) == 0 {
		return nil, fmt.Errorf("at least one connection is required")
	}

	return connections, nil
}

func extractID(s string) string {
	if idx := strings.LastIndex(s, " ("); idx != -1 {
		return s[idx+2 : len(s)-1]
	}
	return s
}
