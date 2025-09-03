// Package action provides CLI actions for managing redundancy classes.
package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateRedundancyClass(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, name, description string
	var innerK, innerN, outerK, outerN, antiAffinityGroup int
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var nexuses []string
	var redundancyClass *api.RedundancyClass

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if innerK, err = cmd.Flags().GetInt("inner-k"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if innerN, err = cmd.Flags().GetInt("inner-n"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if outerK, err = cmd.Flags().GetInt("outer-k"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if outerN, err = cmd.Flags().GetInt("outer-n"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if antiAffinityGroup, err = cmd.Flags().GetInt("anti-affinity-group"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexuses, err = cmd.Flags().GetStringSlice("nexuses"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	bodyRequest := api.CreateRedundancyClassRequestBody{
		Name:              name,
		Description:       description,
		InnerK:            innerK,
		InnerN:            innerN,
		OuterK:            outerK,
		OuterN:            outerN,
		AntiAffinityGroup: antiAffinityGroup,
		Nexuses:           nexuses,
	}

	if redundancyClass, err = api.CreateRedundancyClass(*urls, resolvedProfile.APIKey, swarmID, bodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingRedundancyClassRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.CreateRedundancyClassRequestBody{
			bodyRequest,
		},
		func(rc api.CreateRedundancyClassRequestBody) []string {
			return []string{
				redundancyClass.ID,
			}
		},
		&utils.SmartOutputConfig[api.CreateRedundancyClassRequestBody]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func DescribeRedundancyClass(cmd *cobra.Command, args []string) error {
	var err error
	var redundancyClassID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var redundancyClass *api.RedundancyClass

	if redundancyClassID, err = cmd.Flags().GetString("rc-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if redundancyClass, err = api.GetRedundancyClass(*urls, resolvedProfile.APIKey, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.RedundancyClass{redundancyClass},
		func(rc *api.RedundancyClass) []string {
			return []string{
				rc.ID,
				rc.Name,
				rc.Description,
			}
		},
		&utils.SmartOutputConfig[*api.RedundancyClass]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func ListRedundancyClasses(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var redundancyClasses *api.GenericPaginatedResponse[*api.RedundancyClass]

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
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

	if redundancyClasses, err = api.ListRedundancyClasses(*urls, resolvedProfile.APIKey, swarmID, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		redundancyClasses.Data,
		func(rc *api.RedundancyClass) []string {
			return []string{
				rc.ID,
			}
		},
		&utils.SmartOutputConfig[*api.RedundancyClass]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func CheckRedundancyClassStatus(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, redundancyClassID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var status *api.SummaryDetailsWithStatusNullable

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassID, err = cmd.Flags().GetString("rc-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if status, err = api.CheckRedundancyClassStatus(*urls, resolvedProfile.APIKey, swarmID, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCheckingRedundancyClassStatusRequest, err)
	}

	humanReadableStatus := status.ToHumanReadableStatus("ring")

	return utils.PrintSmartOutput(
		cmd,
		[]api.HumanReadableStatus{humanReadableStatus},
		func(s api.HumanReadableStatus) []string {
			return []string{
				s.Status,
			}
		},
		&utils.SmartOutputConfig[api.HumanReadableStatus]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func ExpandRedundancyClass(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, redundancyClassID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var redundancyClassExpanded *api.RedundancyClassExpanded
	var dryRun bool

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassID, err = cmd.Flags().GetString("rc-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if dryRun, err = cmd.Flags().GetBool("dry-run"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if redundancyClassExpanded, err = api.ExpandRedundancyClass(*urls, resolvedProfile.APIKey, swarmID, redundancyClassID, dryRun); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorExpandingRedundancyClassRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.RedundancyClassExpanded{redundancyClassExpanded},
		func(rc *api.RedundancyClassExpanded) []string {
			return []string{
				string(rc.Status),
				rc.Message,
			}
		},
		&utils.SmartOutputConfig[*api.RedundancyClassExpanded]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func RecoverRedundancyClass(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, redundancyClassID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var redundancyClassRecover *api.RedundancyClassRecovery
	var dryRun bool

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassID, err = cmd.Flags().GetString("rc-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if dryRun, err = cmd.Flags().GetBool("dry-run"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if redundancyClassRecover, err = api.RecoverRedundancyClass(*urls, resolvedProfile.APIKey, swarmID, redundancyClassID, dryRun); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRecoveringRedundancyClassRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.RedundancyClassRecovery{redundancyClassRecover},
		func(rc *api.RedundancyClassRecovery) []string {
			return []string{
				rc.Message,
			}
		},
		&utils.SmartOutputConfig[*api.RedundancyClassRecovery]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func CheckRedundancyClassRecoveryStatus(cmd *cobra.Command, args []string) error {
	var err error
	var swarmID, redundancyClassID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var status *api.RCProgress

	if swarmID, err = cmd.Flags().GetString("swarm-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassID, err = cmd.Flags().GetString("rc-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if status, err = api.CheckRedundancyClassRecoveryStatus(*urls, resolvedProfile.APIKey, swarmID, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCheckingRedundancyClassRecoveryStatusRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.RCProgress{status},
		func(s *api.RCProgress) []string {
			return []string{string(s.Session.Status)}
		},
		&utils.SmartOutputConfig[*api.RCProgress]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}
