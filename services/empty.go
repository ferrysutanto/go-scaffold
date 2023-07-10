package services

import (
	"context"

	"github.com/pkg/errors"
)

type emptyService struct{}

func (*emptyService) Healthcheck(ctx context.Context) error {
	return errors.New("[services][emptyService][Healthcheck] function is not implemented")
}
