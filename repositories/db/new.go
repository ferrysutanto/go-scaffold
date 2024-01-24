package db

import (
	"context"
	"fmt"

	"github.com/ferrysutanto/go-scaffold/utils"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func validateConfig(ctx context.Context, conf *Config) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[repositories/db][validateConfig] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][validateConfig]")
	defer span.End()

	// 3. validate config
	if err := utils.StructCtx(ctx, conf); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/db][validateConfig] invalid config")
		// record error in span
		span.RecordError(err)

		return err
	}

	return nil
}

func New(ctx context.Context, conf *Config) (DB, error) {
	// 1. check if context is provided
	if ctx == nil {
		return nil, errors.New("[repositories/db][New] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][New]")
	defer span.End()

	// 3. validate config
	if err := validateConfig(ctx, conf); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/db][New] failed on validating config")
		// record error in span
		span.RecordError(err)

		return nil, err
	}

	// 4. check if db type is postgres
	switch conf.DriverName {
	// 4.a. if db type is postgres, initialize postgres model
	case "postgres":
		return newPgModel(ctx, conf)
	// 4.b. if db type is not postgres, return error
	case "dynamodb":
		return newDdbModel(ctx, conf)
	default:
		err := fmt.Errorf("[repositories/db][New] unknown driver name: %s", conf.DriverName)
		// record error in span
		span.RecordError(err)

		return nil, err
	}
}
