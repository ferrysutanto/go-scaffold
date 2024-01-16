package cache

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func (c *redisCache) Set(ctx context.Context, key string, val string) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[cache][redisCache:Set] ctx is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[cache][redisCache:Set]")
	defer span.End()

	// 3. set key
	cmd := c.redisClient.Set(ctx, key, val, 0)
	if err := cmd.Err(); err != nil {
		// wrap error with additional info
		err = errors.Wrapf(err, "[cache][redisCache:Set] failed to set key '%s'", key)
		// record error in span
		span.RecordError(err)

		return err
	}

	return nil
}
