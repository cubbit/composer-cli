// Package action provides CLI actions for managing swarms.
package action

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var name, description string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var response *api.GenericIDResponseModel

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	req := api.CreateSwarmRequest{
		Name:          name,
		Description:   &description,
		Configuration: map[string]interface{}{},
	}

	if response, err = api.CreateSwarm(*urls, resolvedProfile.APIKey, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingSwarmRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.CreateSwarmRequest{req},
		func(s api.CreateSwarmRequest) []string { return []string{response.ID} },
		&utils.SmartOutputConfig[api.CreateSwarmRequest]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func DescribeSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var id string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var operator *api.Operator
	var swarm *api.Swarm

	if id, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if operator, err = api.GetIAMUserSelf(*urls, "", resolvedProfile.APIKey); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if swarm, err = api.GetSwarm(*urls, resolvedProfile.APIKey, operator.ID, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.Swarm{swarm},
		func(s *api.Swarm) []string { return []string{s.ID, s.Name, utils.StringOrEmpty(s.Description)} },
		&utils.SmartOutputConfig[*api.Swarm]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func EditSwarm(cmd *cobra.Command, args ...string) error {
	var err error
	var id, name, description string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var req api.UpdateSwarmRequest

	if id, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if name != "" {
		req.Name = &name
	}

	if description != "" {
		req.Description = &description
	}

	if err = api.EditSwarm(*urls, resolvedProfile.APIKey, id, req); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.UpdateSwarmRequest{&req},
		func(s *api.UpdateSwarmRequest) []string { return []string{id} },
		&utils.SmartOutputConfig[*api.UpdateSwarmRequest]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func CheckSwarmStatus(cmd *cobra.Command, args []string) error {
	var err error
	var id string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var status *api.SummaryDetailsWithStatusNullable

	if id, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if status, err = api.GetSwarmStatus(*urls, resolvedProfile.APIKey, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmStatusRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.SummaryDetailsWithStatusNullable{status},
		func(s *api.SummaryDetailsWithStatusNullable) []string {
			status := s.SummaryStatusNullable
			if status.EvaluatedStatus == nil {
				return []string{"No status available"}
			}

			return []string{string(*s.SummaryStatusNullable.EvaluatedStatus)}
		},
		&utils.SmartOutputConfig[*api.SummaryDetailsWithStatusNullable]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func ListSwarms(cmd *cobra.Command, args []string) error {
	var err error
	var sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var swarms *api.GenericPaginatedResponse[*api.Swarm]

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

	if swarms, err = api.ListSwarmsV2(*urls, resolvedProfile.APIKey, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		swarms.Data,
		func(s *api.Swarm) []string { return []string{s.ID} },
		&utils.SmartOutputConfig[*api.Swarm]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func RemoveSwarm(cmd *cobra.Command, args []string) error {
	var err error
	var id string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if id, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.RemoveSwarm(*urls, resolvedProfile.APIKey, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingSwarmRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{id},
		func(s string) []string { return []string{s} },
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}
