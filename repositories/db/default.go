package db

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

var (
	def           DB = &emptyModel{}
	isDefaultInit bool
)

// Init initializes default db
func Init(ctx context.Context, conf *Config) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[repositories/db][Init] context is nil")
	}

	// 2. start span and defer span end
	_, span := otel.Tracer("").Start(ctx, "[repositories/db][Init]")
	defer span.End()

	// 3. check if default db is already initialized
	var err error
	if isDefaultInit {
		// wrap error with additional info
		err = errors.New("[repositories/db][Init] default model already initialized")
		// record error in span
		span.RecordError(err)

		return err
	}

	// 4. if default db is not initialized, initialize it
	def, err = New(ctx, conf)
	if err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/db][Init] failed to init default model")
		// record error in span
		span.RecordError(err)

		return err
	}

	return nil
}

// Ping pings default db
func Ping(ctx context.Context) error {
	return def.Ping(ctx)
}
