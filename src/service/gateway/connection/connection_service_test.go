package connection

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

const (
	testSwarmID       = "aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa"
	testOtherSwarmID  = "bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbbb"
	testRedundancyIDA = "cccccccc-cccc-4ccc-8ccc-cccccccccccc"
	testRedundancyIDB = "dddddddd-dddd-4ddd-8ddd-dddddddddddd"
	testRedundancyIDC = "eeeeeeee-eeee-4eee-8eee-eeeeeeeeeeee"
	testConnectionID  = "12345678-1234-4789-8abc-def012345678"
)

func setupConnectionTestCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.Flags().String("output", "human", "Output format")
	cmd.Flags().Bool("quiet", false, "Quiet mode")
	cmd.Flags().Bool("no-headers", false, "No headers")
	cmd.Flags().String("gateway-id", "", "Gateway ID")
	cmd.Flags().StringArray("connection", []string{}, "Connection")
	cmd.Flags().String("id", "", "Connection ID")
	return cmd
}

func connectionTestProfile() configuration_models.ProfileV2 {
	return configuration_models.ProfileV2{
		Output:         configuration_models.OutputHuman,
		APIKey:         "test-api-key",
		OrganizationID: "test-org-id",
	}
}

func TestList_SendsGatewayIDFromFlag(t *testing.T) {
	var requestedGatewayID string

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			requestedGatewayID = gatewayID
			return &api.GatewayConnectionsV5ListResponse{Data: []api.GatewayConnectionV5{}}, nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")

	err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gateway-123" {
		t.Fatalf("Expected gateway ID gateway-123, got %q", requestedGatewayID)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if output != "No connections found.\n" {
		t.Fatalf("Expected empty output, got %q", output)
	}
}

func TestList_TrimsGatewayID(t *testing.T) {
	var requestedGatewayID string

	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			requestedGatewayID = gatewayID
			return &api.GatewayConnectionsV5ListResponse{Data: []api.GatewayConnectionV5{}}, nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "  gateway-123  ")

	err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gateway-123" {
		t.Fatalf("Expected gateway ID gateway-123, got %q", requestedGatewayID)
	}
}

func TestList_RequiresGatewayID(t *testing.T) {
	cmd := setupConnectionTestCommand()

	err := List(Dependencies{GatewayAPI: &api.MockGatewayAPI{}}, cmd, nil, connectionTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--gateway-id must not be empty" {
		t.Fatalf("Expected empty gateway-id error, got %q", err.Error())
	}
}

