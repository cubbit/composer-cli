package service

import (
	"fmt"

	"encoding/json"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

type LocationServiceInterface interface {
	List(cmd *cobra.Command, args []string) error
	ListAggregated(cmd *cobra.Command, args []string) error
	CreateVirtual(cmd *cobra.Command, args []string) error
	CreateVirtualNode(cmd *cobra.Command, args []string) error
}

type LocationService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	locationAPI   api.LocationAPIInterface
	userAPI       api.UserAPIInterface
}

func NewLocationService(
	configuration configuration_handler.ConfigurationHandlerInterface,
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
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	locations, err := s.locationAPI.List(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
	)
	if err != nil {
		return fmt.Errorf("failed to list locations: %w", err)
	}

	return PrintClusters(cmd, s.configuration, locations)
}

func (s LocationService) ListAggregated(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	clusters, err := s.locationAPI.ListAggregated(profile.Endpoints, profile.APIKey, profile.OrganizationID)
	if err != nil {
		return fmt.Errorf("failed to list aggregated locations: %w", err)
	}

	clusterName, err := cmd.Flags().GetString("cluster-name")
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}
	clusterID, err := cmd.Flags().GetString("cluster-id")
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if clusterName != "" && clusterID != "" {
		return fmt.Errorf("cluster-name and cluster-id filters are mutually exclusive; please provide only one of them")
	}

	if clusterName == "" && clusterID == "" {
		return fmt.Errorf("either cluster-name or cluster-id must be provided as a filter")
	}

	var filteredCluster *api.InfraAggregateCluster

	if clusterID != "" {
		if cluster, found := utils.Find(clusters, func(c api.InfraAggregateCluster) bool {
			return c.ClusterID == clusterID
		}); found {
			filteredCluster = &cluster
		}
	}

	if clusterName != "" {
		if cluster, found := utils.Find(clusters, func(c api.InfraAggregateCluster) bool {
			return c.Name == clusterName
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

	return PrintClusterDetails(cmd, s.configuration, *filteredCluster)
}

func (s LocationService) CreateVirtual(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
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

	location, err := s.locationAPI.CreateVirtualCluster(profile.Endpoints, profile.APIKey, profile.OrganizationID, name, description)
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
			DefaultOutput:               profile.Output,
		},
	)
}

func (s LocationService) CreateVirtualNode(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
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

	node, err := s.locationAPI.CreateVirtualNode(profile.Endpoints, profile.APIKey, profile.OrganizationID, clusterID, name, storageType, configuration)
	if err != nil {
		return fmt.Errorf("failed to create virtual node: %w", err)
	}

	return PrintVirtualNodes(cmd, s.configuration, []api.InfraAggregateVirtualNodeDetail{*node})
}
