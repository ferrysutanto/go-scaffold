package services

import (
	"context"

	"github.com/ferrysutanto/go-errors"
)

var errNotImplemented = errors.NewWithCode("function not implemented", 501)

type svcPlaceholder struct{}

func placeholder() IService {
	return &svcPlaceholder{}
}

// Healthcheck empty placeholder
func (*svcPlaceholder) Healthcheck(ctx context.Context) error {
	return errNotImplemented
}
