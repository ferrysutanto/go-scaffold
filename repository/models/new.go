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
		return errors.Wrap(err, "[models][validateConfig] invalid config")
	}

	return nil
}

func New(ctx context.Context, conf Config) (Model, error) {
	if err := validateConfig(ctx, conf); err != nil {
		return nil, errors.Wrap(err, "[models][New] failed on validating config")
	}

	switch conf.DriverName {
	case "postgres":
		return newPgModel(ctx, conf)
	default:
		return nil, fmt.Errorf("[models][New] failed to create models: unknown driver name %s", conf.DriverName)
	}
}
