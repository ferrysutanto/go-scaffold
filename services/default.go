package services

import (
	"context"
)

var g IService = placeholder()

// Set the global service
func Set(s IService) {
	g = s
}

// Get the global service
func Get() IService {
	return g
}

func Healthcheck(ctx context.Context) error {
	return g.Healthcheck(ctx)
}
