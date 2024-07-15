package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// 1. DB RELATED CONFIG

type dbConfig struct {
	Driver string `env:"DRIVER" envDefault:"postgres" validate:"required"`

	PG  *pgDBConfig     `env:", prefix="`
	DDB *dynamoDBConfig `env:", prefix="`
}

type pgDBConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     uint   `env:"PORT" envDefault:"5432"`
	Username string `env:"USERNAME" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Name     string `env:"NAME" envDefault:"postgres"`
	SslMode  string `env:"SSL_MODE" envDefault:"disable"`

	ReplicaHost     *string `env:"REPLICA_HOST"`
	ReplicaPort     *uint   `env:"REPLICA_PORT"`
	ReplicaUsername *string `env:"REPLICA_USERNAME"`
	ReplicaPassword *string `env:"REPLICA_PASSWORD"`
	ReplicaDatabase *string `env:"REPLICA_NAME"`
	ReplicaSslMode  *string `env:"REPLICA_SSL_MODE"`
}

type dynamoDBConfig struct {
	AccessKeyID     *string `env:"AWS_ACCESS_KEY"`
	SecretAccessKey *string `env:"AWS_SECRET"`
	Endpoint        *string `env:"AWS_ENDPOINT" envDefault:"http://localhost:8000"`
	Region          *string `env:"AWS_REGION" envDefault:"ap-southeast-3"`
}

// 2. CACHE RELATED CONFIG

type cacheConfig struct {
	Redis *redisConfig `env:", prefix=REDIS_"`
}

type redisConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     uint   `env:"PORT" envDefault:"6379"`
	Username string `env:"USERNAME" envDefault:""`
	Password string `env:"PASSWORD" envDefault:""`
	DB       uint   `env:"DB" envDefault:"0"`
}

// 3. TRACER RELATED CONFIG

type tracerConfig struct {
	AgentType   string `env:"AGENT_TYPE" envDefault:"jaeger" validate:"required"`
	IsEnabled   bool   `env:"ENABLED" envDefault:"true" validate:"required"`
	Host        string `env:"AGENT_HOST" envDefault:"localhost" validate:"required"`
	Port        string `env:"AGENT_PORT" envDefault:"4317" validate:"required"`
	IsSecure    bool   `env:"AGENT_IS_SECURE" envDefault:"false" validate:"required"`
	ServiceName string `env:"SERVICE_NAME" envDefault:"go-scaffold" validate:"required"`
}

// 4. OTHER INTEGRATION RELATED CONFIG

// 4.a. COGNITO RELATED CONFIG
type cognitoConfig struct {
	AwsRegion  string `env:"AWS_REGION" envDefault:"ap-southeast-3"`
	UserPoolID string `env:"USER_POOL_ID" envDefault:""`
	ClientID   string `env:"CLIENT_ID" envDefault:""`
}

// helper function to load environment variables from .env file and find it in the parent directory recursively
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
