package services

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type emptyService struct{}

// Healthcheck empty placeholder
func (*emptyService) Healthcheck(ctx context.Context) error {
	// 1. validate context
	if ctx == nil {
		return errors.New("[services][Healthcheck] context is nil")
	}

	// 2. init tracer span and defer span end
	_, span := otel.Tracer("").Start(ctx, "services:Healthcheck")
	defer span.End()

	// 3. declare error and record error
	err := errors.New("[services][Healthcheck] function is not implemented")
	span.RecordError(err)

	return err
}
