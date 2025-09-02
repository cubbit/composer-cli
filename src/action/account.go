// Package action provides CLI actions for managing tenant accounts.
package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func ListTenantAccounts(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
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

	if accounts, err = api.ListTenantAccounts(*urls, resolvedProfile.APIKey, tenantID, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		accounts.Data,
		func(a *api.Account) []string {
			return []string{
				a.ID,
			}
		},
		&utils.SmartOutputConfig[*api.Account]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func DescribeTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var account *api.Account

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if account, err = api.GetTenantAccount(*urls, resolvedProfile.APIKey, tenantID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.Account{account},
		func(a *api.Account) []string {
			return []string{
				a.ID,
				a.FirstName,
				a.LastName,
			}
		},
		&utils.SmartOutputConfig[*api.Account]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func RemoveTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.RemoveTenantAccount(*urls, resolvedProfile.APIKey, tenantID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantAccountRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{userID},
		func(userID string) []string {
			return []string{
				userID,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func BanTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.ToggleBanAccount(*urls, resolvedProfile.APIKey, tenantID, userID, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorFreezingTenantAccountRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{userID},
		func(userID string) []string {
			return []string{
				userID,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func UnbanTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.ToggleBanAccount(*urls, resolvedProfile.APIKey, tenantID, userID, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUnfreezingTenantAccountRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{userID},
		func(userID string) []string {
			return []string{
				userID,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func RestoreTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.RestoreTenantAccount(*urls, resolvedProfile.APIKey, tenantID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRestoringTenantAccountRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{userID},
		func(userID string) []string {
			return []string{
				userID,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func DeleteTenantAccountSessions(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.DeleteTenantAccountSessions(*urls, resolvedProfile.APIKey, tenantID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantAccountSessionsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{userID},
		func(userID string) []string {
			return []string{
				userID,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func UpdateTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID, firstName, lastName, endpointGateway string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var internal bool
	var maxAllowedProjects int

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if firstName, err = cmd.Flags().GetString("first-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if lastName, err = cmd.Flags().GetString("last-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if internal, err = cmd.Flags().GetBool("internal"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if maxAllowedProjects, err = cmd.Flags().GetInt("max-allowed-projects"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if endpointGateway, err = cmd.Flags().GetString("endpoint-gateway"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	internalPtr := &internal
	isChanged := cmd.Flags().Changed("internal")
	if !isChanged {
		internalPtr = nil
	}

	maxAllowedProjectsPtr := &maxAllowedProjects
	isChanged = cmd.Flags().Changed("max-allowed-projects")
	if !isChanged {
		maxAllowedProjectsPtr = nil
	}

	var firstNamePtr *string
	if firstName != "" {
		firstNamePtr = &firstName
	}

	var lastNamePtr *string
	if lastName != "" {
		lastNamePtr = &lastName
	}

	var endpointGatewayPtr *string
	if endpointGateway != "" {
		endpointGatewayPtr = &endpointGateway
	}

	requestBody := api.UpdateAccountRequest{
		FirstName:          firstNamePtr,
		LastName:           lastNamePtr,
		Internal:           internalPtr,
		EndpointGateway:    endpointGatewayPtr,
		MaxAllowedProjects: maxAllowedProjectsPtr,
	}

	if err = api.UpdateAccount(*urls, resolvedProfile.APIKey, tenantID, userID, requestBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUpdatingTenantAccountRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.UpdateAccountRequest{requestBody},
		func(req api.UpdateAccountRequest) []string {
			return []string{
				userID,
			}
		},
		&utils.SmartOutputConfig[api.UpdateAccountRequest]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		})
}

func CreateTenantAccounts(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var emails []string

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if emails, err = cmd.Flags().GetStringSlice("emails"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.CreateTenantAccounts(*urls, resolvedProfile.APIKey, tenantID, emails); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenantAccountsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		emails,
		func(email string) []string {
			return []string{email}
		},
		&utils.SmartOutputConfig[string]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}
