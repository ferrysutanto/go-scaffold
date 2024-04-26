package utils

import (
	"context"
	"fmt"

	"log"

	"os"
	"path/filepath"

	"github.com/ferrysutanto/go-errors"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"go.opentelemetry.io/otel"
)

type EnvironmentVariables struct {
	AppName     string         `env:"APP_NAME,required"`
	AppHost     string         `env:"APP_HOST" envDefault:"localhost"`
	AppPort     int            `env:"APP_PORT" envDefault:"8080"`
	Environment string         `env:"APP_ENV,required"`
	IsDebug     bool           `env:"APP_DEBUG" envDefault:"false"`
	DB          *dbConfig      `env:", prefix="`
	Cache       *cacheConfig   `env:", prefix=CACHE_"`
	Tracer      *tracerConfig  `env:", prefix=TRACING_"`
	Cognito     *cognitoConfig `env:", prefix=COGNITO_"`
}

type redisConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     uint   `env:"PORT" envDefault:"6379"`
	Username string `env:"USERNAME" envDefault:""`
	Password string `env:"PASSWORD" envDefault:""`
	DB       uint   `env:"DB" envDefault:"0"`
}

type cacheConfig struct {
	Redis *redisConfig `env:", prefix=REDIS_"`
}

type cognitoConfig struct {
	AwsRegion  string `env:"AWS_REGION" envDefault:"ap-southeast-3"`
	UserPoolID string `env:"USER_POOL_ID" envDefault:""`
	ClientID   string `env:"CLIENT_ID" envDefault:""`
}

type dbConfig struct {
	Driver string `env:"DB_DRIVER" envDefault:"dynamodb" validate:"required"`

	PG  *pgDBConfig     `env:", prefix=DB_PG_"`
	DDB *dynamoDBConfig `env:", prefix=DB_DDB_"`
}

type pgDBConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     uint   `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"postgres"`
	SslMode  string `env:"SSL_MODE" envDefault:"disable"`

	ReplicaHost     *string `env:"REPLICA_HOST"`
	ReplicaPort     *uint   `env:"REPLICA_PORT"`
	ReplicaUser     *string `env:"REPLICA_USER"`
	ReplicaPassword *string `env:"REPLICA_PASSWORD"`
	ReplicaDatabase *string `env:"REPLICA_DATABASE"`
	ReplicaSslMode  *string `env:"REPLICA_SSL_MODE"`
}

type dynamoDBConfig struct {
	AccessKeyID     *string `env:"AWS_ACCESS_KEY"`
	SecretAccessKey *string `env:"AWS_SECRET"`
	Endpoint        *string `env:"AWS_ENDPOINT" envDefault:"http://localhost:8000"`
	Region          *string `env:"AWS_REGION" envDefault:"ap-southeast-3"`
}

type tracerConfig struct {
	AgentType   string `env:"AGENT_TYPE" envDefault:"jaeger" validate:"required"`
	IsEnabled   bool   `env:"ENABLED" envDefault:"true" validate:"required"`
	Host        string `env:"AGENT_HOST" envDefault:"localhost" validate:"required"`
	Port        string `env:"AGENT_PORT" envDefault:"4317" validate:"required"`
	IsSecure    bool   `env:"AGENT_IS_SECURE" envDefault:"false" validate:"required"`
	ServiceName string `env:"SERVICE_NAME" envDefault:"go-scaffold" validate:"required"`
}

func GetEnv(ctx context.Context) (*EnvironmentVariables, error) {
	if ctx == nil {
		return nil, errors.NewWithCode("context is required", 400)
	}

	ctx, span := otel.Tracer("").Start(ctx, "[utils][GetEnv]")
	defer span.End()

	wd, err := os.Getwd()
	if err != nil {
		err = errors.ErrorfWithCode(errors.ErrUnexpected, "failed to get working directory: %v", err)
		span.RecordError(err)
		return nil, err
	}

	if err := findAndLoadEnv(wd); err != nil {
		log.Printf("failed to load .env file: %v...\nprocessing with default environment variables", err)
	}

	resp := EnvironmentVariables{}
	if err := envconfig.Process(ctx, &resp); err != nil {
		err = errors.ErrorfWithCode(errors.ErrUnexpected, "failed to process environment variables: %v", err)
		span.RecordError(err)
		return nil, err
	}

	return &resp, nil
}

func findAndLoadEnv(currentDir string) error {
	_, err := os.Stat(filepath.Join(currentDir, ".env"))
	if os.IsNotExist(err) {
		parentDir := filepath.Dir(currentDir)
		// Stop at some reasonable level or the filesystem root
		if parentDir == currentDir || parentDir == "/" {
			return fmt.Errorf("no .env file found")
		}
		return findAndLoadEnv(parentDir)
	}
	return godotenv.Load(filepath.Join(currentDir, ".env"))
}
