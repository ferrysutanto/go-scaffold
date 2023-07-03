package services

import (
	"context"

	"github.com/pkg/errors"
)

var (
	def           Service = &emptyService{}
	isDefaultInit bool
)

func Init(ctx context.Context, conf Config) error {
	var err error
	if isDefaultInit {
		return errors.New("[services][Init] default service already initialized")
	}

	def, err = New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "[services][Init] failed to init default service")
	}

	isDefaultInit = true

	return nil
}

func Hello() error {
	return def.Hello()
}
