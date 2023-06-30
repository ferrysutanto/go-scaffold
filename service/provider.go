package service

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func validateConfig(ctx context.Context, conf Config) error {
	if ctx == nil {
		return fmt.Errorf("[service] failed to create service: context is nil")
	}

	validator := validator.New()
	log.Println(conf.DB.DbUsername)
	if err := validator.StructCtx(ctx, conf); err != nil {
		return errors.Wrap(err, "[service] failed to create service: invalid config")
	}

	return nil
}

func New(ctx context.Context, conf Config) (Service, error) {
	if err := validateConfig(ctx, conf); err != nil {
		return nil, err
	}

	switch conf.Type {
	case BASIC_SERVICE:
		return newBasicService(ctx, conf)
	default:
		return nil, errors.New("[service] failed to create service: unknown service type")
	}
}
