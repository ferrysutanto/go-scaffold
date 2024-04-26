package tracers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ferrysutanto/go-errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

type Config struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	IsSecure bool   `json:"is_secure" yaml:"is_secure"`
	AppName  string `json:"app_name" yaml:"app_name"`
}

func Init(ctx context.Context, cfg *Config) error {
	// 1. declare tracing agent configuration
	tracingAgentIsSecure := cfg.IsSecure
	tracingAgentHost := cfg.Host
	tracingAgentPort := cfg.Port
	tracingAppName := cfg.AppName

	// 2. validate tracing agent configuration
	if tracingAgentHost == "" {
		tracingAgentHost = "localhost"
	}

	if tracingAgentPort == 0 {
		tracingAgentPort = 6831
	}

	tracingAgentEndpoint := fmt.Sprintf("%s:%d", tracingAgentHost, tracingAgentPort)

	// 3. declare exporter options
	exporterOptions := make([]otlptracegrpc.Option, 0)
	if !tracingAgentIsSecure {
		exporterOptions = append(exporterOptions, otlptracegrpc.WithInsecure())
	}
	exporterOptions = append(exporterOptions, otlptracegrpc.WithTimeout(2*time.Second))
	exporterOptions = append(exporterOptions, otlptracegrpc.WithEndpoint(tracingAgentEndpoint))

	// 4. create a new collector exporter
	exporter, err := otlptracegrpc.New(
		ctx,
		exporterOptions...,
	)
	if err != nil {
		return errors.WrapWithCode(err, "failed to create tracing agent exporter", 500)
	}

	log.Printf("[tracers][Init] tracing agent is running on %s, secure conn: %t", tracingAgentEndpoint, tracingAgentIsSecure)

	// 5. create a new tracer provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(tracingAppName),
		)),
	)

	// 6. set the global tracer provider
	otel.SetTracerProvider(tp)

	return nil
}
