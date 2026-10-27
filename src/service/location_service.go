package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/user/shared"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type LocationServiceInterface interface {
	List(cmd *cobra.Command, args []string) error
	Describe(cmd *cobra.Command, args []string) error
	Update(cmd *cobra.Command, args []string) error
	Delete(cmd *cobra.Command, args []string) error
	CreateVirtual(cmd *cobra.Command, args []string) error
	CreateVirtualNode(cmd *cobra.Command, args []string) error
}

type LocationService struct {
	configuration configuration_handler.ConfigurationHandlerInterface
	locationAPI   api.LocationAPIInterface
	userAPI       api.UserAPIInterface
	processAPI    api.ProcessAPIInterface
}

func NewLocationService(
	configuration configuration_handler.ConfigurationHandlerInterface,
	locationAPI api.LocationAPIInterface,
	userAPI api.UserAPIInterface,
	processAPI api.ProcessAPIInterface,
) LocationService {
	return LocationService{
		configuration: configuration,
		locationAPI:   locationAPI,
		userAPI:       userAPI,
		processAPI:    processAPI,
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

func (s LocationService) Describe(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
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

	// Describe reads the location through the single-location endpoint. When the user filters by
	// name (there is no lookup-by-name endpoint) the aggregated list is used only to resolve the id.
	if clusterName != "" {
		clusters, err := s.locationAPI.ListAggregated(profile.Endpoints, profile.APIKey, profile.OrganizationID)
		if err != nil {
			return fmt.Errorf("failed to list aggregated locations: %w", err)
		}

		clusterID, err = resolveClusterIDByName(clusters, clusterName)
		if err != nil {
			return err
		}
	}

	cluster, err := s.locationAPI.GetAggregated(profile.Endpoints, profile.APIKey, profile.OrganizationID, clusterID)
	if err != nil {
		return err
	}

	return PrintClusterDetails(cmd, s.configuration, *cluster)
}

// resolveClusterIDByName resolves a location from its name. The name is the user-facing label of the
// location: the display name when set, the Kubernetes cluster name otherwise.
func resolveClusterIDByName(clusters []api.InfraAggregateCluster, name string) (string, error) {
	matches := make([]api.InfraAggregateCluster, 0, 1)

	for _, cluster := range clusters {
		if cluster.Name == name {
			matches = append(matches, cluster)
		}
	}

	switch len(matches) {
	case 0:
		return "", fmt.Errorf("cluster with name '%s' not found", name)
	case 1:
		return matches[0].ClusterID, nil
	default:
		ids := make([]string, len(matches))
		for i, match := range matches {
			ids[i] = match.ClusterID
		}

		return "", fmt.Errorf("several clusters match the name '%s': %s", name, strings.Join(ids, ", "))
	}
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

	// The creation output keeps only the identity of the new location: the metadata fields returned by
	// the list endpoint are not part of it.
	created := []createdLocationOutput{{
		ClusterID: location.ClusterID,
		Name:      location.Name,
		Type:      location.Type,
	}}

	return utils.PrintSmartOutput(
		cmd,
		created,
		nil,
		&utils.SmartOutputConfig[createdLocationOutput]{
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

func (s LocationService) Update(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	clusterID, err := cmd.Flags().GetString("cluster-id")
	if err != nil {
		return fmt.Errorf("%s cluster-id: %w", constants.ErrorRetrievingField, err)
	}

	request := api.UpdateInfraClusterRequest{}

	if request.DisplayName, err = optionalStringFlag(cmd, "name"); err != nil {
		return err
	}

	if request.Description, err = optionalStringFlag(cmd, "description"); err != nil {
		return err
	}

	if request.Position, err = optionalStringFlag(cmd, "position"); err != nil {
		return err
	}

	if cmd.Flags().Changed("contact-user") {
		contactUser, err := optionalStringFlag(cmd, "contact-user")
		if err != nil {
			return err
		}

		contactUserID := *contactUser
		// An empty value clears the contact: only a provided user has to be resolved.
		if contactUserID != "" {
			if contactUserID, err = s.resolveContactUserID(profile, contactUserID); err != nil {
				return err
			}
		}

		request.ContactIAMUserID = &contactUserID
	}

	if request.IsEmpty() {
		return fmt.Errorf("no fields to update: provide at least one of --name, --description, --position or --contact-user")
	}

	if err := s.locationAPI.Update(profile.Endpoints, profile.APIKey, profile.OrganizationID, clusterID, &request); err != nil {
		var conflict *api.InfraClusterConflictError
		if errors.As(err, &conflict) {
			return formatInfraClusterConflict(clusterID, conflict)
		}

		return err
	}

	return printer.PrintText(cmd, s.configuration, fmt.Sprintf("Location %s updated successfully\n", clusterID))
}

func (s LocationService) Delete(cmd *cobra.Command, args []string) error {
	profile, err := s.configuration.GetActiveProfile()
	if err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	clusterID, err := cmd.Flags().GetString("cluster-id")
	if err != nil {
		return fmt.Errorf("%s cluster-id: %w", constants.ErrorRetrievingField, err)
	}

	response, err := s.locationAPI.Delete(profile.Endpoints, profile.APIKey, profile.OrganizationID, clusterID)
	if err != nil {
		var conflict *api.InfraClusterConflictError
		if errors.As(err, &conflict) {
			return formatInfraClusterConflict(clusterID, conflict)
		}

		return err
	}

	if response.ID == "" {
		return fmt.Errorf("location %s deletion did not return a process id", clusterID)
	}

	// The progress line is only emitted in human output; json/yaml emit a single document at the
	// end (see waitForClusterDeletion).
	if printer.IsHumanOutput(cmd, s.configuration) {
		if err := printer.PrintText(cmd, s.configuration, fmt.Sprintf("Location %s deletion started (process %s)\n", clusterID, response.ID)); err != nil {
			return err
		}
	}

	return s.waitForClusterDeletion(cmd, profile, clusterID, response.ID)
}

// optionalStringFlag returns the value of a string flag when it was explicitly provided, or nil when
// it was not. An explicitly empty value is a valid clear: the backend stores it as an absent field.
func optionalStringFlag(cmd *cobra.Command, name string) (*string, error) {
	if !cmd.Flags().Changed(name) {
		return nil, nil
	}

	value, err := cmd.Flags().GetString(name)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %w", constants.ErrorRetrievingField, name, err)
	}

	return &value, nil
}

// resolveContactUserID accepts either the IAM user id (uuid) or the username (both unique within
// an organization) and returns the user id expected by the update endpoint.
func (s LocationService) resolveContactUserID(profile configuration_models.ProfileV2, value string) (string, error) {
	if _, err := uuid.Parse(value); err == nil {
		return value, nil
	}

	return shared.ResolveUserIDByUsername(s.userAPI, profile, value)
}

var (
	clusterDeletionPollInterval = 3 * time.Second
	clusterDeletionTimeout      = 1 * time.Hour
)

// waitForClusterDeletion follows the asynchronous cluster deletion process until it reaches a
// terminal state. It is intentionally non-interactive so the command can be used in scripts.
func (s LocationService) waitForClusterDeletion(cmd *cobra.Command, profile configuration_models.ProfileV2, clusterID string, processID string) error {
	deadline := time.Now().Add(clusterDeletionTimeout)

	for {
		process, err := s.processAPI.GetProcess(profile.Endpoints, profile.APIKey, profile.OrganizationID, processID)
		if err != nil {
			return fmt.Errorf("failed to poll the location deletion process: %w", err)
		}

		if process.Type != api.ProcessTypeClusterDeletion {
			return fmt.Errorf("unexpected process type: expected %s, got %s", api.ProcessTypeClusterDeletion, process.Type)
		}

		switch process.Status {
		case api.ProcessStatusSuccess:
			return printer.PrintText(cmd, s.configuration, fmt.Sprintf("Location %s deleted successfully\n", clusterID))
		case api.ProcessStatusFailed:
			if deletion, ok := process.CastToClusterDeletionProcess(); ok && deletion.Data.Error != nil {
				return fmt.Errorf("location deletion failed: %s (%s)", deletion.Data.Error.Message, deletion.Data.Error.Code)
			}

			return fmt.Errorf("location deletion failed")
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("location deletion timed out after %s", clusterDeletionTimeout)
		}

		time.Sleep(clusterDeletionPollInterval)
	}
}

// formatInfraClusterConflict maps the Composer Hub 409 conflict reasons to actionable messages.
func formatInfraClusterConflict(clusterID string, conflict *api.InfraClusterConflictError) error {
	switch conflict.Reason {
	case api.InfraClusterConflictReasonProfileInstalled:
		if conflict.ConflictsWith != "" {
			return fmt.Errorf("location %s cannot be deleted: the profile %s is still installed; delete the gateway or swarm using it first", clusterID, conflict.ConflictsWith)
		}

		return fmt.Errorf("location %s cannot be deleted: a profile is still installed; delete the gateway or swarm using it first", clusterID)
	case api.InfraClusterConflictReasonDeletionInProgress:
		return fmt.Errorf("a deletion is already in progress for location %s", clusterID)
	default:
		return fmt.Errorf("location %s is in a conflicting state: %s", clusterID, conflict.Reason)
	}
}
