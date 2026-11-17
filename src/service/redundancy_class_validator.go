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

type RCValidationErrors struct {
	AntiAffinityGroup string
	OuterK            string
	LocalNK           string
	LocalK            string
	Name              string
}

func ValidateAAG(aag, minDisksPerNode int) string {
	if aag < 1 {
		return "AAG must be at least 1"
	}
	if aag > minDisksPerNode {
		return fmt.Sprintf("AAG cannot exceed minimum disks per node (%d)", minDisksPerNode)
	}
	return ""
}

func ValidateOuterK(outerK, numLocations int) string {
	if outerK < 0 {
		return "Outer K must be at least 0"
	}
	if outerK >= numLocations {
		return fmt.Sprintf("Outer K must be less than number of locations (%d)", numLocations)
	}
	if numLocations-outerK < 1 {
		return "Outer N (locations - outer K) must be at least 1"
	}
	return ""
}

func ValidateLocalNK(localNK, aag, minNodesPerLocation, minLocationDisks int) string {
	if localNK < 1 {
		return "Local N+K must be at least 1"
	}
	if localNK > minLocationDisks {
		return fmt.Sprintf("Local N+K cannot exceed minimum disks across locations (%d)", minLocationDisks)
	}
	return ""
}

func ValidateLocalK(localK, localNK int) string {
	if localK < 0 {
		return "Local K must be at least 0"
	}
	if localK >= localNK {
		return fmt.Sprintf("Local K must be less than Local N+K (%d)", localNK)
	}
	if localNK-localK < 1 {
		return "Local N (Local N+K - Local K) must be at least 1"
	}
	return ""
}

func ValidateRCName(name string, existingNames []string) string {
	if len(name) < 1 {
		return "RC name is required"
	}
	for _, existing := range existingNames {
		if name == existing {
			return fmt.Sprintf("RC name '%s' is already used", name)
		}
	}
	return ""
}
