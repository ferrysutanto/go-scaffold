package servers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ferrysutanto/go-scaffold/build/api/middlewares"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

type Config struct {
	Host string
	Port int
}

func New(ctx context.Context, conf *Config) (*http.Server, error) {
	// 1. validate context
	if ctx == nil {
		return nil, fmt.Errorf("[api/servers][New] context is required")
	}

	// 2. start a span and defer its closure
	_, span := otel.Tracer("").Start(ctx, "[api/servers][New]")
	defer span.End()

	// 3. declare router
	r := gin.Default()

	// 4. register middlewares
	r.Use(middlewares.RequestIdMiddleware)

	// 5. register handlers
	r.GET("/ping", Healthcheck)

	// 6. create the server
	resp := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Handler: r,
	}

	return resp, nil
}
