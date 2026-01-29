package service

import "github.com/spf13/cobra"

type AuthServiceMock struct {
	LoginFunc  func(cmd *cobra.Command, args []string) error
	LogoutFunc func(cmd *cobra.Command, args []string) error
}

func NewAuthServiceMock() *AuthServiceMock {
	return &AuthServiceMock{
		LoginFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		LogoutFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *AuthServiceMock) Login(cmd *cobra.Command, args []string) error {
	if m.LoginFunc != nil {
		return m.LoginFunc(cmd, args)
	}
	return nil
}

func (m *AuthServiceMock) Logout(cmd *cobra.Command, args []string) error {
	if m.LogoutFunc != nil {
		return m.LogoutFunc(cmd, args)
	}
	return nil
}
