package action

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, description, imageUrl, settingsString, couponCode, configPath, zone string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config
	var choices []string
	var choice string

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs("Fill in the form below", false, tui.Input{Placeholder: "Name*", IsPassword: false, Value: &name}, tui.Input{Placeholder: "Description", IsPassword: false, Value: &description}, tui.Input{Placeholder: "Coupon code*", IsPassword: false, Value: &couponCode}, tui.Input{Placeholder: "Image URL", IsPassword: false, Value: &imageUrl}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if _, err = tui.TextAreas("Fill in the tenant settings", false, tui.TextArea{InitialValue: "{}", Value: &settingsString}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if settingsString == "" {
		settingsString = "{}"
	}

	var settings api.TenantSettings

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	var zones *api.ZoneMap
	if zones, err = api.GetGatwayZones(conf.Urls); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingZonesRequest, err)
	}

	if len(zones.Zones) != 0 {
		choices = append(choices, fmt.Sprintf("• %s", "Default"))
		for _, zn := range zones.Zones {
			choices = append(choices, fmt.Sprintf("• %s", zn.Name))
		}

		if choice, err = tui.ChooseOne("Which zone would you like to create your tenant?", true, true, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingZonesRequest, err)
		}

		if choice != "" {
			_, value, _ := strings.Cut(choice, " ")
			if value != "Default" {
				for _, zn := range zones.Zones {
					if value == zn.Name {
						zone = zn.Key
						break
					}
				}
			}
		}
	}

	if response, err = api.CreateTenant(conf.Urls, *accessToken, name, &description, &imageUrl, settings, couponCode, zone); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant  %s created successfully", response.ID))

	return nil
}

func RemoveTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, email, password, code, configPath, deleteTenantToken, choice string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	if len(tenants.Data) == 0 {
		utils.PrintNotFound("No tenants found")
		return nil
	}

	var choices []string

	for _, tenant := range tenants.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No tenants found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which tenant would you like to delete?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
	}

	splits := strings.Split(choice, ",")
	_, id, _ = strings.Cut(splits[0], " ")
	_, name, _ := strings.Cut(splits[1], " ")

	if _, err = tui.TextInputs(fmt.Sprintf("Confirm your login to delete the tenant %s - %s 🚮", utils.RedBg.Render(name), utils.RedBg.Render(id)), true, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if deleteTenantToken, err = api.ForgeOperatorDeleteTenantToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteTokenRequest, err)
	}

	if err = api.RemoveTenant(conf.Urls, *accessToken, id, deleteTenantToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("tenant %s removed successfully", id))

	return nil
}

func DescribeTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, format, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant

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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, tenant := range tenants.Data {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No tenants found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which tenant would you like to retrieve?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
		utils.PrintFormattedData(*tenant, format)
		return nil

	}

	if tenant, err = getTenantByNameOrId(conf, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmRequest, err)
	}

	utils.PrintFormattedData(*tenant, format)

	return nil
}

func EditTenantDescriptionInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, description string
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if _, err = tui.TextInputs("Enter your new tenant description", true, tui.Input{Placeholder: "New Tenant Description", Value: &description}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if err = api.EditTenantDescription(conf.Urls, *accessToken, id, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s description updated successfully", id))

	return nil
}

func EditTenantImageInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, image string
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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if _, err = tui.TextInputs("Enter your new tenant image URL", true, tui.Input{Placeholder: "New Tenant Image", Value: &image}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if image != "" {
		if _, err := url.ParseRequestURI(image); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if err = api.EditTenantImage(conf.Urls, *accessToken, id, image); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s image updated successfully", id))

	return nil
}

func ListTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var tenants *api.GenericPaginatedResponse[*api.Tenant]
	var sort string

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if sort, err = tui.ChooseOne("Choose your sorting option", false, true, []string{"id", "name", "owner_id", "coupon_id", "created_at", "deleted_at"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, sort, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	utils.PrintList("Your Tenants List")

	if len(tenants.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string
	for _, tenant := range tenants.Data {
		list = append(list, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
	}

	tui.List(list)

	return nil
}

func ListAvailableSwarmsTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var swarms *api.SwarmList

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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if swarms, err = api.ListAvailableTenantSwarms(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	utils.PrintList("Your Tenant Connected Swarms List")

	if len(swarms.Swarms) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string
	for _, swarm := range swarms.Swarms {
		defaultString := " "
		if swarm.Default {
			defaultString = "[DEFAULT]"
		}
		list = append(list, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, defaultString))

	}

	tui.List(list)

	return nil
}

func AddOperatorToTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, role, email, first_name, last_name, configPath, choice string
	var conf *configuration.Config
	var policies *api.PolicyList

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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if _, err = tui.TextInputs("Fill in the form for the operator to invite", false, tui.Input{Placeholder: "Email*", Value: &email}, tui.Input{Placeholder: "First Name", Value: &first_name}, tui.Input{Placeholder: "Last Name", Value: &last_name}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if policies, err = api.ListTenantPolicies(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingPoliciesRequest, err)
	}

	var choices []string

	for _, policy := range policies.Policies {
		choices = append(choices, fmt.Sprintf("• %s", policy.Name))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No policies found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which policy would you like to assign to the operator?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingPolicies, err)
	}
	_, choice, _ = strings.Cut(choice, " ")

	for _, policy := range policies.Policies {
		if policy.Name == choice {
			role = policy.ID
		}
	}

	if err = api.InviteOperatorToTenant(conf.Urls, *accessToken, id, email, role, first_name, last_name); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s invited successfully", email))

	return nil
}

func ListTenantOperatorsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var operators *api.OperatorList

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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if operators, err = api.ListTenantOperators(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	utils.PrintList("Your Tenant Operators List")

	if len(operators.Operators) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string
	for _, operator := range operators.Operators {
		list = append(list, fmt.Sprintf("• %s, %s, %s %s", operator.ID, operator.Email, operator.FirstName, operator.LastName))
	}

	tui.List(list)

	return nil
}

func RemoveTenantOperatorInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, operatorID string
	var conf *configuration.Config
	var operators *api.OperatorList
	var operator *api.Operator

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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if operators, err = api.ListTenantOperators(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if len(operators.Operators) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	var choices []string

	for _, op := range operators.Operators {
		if op.ID != operator.ID {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s %s", op.ID, op.Email, op.FirstName, op.LastName))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which operator would you like to remove?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	operatorID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.RemoveTenantOperator(conf.Urls, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingOperatorRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("operator %s removed successfully", operatorID))

	return nil
}

func ConnectSwarmInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, swarm string
	var conf *configuration.Config
	var operator *api.Operator
	var swarms []api.Swarm
	var swarms1 *api.SwarmList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {

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

	if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	if swarms1, err = api.GetTenantCouponSwarms(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	mergedSwarms := make(map[string]api.Swarm)

	for _, swarm := range swarms {

		mergedSwarms[swarm.ID] = swarm
	}

	for _, swarm := range swarms1.Swarms {

		if _, ok := mergedSwarms[swarm.ID]; !ok {
			mergedSwarms[swarm.ID] = swarm
		}
	}

	if len(mergedSwarms) == 0 {
		utils.PrintNotFound("No swarms found")
		return nil
	}

	var choices []string

	for _, sw := range mergedSwarms {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", sw.ID, sw.Name, sw.Description))
	}

	if choice, err = tui.ChooseOne("Which swarm would you like to connect?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	swarm, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.ConnectSwarm(conf.Urls, *accessToken, id, swarm); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConnectingSwarmRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s connected to swarm %s successfully", id, swarm))

	return nil
}

func EditTenantSettingsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, settingsString string
	var conf *configuration.Config
	var tenant *api.Tenant

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

		for _, tn := range tenants.Data {
			if tn.ID == id {
				tenant = tn
				break
			}
		}

	}

	if id != "" {
		if tenant, err = api.GetTenant(conf.Urls, *accessToken, id); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
	}

	if id == "" && name != "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID

	}

	settingsJSON, err := json.MarshalIndent(tenant.Settings, "", "    ")
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorMarshallingJsonSettings, err)
	}
	if _, err = tui.TextAreas("Fill in the tenant settings", false, tui.TextArea{InitialValue: string(settingsJSON), Value: &settingsString}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if settingsString == "" {
		settingsString = "{}"
	}

	var settings api.TenantSettings

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if err = api.EditTenantSettings(conf.Urls, *accessToken, id, settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s settings updated successfully", id))

	return nil
}

func DescribeTenantOperatorInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, choice, operatorID, format string
	var conf *configuration.Config
	var operators *api.OperatorList
	var operator *api.Operator

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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if operators, err = api.ListTenantOperators(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if len(operators.Operators) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	var choices []string

	for _, op := range operators.Operators {
		if op.ID != operator.ID {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s %s", op.ID, op.Email, op.FirstName, op.LastName))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which operator would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	operatorID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, op := range operators.Operators {
		if op.ID == operatorID {
			utils.PrintFormattedData(op, format)
			break
		}
	}

	return nil
}

func EditTenantOperatorRoleInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, role, operatorID, configPath, choice string
	var conf *configuration.Config
	var policies *api.PolicyList
	var operators *api.OperatorList
	var operator *api.Operator

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

		if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
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

	if operators, err = api.ListTenantOperators(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if len(operators.Operators) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	var choices []string
	for _, op := range operators.Operators {
		if op.ID != operator.ID {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s %s", op.ID, op.Email, op.FirstName, op.LastName))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No operators found")
		return nil
	}

	if choice, err = tui.ChooseOne("choose your operator", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	operatorID, _, _ = strings.Cut(withoutPrefix, ",")

	if policies, err = api.ListTenantPolicies(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingPoliciesRequest, err)
	}

	choices = make([]string, 0)
	for _, policy := range policies.Policies {
		choices = append(choices, fmt.Sprintf("• %s", policy.Name))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No policies found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which policy would you like to assign to the operator?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingPolicies, err)
	}
	_, choice, _ = strings.Cut(choice, " ")

	for _, policy := range policies.Policies {
		if policy.Name == choice {
			role = policy.ID
		}
	}

	if err = api.EditOperatorRoleInTenant(conf.Urls, *accessToken, id, operatorID, role); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingOperatorRoleRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s 's role updated successfully", operatorID))

	return nil
}

func AssignTenantToCouponInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var tenantID, couponCode, configPath string
	var conf *configuration.Config
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributorCoupons *api.DistributorCouponList
	var tenants *api.GenericPaginatedResponse[*api.Tenant]
	var response *api.GenericIDResponseModel

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	for _, tenant := range tenants.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, utils.StringOrEmpty(tenant.Description)))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No tenants found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which tenant would you like to select?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	tenantID, _, _ = strings.Cut(withoutPrefix, ",")

	choices = []string{}
	if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
	}

	for _, distributor := range distributors.Distributors {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributor found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor would you like to select?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
	}

	_, withoutPrefix, _ = strings.Cut(choice, " ")
	id, _, _ := strings.Cut(withoutPrefix, ",")

	choices = []string{}
	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	for _, coupon := range distributorCoupons.Coupons {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", coupon.ID, coupon.Name, coupon.Code))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributor code found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor code would you like to assign?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponList, err)
	}

	parts := strings.Split(choice, ",")

	couponCode = strings.TrimSpace(parts[2])

	if response, err = api.AssignTenantToCoupon(conf.Urls, *accessToken, tenantID, couponCode); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s assigned successfully to %s", response.ID, couponCode))

	return nil
}
