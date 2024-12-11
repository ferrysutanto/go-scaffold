package db

import (
	"context"

	"github.com/ferrysutanto/go-errors"
)

var errNotImplemented = errors.NewWithCode("function not implemented", 501)

// phDB is a model that implements DB interface
type phDB struct{}

func ph() IDB {
	return &phDB{}
}

// Ping by placeholderModel is just displaying that the function is not implemented yet
func (*phDB) Ping(ctx context.Context) error {
	return errNotImplemented
}

func (*phDB) BeginTx(ctx context.Context) (ITx, error) {
	return &phTx{
		phAccTx: &phAccTx{},
	}, nil
}

type phTx struct {
	*phAccTx
}

func (*phTx) Commit(ctx context.Context) error {
	return errNotImplemented
}

func (*phTx) Rollback(ctx context.Context) error {
	return errNotImplemented
}
