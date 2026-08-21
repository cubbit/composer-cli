package apikey

import "github.com/spf13/cobra"

type APIKeyServiceMock struct {
	CreateAPIKeyFunc   func(cmd *cobra.Command, args []string) error
	ListAPIKeysFunc    func(cmd *cobra.Command, args []string) error
	DescribeAPIKeyFunc func(cmd *cobra.Command, args []string) error
	EditAPIKeyFunc     func(cmd *cobra.Command, args []string) error
	RevokeAPIKeyFunc   func(cmd *cobra.Command, args []string) error
}

func NewAPIKeyServiceMock() *APIKeyServiceMock {
	return &APIKeyServiceMock{
		CreateAPIKeyFunc: func(cmd *cobra.Command, args []string) error { return nil },
		ListAPIKeysFunc:  func(cmd *cobra.Command, args []string) error { return nil },
		DescribeAPIKeyFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		EditAPIKeyFunc:   func(cmd *cobra.Command, args []string) error { return nil },
		RevokeAPIKeyFunc: func(cmd *cobra.Command, args []string) error { return nil },
	}
}

func (m *APIKeyServiceMock) CreateAPIKey(cmd *cobra.Command, args []string) error {
	if m.CreateAPIKeyFunc != nil {
		return m.CreateAPIKeyFunc(cmd, args)
	}
	return nil
}

func (m *APIKeyServiceMock) ListAPIKeys(cmd *cobra.Command, args []string) error {
	if m.ListAPIKeysFunc != nil {
		return m.ListAPIKeysFunc(cmd, args)
	}
	return nil
}

func (m *APIKeyServiceMock) DescribeAPIKey(cmd *cobra.Command, args []string) error {
	if m.DescribeAPIKeyFunc != nil {
		return m.DescribeAPIKeyFunc(cmd, args)
	}
	return nil
}

func (m *APIKeyServiceMock) EditAPIKey(cmd *cobra.Command, args []string) error {
	if m.EditAPIKeyFunc != nil {
		return m.EditAPIKeyFunc(cmd, args)
	}
	return nil
}

func (m *APIKeyServiceMock) RevokeAPIKey(cmd *cobra.Command, args []string) error {
	if m.RevokeAPIKeyFunc != nil {
		return m.RevokeAPIKeyFunc(cmd, args)
	}
	return nil
}

var _ APIKeyServiceInterface = (*APIKeyServiceMock)(nil)
