package listconnections

import (
	"fmt"
	"sync"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/gateway/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	ConnectionAPI api.ConnectionAPIInterface
	TenantAPI     api.TenantAPIInterface
}

func ListConnections(deps Dependencies, cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, profile configuration_models.ProfileV2, args []string) error {
	tenantID, err := resolveTenantID(deps, cmd, profile)
	if err != nil {
		return err
	}

	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		return fmt.Errorf("%s page: %w", constants.ErrorRetrievingField, err)
	}

	items, err := cmd.Flags().GetInt("items")
	if err != nil {
		return fmt.Errorf("%s items: %w", constants.ErrorRetrievingField, err)
	}

	opts := buildListOpts(cmd, page, items)

	connections, err := deps.ConnectionAPI.ListConnectionsV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		tenantID,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("failed to list tenant connections: %w", err)
	}

	skipVerify, err := cmd.Flags().GetBool("skip-verify")
	if err != nil {
		return fmt.Errorf("%s skip-verify: %w", constants.ErrorRetrievingField, err)
	}

	if !skipVerify && len(connections.Data) > 0 {
		verifiedMap, warnings := verifyConnections(deps, cmd, profile, tenantID, connections.Data)

		for _, w := range warnings {
			fmt.Fprintln(cmd.ErrOrStderr(), w)
		}

		connections.Data = mergeVerifiedConnections(connections.Data, verifiedMap)
	}

	output, err := shared.ResolveCommandOutput(cmd, profile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration_models.OutputHuman) {
		return PrintConnectionsList(cmd, handler, connections.Data)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), connections.Data, output)
	return nil
}

func buildListOpts(cmd *cobra.Command, page, items int) []api.ListConnectionsV5Option {
	var opts []api.ListConnectionsV5Option
	if cmd.Flags().Changed("page") || cmd.Flags().Changed("items") {
		if cmd.Flags().Changed("page") {
			opts = append(opts, api.WithConnectionsPage(page))
		}
		if cmd.Flags().Changed("items") {
			opts = append(opts, api.WithConnectionsItems(items))
		}
		opts = append(opts, api.WithConnectionsFetchAllPages(false))
	} else {
		opts = append(opts, api.WithConnectionsFetchAllPages(true))
	}

	return opts
}

const maxConcurrentVerifications = 5

func verifyConnections(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2, tenantID string, connections []api.ConnectionV5DTO) (map[string]api.ConnectionV5DTO, []string) {
	verifiedByID := make(map[string]api.ConnectionV5DTO, len(connections))
	var warnings []string
	var mu sync.Mutex

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrentVerifications)
	errCh := make(chan string, len(connections))

	for _, conn := range connections {
		wg.Add(1)
		go func(conn api.ConnectionV5DTO) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			vc, err := deps.ConnectionAPI.VerifyConnectionV5(
				profile.Endpoints,
				profile.APIKey,
				profile.OrganizationID,
				tenantID,
				conn.ID,
			)
			if err != nil {
				errCh <- fmt.Sprintf("connection %s (%s) verification failed: %s", conn.Domain.DomainName, conn.ID, err)
				return
			}

			if vc != nil {
				mu.Lock()
				verifiedByID[conn.ID] = *vc
				mu.Unlock()
			}
		}(conn)
	}

	wg.Wait()
	close(errCh)

	for msg := range errCh {
		warnings = append(warnings, msg)
	}

	return verifiedByID, warnings
}

func mergeVerifiedConnections(original []api.ConnectionV5DTO, verified map[string]api.ConnectionV5DTO) []api.ConnectionV5DTO {
	merged := make([]api.ConnectionV5DTO, len(original))
	for i, conn := range original {
		if vc, ok := verified[conn.ID]; ok {
			merged[i] = vc
		} else {
			merged[i] = conn
		}
	}
	return merged
}

func resolveTenantID(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2) (string, error) {
	tenantIDFlag, err := cmd.Flags().GetString("tenant-id")
	if err != nil {
		return "", fmt.Errorf("%s tenant-id: %w", constants.ErrorRetrievingField, err)
	}

	tenantNameFlag, err := cmd.Flags().GetString("tenant-name")
	if err != nil {
		return "", fmt.Errorf("%s tenant-name: %w", constants.ErrorRetrievingField, err)
	}

	if tenantIDFlag != "" && tenantNameFlag != "" {
		return "", fmt.Errorf("specify either --tenant-id or --tenant-name, not both")
	}

	if tenantIDFlag != "" {
		return tenantIDFlag, nil
	}

	if tenantNameFlag != "" {
		return resolveTenantIDByName(deps, profile, tenantNameFlag)
	}

	return "", fmt.Errorf("specify either --tenant-id or --tenant-name")
}

func resolveTenantIDByName(deps Dependencies, profile configuration_models.ProfileV2, tenantName string) (string, error) {
	const itemsPerPage = 1000
	page := 1

	for {
		response, err := deps.TenantAPI.ListTenantsV5(
			profile.Endpoints,
			profile.APIKey,
			profile.OrganizationID,
			api.WithTenantPage(page),
			api.WithTenantItems(itemsPerPage),
		)
		if err != nil {
			return "", fmt.Errorf("failed to resolve tenant name '%s': %w", tenantName, err)
		}

		for _, t := range response.Data {
			if t.Name == tenantName {
				return t.ID, nil
			}
		}

		if response.NextPage == nil {
			break
		}
		page = *response.NextPage
	}

	return "", fmt.Errorf("tenant with name '%s' not found", tenantName)
}