func TestList_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		ListGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			return nil, errors.New("gateway not found")
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")

	err := List(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "failed to list connections: gateway not found" {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}

func TestCreate_SendsGroupedPairs(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			if gatewayID != "gateway-123" {
				t.Fatalf("Expected gateway ID gateway-123, got %q", gatewayID)
			}

			expected := []api.CreateGatewayConnectionV5RequestItem{
				{SwarmID: testSwarmID, RedundancyClassIDs: []string{testRedundancyIDA, testRedundancyIDB}},
				{SwarmID: testOtherSwarmID, RedundancyClassIDs: []string{testRedundancyIDC}},
			}
			if len(request.Connections) != len(expected) {
				t.Fatalf("Expected %d request items, got %+v", len(expected), request.Connections)
			}
			for i := range expected {
				if request.Connections[i].SwarmID != expected[i].SwarmID {
					t.Fatalf("Expected swarm %q at position %d, got %q", expected[i].SwarmID, i, request.Connections[i].SwarmID)
				}
				if strings.Join(request.Connections[i].RedundancyClassIDs, ",") != strings.Join(expected[i].RedundancyClassIDs, ",") {
					t.Fatalf("Expected redundancy classes %v at position %d, got %v", expected[i].RedundancyClassIDs, i, request.Connections[i].RedundancyClassIDs)
				}
			}

			return &api.GatewayConnectionsV5ListResponse{}, nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("connection", testSwarmID+":"+testRedundancyIDA)
	_ = cmd.Flags().Set("connection", testSwarmID+":"+testRedundancyIDB)
	_ = cmd.Flags().Set("connection", testOtherSwarmID+":"+testRedundancyIDC)

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCreate_CanonicalizesUUIDsBeforeSending(t *testing.T) {
	var request *api.CreateGatewayConnectionsV5Request

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			createRequest *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			request = createRequest
			return &api.GatewayConnectionsV5ListResponse{}, nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set(
		"connection",
		strings.ToUpper(testSwarmID)+":"+strings.ToUpper(testRedundancyIDA),
	)

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if request == nil || len(request.Connections) != 1 {
		t.Fatalf("Expected one request item, got %+v", request)
	}
	if request.Connections[0].SwarmID != testSwarmID {
		t.Fatalf("Expected canonical swarm ID %s, got %q", testSwarmID, request.Connections[0].SwarmID)
	}
	if len(request.Connections[0].RedundancyClassIDs) != 1 ||
		request.Connections[0].RedundancyClassIDs[0] != testRedundancyIDA {
		t.Fatalf("Expected canonical redundancy class ID %s, got %v", testRedundancyIDA, request.Connections[0].RedundancyClassIDs)
	}
}

func TestCreate_ValidatesConnectionsBeforeGatewayID(t *testing.T) {
	validPair := testSwarmID + ":" + testRedundancyIDA

	tests := []struct {
		name        string
		connections []string
		expected    string
	}{
		{
			name:        "missing flag",
			connections: nil,
			expected:    "at least one --connection flag is required",
		},
		{
			name:        "invalid format",
			connections: []string{"not-a-pair"},
			expected:    "invalid --connection format in connection 1: expected '<swarm-id>:<redundancy-class-id>[,<redundancy-class-id>...]', got \"not-a-pair\"",
		},
		{
			name:        "empty swarm",
			connections: []string{":" + testRedundancyIDA},
			expected:    "swarm ID cannot be empty in connection 1",
		},
		{
			name:        "invalid swarm uuid",
			connections: []string{"not-a-uuid:" + testRedundancyIDA},
			expected:    "swarm ID in connection 1 is not a valid UUID: \"not-a-uuid\"",
		},
		{
			name:        "empty redundancy class",
			connections: []string{testSwarmID + ":"},
			expected:    "connection 1, redundancy class 1: redundancy class ID cannot be empty",
		},
		{
			name:        "invalid redundancy class uuid",
			connections: []string{testSwarmID + ":not-a-uuid"},
			expected:    "connection 1, redundancy class 1: redundancy class ID is not a valid UUID: \"not-a-uuid\"",
		},
		{
			name:        "duplicate pair",
			connections: []string{validPair, validPair},
			expected:    "connection 2, redundancy class 1: duplicate connection for swarm " + testSwarmID + " and redundancy class " + testRedundancyIDA,
		},
		{
			name:        "duplicate pair with different case",
			connections: []string{validPair, strings.ToUpper(validPair)},
			expected:    "connection 2, redundancy class 1: duplicate connection for swarm " + testSwarmID + " and redundancy class " + testRedundancyIDA,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			createCalled := false

			mockGatewayAPI := &api.MockGatewayAPI{
				CreateGatewayConnectionsV5Func: func(
					endpoints configuration_models.EndpointsV2,
					apiKey string,
					organizationID string,
					gatewayID string,
					request *api.CreateGatewayConnectionsV5Request,
				) (*api.GatewayConnectionsV5ListResponse, error) {
					createCalled = true
					return &api.GatewayConnectionsV5ListResponse{}, nil
				},
			}

			cmd := setupConnectionTestCommand()
			for _, value := range test.connections {
				_ = cmd.Flags().Set("connection", value)
			}

			err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("Expected error %q, got %q", test.expected, err.Error())
			}
			if createCalled {
				t.Fatal("CreateGatewayConnectionsV5 should not be called when connection validation fails")
			}
		})
	}
}

func TestCreate_RequiresGatewayID(t *testing.T) {
	createCalled := false

	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			createCalled = true
			return &api.GatewayConnectionsV5ListResponse{}, nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("connection", testSwarmID+":"+testRedundancyIDA)

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--gateway-id must not be empty" {
		t.Fatalf("Expected empty gateway-id error, got %q", err.Error())
	}
	if createCalled {
		t.Fatal("CreateGatewayConnectionsV5 should not be called when the gateway ID is empty")
	}
}

