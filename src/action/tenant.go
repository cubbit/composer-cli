package action

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, description, settingsString, couponCode, configPath, zone string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config
	var zones *api.ZoneMap
	var settings api.TenantSettings

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if settingsString, err = cmd.Flags().GetString("settings"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if settingsString == "" {
		settingsString = "{}"
	}

	if couponCode, err = cmd.Flags().GetString("distributor-code"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if zone, err = cmd.Flags().GetString("zone"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if zone != "" {
		if zones, err = api.GetGatwayZones(conf.Urls); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingZonesRequest, err)
		}
		if _, ok := zones.Zones[zone]; !ok {
			return fmt.Errorf(constants.ErrorInvalidZone)
		}
	}

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if response, err = api.CreateTenant(conf.Urls, *accessToken, name, &description, settings, couponCode, zone); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant: %s created successfully", response.ID))

	return nil
}

func ListTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var tenants *api.GenericPaginatedResponse[*api.Tenant]
	var sort, filter string
	var verbose, l bool

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
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

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintList("Your Tenants List")

	if len(tenants.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	if verbose {
		utils.PrintVerbose(tenants.Data, l)
		return nil
	}

	for _, tenant := range tenants.Data {
		fmt.Printf(" • %s\n", tenant.Name)
		if l {
			fmt.Println()
		}
	}

	return nil
}

func RemoveTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, email, password, code, configPath, deleteTenantToken string
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

	if deleteTenantToken, err = api.ForgeOperatorDeleteTenantToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingTenantDeleteTokenRequest, err)
	}

	if err = api.RemoveTenant(conf.Urls, *accessToken, id, deleteTenantToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("tenant %s removed successfully", id))

	return nil
}

func DescribeTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, format, configPath string
	var conf *configuration.Config
	var tenants *api.GenericPaginatedResponse[*api.Tenant]

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
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

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken, "", ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	switch {
	case name == "":
		for _, tenant := range tenants.Data {
			if id == tenant.ID {
				utils.PrintFormattedData(*tenant, format)
				break
			}
		}
	case id == "":
		for _, tenant := range tenants.Data {
			if name == tenant.Name {
				utils.PrintFormattedData(*tenant, format)
				break
			}
		}
	default:
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	return nil
}

func EditTenantDescription(cmd *cobra.Command, args ...string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant

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
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	description := args[0]

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if err = api.EditTenantDescription(conf.Urls, *accessToken, id, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s description updated successfully", id))

	return nil
}

func EditTenantImage(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant

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

	image := args[0]

	if image != "" {
		if _, err := url.ParseRequestURI(image); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if err = api.EditTenantImage(conf.Urls, *accessToken, id, image); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s image updated successfully", id))

	return nil
}

func ListAvailableSwarmsTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var swarms *api.SwarmList
	var tenant *api.Tenant

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

	if swarms, err = api.ListAvailableTenantSwarms(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	utils.PrintList("Your Tenant Connected Swarms List")

	if len(swarms.Swarms) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	for _, swarm := range swarms.Swarms {
		cross := " "
		if swarm.Default {
			cross = "x"
		}
		fmt.Printf("[%s] %s, %s\n", cross, swarm.SwarmID, swarm.SwarmName)
	}
	return nil
}

func AddOperatorToTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, email, role, firstName, lastName, configPath string
	var conf *configuration.Config
	var policies *api.PolicyList
	var found bool
	var tenant *api.Tenant

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if firstName, err = cmd.Flags().GetString("first-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if lastName, err = cmd.Flags().GetString("last-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if role, err = cmd.Flags().GetString("role"); err != nil {
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

	if policies, err = api.ListTenantPolicies(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingPoliciesRequest, err)
	}

	for _, policy := range policies.Policies {
		if policy.Name == role {
			role = policy.ID
			found = true
		}
	}

	if !found {
		utils.PrintNotFound(fmt.Sprintf("Policy %s not found", role))
		return nil
	}

	if err = api.InviteOperatorToTenant(conf.Urls, *accessToken, id, email, role, firstName, lastName); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s invited successfully", email))

	return nil
}

func ListTenantOperators(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var operators *api.OperatorList
	var tenant *api.Tenant
	var verbose, l bool

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

	if operators, err = api.ListTenantOperators(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	utils.PrintList("Your Tenant Operators List")

	if len(operators.Operators) == 0 {
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
		utils.PrintVerbose(operators.Operators, l)
		return nil
	}

	for _, operator := range operators.Operators {
		fmt.Printf(" • %s\n", operator.Email)
		if l {
			fmt.Println()
		}
	}

	return nil
}

func RemoveTenantOperator(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant
	var operator *api.Operator

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

	operatorID := args[0]

	if operator, err = getTenantOperatorByEmailOrId(conf, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	operatorID = operator.ID
	if err = api.RemoveTenantOperator(conf.Urls, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingOperatorRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("operator %s removed successfully", operatorID))

	return nil
}

func ConnectSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var tenant *api.Tenant

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
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	swarm := args[0]
	if swarm, err = getSwarmByNameOrId(conf, *accessToken, swarm); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
	}

	if err = api.ConnectSwarm(conf.Urls, *accessToken, id, swarm); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConnectingSwarmRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s connected to swarm %s successfully", id, swarm))

	return nil
}

func CreateTenantAccounts(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var emails []string
	var tenant *api.Tenant

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if emails, err = cmd.Flags().GetStringSlice("emails"); err != nil {
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

	if err = api.CreateTenantAccounts(conf.Urls, *accessToken, id, emails); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingTenantAccountsRequest, err)
	}

	utils.PrintSuccess("users created successfully")

	return nil
}

func UpdateTenantAccount(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, firstName, lastName, endpointGateway string
	var internal bool
	var maxAllowedProjects int
	var conf *configuration.Config
	var tenant *api.Tenant

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
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

	if err = api.UpdateAccount(conf.Urls, *accessToken, id, accountID, requestBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUpdatingTenantAccountRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("user %s updated successfully", accountID))

	return nil
}

func getTenantByNameOrId(conf *configuration.Config, accessToken string, tenantID string) (*api.Tenant, error) {
	var err error
	var tenants *api.GenericPaginatedResponse[*api.Tenant]
	var tenant *api.Tenant

	if tenants, err = api.ListTenants(conf.Urls, accessToken, "", ""); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingTenantsRequest, err)
	}

	for _, tn := range tenants.Data {
		if tenantID == tn.Name || tenantID == tn.ID {
			tenant = tn

		}
	}

	if tenant == nil {
		return nil, fmt.Errorf("tenant %s not found", tenant.ID)
	}

	return tenant, nil
}

