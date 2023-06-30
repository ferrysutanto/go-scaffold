package service

import (
	"context"

	"github.com/pkg/errors"
)

var (
	def           Service = &emptyService{}
	isDefaultInit bool
)

func InitDefaultService(ctx context.Context, conf Config) error {
	var err error
	if isDefaultInit {
		return errors.New("[service] default service already initialized")
	}

	def, err = New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "[service] failed to init default service")
	}

	return nil
}

func Hello() error {
	return def.Hello()
}
