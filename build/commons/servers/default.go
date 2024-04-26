package servers

import (
	"context"
	"net/http"
)

func init() {
	ctx := context.Background()

	srv, err := New(ctx, &Config{})
	if err != nil {
		panic(err)
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	g = srv
}

var g *http.Server

func Global() *http.Server {
	return g
}

func SetGlobal(srv *http.Server) {
	g = srv
}
