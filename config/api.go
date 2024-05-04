package config

import (
	"context"

	"log"

	"os"

	"github.com/ferrysutanto/go-errors"
	"github.com/sethvargo/go-envconfig"
	"go.opentelemetry.io/otel"
)

var g *Config = &Config{}

func Get() *Config {
	return g
}

func init() {
	if err := LoadApiConfig(context.Background()); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}

type Config struct {
	AppName string         `env:"APP_NAME,required"`
	AppHost string         `env:"APP_HOST" envDefault:"localhost"`
	AppPort int            `env:"APP_PORT" envDefault:"8080"`
	Env     string         `env:"APP_ENV,required"`
	IsDebug bool           `env:"APP_DEBUG" envDefault:"false"`
	DB      *dbConfig      `env:", prefix=DB_"`
	Cache   *cacheConfig   `env:", prefix=CACHE_"`
	Tracer  *tracerConfig  `env:", prefix=TRACING_"`
	Cognito *cognitoConfig `env:", prefix=COGNITO_"`
}

func LoadApiConfig(ctx context.Context) error {
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	ctx, span := otel.Tracer("").Start(ctx, "[utils][GetEnv]")
	defer span.End()

	wd, err := os.Getwd()
	if err != nil {
		err = errors.ErrorfWithCode(errors.ErrUnexpected, "failed to get working directory: %v", err)
		span.RecordError(err)
		return err
	}

	if err := findAndLoadEnv(wd); err != nil {
		log.Printf("failed to load .env file: %v...\nprocessing with default environment variables", err)
	}

	if err := envconfig.Process(ctx, g); err != nil {
		err = errors.ErrorfWithCode(errors.ErrUnexpected, "failed to process environment variables: %v", err)
		span.RecordError(err)
		return err
	}

	return nil
}
