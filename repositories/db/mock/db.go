package mock

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

type DB struct{}

func (this *DB) Ping(context.Context) error {
	return nil
}
func (this *DB) BeginTx(ctx context.Context) (db.ITx, error) {
	return nil, nil
}
func (this *DB) Account(ctx context.Context) db.IAccountRepository {
	return nil
}
func (this *DB) Profile(ctx context.Context) db.IProfileRepository {
	return nil
}
