package ddb

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
	if dbDriver != "ddb" && dbDriver != "dynamodb" {
		log.Printf("db driver is %s. skipping...\n", dbDriver)
		return
	}

	defaultDB, err := New(ctx, &Config{
		AwsAccessKeyID:     cfg.DB.DDB.AccessKeyID,
		AwsSecretAccessKey: cfg.DB.DDB.SecretAccessKey,
		AwsRegion:          cfg.DB.DDB.Region,
		AwsEndpointURL:     cfg.DB.DDB.Endpoint,
	})
	if err != nil {
		log.Println("failed to init default db. skipping...")
		return
	}

	db.Set(defaultDB)
}
