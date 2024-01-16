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
	Env       string    `json:"env" yaml:"env" env-default:"development" validate:"required"`
	DB        *DbConfig `validate:"required"`
	ReplicaDB *DbConfig `validate:"required"`
	Cache     *CacheConfig
	Tracer    *TracerConfig
}

type DbConfig struct {
	// DriverName is the name of the database driver and it's mandatory
	DriverName string `json:"driver_name" yaml:"driver_name" validate:"required"`

	// You can supply either an sql.DB object or config of the database connection, but not both (mutually exclusive)
	DB       *sql.DB
	Host     string `json:"db_host" yaml:"db_host" validate:"hostname|ip"`
	Port     int    `json:"db_port" yaml:"db_port" validate:"numeric"`
	Name     string `json:"db_name" yaml:"db_name"`
	Username string `json:"db_username" yaml:"db_username"`
	Password string `json:"db_password" yaml:"db_password"`
	SslMode  bool   `json:"db_ssl_mode" yaml:"db_ssl_mode"`
}

type CacheConfig struct {
	// You can supply either an redis.Client object or config of the redis connection, but not both (mutually exclusive)
	Client   *redis.Client
	Driver   string `json:"driver_name" yaml:"driver_name" validate:"required"`
	Host     string `json:"host" yaml:"host" validate:"hostname|ip"`
	Port     int    `json:"port" yaml:"port" validate:"numeric"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db" validate:"numeric"`
}

type TracerConfig struct {
	AgentType        string  `json:"tracer_agent_type" yaml:"tracer_agent_type" validate:"required"`
	IsEnabled        bool    `json:"tracer_is_enabled" yaml:"tracer_is_enabled"`
	Host             string  `json:"tracer_host" yaml:"tracer_host" validate:"hostname|ip"`
	Port             int     `json:"tracer_port" yaml:"tracer_port" validate:"numeric"`
	AppName          string  `json:"tracer_app_name" yaml:"tracer_app_name"`
	IsSecure         bool    `json:"tracer_is_secure" yaml:"tracer_is_secure"`
	ApiKey           string  `json:"tracer_api_key" yaml:"tracer_api_key"`
	TracerSampleRate float64 `json:"tracer_sample_rate" yaml:"tracer_sample_rate" validate:"numeric"`
}
