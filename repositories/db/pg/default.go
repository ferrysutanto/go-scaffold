package pg

import (
	"context"
	"log"
	"strings"

	"github.com/ferrysutanto/go-scaffold/config"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

func init() {
	ctx := context.Background()

	cfg := config.Get()

	dbDriver := strings.ToLower(cfg.DB.Driver)
	if dbDriver != "pg" && dbDriver != "postgres" {
		log.Printf("db driver is %s. skipping...\n", dbDriver)
		return
	}

	defaultDB, err := New(ctx, &Config{
		Host: &cfg.DB.PG.Host,
	})
	if err != nil {
		log.Println("failed to init default db. skipping...")
		return
	}

	db.Set(defaultDB)
}
