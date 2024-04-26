package services

import (
	"context"
)

var g IService

func Healthcheck(ctx context.Context) error {
	return g.Healthcheck(ctx)
}
