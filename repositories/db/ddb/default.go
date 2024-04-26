package ddb

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
	if dbDriver != "ddb" && dbDriver != "dynamodb" {
		log.Printf("db driver is %s. skipping...\n", dbDriver)
		return
	}

	defaultDB, err := New(ctx, &Config{
		AwsAccessKeyID:     env.DB.DDB.AccessKeyID,
		AwsSecretAccessKey: env.DB.DDB.SecretAccessKey,
		AwsRegion:          env.DB.DDB.Region,
		AwsEndpointURL:     env.DB.DDB.Endpoint,
	})
	if err != nil {
		log.Println("failed to init default db. skipping...")
		return
	}

	db.SetGlobal(defaultDB)
}
