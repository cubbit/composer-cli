package tenant

import "github.com/spf13/cobra"

type TenantServiceMock struct {
	CreateFunc   func(cmd *cobra.Command, args []string) error
	ListFunc     func(cmd *cobra.Command, args []string) error
	DescribeFunc func(cmd *cobra.Command, args []string) error
}

func NewTenantServiceMock() *TenantServiceMock {
	return &TenantServiceMock{
		CreateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DescribeFunc: func(cmd *cobra.Command, args []string) error {
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

func (m *TenantServiceMock) List(cmd *cobra.Command, args []string) error {
	if m.ListFunc != nil {
		return m.ListFunc(cmd, args)
	}
	return nil
}

func (m *TenantServiceMock) Describe(cmd *cobra.Command, args []string) error {
	if m.DescribeFunc != nil {
		return m.DescribeFunc(cmd, args)
	}
	return nil
}
