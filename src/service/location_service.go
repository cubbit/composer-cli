package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type LocationServiceInterface interface {
	List(cmd *cobra.Command, args []string) error
	ListAggregated(cmd *cobra.Command, args []string) error
}

type LocationService struct {
	configuration *configuration.Config
	locationAPI   api.LocationAPIInterface
	userAPI       api.UserAPIInterface
}

func NewLocationService(
	configuration *configuration.Config,
	locationAPI api.LocationAPIInterface,
	userAPI api.UserAPIInterface,
) LocationService {
	return LocationService{
		configuration: configuration,
		locationAPI:   locationAPI,
		userAPI:       userAPI,
	}
}

func (s LocationService) List(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	user, err := s.userAPI.GetIAMUserSelf(*urls, "", resolvedProfile.APIKey)
	if err != nil {
		return fmt.Errorf("failed to get iam user information: %w", err)
	}

	if user.OrganizationID == nil {
		return fmt.Errorf("user does not belong to an organization")
	}

	locations, err := s.locationAPI.List(*urls, resolvedProfile.APIKey, *user.OrganizationID)
	if err != nil {
		return fmt.Errorf("failed to list locations: %w", err)
	}

	return utils.PrintSmartOutput(
		cmd,
		locations,
		nil,
		&utils.SmartOutputConfig[api.InfrastructureCluster]{
			SingleResourceCompactOutput: false,
			SingleResource:              false,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func (s LocationService) ListAggregated(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	user, err := s.userAPI.GetIAMUserSelf(*urls, "", resolvedProfile.APIKey)
	if err != nil {
		return fmt.Errorf("failed to get iam user information: %w", err)
	}

	if user.OrganizationID == nil {
		return fmt.Errorf("user does not belong to an organization")
	}

	clusters, err := s.locationAPI.ListAggregated(*urls, resolvedProfile.APIKey, *user.OrganizationID)
	if err != nil {
		return fmt.Errorf("failed to list aggregated locations: %w", err)
	}

	if len(clusters) == 0 {
		utils.PrintInfo("No aggregated locations found")
		return nil
	}

	clusterName, err := cmd.Flags().GetString("cluster-name")
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	clusterID, err := cmd.Flags().GetString("cluster-id")
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if clusterName != "" || clusterID != "" {
		var filteredCluster *api.InfraAggregateCluster

		if clusterName != "" {
			if cluster, found := utils.Find(clusters, func(c api.InfraAggregateCluster) bool {
				return c.Name == clusterName
			}); found {
				filteredCluster = &cluster
			}
		} else {
			if cluster, found := utils.Find(clusters, func(c api.InfraAggregateCluster) bool {
				return c.ClusterID == clusterID
			}); found {
				filteredCluster = &cluster
			}
		}

		if filteredCluster == nil {
			if clusterName != "" {
				return fmt.Errorf("cluster with name '%s' not found", clusterName)
			}
			return fmt.Errorf("cluster with ID '%s' not found", clusterID)
		}

		return PrintClusterDetails(cmd, *filteredCluster)
	}

	return PrintClusters(cmd, clusters)
}
