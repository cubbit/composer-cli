package gateway

import "github.com/spf13/cobra"

type GatewayServiceMock struct {
	CreateFunc func(cmd *cobra.Command, args []string) error
}

func NewGatewayServiceMock() *GatewayServiceMock {
	return &GatewayServiceMock{
		CreateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *GatewayServiceMock) Create(cmd *cobra.Command, args []string) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(cmd, args)
	}
	return nil
}
