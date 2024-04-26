package redis

import (
	"context"
	"fmt"

	"github.com/ferrysutanto/go-errors"
	"github.com/ferrysutanto/go-scaffold/repositories/cache"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
)

type redisCache struct {
	redisClient *redis.Client
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

func validateConfig(ctx context.Context, cfg *Config) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	// 2. check if config is provided
	if cfg == nil {
		return errors.NewWithCode("config is required", 400)
	}

	return nil
}

func New(ctx context.Context, cfg *Config) (cache.ICache, error) {
	// 1. check if context is provided
	if err := validateConfig(ctx, cfg); err != nil {
		return nil, errors.WrapWithCode(err, "bad config", 400)
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/cache/redis][New]")
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
		return nil, errors.WrapWithCode(err, "failed to ping", 500)
	}

	// 6. return redis cache
	return &redisCache{
		redisClient,
	}, nil
}

func (c *redisCache) Get(ctx context.Context, key string) (string, error) {
	// 1. check if context is provided
	if ctx == nil {
		return "", errors.NewWithCode("context is required", 400)
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[cache][redisCache:Get]")
	defer span.End()

	// 3. get value by key
	cmd := c.redisClient.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		// wrap error with additional info
		err = errors.WrapWithCode(err, "failed to get", 500)
		// record error in span
		span.RecordError(err)

		return "", err
	}

	return cmd.Val(), nil
}

func (c *redisCache) Set(ctx context.Context, key string, val string) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[cache][redisCache:Set]")
	defer span.End()

	// 3. set key
	cmd := c.redisClient.Set(ctx, key, val, 0)
	if err := cmd.Err(); err != nil {
		// wrap error with additional info
		err = errors.WrapWithCode(err, "failed to set", 500)
		// record error in span
		span.RecordError(err)

		return err
	}

	return nil
}
