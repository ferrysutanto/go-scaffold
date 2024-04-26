package servers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ferrysutanto/go-errors"
	"github.com/ferrysutanto/go-scaffold/build/commons/middlewares"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

const (
	errInternalMsg = "internal_server_error"
)

var (
	errInternal = errors.New(errInternalMsg)
)

type Handler interface {
	Healthcheck(c *gin.Context)
}

type HealthcheckResponse struct {
	Status string `json:"status" example:"OK"`
}

type GenericResponse struct {
	Status *string  `json:"status,omitempty" yaml:"status,omitempty" example:"OK"`
	Errors []string `json:"errors,omitempty" yaml:"errors,omitempty" example:"[\"error1\", \"error2\"]"`
}

type PageInfo struct {
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
	TotalRecord int  `json:"total_record"`
	TotalPage   *int `json:"total_page,omitempty"`
}

type Config struct {
	Host string
	Port int
}

func New(ctx context.Context, conf *Config) (*http.Server, error) {
	// 1. validate context
	if ctx == nil {
		return nil, errors.NewWithCode("context is required", 400)
	}

	// 2. start a span and defer its closure
	_, span := otel.Tracer("").Start(ctx, "[api/servers][New]")
	defer span.End()

	// 3. declare router
	r := gin.Default()
	v1 := r.Group("/v1")

	// 4. register middlewares
	v1.Use(middlewares.CORSMiddleware, middlewares.RequestIdMiddleware)

	// 5. register handlers
	v1.GET("/ping", Healthcheck)

	// 6. create the server
	resp := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Handler: r,
	}

	return resp, nil
}
