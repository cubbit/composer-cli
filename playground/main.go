package main

import (
	"fmt"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/cmd"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func main() {
	// --- Mock config ---
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			Output:         configuration_models.OutputHuman,
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Endpoints:      configuration_models.EndpointsV2{},
		}, nil
	}

	// --- Mock APIs with simulated delays ---
	gatewayAPI := &api.MockGatewayAPI{}
	swarmAPI := &api.MockSwarmAPI{}
	processAPI := &api.MockProcessAPI{}
	redundancyClassAPI := &api.MockRedundancyClassAPI{}
	locationAPI := &api.MockLocationAPI{}
	tenantAPI := &api.MockTenantAPI{}
	domainAPI := &api.MockDomainAPI{}

	wireGatewayAPI(gatewayAPI)
	wireSwarmAPI(swarmAPI)
	wireProcessAPI(processAPI)
	wireRedundancyClassAPI(redundancyClassAPI)
	wireLocationAPI(locationAPI)
	wireTenantAPI(tenantAPI)
	wireDomainAPI(domainAPI)
	wireGatewayListAPI(gatewayAPI)

	// --- Real services using mocked APIs ---
	gatewayService := servicegateway.NewGatewayService(mockCfg, gatewayAPI, swarmAPI, redundancyClassAPI, processAPI, locationAPI)
	swarmService := service.NewSwarmService(mockCfg, swarmAPI, locationAPI, processAPI, service.NewRedundancyClassValidator())

	tenantService := servicetenant.NewTenantService(mockCfg, tenantAPI, domainAPI, gatewayAPI, processAPI)

	// Trivial stubs for services not under test (avoids needing MockAuthAPI etc.)
	rootCmd := cmd.NewRootCommand(
		mockCfg,
		&trivialAuthService{},
		&trivialOperatorService{},
		&trivialUserService{},
		&trivialLocationService{},
		&trivialConfigService{},
		swarmService,
		&trivialDomainService{},
		tenantService,
		gatewayService,
		"0.0.0-playground",
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// --- Trivial service stubs for commands not under test ---

type trivialAuthService struct{}

func (s *trivialAuthService) Activate(c *cobra.Command, a []string) error { return nil }
func (s *trivialAuthService) SignUp(c *cobra.Command, a []string) error   { return nil }
func (s *trivialAuthService) Login(c *cobra.Command, a []string) error    { return nil }
func (s *trivialAuthService) Logout(c *cobra.Command, a []string) error   { return nil }

type trivialLocationService struct{}

func (s *trivialLocationService) List(c *cobra.Command, a []string) error              { return nil }
func (s *trivialLocationService) ListAggregated(c *cobra.Command, a []string) error    { return nil }
func (s *trivialLocationService) CreateVirtual(c *cobra.Command, a []string) error     { return nil }
func (s *trivialLocationService) CreateVirtualNode(c *cobra.Command, a []string) error { return nil }

type trivialOperatorService struct{}

func (s *trivialOperatorService) Connect(c *cobra.Command, a []string) error { return nil }

type trivialUserService struct{}

func (s *trivialUserService) ImportUsers(c *cobra.Command, a []string) error { return nil }
func (s *trivialUserService) CreateUser(c *cobra.Command, a []string) error  { return nil }

type trivialDomainService struct{}

func (s *trivialDomainService) Create(c *cobra.Command, a []string) error   { return nil }
func (s *trivialDomainService) Describe(c *cobra.Command, a []string) error { return nil }
func (s *trivialDomainService) List(c *cobra.Command, a []string) error     { return nil }
func (s *trivialDomainService) Delete(c *cobra.Command, a []string) error   { return nil }
func (s *trivialDomainService) Verify(c *cobra.Command, a []string) error   { return nil }

type trivialConfigService struct{}

func (s *trivialConfigService) InitConfiguration(c *cobra.Command, a []string) error { return nil }
func (s *trivialConfigService) View(c *cobra.Command, a []string) error              { return nil }
func (s *trivialConfigService) Edit(c *cobra.Command, a []string) error              { return nil }
func (s *trivialConfigService) Profiles(c *cobra.Command, a []string) error          { return nil }
func (s *trivialConfigService) SwitchProfile(c *cobra.Command, a []string) error     { return nil }
func (s *trivialConfigService) Validate(c *cobra.Command, a []string) error          { return nil }

// Wire function variables — defined in gateway.go to keep simulation logic together
var (
	wireGatewayAPI         func(*api.MockGatewayAPI)
	wireSwarmAPI           func(*api.MockSwarmAPI)
	wireProcessAPI         func(*api.MockProcessAPI)
	wireRedundancyClassAPI func(*api.MockRedundancyClassAPI)
	wireLocationAPI        func(*api.MockLocationAPI)
	wireTenantAPI          func(*api.MockTenantAPI)
	wireDomainAPI          func(*api.MockDomainAPI)
	wireGatewayListAPI     func(*api.MockGatewayAPI)
)
