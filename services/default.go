package services

import (
	"context"
	"log"

	"github.com/ferrysutanto/go-scaffold/config"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

func init() {
	ctx := context.Background()

	cfg := config.Get()

	svc, err := New(ctx, &Config{
		Env: cfg.Env,
		DB: &DbConfig{
			DB: db.Get(),
		},
	})
	if err != nil {
		log.Fatalln("failed to init default service. terminating...")
	}

	Set(svc)
}

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
