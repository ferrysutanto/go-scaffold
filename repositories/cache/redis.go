package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type redisCache struct {
	redisClient *redis.Client
}

func validateRedisConfig(ctx context.Context, cfg *Config) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[repositories/cache][validateRedisConfig] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/cache][validateRedisConfig]")
	defer span.End()

	// 3. validate config

	return nil
}

func newRedisCache(ctx context.Context, cfg *Config) (Cache, error) {
	// 1. check if context is provided
	if ctx == nil {
		return nil, errors.New("[repositories/cache][newRedisCache] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/cache][newRedisCache]")
	defer span.End()

	// 3. declare redis client
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	username := cfg.Username
	password := cfg.Password
	db := cfg.DB

	// 4. initialize redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 5. ping redis
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, errors.Wrap(err, "[repositories/cache][newRedisCache] failed to ping redis")
	}

	// 6. return redis cache
	return &redisCache{
		redisClient,
	}, nil
}
