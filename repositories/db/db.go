package db

import (
	"context"
)

type IGenericDB interface {
	Ping(context.Context) error

	BeginTx(ctx context.Context) (ITx, error)

	Account(ctx context.Context) IAccountRepository
	Profile(ctx context.Context) IProfileRepository
}

type ITx interface {
	IAccountTx
	IProfileTx
}

var g IGenericDB = Placeholder()

func Set(db IGenericDB) {
	g = db
}

func Get() IGenericDB {
	return g
}

// Ping pings default db
func Ping(ctx context.Context) error {
	return g.Ping(ctx)
}