func TestCreate_PropagatesAPIError(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		CreateGatewayConnectionsV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			request *api.CreateGatewayConnectionsV5Request,
		) (*api.GatewayConnectionsV5ListResponse, error) {
			return nil, errors.New("code status expected 201, but received 409 instead\nconnection already exists")
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("connection", testSwarmID+":"+testRedundancyIDA)

	err := Create(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "failed to create connections: code status expected 201, but received 409 instead\nconnection already exists"
	if err.Error() != expected {
		t.Fatalf("Expected wrapped API error, got %q", err.Error())
	}
}

func TestDelete_SendsGatewayIDAndConnectionID(t *testing.T) {
	var (
		requestedGatewayID  string
		deletedConnectionID string
	)

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			if apiKey != "test-api-key" || organizationID != "test-org-id" {
				t.Fatalf("Expected credentials to be propagated")
			}
			requestedGatewayID = gatewayID
			deletedConnectionID = connectionID
			return nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("id", testConnectionID)

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if requestedGatewayID != "gateway-123" {
		t.Fatalf("Expected gateway ID gateway-123, got %q", requestedGatewayID)
	}
	if deletedConnectionID != testConnectionID {
		t.Fatalf("Expected connection ID %s, got %q", testConnectionID, deletedConnectionID)
	}
}

func TestDelete_ValidatesConnectionIDBeforeGatewayID(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{
			name:     "empty id",
			id:       "   ",
			expected: "--id must not be empty",
		},
		{
			name:     "invalid uuid",
			id:       "not-a-uuid",
			expected: `--id must be a valid UUID, got "not-a-uuid"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			deleteCalled := false

			mockGatewayAPI := &api.MockGatewayAPI{
				DeleteGatewayConnectionV5Func: func(
					endpoints configuration_models.EndpointsV2,
					apiKey string,
					organizationID string,
					gatewayID string,
					connectionID string,
				) error {
					deleteCalled = true
					return nil
				},
			}

			cmd := setupConnectionTestCommand()
			_ = cmd.Flags().Set("id", test.id)

			err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("Expected error %q, got %q", test.expected, err.Error())
			}
			if deleteCalled {
				t.Fatal("DeleteGatewayConnectionV5 should not be called when the connection ID is invalid")
			}
		})
	}
}

func TestDelete_RequiresGatewayID(t *testing.T) {
	deleteCalled := false

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			deleteCalled = true
			return nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("id", testConnectionID)

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if err.Error() != "--gateway-id must not be empty" {
		t.Fatalf("Expected empty gateway-id error, got %q", err.Error())
	}
	if deleteCalled {
		t.Fatal("DeleteGatewayConnectionV5 should not be called when the gateway ID is empty")
	}
}

func TestDelete_CanonicalizesUUIDBeforeSending(t *testing.T) {
	var deletedConnectionID string

	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			deletedConnectionID = connectionID
			return nil
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("id", strings.ToUpper(testConnectionID))

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deletedConnectionID != testConnectionID {
		t.Fatalf("Expected canonical connection ID %s, got %q", testConnectionID, deletedConnectionID)
	}
}

func TestDelete_DefaultConflictIsForwardedUnchanged(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			return errors.New("code status expected 204, but received 409 instead\nthe default connection cannot be deleted")
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("id", testConnectionID)

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "failed to delete connection: code status expected 204, but received 409 instead\nthe default connection cannot be deleted"
	if err.Error() != expected {
		t.Fatalf("Expected the API error to be forwarded unchanged, got %q", err.Error())
	}
}

func TestDelete_NotFoundPropagates(t *testing.T) {
	mockGatewayAPI := &api.MockGatewayAPI{
		DeleteGatewayConnectionV5Func: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			gatewayID string,
			connectionID string,
		) error {
			return errors.New("code status expected 204, but received 404 instead\nconnection not found")
		},
	}

	cmd := setupConnectionTestCommand()
	_ = cmd.Flags().Set("gateway-id", "gateway-123")
	_ = cmd.Flags().Set("id", testConnectionID)

	err := Delete(Dependencies{GatewayAPI: mockGatewayAPI}, cmd, nil, connectionTestProfile(), nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "failed to delete connection: code status expected 204, but received 404 instead\nconnection not found"
	if err.Error() != expected {
		t.Fatalf("Expected wrapped not found error, got %q", err.Error())
	}
}

func TestBuildCreateConnectionsRequest(t *testing.T) {
	validPair := testSwarmID + ":" + testRedundancyIDA

	tests := []struct {
		name        string
		flags       []string
		expected    []api.CreateGatewayConnectionV5RequestItem
		expectedErr string
	}{
		{
			name:  "single pair",
			flags: []string{validPair},
			expected: []api.CreateGatewayConnectionV5RequestItem{
				{SwarmID: testSwarmID, RedundancyClassIDs: []string{testRedundancyIDA}},
			},
		},
		{
			name:  "comma separated redundancy classes",
			flags: []string{testSwarmID + ":" + testRedundancyIDA + "," + testRedundancyIDB},
			expected: []api.CreateGatewayConnectionV5RequestItem{
				{SwarmID: testSwarmID, RedundancyClassIDs: []string{testRedundancyIDA, testRedundancyIDB}},
			},
		},
		{
			name: "repeated same swarm merges into one item",
			flags: []string{
				testSwarmID + ":" + testRedundancyIDA,
				testSwarmID + ":" + testRedundancyIDB,
				testOtherSwarmID + ":" + testRedundancyIDC,
			},
			expected: []api.CreateGatewayConnectionV5RequestItem{
				{SwarmID: testSwarmID, RedundancyClassIDs: []string{testRedundancyIDA, testRedundancyIDB}},
				{SwarmID: testOtherSwarmID, RedundancyClassIDs: []string{testRedundancyIDC}},
			},
		},
		{
			name:  "uppercase uuids are canonicalized",
			flags: []string{strings.ToUpper(testSwarmID) + ":" + strings.ToUpper(testRedundancyIDA)},
			expected: []api.CreateGatewayConnectionV5RequestItem{
				{SwarmID: testSwarmID, RedundancyClassIDs: []string{testRedundancyIDA}},
			},
		},
		{
			name:        "missing flags",
			flags:       nil,
			expectedErr: "at least one --connection flag is required",
		},
		{
			name:        "malformed format",
			flags:       []string{"swarm-without-class"},
			expectedErr: "invalid --connection format in connection 1: expected '<swarm-id>:<redundancy-class-id>[,<redundancy-class-id>...]', got \"swarm-without-class\"",
		},
		{
			name:        "empty swarm",
			flags:       []string{" :" + testRedundancyIDA},
			expectedErr: "swarm ID cannot be empty in connection 1",
		},
		{
			name:        "invalid swarm uuid",
			flags:       []string{"swarm:" + testRedundancyIDA},
			expectedErr: "swarm ID in connection 1 is not a valid UUID: \"swarm\"",
		},
		{
			name:        "empty redundancy class",
			flags:       []string{testSwarmID + ": , "},
			expectedErr: "connection 1, redundancy class 1: redundancy class ID cannot be empty",
		},
		{
			name:        "second redundancy class invalid",
			flags:       []string{testSwarmID + ":" + testRedundancyIDA + ",class"},
			expectedErr: "connection 1, redundancy class 2: redundancy class ID is not a valid UUID: \"class\"",
		},
		{
			name:        "invalid redundancy class uuid",
			flags:       []string{testSwarmID + ":class"},
			expectedErr: "connection 1, redundancy class 1: redundancy class ID is not a valid UUID: \"class\"",
		},
		{
			name:        "duplicate pair within one flag",
			flags:       []string{testSwarmID + ":" + testRedundancyIDA + "," + testRedundancyIDA},
			expectedErr: "connection 1, redundancy class 2: duplicate connection for swarm " + testSwarmID + " and redundancy class " + testRedundancyIDA,
		},
		{
			name:        "duplicate pair across flags with different case",
			flags:       []string{validPair, strings.ToUpper(validPair)},
			expectedErr: "connection 2, redundancy class 1: duplicate connection for swarm " + testSwarmID + " and redundancy class " + testRedundancyIDA,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := buildCreateConnectionsRequest(test.flags)

			if test.expectedErr != "" {
				if err == nil {
					t.Fatalf("Expected error %q, got nil", test.expectedErr)
				}
				if err.Error() != test.expectedErr {
					t.Fatalf("Expected error %q, got %q", test.expectedErr, err.Error())
				}
				if request != nil {
					t.Fatalf("Expected no request on error, got %+v", request)
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if request == nil {
				t.Fatal("Expected a request, got nil")
			}

			expected := api.CreateGatewayConnectionsV5Request{Connections: test.expected}
			if !reflect.DeepEqual(*request, expected) {
				t.Fatalf("Expected request %+v, got %+v", expected, *request)
			}
		})
	}
}
