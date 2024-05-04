package db

import (
	"context"
)

type IDB interface {
	Ping(context.Context) error
}

var g IDB = placeholder()

func Set(db IDB) {
	g = db
}

func Get() IDB {
	return g
}

// Ping pings default db
func Ping(ctx context.Context) error {
	return g.Ping(ctx)
}
