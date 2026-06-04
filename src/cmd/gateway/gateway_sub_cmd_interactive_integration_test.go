package cmd_gateway

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/cubbit/composer-cli/utils/interactive/interactive_tester"
)

func TestGatewaySubCmd_Create_Interactive_Success(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	var capturedRequest *api.CreateGatewayV5Request

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-001", Name: "prod-swarm"}},
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-002", Name: "staging-swarm"}},
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-003", Name: "dev-swarm"}},
				},
			}, nil
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, swarmID string,
		) ([]api.RedundancyClass, error) {
			switch swarmID {
			case "swarm-001":
				return []api.RedundancyClass{
					{ID: "rc-01-001", Name: "redundancy-alpha"},
					{ID: "rc-01-002", Name: "redundancy-beta"},
					{ID: "rc-01-003", Name: "redundancy-gamma"},
				}, nil
			case "swarm-002":
				return []api.RedundancyClass{
					{ID: "rc-02-001", Name: "redundancy-delta"},
					{ID: "rc-02-002", Name: "redundancy-epsilon"},
					{ID: "rc-02-003", Name: "redundancy-zeta"},
				}, nil
			case "swarm-003":
				return []api.RedundancyClass{
					{ID: "rc-03-001", Name: "redundancy-eta"},
					{ID: "rc-03-002", Name: "redundancy-theta"},
					{ID: "rc-03-003", Name: "redundancy-iota"},
				}, nil
			default:
				return nil, fmt.Errorf("unknown swarm: %s", swarmID)
			}
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{ClusterID: "cluster-eu-01", Name: "EU cluster"},
				{ClusterID: "cluster-us-02", Name: "US cluster"},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			_ configuration.URLs, _ string, _ string, request *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			capturedRequest = request
			return &api.CreateGatewayV5Response{ID: "test-process-id"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration.URLs, _ string, _ string, processID string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gateway-id-001"})
			return &api.Process{
				ID:     processID,
				Step:   api.ProcessStepCompleted,
				Status: api.ProcessStatusSuccess,
				Data:   data,
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select cluster",
		"EU cluster (cluster-eu-01)",
		"US cluster (cluster-us-02)",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectMultipleT(t, []string{"Ingress type", "manual"})
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select swarms",
		"prod-swarm (swarm-001)",
		"staging-swarm (swarm-002)",
		"dev-swarm (swarm-003)",
	})
	h.WriteDataT(t, h.Enter)
	h.ExpectT(t, "at least one option must be selected")
	h.WriteKeysT(t, h.Down, h.Space, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select redundancy class",
		"redundancy-delta (rc-02-001)",
		"redundancy-epsilon (rc-02-002)",
		"redundancy-zeta (rc-02-003)",
	})
	h.WriteKeysT(t, h.Down, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select redundancy class",
		"redundancy-eta (rc-03-001)",
		"redundancy-theta (rc-03-002)",
		"redundancy-iota (rc-03-003)",
	})
	h.WriteKeysT(t, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select the default swarm-RC pair",
		"Swarm: staging-swarm, RC: redundancy-zeta",
		"Swarm: dev-swarm, RC: redundancy-theta",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectT(t, "Gateway deployed — Gateway ID: gateway-id-001", interactive_tester.WithTimeout(30*time.Second))

	if capturedRequest == nil {
		t.Fatal("expected gateway creation to be called")
	}
	if capturedRequest.Name != "test-gateway" {
		t.Fatalf("expected name 'test-gateway', got %q", capturedRequest.Name)
	}
	if capturedRequest.Slug != "test-slug" {
		t.Fatalf("expected slug 'test-slug', got %q", capturedRequest.Slug)
	}
	if capturedRequest.ClusterID != "cluster-us-02" {
		t.Fatalf("expected cluster 'cluster-us-02', got %q", capturedRequest.ClusterID)
	}
	if len(capturedRequest.SwarmsAndRedundancyClass) != 2 {
		t.Fatalf("expected 2 swarm-RC pairs, got %d", len(capturedRequest.SwarmsAndRedundancyClass))
	}
	if capturedRequest.SwarmsAndRedundancyClass[0].SwarmID != "swarm-002" {
		t.Fatalf("expected first swarm 'swarm-002', got %q", capturedRequest.SwarmsAndRedundancyClass[0].SwarmID)
	}
	if capturedRequest.SwarmsAndRedundancyClass[0].RedundancyClassID != "rc-02-003" {
		t.Fatalf("expected first RC 'rc-02-003', got %q", capturedRequest.SwarmsAndRedundancyClass[0].RedundancyClassID)
	}
	if capturedRequest.SwarmsAndRedundancyClass[0].IsDefault {
		t.Fatal("expected first pair not to be default")
	}
	if capturedRequest.SwarmsAndRedundancyClass[1].SwarmID != "swarm-003" {
		t.Fatalf("expected second swarm 'swarm-003', got %q", capturedRequest.SwarmsAndRedundancyClass[1].SwarmID)
	}
	if capturedRequest.SwarmsAndRedundancyClass[1].RedundancyClassID != "rc-03-002" {
		t.Fatalf("expected second RC 'rc-03-002', got %q", capturedRequest.SwarmsAndRedundancyClass[1].RedundancyClassID)
	}
	if !capturedRequest.SwarmsAndRedundancyClass[1].IsDefault {
		t.Fatal("expected second pair to be default")
	}
	if capturedRequest.CubbitIngress.Type != api.IngressTypeManual {
		t.Fatalf("expected ingress type 'manual', got %q", capturedRequest.CubbitIngress.Type)
	}
}

