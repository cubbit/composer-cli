package input

import "errors"

// ErrCancelled is returned when the user cancels a prompt.
var ErrCancelled = errors.New("cancelled")

var ErrAtLeastOneRequired = errors.New("at least one option must be selected")
