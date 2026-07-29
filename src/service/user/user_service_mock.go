package user

import "github.com/spf13/cobra"

type UserServiceMock struct {
	ImportUsersFunc  func(cmd *cobra.Command, args []string) error
	CreateUserFunc   func(cmd *cobra.Command, args []string) error
	ListUsersFunc    func(cmd *cobra.Command, args []string) error
	DescribeUserFunc func(cmd *cobra.Command, args []string) error
	DeleteUserFunc   func(cmd *cobra.Command, args []string) error
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
		DescribeUserFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		DeleteUserFunc: func(cmd *cobra.Command, args []string) error {
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

func (m *UserServiceMock) DescribeUser(cmd *cobra.Command, args []string) error {
	if m.DescribeUserFunc != nil {
		return m.DescribeUserFunc(cmd, args)
	}
	return nil
}

func (m *UserServiceMock) DeleteUser(cmd *cobra.Command, args []string) error {
	if m.DeleteUserFunc != nil {
		return m.DeleteUserFunc(cmd, args)
	}
	return nil
}
