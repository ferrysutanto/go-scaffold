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
	AppName     string `env:"APP_NAME,required"`
	AppHost     string `env:"APP_HOST"`
	AppPort     int    `env:"APP_PORT"`
	Environment string `env:"APP_ENV,required"`
	IsDebug     bool   `env:"APP_DEBUG"`

	DB      *dbConfig      `env:", prefix=DB_"`
	Cache   *cacheConfig   `env:", prefix=CACHE_"`
	Tracer  *tracerConfig  `env:", prefix=TRACING_"`
	Cognito *cognitoConfig `env:", prefix=COGNITO_"`
}

// 1. DB CONFIG SECTION

type dbConfig struct {
	Driver string `env:"DRIVER" validate:"required"`

	PG  *pgDBConfig
	DDB *dynamoDBConfig
}

type pgDBConfig struct {
	Host     string `env:"HOST"`
	Port     uint   `env:"PORT"`
	User     string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Database string `env:"NAME"`
	SslMode  string `env:"SSL_MODE"`

	ReplicaHost     *string `env:"REPLICA_HOST"`
	ReplicaPort     *uint   `env:"REPLICA_PORT"`
	ReplicaUser     *string `env:"REPLICA_USERNAME"`
	ReplicaPassword *string `env:"REPLICA_PASSWORD"`
	ReplicaDatabase *string `env:"REPLICA_NAME"`
	ReplicaSslMode  *string `env:"REPLICA_SSL_MODE"`
}

type dynamoDBConfig struct {
	AccessKeyID     *string `env:"AWS_ACCESS_KEY"`
	SecretAccessKey *string `env:"AWS_SECRET"`
	Endpoint        *string `env:"AWS_ENDPOINT"`
	Region          *string `env:"AWS_REGION"`
}

// 2. CACHE CONFIG SECTION

type cacheConfig struct {
	Driver string `env:"DRIVER"` // redis, memcached

	Redis *redisConfig
}

type redisConfig struct {
	Host     string `env:"HOST"`
	Port     uint   `env:"PORT"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	DB       uint   `env:"DB"`
}

// 3. TRACER CONFIG SECTION

type tracerConfig struct {
	AgentType   string `env:"AGENT_TYPE" validate:"required"`
	IsEnabled   bool   `env:"ENABLED" validate:"required"`
	Host        string `env:"AGENT_HOST" validate:"required"`
	Port        string `env:"AGENT_PORT" validate:"required"`
	IsSecure    bool   `env:"AGENT_IS_SECURE" validate:"required"`
	ServiceName string `env:"SERVICE_NAME" validate:"required"`
}

// 4. 3RD PARTY CONFIG SECTION

// 4.1 AWS Cognito
type cognitoConfig struct {
	AwsRegion  string `env:"AWS_REGION"`
	UserPoolID string `env:"USER_POOL_ID"`
	ClientID   string `env:"CLIENT_ID"`
}

// Signature: func GetEnv(ctx context.Context) (*EnvironmentVariables, error)

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

	// Load env...
	if err := findAndLoadEnv(wd); err != nil {
		log.Printf("failed to load .env file: %v...\nprocessing with default environment variables", err)
	}

	// Map env variables to struct
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
