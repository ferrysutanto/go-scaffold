package pg

import (
	"context"
	"log"
	"strings"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/ferrysutanto/go-scaffold/utils"
)

func init() {
	ctx := context.Background()

	env, err := utils.GetEnv(ctx)
	if err != nil {
		log.Println("failed to get env")
		return
	}

	dbDriver := strings.ToLower(env.DB.Driver)
	if dbDriver != "pg" && dbDriver != "postgres" {
		log.Printf("db driver is %s. skipping...\n", dbDriver)
		return
	}

	defaultDB, err := New(ctx, &Config{
		Host: &env.DB.PG.Host,
	})
	if err != nil {
		log.Println("failed to init default db. skipping...")
		return
	}

	db.SetGlobal(defaultDB)
}
