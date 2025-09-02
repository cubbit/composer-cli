// Package action provides CLI actions for managing operators.
package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func PromoteOperator(cmd *cobra.Command, args []string) error {
	var url *configuration.URLs
	var err error
	var email, policyName, apiServerURL, secret string

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if policyName, err = cmd.Flags().GetString("policy-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if apiServerURL, err = cmd.Flags().GetString("api-server-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if secret, err = cmd.Flags().GetString("secret"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if url, err = configuration.ConfigureAPIServerURL(apiServerURL); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if err = api.PromoteOperator(*url, email, policyName, secret); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorPromotingOperatorRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{email},
		nil,
		nil,
	)
}

func AddOperatorToTenant(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, email, policyID, firstName, lastName string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
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

	if policyID, err = cmd.Flags().GetString("policy-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	req := api.InviteOperatorRequestBody{
		Email:     email,
		PolicyID:  policyID,
		FirstName: &firstName,
		LastName:  &lastName,
	}

	if err = api.InviteOperatorToTenant(*urls, resolvedProfile.APIKey, tenantID, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.InviteOperatorRequestBody{req},
		nil,
		&utils.SmartOutputConfig[api.InviteOperatorRequestBody]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func ListTenantOperators(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var operators *api.OperatorList

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if operators, err = api.ListTenantOperators(*urls, resolvedProfile.APIKey, tenantID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		operators.Operators,
		func(o *api.Operator) []string {
			return []string{o.ID}
		},
		&utils.SmartOutputConfig[*api.Operator]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func RemoveTenantOperator(cmd *cobra.Command, args []string) error {
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

	if err = api.RemoveTenantOperator(*urls, resolvedProfile.APIKey, tenantID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingOperatorRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{userID},
		func(s string) []string {
			return []string{s}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		})
}

func DescribeTenantOperator(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var operator *api.Operator

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

	if operator, err = api.GetTenantOperator(*urls, resolvedProfile.APIKey, tenantID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.Operator{operator},
		func(o *api.Operator) []string {
			return []string{o.ID,
				o.Email,
				o.FirstName,
				o.LastName,
				o.PolicyName,
			}
		},
		&utils.SmartOutputConfig[*api.Operator]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func EditTenantOperatorRole(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, userID, policyID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if userID, err = cmd.Flags().GetString("user-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if policyID, err = cmd.Flags().GetString("policy-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	req := api.ChangeOperatorPolicyRequestBody{
		PolicyID: policyID,
	}

	if err = api.EditOperatorRoleInTenant(*urls, resolvedProfile.APIKey, tenantID, userID, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.ChangeOperatorPolicyRequestBody{req},
		func(s api.ChangeOperatorPolicyRequestBody) []string {
			return []string{
				req.PolicyID,
			}
		},
		&utils.SmartOutputConfig[api.ChangeOperatorPolicyRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func AddOperatorToSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, email, policyID, firstName, lastName string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if policyID, err = cmd.Flags().GetString("policy-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if firstName, err = cmd.Flags().GetString("first-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if lastName, err = cmd.Flags().GetString("last-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	req := api.InviteOperatorRequestBody{
		Email:     email,
		PolicyID:  policyID,
		FirstName: &firstName,
		LastName:  &lastName,
	}

	if err = api.InviteOperatorToSwarm(*urls, resolvedProfile.APIKey, swarmID, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.InviteOperatorRequestBody{req},
		nil,
		&utils.SmartOutputConfig[api.InviteOperatorRequestBody]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func ListSwarmOperators(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var operators *api.OperatorList

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if operators, err = api.ListSwarmOperators(*urls, resolvedProfile.APIKey, swarmID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingOperatorsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		operators.Operators,
		func(o *api.Operator) []string {
			return []string{
				o.ID,
			}
		},
		&utils.SmartOutputConfig[*api.Operator]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func RemoveSwarmOperator(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if err = api.RemoveSwarmOperator(*urls, resolvedProfile.APIKey, swarmID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingOperatorRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{userID},
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

func DescribeSwarmOperator(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var operator *api.Operator

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if operator, err = api.GetSwarmOperator(*urls, resolvedProfile.APIKey, swarmID, userID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.Operator{operator},
		func(o *api.Operator) []string {
			return []string{o.ID, o.Email, o.FirstName, o.LastName, o.PolicyName}
		},
		&utils.SmartOutputConfig[*api.Operator]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func EditSwarmOperatorRole(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, policyID, userID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if policyID, err = cmd.Flags().GetString("policy-id"); err != nil {
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

	req := api.ChangeOperatorPolicyRequestBody{
		PolicyID: policyID,
	}

	if err = api.EditOperatorRoleInSwarm(*urls, resolvedProfile.APIKey, swarmID, userID, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorInvitingOperatorRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.ChangeOperatorPolicyRequestBody{req},
		func(s api.ChangeOperatorPolicyRequestBody) []string {
			return []string{
				req.PolicyID,
			}
		},
		&utils.SmartOutputConfig[api.ChangeOperatorPolicyRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}
