package service

import "github.com/spf13/cobra"

type OperatorServiceMock struct {
	ConnectFunc func(cmd *cobra.Command, args []string) error
}

func NewOperatorServiceMock() *OperatorServiceMock {
	return &OperatorServiceMock{
		ConnectFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *OperatorServiceMock) Connect(cmd *cobra.Command, args []string) error {
	if m.ConnectFunc != nil {
		return m.ConnectFunc(cmd, args)
	}
	return nil
}
