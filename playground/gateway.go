package main

import (
	"encoding/json"
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

var (
	gatewayProcessID string
	pollCount        int
)

func init() {
	wireGatewayAPI = func(m *api.MockGatewayAPI) {
		m.CreateGatewayV5Func = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			orgID string,
			req *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			delay()
			pollCount = 0
			gatewayProcessID = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
			return &api.CreateGatewayV5Response{
				ID: gatewayProcessID,
			}, nil
		}
	}

	wireSwarmAPI = func(m *api.MockSwarmAPI) {
		m.ListSwarmsV5Func = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			orgID string,
			page int,
			items int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			delay()

			swarms := []api.ListSwarmV5ItemPresentation{
				{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:                   "123e4567-e89b-12d3-a456-426614174001",
						Name:                 "eu-west-1",
						RedundancyClassCount: 2,
					},
				},
				{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:                   "123e4567-e89b-12d3-a456-426614174002",
						Name:                 "us-east-1",
						RedundancyClassCount: 2,
					},
				},
				{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:                   "123e4567-e89b-12d3-a456-426614174003",
						Name:                 "ap-southeast-1",
						RedundancyClassCount: 1,
					},
				},
			}

			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: swarms,
			}, nil
		}
	}

	wireRedundancyClassAPI = func(m *api.MockRedundancyClassAPI) {
		m.ListRedundancyClassesBySwarmFunc = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			swarmID string,
		) ([]api.RedundancyClass, error) {
			delay()

			switch swarmID {
			case "123e4567-e89b-12d3-a456-426614174001":
				return []api.RedundancyClass{
					{ID: "223e4567-e89b-12d3-a456-426614174001", Name: "standard"},
					{ID: "223e4567-e89b-12d3-a456-426614174002", Name: "premium"},
				}, nil
			case "123e4567-e89b-12d3-a456-426614174002":
				return []api.RedundancyClass{
					{ID: "223e4567-e89b-12d3-a456-426614174003", Name: "standard"},
					{ID: "223e4567-e89b-12d3-a456-426614174004", Name: "premium"},
				}, nil
			default:
				return []api.RedundancyClass{
					{ID: "223e4567-e89b-12d3-a456-426614174005", Name: "standard"},
				}, nil
			}
		}
	}

	wireProcessAPI = func(m *api.MockProcessAPI) {
		gatewaySteps := []api.ProcessStep{
			api.ProcessStepInitializing,
			api.ProcessStepGatewayProfileDeployment,
			api.ProcessStepGatewayInstallation,
			api.ProcessStepCompleted,
		}
		tenantSteps := []api.ProcessStep{
			api.ProcessStepInitializing,
			api.ProcessStepCreatingTenantGateway,
			api.ProcessStepWaitingForTenantGatewayToBeReady,
			api.ProcessStepCompleted,
		}

		m.GetProcessFunc = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			orgID string,
			processID string,
		) (*api.Process, error) {
			delay()

			p := &api.Process{
				ID:     processID,
				Step:   api.ProcessStepInitializing,
				Status: api.ProcessStatusRunning,
			}

			if processID == gatewayProcessID {
				idx := pollCount
				if idx >= len(gatewaySteps) {
					idx = len(gatewaySteps) - 1
				}
				p.Step = gatewaySteps[idx]
				if p.Step == api.ProcessStepCompleted {
					p.Status = api.ProcessStatusSuccess
				}
				pollCount++

				data := api.GatewayCreationProcessData{
					ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
				}
				if p.Status == api.ProcessStatusFailed {
					data.Error = &api.ProcessError{
						Code:    "DEPLOY_FAILED",
						Message: "gateway deployment encountered an error",
					}
				}

				p.Data, _ = json.Marshal(data)
			} else {
				idx := tenantPollCount
				if idx >= len(tenantSteps) {
					idx = len(tenantSteps) - 1
				}
				p.Step = tenantSteps[idx]
				if p.Step == api.ProcessStepCompleted {
					p.Status = api.ProcessStatusSuccess
				}
				tenantPollCount++

				data := api.TenantCreationProcessData{
					TenantID: "t-tenant-001",
				}
				if p.Status == api.ProcessStatusFailed {
					data.Error = &api.ProcessError{
						Code:    "TENANT_CREATION_FAILED",
						Message: "tenant creation encountered an error",
					}
				}

				p.Data, _ = json.Marshal(data)
			}

			return p, nil
		}
	}

	wireLocationAPI = func(m *api.MockLocationAPI) {
		m.ListFunc = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			orgID string,
			opts ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			delay()

			return []api.InfrastructureCluster{
				{ClusterID: "550e8400-e29b-41d4-a716-446655440001", Name: "eu-west-1a", Type: "aws"},
				{ClusterID: "550e8400-e29b-41d4-a716-446655440002", Name: "us-east-1", Type: "aws"},
			}, nil
		}
	}
}

func delay() {
	time.Sleep(1 * time.Second)
}
