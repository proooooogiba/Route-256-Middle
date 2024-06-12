package errorapp

import "github.com/pkg/errors"

var (
	ErrNotFoundInPS  = errors.New("not found in product service")
	ErrNotFoundUser  = errors.New("not found user")
	ErrInvalidUserId = errors.New("invalid user id")
	ErrInvalidParams = errors.New("invalid params")
	ErrInvalidBody   = errors.New("invalid body")
)
