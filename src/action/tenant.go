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
	var name, description, imageUrl, settingsString, configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if imageUrl, err = cmd.Flags().GetString("image-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if settingsString == "" {
		settingsString = "{}"
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
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

	utils.PrintSuccess(fmt.Sprintf("tenant: %s created successfully", response.ID))

	return nil
}

func ListTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var tenants *api.TenantList

	if conf, configPath, err = configuration.ReadConfig(cmd, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	var verbose, l bool

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintList("Your Tenants List")

	for _, tenant := range tenants.Tenants {
		if verbose {
			fmt.Printf(" • %s, %s, %s\n", tenant.ID, tenant.Name, *tenant.Description)
		} else {
			fmt.Printf(" • %s\n", tenant.Name)
		}
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

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
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

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenants *api.TenantList

		if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
		}

		for _, tenant := range tenants.Tenants {
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
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallenge, err)
	}

	if deleteTenantToken, err = api.ForgeOperatorDeleteTenantToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteToken, err)
	}

	if err = api.RemoveTenant(conf.Urls, *accessToken, id, deleteTenantToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
	}

	utils.PrintDelete(fmt.Sprintf("tenant %s removed successfully", id))

	return nil
}

func DescribeTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, format, configPath string
	var conf *configuration.Config
	var tenants *api.TenantList

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
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

	if tenants, err = api.ListTenants(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	switch {
	case id == "" && name == "":
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantDescription, err)
	case name == "":
		for _, tenant := range tenants.Tenants {
			if id == tenant.ID {
				utils.PrintFormattedData(*tenant, format)
				break
			}
		}
	case id == "":
		for _, tenant := range tenants.Tenants {
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

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if len(args) != 1 {
		return fmt.Errorf("invalid image url: %w", err)
	}

	description := args[0]

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if err = api.EditTenantDescription(conf.Urls, *accessToken, id, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantDescription, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s description updated successfully", id))

	return nil
}

func EditTenantImage(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if len(args) != 1 {
		return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
	}

	image := args[0]

	if image != "" {
		if _, err := url.ParseRequestURI(image); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if err = api.EditTenantImage(conf.Urls, *accessToken, id, image); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingTenant, err)
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

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenant *api.Tenant
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if swarms, err = api.ListAvailableTenantSwarms(conf.Urls, *accessToken, id); err != nil {
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

func AddOperatorToTenant(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, email, role, firstName, lastName, configPath string
	var conf *configuration.Config
	var policies *api.PolicyList
	var found bool

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

	if id == "" && name == "" {
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenant *api.Tenant
		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if policies, err = api.ListTenantPolicies(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
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
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperator, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator: %s invited successfully", email))

	return nil
}

func ListTenantOperators(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var operators *api.OperatorList

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}
		id = tenant.ID
	}

	if operators, err = api.ListTenantOperators(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperators, err)
	}

	utils.PrintList("Your Tenant Operators")

	if len(operators.Operators) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var verbose, l bool

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	for _, operator := range operators.Operators {
		if verbose {
			fmt.Printf(" • %s, %s, %s %s, %s\n", operator.ID, operator.Email, operator.FirstName, operator.LastName, operator.Status)
		} else {
			fmt.Printf(" • %s\n", operator.Email)
		}
		if l {
			fmt.Println()
		}
	}

	return nil
}

func RemoveTenantOperator(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, operator string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("%s: %w", constants.ErrorTenantNameOrID, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var tenant *api.Tenant

		if tenant, err = getTenantByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
		}

		id = tenant.ID
	}

	if len(args) != 1 {
		return fmt.Errorf("invalid operator name or id: %w", err)
	}

	operator = args[0]

	if operator, err = getOperatorByEmailOrId(conf, *accessToken, id, operator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	if err = api.RemoveTenantOperator(conf.Urls, *accessToken, id, operator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingOperator, err)
	}

	utils.PrintDelete(fmt.Sprintf("operator %s removed successfully", operator))

	return nil
}

func ConnectSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		return fmt.Errorf("invalid tenant id or name: %w", err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if len(args) != 1 {
		return fmt.Errorf("invalid swarm id or name: %w", err)
	}

	if id == "" {
		var tenant *api.Tenant

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
		return fmt.Errorf("%s: %w", constants.ErrorConnectingSwarm, err)
	}

	utils.PrintSuccess(fmt.Sprintf("tenant %s connected to swarm %s successfully", id, swarm))

	return nil
}

func getTenantByNameOrId(conf *configuration.Config, accessToken string, tenantID string) (*api.Tenant, error) {
	var err error
	var tenants *api.TenantList
	var tenant *api.Tenant

	if tenants, err = api.ListTenants(conf.Urls, accessToken); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantList, err)
	}

	for _, tn := range tenants.Tenants {
		if tenantID == tn.Name || tenantID == tn.ID {
			tenant = tn

		}
	}

	if tenant == nil {
		return nil, fmt.Errorf("tenant %s not found", tenant.ID)
	}

	return tenant, nil
}

func getOperatorByEmailOrId(conf *configuration.Config, accessToken string, tenantID string, operator string) (string, error) {
	var err error
	var operators *api.OperatorList
	var id string

	if operators, err = api.ListSwarmOperators(conf.Urls, accessToken, tenantID); err != nil {
		return id, fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}
	for _, op := range operators.Operators {
		if operator == op.Email || operator == op.ID {
			id = op.ID
		}
	}

	if id == "" {
		return "", fmt.Errorf("operator %s not found", operator)
	}

	return id, nil
}
