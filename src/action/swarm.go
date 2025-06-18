package action

import (
	"encoding/json"
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarm(cmd *cobra.Command, args []string) error {
	var config *configuration.Config
	var err error
	var name, description, swarmConfig1, configPath string
	var accessToken *string
	var operator *api.Operator
	var response *api.GenericIDResponseModel
	var swarmConfig map[string]interface{}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if swarmConfig1, err = cmd.Flags().GetString("configuration"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if err = json.Unmarshal([]byte(swarmConfig1), &swarmConfig); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if config, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if response, err = api.CreateSwarm(config.Urls, *accessToken, operator.ID, name, description, swarmConfig); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingSwarmRequest, err)
	}

	utils.PrintCreateSuccess("swarm", response.ID)

	return nil
}

func DescribeSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var id, name, format, configPath string
	var accessToken *string
	var config *configuration.Config
	var operator *api.Operator
	var swarm *api.Swarm
	var swarms []*api.Swarm

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if config, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if name != "" {
		if swarms, err = api.ListSwarms(config.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, sw := range swarms {
			if sw.Name == name {
				utils.PrintFormattedData(sw, format)
				return nil
			}
		}

	}
	if swarm, err = api.GetSwarm(config.Urls, *accessToken, operator.ID, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmRequest, err)
	}

	utils.PrintFormattedData(*swarm, format)

	return nil
}

func ListSwarms(cmd *cobra.Command, args []string) error {
	var err error
	var configPath string
	var accessToken *string
	var config *configuration.Config
	var operator *api.Operator
	var swarms []*api.Swarm
	var verbose, l bool

	if config, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}
	if swarms, err = api.ListSwarms(config.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(swarms) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	utils.PrintList("Your Swarms List")

	if verbose {
		utils.PrintVerbose(swarms, l)
	} else {
		var IDs []string
		for _, swarm := range swarms {
			IDs = append(IDs, swarm.ID)
		}
		utils.PrintSimpleList(IDs)
	}

	return nil
}

func RemoveSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, email, password, code, configPath, deleteSwarmToken string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel

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
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if deleteSwarmToken, err = api.ForgeOperatorDeleteSwarmToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteTokenRequest, err)
	}

	if err = api.RemoveSwarm(conf.Urls, *accessToken, id, deleteSwarmToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingSwarmRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("swarm %s removed successfully", id))

	return nil
}

func EditSwarmDescription(cmd *cobra.Command, args ...string) error {
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

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	description := args[0]

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if err = api.EditSwarmDescription(conf.Urls, *accessToken, id, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("swarm %s description updated successfully", id))

	return nil
}

func EditSwarmName(cmd *cobra.Command, args ...string) error {
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

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	newName := args[0]

	if err = api.EditSwarmName(conf.Urls, *accessToken, id, newName); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("swarm %s name updated successfully", id))

	return nil
}

func AddOperatorToSwarm(cmd *cobra.Command, args []string) error {
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

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if policies, err = api.ListSwarmPolicies(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingPoliciesRequest, err)
	}

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if role, err = cmd.Flags().GetString("role"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
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

	if err = api.InviteOperatorToSwarm(conf.Urls, *accessToken, id, email, role, firstName, lastName); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s invited successfully", email))

	return nil
}

