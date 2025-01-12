package mock

import "errors"

var (
	ErrExpectationNotFound = errors.New("expectation not found")
	ErrFulfilled           = errors.New("all expectation fulfilled")
)
