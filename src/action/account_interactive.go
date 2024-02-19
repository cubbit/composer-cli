package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateAccountInteractive(cmd *cobra.Command) error {
	var urls *configuration.Url
	var apiServerUrl, firstName, lastName, email, password, tenantID string
	var err error

	if _, err = tui.TextInputs("Enter your API server URL", false, tui.Input{Placeholder: "API server url: (default https://api.cubbit.eu)", IsPassword: false, Value: &apiServerUrl}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if urls, err = configuration.ConfigureAPIServerURL(configuration.SessionTypeAccount, apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if _, err = tui.TextInputs(
		"Fill in the form bellow",
		true,
		tui.Input{Placeholder: "First Name*", IsPassword: false, Value: &firstName},
		tui.Input{Placeholder: "Last Name*", IsPassword: false, Value: &lastName},
		tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email},
		tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password},
		tui.Input{Placeholder: "Tenant ID*", IsPassword: false, Value: &tenantID},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if err = api.CreateAccount(*urls, firstName, lastName, email, password, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingOperatorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s created successfully", email))

	return nil
}

func ListTenantAccountsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	utils.PrintList("Your Tenant Accounts List")

	if len(accounts.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	for _, account := range accounts.Data {
		fmt.Printf(" • %s, %s %s\n", account.ID, account.FirstName, account.LastName)
	}

	return nil
}

func DescribeTenantAccountInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, accountID, format string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	if len(accounts.Data) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	var choices []string

	for _, ac := range accounts.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s %s", ac.ID, ac.FirstName, ac.LastName))

	}

	if len(choices) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which account would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDescribingTenantAccountRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	accountID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, ac := range accounts.Data {
		if ac.ID == accountID {
			utils.PrintFormattedData(*ac, format)
			break
		}
	}

	return nil
}

func RemoveTenantAccountInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, accountID, email, password, deleteTenantAccountToken, code string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]
	var challenge *api.ChallengeResponseModel

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	if len(accounts.Data) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	var choices []string

	for _, ac := range accounts.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s %s", ac.ID, ac.FirstName, ac.LastName))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which account would you like to delete?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	accountID, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs(fmt.Sprintf("Confirm your login to delete account %s 🚮", utils.RedBg.Render(accountID)), true, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if deleteTenantAccountToken, err = api.ForgeOperatorDeleteTenantAccountToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, accountID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteTokenRequest, err)
	}

	if err = api.RemoveTenantAccount(conf.Urls, *accessToken, id, accountID, deleteTenantAccountToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantAccountRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("account %s removed successfully", accountID))

	return nil
}

func BanTenantAccountInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, accountID string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	if len(accounts.Data) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	var choices []string

	for _, ac := range accounts.Data {
		if !ac.Banned {
			choices = append(choices, fmt.Sprintf("• %s, %s %s", ac.ID, ac.FirstName, ac.LastName))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which account would you like to ban?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	accountID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.ToggleBanAccount(conf.Urls, *accessToken, id, accountID, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorFreezingTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s banned successfully", accountID))

	return nil
}

func UnbanTenantAccountInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, accountID string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	if len(accounts.Data) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	var choices []string

	for _, ac := range accounts.Data {
		if ac.Banned {
			choices = append(choices, fmt.Sprintf("• %s, %s %s", ac.ID, ac.FirstName, ac.LastName))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which account would you like to unban?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	accountID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.ToggleBanAccount(conf.Urls, *accessToken, id, accountID, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s unbanned successfully", accountID))

	return nil
}

func RestoreTenantAccountInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, accountID string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	if len(accounts.Data) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	var choices []string

	for _, ac := range accounts.Data {
		if ac.DeletedAt != nil {
			choices = append(choices, fmt.Sprintf("• %s, %s %s", ac.ID, ac.FirstName, ac.LastName))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which account would you like to restore?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	accountID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.RestoreTenantAccount(conf.Urls, *accessToken, id, accountID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRestoringTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s restored successfully", accountID))

	return nil
}

func DeleteTenantAccountSessionsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, accountID string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	if len(accounts.Data) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	var choices []string

	for _, ac := range accounts.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s %s", ac.ID, ac.FirstName, ac.LastName))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which account would you like to delete its sessions?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	accountID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.DeleteTenantAccountSessions(conf.Urls, *accessToken, id, accountID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s sessions deleted successfully", accountID))

	return nil
}

func CreateTenantAccountsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {

			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	var emailsString string

	if _, err = tui.EmailBulkTextAreas("Fill in with the list of emails", true, tui.EmailBulkTextArea{Placeholder: "email1,email2,email3...", Value: &emailsString}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	emails := strings.Split(emailsString, ",")
	for i, email := range emails {
		emails[i] = strings.TrimSpace(email)
	}

	if err = api.CreateTenantAccounts(conf.Urls, *accessToken, id, emails); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenantAccountsRequest, err)
	}

	utils.PrintSuccess("accounts created successfully")

	return nil
}

func UpdateTenantAccountInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, firstName, lastName, endpointGateway, internal, maxAllowedProjects string
	var conf *configuration.Config
	var accounts *api.GenericPaginatedResponse[*api.Account]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		var choice string
		var choices []string
		var tenants *api.GenericPaginatedResponse[*api.Tenant]

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Choose your tenant", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if accounts, err = api.ListTenantAccounts(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantAccountsRequest, err)
	}

	if len(accounts.Data) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	var choices []string

	for _, ac := range accounts.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s %s", ac.ID, ac.FirstName, ac.LastName))

	}

	if len(choices) == 0 {
		utils.PrintNotFound("No accounts found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which account would you like to update?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUpdatingTenantAccountRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	accountID, _, _ := strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Fill in the form for the account to update", true, tui.Input{Placeholder: "First Name", Value: &firstName}, tui.Input{Placeholder: "Last Name", Value: &lastName}, tui.Input{Placeholder: "Endpoint Gateway", Value: &endpointGateway}, tui.Input{Placeholder: "Internal", Value: &internal}, tui.Input{Placeholder: "Max Allowed Projects", Value: &maxAllowedProjects}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	var internalPtr *bool
	if internal != "" && internal != "true" && internal != "false" {
		return fmt.Errorf("%s: %w", constants.ErrorInvalidInternalValue, err)
	}
	v := internal == "true"
	internalPtr = &v

	var maxAllowedProjectsPtr *int
	if maxAllowedProjects != "" {

		v, err := strconv.Atoi(maxAllowedProjects)
		if err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorParsingMaxAllowedProjects, err)
		}

		maxAllowedProjectsPtr = &v
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
		EndpointGateway:    endpointGatewayPtr,
		Internal:           internalPtr,
		MaxAllowedProjects: maxAllowedProjectsPtr,
	}

	if err = api.UpdateAccount(conf.Urls, *accessToken, id, accountID, requestBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUpdatingTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("account %s updated successfully", accountID))

	return nil
}
