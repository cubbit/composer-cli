package input

import tea "github.com/charmbracelet/bubbletea"

// Secret shows a password-style masked input and returns the entered value.
func Secret(title string, opts ...FieldOption) (string, error) {
	return SecretWithOpts(title, nil, opts...)
}

// SecretWithOpts is like Secret but accepts tea.ProgramOption values for testing.
func SecretWithOpts(title string, progOpts []tea.ProgramOption, opts ...FieldOption) (string, error) {
	f := Field{Name: "secret", Placeholder: title, Password: true}
	for _, o := range opts {
		o(&f)
	}
	res, err := MultiInputWithOpts(title, []Field{f}, progOpts...)
	if err != nil {
		return "", err
	}
	return res["secret"], nil
}
