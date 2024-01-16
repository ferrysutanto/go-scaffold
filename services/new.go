package services

import (
	"context"
	"fmt"

	"github.com/ferrysutanto/go-scaffold/utils"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

func validateConfig(ctx context.Context, conf *Config) error {
	if ctx == nil {
		return fmt.Errorf("[services][validateConfig] context is required")
	}

	ctx, span := otel.Tracer("").Start(ctx, "[services][validateConfig]")
	defer span.End()

	if err := utils.StructCtx(ctx, conf); err != nil {
		return errors.Wrap(err, "[services][validateConfig] invalid config.")
	}

	return nil
}

// New creates a new Service instance
func New(ctx context.Context, conf *Config) (Service, error) {
	// 1. Validate context
	if ctx == nil {
		return nil, errors.New("[services][New] context is required")
	}

	// 2. Validate config
	ctx, span := otel.Tracer("").Start(ctx, "[services][New]")
	defer span.End()

	// 3. Validate config
	if err := validateConfig(ctx, conf); err != nil {
		span.RecordError(err)

		return nil, errors.Wrap(err, "[services][New] failed on validating config")
	}

	// 4. Create new service
	return newBasicService(ctx, conf)
}
