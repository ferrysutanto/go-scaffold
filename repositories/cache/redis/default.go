package redis

import (
	"context"
	"log"

	"github.com/ferrysutanto/go-scaffold/config"
	"github.com/ferrysutanto/go-scaffold/repositories/cache"
)

func init() {
	ctx := context.Background()

	// Load .cfg file
	cfg := config.Get()

	if cfg.Cache == nil || cfg.Cache.Redis == nil {
		log.Printf("no cache configuration found. skipping...")
		return
	}

	redisCache, err := New(ctx, &Config{
		Host:     cfg.Cache.Redis.Host,
		Port:     int(cfg.Cache.Redis.Port),
		Username: cfg.Cache.Redis.Username,
		Password: cfg.Cache.Redis.Password,
		DB:       int(cfg.Cache.Redis.DB),
	})
	if err != nil {
		log.Fatalf("failed to init redis cache: %v", err)
	}

	cache.SetGlobal(redisCache)
}
