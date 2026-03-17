package service

import (
	"fmt"

	"encoding/json"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type LocationServiceInterface interface {
	List(cmd *cobra.Command, args []string) error
	ListAggregated(cmd *cobra.Command, args []string) error
	CreateVirtual(cmd *cobra.Command, args []string) error
	CreateVirtualNode(cmd *cobra.Command, args []string) error
}

type LocationService struct {
	configuration configuration.ConfigInterface
	locationAPI   api.LocationAPIInterface
	userAPI       api.UserAPIInterface
}

func NewLocationService(
	configuration configuration.ConfigInterface,
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

	locations, err := s.locationAPI.List(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID)
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

func (s LocationService) CreateVirtual(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
	}

	description, err := utils.GetOptionalStringFlag(cmd, "description")
	if err != nil {
		return fmt.Errorf("%s description: %w", constants.ErrorRetrievingField, err)
	}

	location, err := s.locationAPI.CreateVirtualCluster(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, name, description)
	if err != nil {
		return fmt.Errorf("failed to create virtual location: %w", err)
	}

	return utils.PrintSmartOutput(
		cmd,
		[]api.InfrastructureCluster{*location},
		nil,
		&utils.SmartOutputConfig[api.InfrastructureCluster]{
			SingleResourceCompactOutput: false,
			SingleResource:              false,
			DefaultOutput:               resolvedProfile.Output,
		},
	)
}

func (s LocationService) CreateVirtualNode(cmd *cobra.Command, args []string) error {
	resolvedProfile, urls, err := s.configuration.ResolveProfileAndURLs(cmd, configuration.ProfileTypeComposer)
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
	}
	clusterID, err := cmd.Flags().GetString("cluster-id")
	if err != nil {
		return fmt.Errorf("%s cluster-id: %w", constants.ErrorRetrievingField, err)
	}
	storageType, err := cmd.Flags().GetString("storage-type")
	if err != nil {
		return fmt.Errorf("%s storage-type: %w", constants.ErrorRetrievingField, err)
	}
	configurationStr, err := cmd.Flags().GetString("configuration")
	if err != nil {
		return fmt.Errorf("%s configuration: %w", constants.ErrorRetrievingField, err)
	}

	var configuration map[string]any
	err = json.Unmarshal([]byte(configurationStr), &configuration)
	if err != nil {
		return fmt.Errorf("%s configuration: %w", constants.ErrorParsingJSONConfiguration, err)
	}

	node, err := s.locationAPI.CreateVirtualNode(*urls, resolvedProfile.APIKey, resolvedProfile.OrganizationID, clusterID, name, storageType, configuration)
	if err != nil {
		return fmt.Errorf("failed to create virtual node: %w", err)
	}

	printFuncs := []func() error{
		func() error { return PrintVirtualNodes(cmd, []api.InfraAggregateVirtualNodeDetail{*node}) },
	}

	return printer.Compose(cmd, printFuncs...)
}
