package services

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sethvargo/go-envconfig"
	"go.opentelemetry.io/otel"
)

type environmentVariables struct {
	AppName     string        `env:"APP_NAME,required"`
	Environment string        `env:"APP_ENV,required"`
	DB          *dbConfig     `env:", prefix="`
	ReplicaDB   *dbConfig     `env:", prefix=REPLICA_"`
	Cache       *cacheConfig  `env:", prefix=CACHE_"`
	Tracer      *tracerConfig `env:", prefix=TRACING_"`
}

type dbConfig struct {
	DriverName string        `env:"DB_DRIVER" envDefault:"postgres" validate:"required"`
	Host       string        `env:"DB_HOST" envDefault:"localhost" validate:"required"`
	Port       string        `env:"DB_PORT" envDefault:"5432" validate:"required"`
	Username   string        `env:"DB_USERNAME" envDefault:"" validate:"required"`
	Password   string        `env:"DB_PASSWORD" envDefault:"" validate:"required"`
	Name       string        `env:"DB_NAME" envDefault:"" validate:"required"`
	Timeout    time.Duration `env:"DB_TIMEOUT" envDefault:"5s" validate:"required"`
	SslMode    string        `env:"DB_SSL_MODE" envDefault:"disable" validate:"required"`
}

type cacheConfig struct {
	Driver   string        `env:"DRIVER" envDefault:"redis" validate:"required"`
	Host     string        `env:"HOST" envDefault:"localhost" validate:"required"`
	Port     string        `env:"PORT" envDefault:"6379" validate:"required"`
	Username string        `env:"USERNAME" envDefault:"" validate:"required"`
	Password string        `env:"PASSWORD" envDefault:"" validate:"required"`
	DB       string        `env:"DB" envDefault:"0" validate:"required"`
	Timeout  time.Duration `env:"TIMEOUT" envDefault:"5s" validate:"required"`
}

type tracerConfig struct {
	AgentType   string `env:"AGENT_TYPE" envDefault:"jaeger" validate:"required"`
	IsEnabled   bool   `env:"ENABLED" envDefault:"true" validate:"required"`
	Host        string `env:"AGENT_HOST" envDefault:"localhost" validate:"required"`
	Port        string `env:"AGENT_PORT" envDefault:"4317" validate:"required"`
	IsSecure    bool   `env:"AGENT_IS_SECURE" envDefault:"false" validate:"required"`
	ServiceName string `env:"SERVICE_NAME" envDefault:"go-scaffold" validate:"required"`
}

func GenerateEnvironmentVariables(ctx context.Context) (*environmentVariables, error) {
	if ctx == nil {
		return nil, errors.New("[services][GenerateEnvironmentVariables] context is required")
	}

	ctx, span := otel.Tracer("").Start(ctx, "[services][GenerateEnvironmentVariables]")
	defer span.End()

	resp := environmentVariables{}
	if err := envconfig.Process(ctx, &resp); err != nil {
		err = errors.Wrap(err, "[services][GenerateEnvironmentVariables] failed to process environment variables")
		span.RecordError(err)
		return nil, err
	}

	return &resp, nil
}