func getTenantOperatorByEmailOrId(conf *configuration.Config, accessToken string, tenantID string, operator string) (*api.Operator, error) {
	var err error
	var operators *api.OperatorList

	if operators, err = api.ListTenantOperators(conf.Urls, accessToken, tenantID); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}
	for _, op := range operators.Operators {
		if operator == op.Email || operator == op.ID {
			return op, nil
		}
	}

	return nil, fmt.Errorf("operator %s not found", operator)
}

func EditTenantSettings(cmd *cobra.Command, args ...string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var settings api.TenantSettings
	var tenant *api.Tenant

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
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	settingsString := args[0]

	if err = json.Unmarshal([]byte(settingsString), &settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if err = api.EditTenantSettings(conf.Urls, *accessToken, id, settings); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s settings updated successfully", id))

	return nil
}

func DescribeTenantOperator(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, format string
	var conf *configuration.Config
	var tenant *api.Tenant
	var operator *api.Operator

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

	operatorID := args[0]
	if operator, err = getTenantOperatorByEmailOrId(conf, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintFormattedData(*operator, format)

	return nil
}

func EditTenantOperatorRole(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, role, configPath string
	var conf *configuration.Config
	var policies *api.PolicyList
	var found bool
	var tenant *api.Tenant
	var operator *api.Operator

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if role, err = cmd.Flags().GetString("role"); err != nil {
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

	operatorID := args[0]
	if operator, err = getTenantOperatorByEmailOrId(conf, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	operatorID = operator.ID

	if policies, err = api.ListTenantPolicies(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingPoliciesRequest, err)
	}

	for _, policy := range policies.Policies {
		if policy.Name == role {
			role = policy.ID
			found = true
		}
	}

	if !found {
		utils.PrintNotFound(fmt.Sprintf("Policy %s not found", role))
		return nil
	}

	if err = api.EditOperatorRoleInTenant(conf.Urls, *accessToken, id, operatorID, role); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s role updated successfully", operatorID))

	return nil
}

func AssignTenantToCoupon(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, couponCode, configPath string
	var conf *configuration.Config
	var response *api.GenericIDResponseModel
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

	if id == "" {
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	couponCode = args[0]

	if response, err = api.AssignTenantToCoupon(conf.Urls, *accessToken, id, couponCode); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenantRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s assigned successfully to %s", response.ID, couponCode))

	return nil
}
