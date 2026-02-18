package service

import "github.com/spf13/cobra"

type LocationServiceMock struct {
	ListFunc func(cmd *cobra.Command, args []string) error
}

func NewLocationServiceMock() *LocationServiceMock {
	return &LocationServiceMock{
		ListFunc: func(cmd *cobra.Command, args []string) error {
			cmd.Println("Mock: Locations list: [location1, location2, location3]")
			return nil
		},
	}
}

func (m *LocationServiceMock) List(cmd *cobra.Command, args []string) error {
	if m.ListFunc != nil {
		return m.ListFunc(cmd, args)
	}
	return nil
}
