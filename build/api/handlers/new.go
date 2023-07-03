package handlers

import (
	"context"
	"errors"
)

func New(ctx context.Context, conf Config) (Handler, error) {
	if conf.Type == TypeBasic {
		return newBasicHandler(ctx, conf)
	}

	return nil, errors.New("[handlers][New] unknown handler type")
}