func ListSwarmOperators(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var operators *api.OperatorList
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
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if operators, err = api.ListSwarmOperators(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(operators.Operators) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	if len(operators.Operators) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	utils.PrintList("Your Swarm Operators List")

	if verbose {
		utils.PrintVerbose(operators.Operators, l)
	} else {
		var IDs []string
		for _, operator := range operators.Operators {
			IDs = append(IDs, operator.ID)
		}
		utils.PrintSimpleList(IDs)
	}

	return nil
}

func RemoveSwarmOperator(cmd *cobra.Command, args []string) error {
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

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	operatorID := args[0]

	if _, err = getSwarmOperatorByEmailOrId(conf, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if err = api.RemoveSwarmOperator(conf.Urls, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingOperatorRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("operator %s removed successfully", operatorID))

	return nil
}

func DescribeSwarmOperator(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, format string
	var conf *configuration.Config
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
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	operatorID := args[0]
	if operator, err = getSwarmOperatorByEmailOrId(conf, *accessToken, id, operatorID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintFormattedData(*operator, format)

	return nil
}

func EditSwarmOperatorRole(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, role, configPath string
	var conf *configuration.Config
	var policies *api.PolicyList
	var found bool
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
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	operatorID := args[0]
	if operator, err = getSwarmOperatorByEmailOrId(conf, *accessToken, id, operatorID); err != nil {
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

func CreateSwarmNexus(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, nexusName, description, location, configPath string
	var conf *configuration.Config
	var nexus *api.Nexus

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusName, err = cmd.Flags().GetString("nexus-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if location, err = cmd.Flags().GetString("location"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	nexusBodyRequest := api.CreateNexusRequestBody{
		Name:        nexusName,
		Description: description,
		Location:    location,
	}

	if nexus, err = api.CreateNexus(conf.Urls, *accessToken, id, nexusBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingNexusRequest, err)
	}

	utils.PrintCreateSuccess("nexus", nexus.ID)

	return nil
}

func DescribeSwarmNexus(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath, format string
	var conf *configuration.Config
	var nexus *api.Nexus

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	nexusID := args[0]
	if nexus, err = api.GetNexus(conf.Urls, *accessToken, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingNexusRequest, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintFormattedData(*nexus, format)

	return nil
}

func EditSwarmNexus(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var nexusName, description, configPath string
	var conf *configuration.Config

	if nexusName, err = cmd.Flags().GetString("nexus-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	nexusID := args[0]

	nexusBodyRequest := api.UpdateNexusRequestBody{
		Name:        nexusName,
		Description: description,
	}

	if err = api.UpdateNexus(conf.Urls, *accessToken, nexusID, nexusBodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingNexusRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("nexus %s updated successfully", nexusID))

	return nil
}

func ListSwarmNexuses(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var nexuses *api.NexusList
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
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	utils.PrintList("Your Swarm Nexuses List")

	if verbose {
		utils.PrintVerbose(nexuses.Nexuses, l)
	} else {
		var IDs []string
		for _, nexus := range nexuses.Nexuses {
			IDs = append(IDs, nexus.ID)
		}
		utils.PrintSimpleList(IDs)
	}

	return nil
}

func RemoveSwarmNexus(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	nexusID := args[0]

	if err = api.DeleteNexus(conf.Urls, *accessToken, nexusID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingNexusRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("nexus %s removed successfully", nexusID))

	return nil
}

func getSwarmByNameOrId(conf *configuration.Config, accessToken string, swarm string) (string, error) {
	var err error
	var operator *api.Operator
	var swarms []*api.Swarm
	var id string

	if operator, err = api.GetOperatorSelf(conf.Urls, accessToken); err != nil {
		return id, fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if swarms, err = api.ListSwarms(conf.Urls, accessToken, operator.ID); err != nil {
		return "", fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	for _, sw := range swarms {
		if swarm == sw.Name || swarm == sw.ID {
			id = sw.ID
		}
	}

	if id == "" {
		return "", fmt.Errorf("swarm %s not found", swarm)
	}

	return id, nil
}

func getSwarmOperatorByEmailOrId(conf *configuration.Config, accessToken string, tenantID string, operator string) (*api.Operator, error) {
	var err error
	var operators *api.OperatorList

	if operators, err = api.ListSwarmOperators(conf.Urls, accessToken, tenantID); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}
	for _, op := range operators.Operators {
		if operator == op.Email || operator == op.ID {
			return op, nil
		}
	}

	return nil, fmt.Errorf("operator %s not found", operator)
}