func TestGatewaySubCmd_Create_Interactive_ClusterAPIError(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{Data: []api.ListSwarmV5ItemPresentation{}}, nil
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) ([]api.RedundancyClass, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return nil, fmt.Errorf("connection refused")
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "connection refused", interactive_tester.WithTimeout(15*time.Second))
}

func TestGatewaySubCmd_Create_Interactive_NoClusters(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{Data: []api.ListSwarmV5ItemPresentation{}}, nil
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) ([]api.RedundancyClass, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "no clusters available", interactive_tester.WithTimeout(15*time.Second))
}

func TestGatewaySubCmd_Create_Interactive_SwarmAPIError(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return nil, fmt.Errorf("internal server error")
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) ([]api.RedundancyClass, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{ClusterID: "cluster-eu-01", Name: "EU cluster"},
				{ClusterID: "cluster-us-02", Name: "US cluster"},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select cluster",
		"EU cluster (cluster-eu-01)",
		"US cluster (cluster-us-02)",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectMultipleT(t, []string{"Ingress type", "manual"})
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "internal server error", interactive_tester.WithTimeout(15*time.Second))
}

func TestGatewaySubCmd_Create_Interactive_NoSwarms(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{Data: []api.ListSwarmV5ItemPresentation{}}, nil
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, _ string,
		) ([]api.RedundancyClass, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{ClusterID: "cluster-eu-01", Name: "EU cluster"},
				{ClusterID: "cluster-us-02", Name: "US cluster"},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select cluster",
		"EU cluster (cluster-eu-01)",
		"US cluster (cluster-us-02)",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectMultipleT(t, []string{"Ingress type", "manual"})
	h.WriteDataT(t, h.Enter)

	h.ExpectT(t, "no swarms available", interactive_tester.WithTimeout(15*time.Second))
}

