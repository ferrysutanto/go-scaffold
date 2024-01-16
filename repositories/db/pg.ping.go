package db

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func (m *pgModel) Ping(ctx context.Context) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[repositories/db][pgModel:Ping] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][pgModel:Ping]")
	defer span.End()

	// 3. ping main db
	if err := m.db.PingContext(ctx); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/db][pgModel:Ping] failed to ping db")
		// record error in span
		span.RecordError(err)

		return err
	}

	// 4. ping replica db, same as above but for replica
	if err := m.replDb.PingContext(ctx); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/db][pgModel:Ping] failed to ping db replica")
		// record error in span
		span.RecordError(err)

		return err
	}

	return nil
}
