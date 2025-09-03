// Package action provides CLI actions for managing tenantprojects.
package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func ListTenantProjects(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, sort, filter string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var projects *api.GenericPaginatedResponse[*api.ProjectItem]

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter, err = cmd.Flags().GetString("filter"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if filter != "" {
		filter = utils.BuildFilterQuery(filter)
	}

	if sort, err = cmd.Flags().GetString("sort"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if projects, err = api.ListTenantProjects(*urls, resolvedProfile.APIKey, tenantID, sort, filter); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingTenantProjectsRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		projects.Data,
		func(p *api.ProjectItem) []string {
			return []string{
				p.ProjectID,
			}
		},
		&utils.SmartOutputConfig[*api.ProjectItem]{
			DefaultOutput: resolvedProfile.Output,
		},
	)
}

func DescribeTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, projectID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs
	var project *api.ProjectItem

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if projectID, err = cmd.Flags().GetString("project-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if project, err = api.GetTenantProject(*urls, resolvedProfile.APIKey, tenantID, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenantProjectRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]*api.ProjectItem{project},
		func(p *api.ProjectItem) []string {
			return []string{
				p.ProjectID,
				p.ProjectName,
				p.ProjectDescription,
			}
		},
		&utils.SmartOutputConfig[*api.ProjectItem]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}

func RemoveTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, projectID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if projectID, err = cmd.Flags().GetString("project-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.RemoveTenantProject(*urls, resolvedProfile.APIKey, tenantID, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingTenantProjectRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{projectID},
		func(s string) []string {
			return []string{
				s,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func BanTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, projectID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if projectID, err = cmd.Flags().GetString("project-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.ToggleBanProject(*urls, resolvedProfile.APIKey, tenantID, projectID, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorFreezingTenantProjectRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{projectID},
		func(s string) []string {
			return []string{
				s,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func UnbanTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, projectID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if projectID, err = cmd.Flags().GetString("project-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.ToggleBanProject(*urls, resolvedProfile.APIKey, tenantID, projectID, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUnfreezingTenantProjectRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{projectID},
		func(s string) []string {
			return []string{
				s,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func RestoreTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, projectID string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if projectID, err = cmd.Flags().GetString("project-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if err = api.RestoreTenantProject(*urls, resolvedProfile.APIKey, tenantID, projectID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRestoringTenantProjectRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]string{projectID},
		func(s string) []string {
			return []string{
				s,
			}
		},
		&utils.SmartOutputConfig[string]{
			SingleResource:              true,
			SingleResourceCompactOutput: true,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func UpdateTenantProject(cmd *cobra.Command, args []string) error {
	var err error
	var tenantID, projectID, name, description, imageURL string
	var conf *configuration.Config
	var resolvedProfile *configuration.ResolvedProfile
	var urls *configuration.URLs

	if tenantID, err = cmd.Flags().GetString("tenant-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if projectID, err = cmd.Flags().GetString("project-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if imageURL, err = cmd.Flags().GetString("image-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, err = configuration.LoadConfig(); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if resolvedProfile, urls, err = conf.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	var namePtr *string
	if name != "" {
		namePtr = &name
	}

	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}

	var imageURLPtr *string
	if imageURL != "" {
		imageURLPtr = &imageURL
	}

	requestBody := api.UpdateTenantProjectRequestBody{
		Name:        namePtr,
		Description: descriptionPtr,
		ImageURL:    imageURLPtr,
	}

	if err = api.UpdateProject(*urls, resolvedProfile.APIKey, tenantID, projectID, requestBody); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorUpdatingTenantProjectRequest, err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.UpdateTenantProjectRequestBody{requestBody},
		func(r api.UpdateTenantProjectRequestBody) []string {
			return []string{
				projectID,
			}
		},
		&utils.SmartOutputConfig[api.UpdateTenantProjectRequestBody]{
			SingleResource: true,
			DefaultOutput:  resolvedProfile.Output,
		},
	)
}
