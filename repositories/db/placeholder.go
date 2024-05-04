package db

import (
	"context"

	"github.com/ferrysutanto/go-errors"
)

var errNotImplemented = errors.NewWithCode("function not implemented", 501)

// placeholderDB is a model that implements DB interface
type placeholderDB struct{}

func placeholder() IDB {
	return &placeholderDB{}
}

// Ping by placeholderModel is just displaying that the function is not implemented yet
func (*placeholderDB) Ping(ctx context.Context) error {
	return errNotImplemented
}
