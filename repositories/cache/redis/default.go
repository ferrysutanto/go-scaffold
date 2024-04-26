package redis

import (
	"context"
	"log"

	"github.com/ferrysutanto/go-scaffold/repositories/cache"
	"github.com/ferrysutanto/go-scaffold/utils"
)

func init() {
	ctx := context.Background()

	// Load .env file
	env, err := utils.GetEnv(ctx)
	if err != nil {
		log.Printf("no .env file found. skipping...")
		return
	}

	if env.Cache == nil || env.Cache.Redis == nil {
		log.Printf("no cache configuration found. skipping...")
		return
	}

	redisCache, err := New(ctx, &Config{
		Host:     env.Cache.Redis.Host,
		Port:     int(env.Cache.Redis.Port),
		Username: env.Cache.Redis.Username,
		Password: env.Cache.Redis.Password,
		DB:       int(env.Cache.Redis.DB),
	})
	if err != nil {
		log.Fatalf("failed to init redis cache: %v", err)
	}

	cache.SetGlobal(redisCache)
}