func TestGatewaySubCmd_Create_Interactive_NoRCForSwarm(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-001", Name: "no-rc-swarm"}},
				},
			}, nil
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, swarmID string,
		) ([]api.RedundancyClass, error) {
			if swarmID == "swarm-001" {
				return []api.RedundancyClass{}, nil
			}
			return nil, fmt.Errorf("unknown swarm: %s", swarmID)
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{ClusterID: "cluster-eu-01", Name: "EU cluster"},
				{ClusterID: "cluster-us-02", Name: "US cluster"},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select cluster",
		"EU cluster (cluster-eu-01)",
		"US cluster (cluster-us-02)",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectMultipleT(t, []string{"Ingress type", "manual"})
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{"Select swarms", "no-rc-swarm (swarm-001)"})
	h.WriteKeysT(t, h.Space, h.Enter)

	h.ExpectT(t, "no redundancy classes available", interactive_tester.WithTimeout(15*time.Second))
}

func TestGatewaySubCmd_Create_Interactive_GatewayAPIFailure(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-001", Name: "prod-swarm"}},
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-002", Name: "staging-swarm"}},
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-003", Name: "dev-swarm"}},
				},
			}, nil
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, swarmID string,
		) ([]api.RedundancyClass, error) {
			switch swarmID {
			case "swarm-002":
				return []api.RedundancyClass{
					{ID: "rc-02-001", Name: "redundancy-delta"},
					{ID: "rc-02-002", Name: "redundancy-epsilon"},
					{ID: "rc-02-003", Name: "redundancy-zeta"},
				}, nil
			default:
				return nil, fmt.Errorf("unknown swarm: %s", swarmID)
			}
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{ClusterID: "cluster-eu-01", Name: "EU cluster"},
				{ClusterID: "cluster-us-02", Name: "US cluster"},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			_ configuration.URLs, _ string, _ string, _ *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			return nil, fmt.Errorf("rate limit exceeded")
		},
	}

	mockProcessAPI := &api.MockProcessAPI{}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select cluster",
		"EU cluster (cluster-eu-01)",
		"US cluster (cluster-us-02)",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectMultipleT(t, []string{"Ingress type", "manual"})
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select swarms",
		"prod-swarm (swarm-001)",
		"staging-swarm (swarm-002)",
		"dev-swarm (swarm-003)",
	})
	h.WriteKeysT(t, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select redundancy class",
		"redundancy-delta (rc-02-001)",
		"redundancy-epsilon (rc-02-002)",
		"redundancy-zeta (rc-02-003)",
	})
	h.WriteKeysT(t, h.Down, h.Down, h.Space, h.Enter)

	h.ExpectT(t, "rate limit exceeded", interactive_tester.WithTimeout(15*time.Second))
}

