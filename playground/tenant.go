package main

import (
	"time"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

var (
	tenantProcessID string
	tenantPollCount int
)

func init() {
	wireTenantAPI = func(m *api.MockTenantAPI) {
		m.CreateTenantV5Func = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			orgID string,
			req *api.CreateTenantV5Request,
		) (*api.GenericIDResponseModel, error) {
			delay()
			tenantPollCount = 0
			tenantProcessID = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
			return &api.GenericIDResponseModel{
				ID: tenantProcessID,
			}, nil
		}
	}

	wireDomainAPI = func(m *api.MockDomainAPI) {
		m.ListFunc = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			orgID string,
			page int,
			items int,
		) (*api.GenericPaginatedResponse[api.DomainDTO], error) {
			delay()
			now := time.Now()
			domains := []api.DomainDTO{
				{
					ID:             "d1e2f3a4-b5c6-7890-abcd-ef1234567890",
					DomainName:     "example.com",
					CreatedAt:      now,
					VerifiedAt:     &now,
					OrganizationID: orgID,
				},
				{
					ID:             "d2e3f4a5-b6c7-8901-bcde-f12345678901",
					DomainName:     "test.org",
					CreatedAt:      now,
					VerifiedAt:     &now,
					OrganizationID: orgID,
				},
				{
					ID:             "d3e4f5a6-b7c8-9012-cdef-123456789012",
					DomainName:     "unverified.dev",
					CreatedAt:      now,
					VerifiedAt:     nil,
					OrganizationID: orgID,
				},
			}
			return &api.GenericPaginatedResponse[api.DomainDTO]{
				Data: domains,
			}, nil
		}
	}

	wireGatewayListAPI = func(m *api.MockGatewayAPI) {
		m.ListGatewaysV5Func = func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			orgID string,
			opts ...api.ListGatewaysV5Option,
		) (*api.GenericPaginatedResponse[api.GatewayV5ListItemResponse], error) {
			delay()
			gateways := []api.GatewayV5ListItemResponse{
				{ID: "g1a2b3c4-d5e6-7890-abcd-ef1234567890", Name: "gw-eu-west-1", Slug: "gw-eu-west-1"},
				{ID: "g2a3b4c5-d6e7-8901-bcde-f12345678901", Name: "gw-us-east-1", Slug: "gw-us-east-1"},
			}
			return &api.GenericPaginatedResponse[api.GatewayV5ListItemResponse]{
				Data: gateways,
			}, nil
		}
	}
}
