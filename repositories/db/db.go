package db

import (
	"context"
)

type IDB interface {
	Ping(context.Context) error
}

var g IDB = &emptyModel{}

func SetGlobal(db IDB) {
	g = db
}

func GetGlobal() IDB {
	return g
}

// Ping pings default db
func Ping(ctx context.Context) error {
	return g.Ping(ctx)
}
