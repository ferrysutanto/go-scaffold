package services

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
)

var (
	def           Service = &emptyService{}
	isDefaultInit bool
)

func Init(ctx context.Context) error {
	var err error
	if isDefaultInit {
		return errors.New("[services][Init] default service already initialized")
	}

	env, err := GenerateEnvironmentVariables(ctx)
	if err != nil {
		return errors.Wrap(err, "[services][Init] failed to generate environment variables")
	}

	conf := &Config{
		Env: env.Environment,
		DB: &DbConfig{
			DriverName: env.DB.DriverName,
			Host:       env.DB.Host,
			Port:       func() int { port, _ := strconv.Atoi(env.DB.Port); return port }(),
			Username:   env.DB.Username,
			Password:   env.DB.Password,
			Name:       env.DB.Name,
			SslMode:    func() bool { sslMode, _ := strconv.ParseBool(env.DB.SslMode); return sslMode }(),
		},
		ReplicaDB: &DbConfig{
			DriverName: env.ReplicaDB.DriverName,
			Host:       env.ReplicaDB.Host,
			Port:       func() int { port, _ := strconv.Atoi(env.ReplicaDB.Port); return port }(),
			Username:   env.ReplicaDB.Username,
			Password:   env.ReplicaDB.Password,
			Name:       env.ReplicaDB.Name,
			SslMode:    func() bool { sslMode, _ := strconv.ParseBool(env.ReplicaDB.SslMode); return sslMode }(),
		},
		Cache: &CacheConfig{
			Driver:   env.Cache.Driver,
			Host:     env.Cache.Host,
			Port:     func() int { port, _ := strconv.Atoi(env.Cache.Port); return port }(),
			Username: env.Cache.Username,
			Password: env.Cache.Password,
			DB:       func() int { db, _ := strconv.Atoi(env.Cache.DB); return db }(),
		},
		Tracer: &TracerConfig{
			AgentType: env.Tracer.AgentType,
			IsEnabled: env.Tracer.IsEnabled,
			Host:      env.Tracer.Host,
			Port:      func() int { port, _ := strconv.Atoi(env.Tracer.Port); return port }(),
			IsSecure:  env.Tracer.IsSecure,
			AppName:   env.Tracer.ServiceName,
		},
	}

	if err := validateBasicServiceConfig(ctx, conf); err != nil {
		return errors.Wrap(err, "[services][Init] failed to validate basic service config")
	}

	def, err = New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "[services][Init] failed to init default service")
	}

	isDefaultInit = true

	return nil
}

func Healthcheck(ctx context.Context) error {
	return def.Healthcheck(ctx)
}
