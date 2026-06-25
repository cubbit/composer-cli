package tenant

import "github.com/spf13/cobra"

type TenantServiceMock struct {
	CreateFunc func(cmd *cobra.Command, args []string) error
}

func NewTenantServiceMock() *TenantServiceMock {
	return &TenantServiceMock{
		CreateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *TenantServiceMock) Create(cmd *cobra.Command, args []string) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(cmd, args)
	}
	return nil
}
