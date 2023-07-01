package handlers

import (
	"context"
	"errors"
)

func New(ctx context.Context, conf Config) (Handler, error) {
	if conf.Type == TypeBasic {
		return newBasicHandler(ctx, conf)
	}

	return nil, errors.New("[handler] failed to create handler: unknown handler type")
}
