package models

import (
	"context"

	"github.com/pkg/errors"
)

var (
	def           Models = &emptyModel{}
	isDefaultInit bool
)

func InitDefaultModel(ctx context.Context, conf Config) error {
	var err error
	if isDefaultInit {
		return errors.New("[models] default model already initialized")
	}

	def, err = New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "[models] failed to init default model")
	}

	return nil
}
