package service

import "github.com/spf13/cobra"

type UserServiceMock struct {
	ImportUsersFunc func(cmd *cobra.Command, args []string) error
	CreateUserFunc  func(cmd *cobra.Command, args []string) error
	ListUsersFunc   func(cmd *cobra.Command, args []string) error
}

func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{
		ImportUsersFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		CreateUserFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ListUsersFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *UserServiceMock) ImportUsers(cmd *cobra.Command, args []string) error {
	if m.ImportUsersFunc != nil {
		return m.ImportUsersFunc(cmd, args)
	}
	return nil
}

func (m *UserServiceMock) CreateUser(cmd *cobra.Command, args []string) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(cmd, args)
	}
	return nil
}

func (m *UserServiceMock) ListUsers(cmd *cobra.Command, args []string) error {
	if m.ListUsersFunc != nil {
		return m.ListUsersFunc(cmd, args)
	}
	return nil
}
