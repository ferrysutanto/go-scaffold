package cache

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func (c *redisCache) Get(ctx context.Context, key string) (string, error) {
	// 1. check if context is provided
	if ctx == nil {
		return "", errors.New("[cache][redisCache:Get] ctx is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[cache][redisCache:Get]")
	defer span.End()

	// 3. get value by key
	cmd := c.redisClient.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		// wrap error with additional info
		err = errors.Wrapf(err, "[cache][redisCache:Get] failed to get key '%s'", key)
		// record error in span
		span.RecordError(err)

		return "", err
	}

	return cmd.Val(), nil
}
