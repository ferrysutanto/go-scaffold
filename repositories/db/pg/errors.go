package pg

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrIdRequired       = errors.New("id is required")
	ErrInvalidUUID      = errors.New("invalid UUID")
	ErrValidationFailed = func(errs []string) error {
		return fmt.Errorf("validation failed:\n- %v", strings.Join(errs, "\n- "))
	}
	ErrUsernameRequired     = errors.New("username is required")
	ErrEmailOrPhoneRequired = errors.New("either email or phone is required")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPhone         = errors.New("invalid phone")
)

var (
	ErrAccountNotFound = errors.New("account not found")
)
