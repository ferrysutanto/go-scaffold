package services

import (
	"context"

	"github.com/ferrysutanto/go-errors"
	"go.opentelemetry.io/otel"
)

func (s *basicService) Healthcheck(ctx context.Context) error {
	// 1. validate context
	if ctx == nil {
		return errors.New("[services][basicService:Healthcheck] context is nil")
	}

	// 2. init tracer span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[services][basicService:Healthcheck]")
	defer span.End()

	// 3. validate model
	if err := s.db.Ping(ctx); err != nil {
		err = errors.Wrap(err, "[services][basicService:Healthcheck] failed to ping model")
		span.RecordError(err)
		return err
	}

	return nil
}
