package service

import "github.com/spf13/cobra"

type SwarmServiceMock struct {
	CreateFunc            func(cmd *cobra.Command, args []string) error
	CreateInteractiveFunc func(cmd *cobra.Command, args []string) error
	DescribeFunc          func(cmd *cobra.Command, args []string) error
	ListFunc              func(cmd *cobra.Command, args []string) error
}

func NewSwarmServiceMock() *SwarmServiceMock {
	return &SwarmServiceMock{
		CreateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		CreateInteractiveFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DescribeFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *SwarmServiceMock) Create(cmd *cobra.Command, args []string) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(cmd, args)
	}
	return nil
}

func (m *SwarmServiceMock) Describe(cmd *cobra.Command, args []string) error {
	if m.DescribeFunc != nil {
		return m.DescribeFunc(cmd, args)
	}
	return nil
}

func (m *SwarmServiceMock) List(cmd *cobra.Command, args []string) error {
	if m.ListFunc != nil {
		return m.ListFunc(cmd, args)
	}
	return nil
}

func (m *SwarmServiceMock) CreateInteractive(cmd *cobra.Command, args []string) error {
	if m.CreateInteractiveFunc != nil {
		return m.CreateInteractiveFunc(cmd, args)
	}
	return nil
}