func TestGatewaySubCmd_Create_Interactive_ExistingGatewayProcess(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	data, _ := json.Marshal(api.GatewayCreationProcessData{ID: "gateway-id-001"})

	mockSwarmAPI := &api.MockSwarmAPI{}
	mockRCApi := &api.MockRedundancyClassAPI{}
	mockLocationAPI := &api.MockLocationAPI{}
	mockGatewayAPI := &api.MockGatewayAPI{}
	mockProcessAPI := &api.MockProcessAPI{
		Processes: []api.Process{
			{
				Type:   api.ProcessTypeGatewayCreation,
				Status: api.ProcessStatusRunning,
				Data:   data,
			},
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectT(t, "a gateway creation is already in progress for gateway gateway-id-001", interactive_tester.WithTimeout(15*time.Second))
}

func TestGatewaySubCmd_Create_Interactive_DeploymentFailed(t *testing.T) {
	mockCfg := api.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(
			_ configuration.URLs, _ string, _ string, _ int, _ int,
		) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-001", Name: "prod-swarm"}},
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-002", Name: "staging-swarm"}},
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-003", Name: "dev-swarm"}},
				},
			}, nil
		},
	}

	mockRCApi := &api.MockRedundancyClassAPI{
		ListRedundancyClassesBySwarmFunc: func(
			_ configuration.URLs, _ string, _ string, swarmID string,
		) ([]api.RedundancyClass, error) {
			switch swarmID {
			case "swarm-002":
				return []api.RedundancyClass{
					{ID: "rc-02-001", Name: "redundancy-delta"},
					{ID: "rc-02-002", Name: "redundancy-epsilon"},
					{ID: "rc-02-003", Name: "redundancy-zeta"},
				}, nil
			case "swarm-003":
				return []api.RedundancyClass{
					{ID: "rc-03-001", Name: "redundancy-eta"},
					{ID: "rc-03-002", Name: "redundancy-theta"},
					{ID: "rc-03-003", Name: "redundancy-iota"},
				}, nil
			default:
				return nil, fmt.Errorf("unknown swarm: %s", swarmID)
			}
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(
			_ configuration.URLs, _ string, _ string, _ ...api.LocationListOption,
		) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{ClusterID: "cluster-eu-01", Name: "EU cluster"},
				{ClusterID: "cluster-us-02", Name: "US cluster"},
			}, nil
		},
	}

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayV5Func: func(
			_ configuration.URLs, _ string, _ string, _ *api.CreateGatewayV5Request,
		) (*api.CreateGatewayV5Response, error) {
			return &api.CreateGatewayV5Response{ID: "proc-fail"}, nil
		},
	}

	mockProcessAPI := &api.MockProcessAPI{
		GetProcessFunc: func(
			_ configuration.URLs, _ string, _ string, processID string,
		) (*api.Process, error) {
			data, _ := json.Marshal(api.GatewayCreationProcessData{
				ID: "gateway-fail-id",
				Error: &api.ProcessError{
					Code:    "DEPLOY_ERR",
					Message: "installation timeout",
				},
			})
			return &api.Process{
				ID:     processID,
				Step:   api.ProcessStepGatewayInstallation,
				Status: api.ProcessStatusFailed,
				Data:   data,
			}, nil
		},
	}

	gatewayService := servicegateway.NewGatewayService(
		mockCfg, mockGatewayAPI, mockSwarmAPI, mockRCApi, mockProcessAPI, mockLocationAPI,
	)
	gatewayCmd := NewGatewayCmd(gatewayService)
	h, err := interactive_tester.New(gatewayCmd, []string{"create", "--interactive"})
	if err != nil {
		t.Fatalf("could not create PTY harness: %v", err)
	}
	h.Start()
	defer h.Close()

	h.ExpectMultipleT(t, []string{"Gateway name", "required"}, interactive_tester.WithTimeout(30*time.Second))
	h.WriteLineT(t, "test-gateway")

	h.ExpectMultipleT(t, []string{"Gateway slug", "required"})
	h.WriteLineT(t, "test-slug")

	h.ExpectT(t, "Description")
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select cluster",
		"EU cluster (cluster-eu-01)",
		"US cluster (cluster-us-02)",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectMultipleT(t, []string{"Ingress type", "manual"})
	h.WriteDataT(t, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select swarms",
		"staging-swarm (swarm-002)",
		"dev-swarm (swarm-003)",
	})
	h.WriteKeysT(t, h.Down, h.Space, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select redundancy class",
		"redundancy-delta (rc-02-001)",
		"redundancy-epsilon (rc-02-002)",
		"redundancy-zeta (rc-02-003)",
	})
	h.WriteKeysT(t, h.Down, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select redundancy class",
		"redundancy-eta (rc-03-001)",
		"redundancy-theta (rc-03-002)",
		"redundancy-iota (rc-03-003)",
	})
	h.WriteKeysT(t, h.Down, h.Space, h.Enter)

	h.ExpectMultipleT(t, []string{
		"Select the default swarm-RC pair",
		"Swarm: staging-swarm, RC: redundancy-zeta",
		"Swarm: dev-swarm, RC: redundancy-theta",
	})
	h.WriteKeysT(t, h.Down, h.Enter)

	h.ExpectT(t, "gateway deployment failed — Gateway ID: gateway-fail-id, Error DEPLOY_ERR: installation timeout", interactive_tester.WithTimeout(15*time.Second))
}
