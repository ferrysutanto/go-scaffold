package services

import (
	"context"

	"github.com/ferrysutanto/go-errors"
	"go.opentelemetry.io/otel"
)

func (s *srvImpl) Healthcheck(ctx context.Context) error {
	// 1. validate context
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	// 2. init tracer span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[services][service:Healthcheck]")
	defer span.End()

	// 3. validate model
	if err := s.db.Ping(ctx); err != nil {
		err = errors.WrapWithCode(err, "failed to ping the service", 500)
		span.RecordError(err)
		return err
	}

	return nil
}
