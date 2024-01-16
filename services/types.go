package services

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type Service interface {
	Healthcheck(ctx context.Context) error
}

type Config struct {
	Env       string    `json:"env" yaml:"env" env:"ENVIRONMENT" env-default:"development" validate:"required"`
	DB        *DbConfig `validate:"required"`
	ReplicaDB *DbConfig `validate:"required" env:", prefix=REPLICA_"`
	Cache     *CacheConfig
	Tracer    *TracerConfig
}

type DbConfig struct {
	// DriverName is the name of the database driver and it's mandatory
	DriverName string `json:"driver_name" yaml:"driver_name" env:"DB_DRIVER_NAME" env-default:"postgres" validate:"required"`

	// You can supply either an sql.DB object or config of the database connection, but not both (mutually exclusive)
	DB       *sql.DB
	Host     string `json:"db_host" yaml:"db_host" env:"DB_HOST" env-default:"localhost" validate:"hostname|ip,required_without=DB,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	Port     int    `json:"db_port" yaml:"db_port" env:"DB_PORT" env-default:"5432" validate:"numeric,required_without=DB,required_with=DbHost,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	Name     string `json:"db_name" yaml:"db_name" env:"DB_NAME" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	Username string `json:"db_username" yaml:"db_username" env:"DB_USERNAME" env-default:"postgres" validate:"required_without=DB"`
	Password string `json:"db_password" yaml:"db_password" env:"DB_PASSWORD" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbSSLMode"`
	SslMode  bool   `json:"db_ssl_mode" yaml:"db_ssl_mode" env:"DB_SSL_MODE" env-default:"disable" validate:""`
}

type CacheConfig struct {
	// You can supply either an redis.Client object or config of the redis connection, but not both (mutually exclusive)
	Client   *redis.Client
	Driver   string `json:"driver_name" yaml:"driver_name" env:"CACHE_DRIVER" env-default:"redis" validate:"required"`
	Host     string `json:"host" yaml:"host" env:"CACHE_HOST" env-default:"localhost" validate:"hostname|ip"`
	Port     int    `json:"port" yaml:"port" env:"CACHE_PORT" env-default:"6379" validate:"numeric"`
	Username string `json:"username" yaml:"username" env:"CACHE_USERNAME" env-default:""`
	Password string `json:"password" yaml:"password" env:"CACHE_PASSWORD" env-default:""`
	DB       int    `json:"db" yaml:"db" env:"CACHE_DB" env-default:"0" validate:"numeric"`
}

type TracerConfig struct {
	AgentType        string  `json:"tracer_agent_type" yaml:"tracer_agent_type" env:"TRACING_AGENT_TYPE" env-default:"jaeger" validate:"required"`
	IsEnabled        bool    `json:"tracer_is_enabled" yaml:"tracer_is_enabled" env:"TRACING_IS_ENABLED" env-default:"true"`
	Host             string  `json:"tracer_host" yaml:"tracer_host" env:"TRACING_AGENT_HOST" env-default:"localhost" validate:"hostname|ip,required_with=Port"`
	Port             int     `json:"tracer_port" yaml:"tracer_port" env:"TRACING_AGENT_PORT" env-default:"6831" validate:"numeric,required_with=Host"`
	AppName          string  `json:"tracer_app_name" yaml:"tracer_app_name" env:"APP_NAME" env-default:"go-scaffold"`
	IsSecure         bool    `json:"tracer_is_secure" yaml:"tracer_is_secure" env:"TRACING_AGENT_IS_SECURE" env-default:"false"`
	ApiKey           string  `json:"tracer_api_key" yaml:"tracer_api_key" env:"TRACING_API_KEY" env-default:""`
	TracerSampleRate float64 `json:"tracer_sample_rate" yaml:"tracer_sample_rate" env:"TRACING_SAMPLE_RATE" env-default:"1.0" validate:"numeric"`
}
