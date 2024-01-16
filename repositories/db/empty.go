package db

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

// emptyModel is a model that implements DB interface
type emptyModel struct{}

// Ping by emptyModel is just displaying that the function is not implemented yet
func (*emptyModel) Ping(ctx context.Context) error {
	if ctx == nil {
		return errors.New("[repositories/db][emptyModel:Ping] context is nil")
	}

	_, span := otel.Tracer("").Start(ctx, "[repositories/db][emptyModel:Ping]")
	defer span.End()

	err := errors.New("[repositories/db][emptyModel:Ping] function is not implemented")
	span.RecordError(err)

	return err
}
