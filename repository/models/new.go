package models

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func validateConfig(ctx context.Context, conf Config) error {
	v := validator.New()

	if err := v.StructCtx(ctx, conf); err != nil {
		return errors.Wrap(err, "[models] failed to validate config")
	}

	return nil
}

func New(ctx context.Context, conf Config) (Models, error) {
	if err := validateConfig(ctx, conf); err != nil {
		return nil, errors.Wrap(err, "[models] failed to create models provider")
	}

	switch conf.DriverName {
	case "postgres":
		return newPgModels(ctx, conf)
	default:
		return nil, fmt.Errorf("[models] failed to create models provider: unknown driver name %s", conf.DriverName)
	}
}
