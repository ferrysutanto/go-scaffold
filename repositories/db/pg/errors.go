package pg

import "errors"

var (
	ErrIdRequired  = errors.New("id is required")
	ErrInvalidUUID = errors.New("invalid UUID")
)

var (
	ErrAccountNotFound = errors.New("account not found")
)
