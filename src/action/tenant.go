// Package action provides CLI actions for managing tenants.
package action

import (
	"encoding/json"
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func CreateTenant(cmd *cobra.Command, args []string) error {
	var err error
	var name, description, couponCode, settingsStr, zoneStr string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var zone *string
	var response *api.GenericIDResponseModel
	var zones *api.ZoneMap
	var settings *api.TenantSettings

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if settingsStr, err = cmd.Flags().GetString("settings"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if couponCode, err = cmd.Flags().GetString("distributor-code"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if zoneStr, err = cmd.Flags().GetString("zone"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if zoneStr != "" {
		if zones, err = api.GetGatewayZones(*urls); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingZonesRequest, err)
		}
		if _, ok := zones.Zones[zoneStr]; !ok {
			return fmt.Errorf(constants.ErrorInvalidZone)
		}

		zone = &zoneStr
	}

	if settingsStr != "" {
		if err = json.Unmarshal([]byte(settingsStr), &settings); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorParsingJSONSettings, err)
		}
	}

	req := api.CreateTenantRequestBody{
		Name:        name,
		Description: &description,
		Settings:    settings,
		CouponCode:  &couponCode,
		Zone:        zone,
	}

	if response, err = api.CreateTenant(*urls, resolvedProfile.APIKey, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenantRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.CreateTenantRequestBody{&req},
		func(t *api.CreateTenantRequestBody) []string {
			return []string{response.ID}
		},
		&utils.SmartOutputConfig[*api.CreateTenantRequestBody]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func ListTenant(cmd *cobra.Command, args []string) error {
	var err error
	var sort, filter string
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var tenants *api.GenericPaginatedResponse[*api.Tenant]
	var conf *configuration.Config

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

	if tenants, err = api.ListTenants(*urls, resolvedProfile.APIKey, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		tenants.Data,
		func(t *api.Tenant) []string {
			return []string{t.ID}
		},
		&utils.SmartOutputConfig[*api.Tenant]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func RemoveTenant(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.RemoveTenant(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{tenantID},
		func(s string) []string {
			return []string{s}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func DescribeTenant(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var tenant *api.Tenant

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if tenant, err = api.GetTenant(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.Tenant{tenant},
		func(t *api.Tenant) []string {
			return []string{t.ID, t.Name, utils.StringOrEmpty(t.Description)}
		},
		&utils.SmartOutputConfig[*api.Tenant]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func EditTenant(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, swarmID, redundancyClassID, description, settingsStr string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var isDefault bool
	var settings api.TenantSettings
	var req api.UpdateTenantRequestBody

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if settingsStr, err = cmd.Flags().GetString("settings"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassID, err = cmd.Flags().GetString("rc-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if isDefault, err = cmd.Flags().GetBool("default"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if settingsStr != "" {
		if err = json.Unmarshal([]byte(settingsStr), &settings); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorParsingJSONSettings, err)
		}
	}

	if swarmID != "" {
		if err = api.ConnectSwarm(*urls, resolvedProfile.APIKey, tenantID, swarmID, redundancyClassID, isDefault); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorConnectingSwarmRequest, err)
		}

		return utils.PrintSmartOutput(
			cmd,
			[]string{tenantID},
			func(s string) []string {
				return []string{s}
			},
			&utils.SmartOutputConfig[string]{
				SingleResource: true,
				DefaultOutput:  resolvedProfile.Output,
			},
		)
	}

	if description != "" {
		req.Description = &description
	}

	if settingsStr != "" {
		req.Settings = &settings
	}

	if req.Description == nil && req.Settings == nil {
		return nil
	}

	if err = api.EditTenant(*urls, resolvedProfile.APIKey, tenantID, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.UpdateTenantRequestBody{&req},
		func(t *api.UpdateTenantRequestBody) []string {
			return []string{tenantID}
		},
		&utils.SmartOutputConfig[*api.UpdateTenantRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func GetTenantReport(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, from, to, outputDir string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var fileName *string
	var tenantReport *api.TenantReportResponseModel

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if from, err = cmd.Flags().GetString("from"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if to, err = cmd.Flags().GetString("to"); err != nil {
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

	if outputDir != "" {
		if fileName, err = api.DownloadTenantReport(*urls, resolvedProfile.APIKey, tenantID, from, to, outputDir); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorDownloadingTenantReportRequest, err)
		}

		utils.PrintSuccess(fmt.Sprintf("report downloaded successfully to : %s", *fileName))
		return nil
	}

	if tenantReport, err = api.GetTenantReport(*urls, resolvedProfile.APIKey, tenantID, from, to); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDownloadingTenantReportRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.TenantReportResponseModel{tenantReport},
		func(tr *api.TenantReportResponseModel) []string {
			return []string{}
		},
		&utils.SmartOutputConfig[*api.TenantReportResponseModel]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func ConfigureTenantDNS(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, domain string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var tenant *api.Tenant
	var force bool

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if domain, err = cmd.Flags().GetString("domain"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if force, err = cmd.Flags().GetBool("force"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if tenant, err = api.GetTenant(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	if tenant.Settings != nil {
		if tenant.Settings.WhiteLabel != nil && tenant.Settings.WhiteLabel.DNS != nil && tenant.Settings.WhiteLabel.DNS.Challenge != "" && !force {
			return fmt.Errorf("%s: %w", constants.ErrorTenantDNSAlreadyConfigured, fmt.Errorf("domain %s is already configured for tenant %s, add --force to override", tenant.Settings.WhiteLabel.DNS.Value, tenant.Name))

		}
	}

	tenantSettings := api.TenantSettings{
		WhiteLabel: &api.WhiteLabel{
			DNS: &api.WhiteLabelDNS{
				Value: domain,
			},
		},
	}

	if err = api.EditTenantSettings(*urls, resolvedProfile.APIKey, tenantID, tenantSettings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorWhileConfiguringTenantDNS, err)
	}

	if tenant, err = api.GetTenant(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	type ChallengeOutput struct {
		Challenge string `json:"_acme-challenge"`
	}
	challenge := ChallengeOutput{
		Challenge: tenant.Settings.WhiteLabel.DNS.Challenge,
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*ChallengeOutput{&challenge},
		func(c *ChallengeOutput) []string {
			return []string{
				c.Challenge,
			}
		},
		&utils.SmartOutputConfig[*ChallengeOutput]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func VerifyTenantDNS(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.VerifyDNS(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorWhileVerifyingTenantDNS, err)
	}

	return nil
}
