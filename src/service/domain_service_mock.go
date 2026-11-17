package service

import "github.com/spf13/cobra"

type DomainServiceMock struct {
	CreateFunc   func(cmd *cobra.Command, args []string) error
	DescribeFunc func(cmd *cobra.Command, args []string) error
	ListFunc     func(cmd *cobra.Command, args []string) error
	DeleteFunc   func(cmd *cobra.Command, args []string) error
	VerifyFunc   func(cmd *cobra.Command, args []string) error
}

func NewDomainServiceMock() *DomainServiceMock {
	return &DomainServiceMock{
		CreateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DescribeFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DeleteFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		VerifyFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *DomainServiceMock) Create(cmd *cobra.Command, args []string) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(cmd, args)
	}
	return nil
}

func (m *DomainServiceMock) Describe(cmd *cobra.Command, args []string) error {
	if m.DescribeFunc != nil {
		return m.DescribeFunc(cmd, args)
	}
	return nil
}

func (m *DomainServiceMock) List(cmd *cobra.Command, args []string) error {
	if m.ListFunc != nil {
		return m.ListFunc(cmd, args)
	}
	return nil
}

func (m *DomainServiceMock) Delete(cmd *cobra.Command, args []string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(cmd, args)
	}
	return nil
}

func (m *DomainServiceMock) Verify(cmd *cobra.Command, args []string) error {
	if m.VerifyFunc != nil {
		return m.VerifyFunc(cmd, args)
	}
	return nil
}
