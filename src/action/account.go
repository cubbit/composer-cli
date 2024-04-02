package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateAccount(cmd *cobra.Command, args []string) error {
	var url *configuration.Url
	var err error
	var email, password, firstName, lastName, apiServerUrl, tenantID string

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if password, err = cmd.Flags().GetString("password"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if firstName, err = cmd.Flags().GetString("first-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if lastName, err = cmd.Flags().GetString("last-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if apiServerUrl, err = cmd.Flags().GetString("api-server-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if url, err = configuration.ConfigureAPIServerURL(configuration.SessionTypeAccount, apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if err = api.CreateAccount(*url, firstName, lastName, email, password, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingOperatorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s created successfully", email))

	return nil
}

func ListTenantAccounts(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, sort, filter string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]
	var tenant *api.Tenant
	var verbose, l bool

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if sort, err = cmd.Flags().GetString("sort"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if filter, err = cmd.Flags().GetString("filter"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter != "" {
		filter = utils.BuildFilterQuery(filter)
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	utils.PrintList("Your Tenant Users List")

	if len(accounts.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if verbose {
		utils.PrintVerbose(accounts.Data, l)
		return nil
	}

	for _, account := range accounts.Data {
		fmt.Printf("• %s\n", account.ID)
		if l {
			fmt.Println()
		}
	}

	return nil
}

func DescribeTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, format string
	var conf *configuration.Config
	var tenant *api.Tenant
	var account *api.Account

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	accountID := args[0]
	if account, err = getTenantAccountById(conf, *accessToken, id, accountID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantAccountRequest, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintFormattedData(*account, format)

	return nil
}

func RemoveTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, email, password, code, configPath, deleteTenantAccountToken string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if password, err = cmd.Flags().GetString("password"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if code, err = cmd.Flags().GetString("code"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	accountID := args[0]
	if deleteTenantAccountToken, err = api.ForgeOperatorDeleteTenantAccountToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, accountID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingTenantAccountDeleteTokenRequest, err)
	}

	if err = api.RemoveTenantAccount(conf.Urls, *accessToken, id, accountID, deleteTenantAccountToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantAccountRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("user %s removed successfully", accountID))

	return nil
}

func BanTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	accountID := args[0]

	if err = api.ToggleBanAccount(conf.Urls, *accessToken, id, accountID, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorFreezingTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("user %s banned successfully", accountID))

	return nil
}

func UnbanTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	accountID := args[0]

	if err = api.ToggleBanAccount(conf.Urls, *accessToken, id, accountID, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUnfreezingTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("user %s unbanned successfully", accountID))

	return nil
}

func RestoreTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	accountID := args[0]

	if err = api.RestoreTenantAccount(conf.Urls, *accessToken, id, accountID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRestoringTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("user %s restored successfully", accountID))

	return nil
}

func DeleteTenantAccountSessions(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				id = tenant.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Tenant %s not found", name))
			return nil
		}
	}

	accountID := args[0]

	if err = api.DeleteTenantAccountSessions(conf.Urls, *accessToken, id, accountID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantAccountSessionsRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("user %s sessions deleted successfully", accountID))

	return nil
}

func getTenantAccountById(conf *configuration.Config, accessToken string, tenantID string, account string) (*api.Account, error) {
	var err error
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if accounts, err = api.ListTenantAccounts(conf.Urls, accessToken, tenantID, "", ""); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}
	for _, ac := range accounts.Data {
		if ac.ID == account {
			return ac, nil
		}
	}

	return nil, fmt.Errorf("operator %s not found", account)
}
