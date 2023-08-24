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
	var name, description, imageUrl, settingsString, configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if _, err = tui.TextInputs("Fill in the form below", false, tui.Input{Placeholder: "Name", IsPassword: false, Value: &name}, tui.Input{Placeholder: "Description", IsPassword: false, Value: &description}, tui.Input{Placeholder: "Image URL", IsPassword: false, Value: &imageUrl}); err != nil {
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

	if _, err = tui.TextAreas("Fill in the tenant settings", true, tui.TextArea{Placeholder: "{}", Value: &settingsString}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if settingsString == "" {
		settingsString = "{}"
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	var settings map[string]interface{}

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if response, err = api.CreateTenant(conf.Urls, *accessToken, name, &description, &imageUrl, settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenant, err)
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
	var operator *api.Operator
	var tenants *api.TenantList

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	if len(tenants.Tenants) == 0 {
		utils.PrintNotFound("No tenants found")
		return nil
	}

	var choices []string

	for _, tenant := range tenants.Tenants {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", tenant.ID, tenant.Name, *tenant.Description))
	}

	if choice, err = tui.ChooseOne("Which tenant would you like to delete?", false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
	}
	_, withoutPrefix, _ := strings.Cut(choice, " ")
	id, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Confirm your login to delete the tenant", true, tui.Input{Placeholder: "Email", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallenge, err)
	}

	if deleteTenantToken, err = api.ForgeOperatorDeleteTenantToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteToken, err)
	}

	if err = api.RemoveTenant(conf.Urls, *accessToken, id, deleteTenantToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
	}

	utils.PrintDelete(fmt.Sprintf("tenant %s removed successfully\n", id))

	return nil
}

func DescribeTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var nameOrId, format, configPath string
	var conf *configuration.Config
	var tenants *api.TenantList
	var operator *api.Operator

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs("Enter your tenant ID or Name", false, tui.Input{Placeholder: "Tenant ID or Name", Value: &nameOrId}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if format, err = tui.ChooseOne("Choose your output format", true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	for _, tenant := range tenants.Tenants {
		if tenant.ID == nameOrId || tenant.Name == nameOrId {
			utils.PrintFormattedData(*tenant, format)
			return nil
		}
	}

	return fmt.Errorf(constants.ErrorRetrievingTenant)
}

func EditTenantDescriptionInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var nameOrId, configPath, description string
	var conf *configuration.Config
	var operator *api.Operator
	var tenants *api.TenantList
	var found bool

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if _, err = tui.TextInputs("Enter your tenant ID or Name", false, tui.Input{Placeholder: "Tenant ID or Name", Value: &nameOrId}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	for _, tenant := range tenants.Tenants {
		if tenant.ID == nameOrId || tenant.Name == nameOrId {
			nameOrId = tenant.ID
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf(constants.ErrorRetrievingTenant)
	}

	if _, err = tui.TextInputs("Enter your new tenant description", true, tui.Input{Placeholder: "New Tenant Description", Value: &description}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if err = api.EditTenantDescription(conf.Urls, *accessToken, nameOrId, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantDescription, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s description updated successfully", nameOrId))
	return nil
}

func EditTenantImageInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var nameOrId, configPath, image string
	var conf *configuration.Config
	var operator *api.Operator
	var tenants *api.TenantList
	var found bool

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs("Enter your tenant ID or Name", false, tui.Input{Placeholder: "Tenant ID or Name", Value: &nameOrId}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	for _, tenant := range tenants.Tenants {
		if tenant.ID == nameOrId || tenant.Name == nameOrId {
			nameOrId = tenant.ID
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf(constants.ErrorRetrievingTenant)
	}

	if _, err = tui.TextInputs("Enter your new tenant image URL", true, tui.Input{Placeholder: "New Tenant Image", Value: &image}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if image != "" {
		if _, err := url.ParseRequestURI(image); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if err = api.EditTenantImage(conf.Urls, *accessToken, nameOrId, image); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenant, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s image updated successfully", nameOrId))

	return nil
}

func ListTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var operator *api.Operator
	var tenants *api.TenantList

	if conf, configPath, err = configuration.ReadConfig(cmd, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	utils.PrintList("Your Tenants List")
	for _, tenant := range tenants.Tenants {
		fmt.Printf("• %s, %s, %s\n", tenant.ID, tenant.Name, *tenant.Description)
	}

	return nil
}

func ListAvailableSwarmsTenantInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var nameOrId, configPath string
	var conf *configuration.Config
	var swarms *api.SwarmList
	var tenants *api.TenantList
	var operator *api.Operator
	var found bool

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs("Enter your tenant ID or Name", false, tui.Input{Placeholder: "Tenant ID or Name", Value: &nameOrId}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	for _, tenant := range tenants.Tenants {
		if tenant.ID == nameOrId || tenant.Name == nameOrId {
			nameOrId = tenant.ID
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf(constants.ErrorRetrievingTenant)
	}

	if swarms, err = api.ListAvailableTenantSwarms(conf.Urls, *accessToken, nameOrId); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingAvailableTenantSwarms, err)
	}

	utils.PrintList("Your Tenant Connected Swarms")

	if len(swarms.Swarms) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	for _, swarm := range swarms.Swarms {
		cross := " "
		if swarm.Default {
			cross = "x"
		}
		fmt.Printf("[%s] %s\n", cross, swarm.SwarmID)
	}

	return nil
}

func AddOperatorInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, tenant, role, email, first_name, last_name, configPath, choice string
	var conf *configuration.Config
	var policies *api.PolicyList

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if _, err = tui.TextInputs("Enter your tenant ID or Name", false, tui.Input{Placeholder: "Tenant ID or Name", Value: &tenant}); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if tenant != "" {
		if id, err = getTenantByNameOrId(conf, *accessToken, tenant); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
	}

	if name != "" {
		if id, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
	}

	if _, err = tui.TextInputs("Fill in the form for the operator to invite", false, tui.Input{Placeholder: "Operator Email", Value: &email}, tui.Input{Placeholder: "Operator First Name", Value: &first_name}, tui.Input{Placeholder: "Operator Last Name", Value: &last_name}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if policies, err = api.ListTenantPolicies(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	var choices []string

	for _, policy := range policies.Policies {
		choices = append(choices, fmt.Sprintf("• %s", policy.Name))
	}

	if choice, err = tui.ChooseOne("Which policy would you like to assign to the operator?", true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
	}
	_, choice, _ = strings.Cut(choice, " ")

	for _, policy := range policies.Policies {
		if policy.Name == choice {
			role = policy.ID
		}
	}

	if err = api.InviteOperator(conf.Urls, *accessToken, id, email, role, first_name, last_name); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperator, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator: %s invited successfully", email))

	return nil
}
