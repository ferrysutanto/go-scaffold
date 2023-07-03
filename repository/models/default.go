package models

import (
	"context"

	"github.com/pkg/errors"
)

var (
	def           Model = &emptyModel{}
	isDefaultInit bool
)

func Init(ctx context.Context, conf Config) error {
	var err error
	if isDefaultInit {
		return errors.New("[models][Init] default model already initialized")
	}

	def, err = New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "[models][Init] failed to init default model")
	}

	return nil
}
