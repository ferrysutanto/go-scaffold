package services

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/repositories/cache"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type IService interface {
	Healthcheck(ctx context.Context) error
}

type Config struct {
	Env string `json:"env" yaml:"env" env-default:"development" validate:"required"`

	DbClient db.IDB
	DB       *DbConfig

	CacheClient cache.ICache
	Cache       *CacheConfig
	Tracer      *TracerConfig
}

type RDBMSConfig struct {
	DbClient *sqlx.DB

	Host     *string
	Port     *uint
	Username *string
	Password *string
	Name     *string
	SslMode  *string

	ReplicaDbClient *sqlx.DB

	ReplicaHost     *string
	ReplicaPort     *uint
	ReplicaUsername *string
	ReplicaPassword *string
	ReplicaName     *string
	ReplicaSslMode  *string
}

type DbConfig struct {
	// DriverName is the name of the database driver and it's mandatory
	DriverName string `json:"driver_name" yaml:"driver_name" validate:"required"`

	// You can supply either an sql.DB object or config of the database connection, but not both (mutually exclusive)
	DB db.IDB

	RDBMSConfig *RDBMSConfig
}

type RedisConfig struct {
	RedisClient *redis.Client

	Host     *string `json:"host" yaml:"host" validate:"hostname|ip"`
	Port     *uint   `json:"port" yaml:"port" validate:"numeric"`
	Username *string `json:"username" yaml:"username"`
	Password *string `json:"password" yaml:"password"`
	DB       *uint   `json:"db" yaml:"db" validate:"numeric"`
}

type CacheConfig struct {
	Driver string `json:"driver_name" yaml:"driver_name" validate:"required"`

	*RedisConfig
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
