package service

import "github.com/spf13/cobra"

type ConfigServiceMock struct {
	ViewFunc          func(cmd *cobra.Command, args []string) error
	EditFunc          func(cmd *cobra.Command, args []string) error
	ProfilesFunc      func(cmd *cobra.Command, args []string) error
	SwitchProfileFunc func(cmd *cobra.Command, args []string) error
	ValidateFunc      func(cmd *cobra.Command, args []string) error
}

func NewConfigServiceMock() *ConfigServiceMock {
	return &ConfigServiceMock{
		ViewFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		EditFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ProfilesFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		SwitchProfileFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		ValidateFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func (m *ConfigServiceMock) View(cmd *cobra.Command, args []string) error {
	if m.ViewFunc != nil {
		return m.ViewFunc(cmd, args)
	}
	return nil
}

func (m *ConfigServiceMock) Edit(cmd *cobra.Command, args []string) error {
	if m.EditFunc != nil {
		return m.EditFunc(cmd, args)
	}
	return nil
}

func (m *ConfigServiceMock) Profiles(cmd *cobra.Command, args []string) error {
	if m.ProfilesFunc != nil {
		return m.ProfilesFunc(cmd, args)
	}
	return nil
}

func (m *ConfigServiceMock) SwitchProfile(cmd *cobra.Command, args []string) error {
	if m.SwitchProfileFunc != nil {
		return m.SwitchProfileFunc(cmd, args)
	}
	return nil
}

func (m *ConfigServiceMock) Validate(cmd *cobra.Command, args []string) error {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(cmd, args)
	}
	return nil
}
