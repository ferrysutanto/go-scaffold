package db

import (
	"context"

	"github.com/ferrysutanto/go-errors"
)

var errNotImplemented = errors.NewWithCode("function not implemented", 501)

// phGenericDB is a placeholder for IGenericDB
type phGenericDB struct{}

func Placeholder() IGenericDB {
	return &phGenericDB{}
}

// Ping by placeholderModel is just displaying that the function is not implemented yet
func (*phGenericDB) Ping(ctx context.Context) error {
	return errNotImplemented
}

// BeginTx by placeholderModel is just displaying that the function is not implemented yet
func (*phGenericDB) BeginTx(ctx context.Context) (ITx, error) {
	return nil, errNotImplemented
}

// Account by placeholderModel is just displaying that the function is not implemented yet
func (*phGenericDB) Account(ctx context.Context) IAccountRepository {
	return &phAccDB{}
}

// Profile by placeholderModel is just displaying that the function is not implemented yet
func (*phGenericDB) Profile(ctx context.Context) IProfileRepository {
	return &phProfDB{}
}
