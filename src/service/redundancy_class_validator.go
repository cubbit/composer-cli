package service

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/api"
)

type RedundancyClassValidatorInterface interface {
	ValidateRedundancyClasses(
		redundancyClasses []api.RedundancyClassRequest,
		nexuses []api.NexusV5Request,
	) error
}

type RedundancyClassValidator struct{}

func NewRedundancyClassValidator() *RedundancyClassValidator {
	return &RedundancyClassValidator{}
}

func (v *RedundancyClassValidator) ValidateRedundancyClasses(
	redundancyClasses []api.RedundancyClassRequest,
	nexuses []api.NexusV5Request,
) error {
	if err := v.validateUniqueNames(redundancyClasses); err != nil {
		return err
	}

	allClusterIDs := v.collectAllClusterIDs(nexuses)

	for _, rc := range redundancyClasses {
		if err := v.validateClusterIDsSubset(rc, allClusterIDs); err != nil {
			return err
		}

		if err := v.validateClusterIDsCount(rc); err != nil {
			return err
		}
	}

	return nil
}

func (v *RedundancyClassValidator) validateUniqueNames(redundancyClasses []api.RedundancyClassRequest) error {
	seen := make(map[string]bool)
	for _, rc := range redundancyClasses {
		if seen[rc.Name] {
			return fmt.Errorf("redundancy class name '%s' is duplicated - all names must be unique", rc.Name)
		}
		seen[rc.Name] = true
	}
	return nil
}

func (v *RedundancyClassValidator) collectAllClusterIDs(nexuses []api.NexusV5Request) map[string]bool {
	allClusterIDs := make(map[string]bool)
	for _, nexus := range nexuses {
		allClusterIDs[nexus.ClusterID] = true
	}
	return allClusterIDs
}

func (v *RedundancyClassValidator) validateClusterIDsSubset(rc api.RedundancyClassRequest, allClusterIDs map[string]bool) error {
	for _, clusterID := range rc.ClusterIDs {
		if !allClusterIDs[clusterID] {
			return fmt.Errorf("redundancy class '%s' references cluster '%s' which is not in any nexus - all RC cluster_ids must be a subset of nexus cluster-ids", rc.Name, clusterID)
		}
	}
	return nil
}

func (v *RedundancyClassValidator) validateClusterIDsCount(rc api.RedundancyClassRequest) error {
	expectedCount := rc.OuterN + rc.OuterK
	actualCount := len(rc.ClusterIDs)
	if actualCount != expectedCount {
		return fmt.Errorf("redundancy class '%s' has %d cluster_ids but outer_n + outer_k = %d - cluster_ids count must equal outer_n + outer_k", rc.Name, actualCount, expectedCount)
	}
	return nil
}
