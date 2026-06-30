package describe

import (
	"fmt"
	"strings"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/gateway/shared"
	"github.com/cubbit/composer-cli/utils"
	printerutils "github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	TenantAPI api.TenantAPIInterface
}

func Describe(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2, args []string) error {
	tenantID, err := resolveTenantID(args)
	if err != nil {
		return err
	}

	tenant, err := deps.TenantAPI.GetTenantV5(profile.Endpoints, profile.APIKey, profile.OrganizationID, tenantID)
	if err != nil {
		return fmt.Errorf("failed to describe tenant: %w", err)
	}

	output, err := shared.ResolveCommandOutput(cmd, profile.Output)
	if err != nil {
		return err
	}

	if output == string(configuration_models.OutputHuman) {
		return PrintTenantDetails(cmd, *tenant)
	}

	utils.PrintFormattedData(cmd.OutOrStdout(), tenant, output)
	return nil
}

func resolveTenantID(args []string) (string, error) {
	if len(args) == 0 || args[0] == "" {
		return "", fmt.Errorf("tenant ID is required")
	}

	return args[0], nil
}

func formatBytesDetailPtr(project *api.TenantSettingsProject) string {
	if project == nil || project.DefaultMaxProjectStorageGB == nil {
		return "none"
	}
	return printerutils.FormatBytes(*project.DefaultMaxProjectStorageGB)
}

func formatBandwidthDetailPtr(project *api.TenantSettingsProject) string {
	if project == nil || project.DefaultMaxProjectEgressBandwidthGB == nil {
		return "none"
	}
	return printerutils.FormatBytes(*project.DefaultMaxProjectEgressBandwidthGB)
}

func formatMaxProjectPtr(account *api.TenantSettingsAccount) string {
	if account == nil || account.DefaultMaxProject == nil {
		return "none"
	}
	return fmt.Sprintf("%d", *account.DefaultMaxProject)
}

func formatAuthProvidersPtr(account *api.TenantSettingsAccount) string {
	if account == nil || account.EnabledAuthProviders == nil || len(*account.EnabledAuthProviders) == 0 {
		return "none"
	}
	providers := make([]string, len(*account.EnabledAuthProviders))
	for i, p := range *account.EnabledAuthProviders {
		providers[i] = string(p)
	}
	return strings.Join(providers, ", ")
}
