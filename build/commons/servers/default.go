package servers

import (
	"context"
	"net/http"

	"github.com/ferrysutanto/go-scaffold/config"
)

func init() {
	ctx := context.Background()

	cfg := config.Get()

	srv, err := New(ctx, &Config{
		Host: cfg.AppHost,
		Port: cfg.AppPort,
	})
	if err != nil {
		panic(err)
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	g = srv
}

var g *http.Server

func Get() *http.Server {
	return g
}

func SetGet(srv *http.Server) {
	g = srv
}
