package errorapp

import "github.com/pkg/errors"

var (
	ErrNotFoundInPS = errors.New("not found in product service")
	ErrNotFoundUser = errors.New("not found user")
)
