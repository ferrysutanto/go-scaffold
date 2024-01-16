package repositories

import (
	"context"
	"database/sql"

	"github.com/ferrysutanto/go-scaffold/repositories/cache"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type Repository interface {
	Cache() cache.Cache
	DB() db.DB
}

type repository struct {
	cache cache.Cache
	db    db.DB
}

func (r *repository) Cache() cache.Cache {
	return r.cache
}

func (r *repository) DB() db.DB {
	return r.db
}

type Config struct {
	DB        *DbConfig    `json:"db" yaml:"db"`
	ReplicaDB *DbConfig    `json:"replica_db" yaml:"replica_db"`
	Cache     *CacheConfig `json:"cache" yaml:"cache"`
}

type DbConfig struct {
	DriverName string  `json:"driver_name" yaml:"driver_name" env:"DB_DRIVER_NAME" env-default:"postgres"`
	DB         *sql.DB `validate:"required_without=DbHost"`
	Host       string  `json:"host" yaml:"host" env:"DB_HOST" env-default:"localhost" validate:"hostname|ip,required_without=DB,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	Port       int     `json:"port" yaml:"port" env:"DB_PORT" env-default:"5432" validate:"numeric,required_without=DB,required_with=DbHost,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	Name       string  `json:"name" yaml:"name" env:"DB_NAME" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	Username   string  `json:"username" yaml:"username" env:"DB_USERNAME" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbPassword,required_with=DbSSLMode"`
	Password   string  `json:"password" yaml:"password" env:"DB_PASSWORD" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbSSLMode"`
	SslMode    bool    `json:"ssl_mode" yaml:"ssl_mode" env:"DB_SSL_MODE" env-default:"disable" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbPassword"`
}

type CacheConfig struct {
	DriverName string `json:"driver_name" yaml:"driver_name"`
	Host       string `json:"host" yaml:"host"`
	Port       int    `json:"port" yaml:"port"`
	Username   string `json:"username" yaml:"username"`
	Password   string `json:"password" yaml:"password"`
	DB         int    `json:"db" yaml:"db"`
}

func newRepository(ctx context.Context, cfg *Config) (Repository, error) {
	if ctx == nil {
		return nil, errors.New("[repositories][newRepository] ctx is nil")
	}

	ctx, span := otel.Tracer("").Start(ctx, "[repositories][newRepository]")
	defer span.End()

	c, err := cache.New(ctx, &cache.Config{
		DriverName: cfg.Cache.DriverName,
		Host:       cfg.Cache.Host,
		Port:       cfg.Cache.Port,
		Username:   cfg.Cache.Username,
		Password:   cfg.Cache.Password,
		DB:         cfg.Cache.DB,
	})
	if err != nil {
		span.RecordError(err)
		return nil, errors.Wrap(err, "[repositories][newRepository] failed to initialize cache")
	}

	d, err := db.New(ctx, &db.Config{
		DriverName: cfg.DB.DriverName,
		DbHost:     cfg.DB.Host,
		DbPort:     cfg.DB.Port,
		DbName:     cfg.DB.Name,
		DbUsername: cfg.DB.Username,
		DbPassword: cfg.DB.Password,
		DbSslMode: func() string {
			if cfg.DB.SslMode {
				return "enable"
			}
			return "disable"
		}(),
		ReplicaDbHost:     cfg.ReplicaDB.Host,
		ReplicaDbPort:     cfg.ReplicaDB.Port,
		ReplicaDbName:     cfg.ReplicaDB.Name,
		ReplicaDbUsername: cfg.ReplicaDB.Username,
		ReplicaDbPassword: cfg.ReplicaDB.Password,
		ReplicaDbSslMode: func() string {
			if cfg.ReplicaDB.SslMode {
				return "enable"
			}
			return "disable"
		}(),
	})
	if err != nil {
		span.RecordError(err)
		return nil, errors.Wrap(err, "[repositories][newRepository] failed to initialize db")
	}

	return &repository{
		cache: c,
		db:    d,
	}, nil
}

func New(ctx context.Context, cfg *Config) (Repository, error) {
	return newRepository(ctx, cfg)
}
