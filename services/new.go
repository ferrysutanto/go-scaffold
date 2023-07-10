package services

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func validateConfig(ctx context.Context, conf *Config) error {
	if ctx == nil {
		return fmt.Errorf("[services][validateConfig] context is required")
	}

	validator := validator.New()

	if err := validator.StructCtx(ctx, conf); err != nil {
		return errors.Wrap(err, "[services][validateConfig] invalid config")
	}

	return nil
}

func New(ctx context.Context, conf *Config) (Service, error) {
	if err := validateConfig(ctx, conf); err != nil {
		return nil, errors.Wrap(err, "[services][New] failed on validating config")
	}

	switch conf.Type {
	case BASIC_SERVICE:
		return newBasicService(ctx, conf)
	default:
		return nil, fmt.Errorf("[services][New] unknown service type: %s", conf.Type)
	}
}
