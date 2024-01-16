package cache

import (
	"context"
	"fmt"

	"github.com/ferrysutanto/go-scaffold/utils"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

/*
	validateConfig is a helper function to validate config

it returns error if config is invalid
it returns nil if config is valid
it also records error in span
*/
func validateConfig(ctx context.Context, conf *Config) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[repositories/cache][validateConfig] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/cache][validateConfig]")
	defer span.End()

	// 3. validate config
	if err := utils.StructCtx(ctx, conf); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/cache][validateConfig] failed on validating config")
		// record error in span
		span.RecordError(err)

		return err
	}

	return nil
}

/*
	New is a factory function that returns Cache interface

it returns Cache interface if success
it returns error if failed
*/
func New(ctx context.Context, conf *Config) (Cache, error) {
	// 1. check if context is provided
	if ctx == nil {
		return nil, errors.New("[repositories/cache][New] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/cache][New]")
	defer span.End()

	// 3. validate config
	if err := validateConfig(ctx, conf); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/cache][New] failed on validating config")
		// record error in span
		span.RecordError(err)

		return nil, err
	}

	// 4. check if db type is postgres
	switch conf.DriverName {
	// 4.a. if db type is postgres, initialize postgres model
	case "redis":
		return newRedisCache(ctx, conf)
	// 4.b. if db type is not postgres, return error
	default:
		// wrap error with additional info
		err := fmt.Errorf("[repositories/cache][New] unknown driver name: %s", conf.DriverName)
		// record error in span
		span.RecordError(err)

		return nil, err
	}
}
