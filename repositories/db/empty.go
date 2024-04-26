package db

import (
	"context"

	"github.com/ferrysutanto/go-errors"
	"go.opentelemetry.io/otel"
)

// emptyModel is a model that implements DB interface
type emptyModel struct{}

func newEmpty() IDB {
	return &emptyModel{}
}

// Ping by emptyModel is just displaying that the function is not implemented yet
func (*emptyModel) Ping(ctx context.Context) error {
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	_, span := otel.Tracer("").Start(ctx, "[repositories/db][emptyModel:Ping]")
	defer span.End()

	err := errors.NewWithCode("function is not implemented", 501)
	span.RecordError(err)

	return err
}
