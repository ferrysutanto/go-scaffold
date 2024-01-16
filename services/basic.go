package services

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/ferrysutanto/go-scaffold/integrations/tracers"
	"github.com/ferrysutanto/go-scaffold/repositories"
)

type basicService struct {
	r repositories.Repository
}

func validateBasicServiceConfig(ctx context.Context, cfg *Config) error {
	return nil
}

func newBasicService(ctx context.Context, conf *Config) (*basicService, error) {
	if ctx == nil {
		return nil, errors.New("[services][newBasicService] context is nil")
	}

	// 3. init tracer provider
	if conf.Tracer != nil {
		if err := tracers.Init(ctx, &tracers.Config{
			Host:     conf.Tracer.Host,
			Port:     conf.Tracer.Port,
			IsSecure: conf.Tracer.IsSecure,
			AppName:  conf.Tracer.AppName,
		}); err != nil {
			log.Println(ctx, "[services][newBasicService] failed to init tracers", err)
		}
	}

	// 4. init tracer
	ctx, span := otel.Tracer("").Start(ctx, "[services][newBasicService]")
	defer span.End()

	// 5. init repositories
	repo, err := repositories.New(ctx, &repositories.Config{
		DB: &repositories.DbConfig{
			DriverName: conf.DB.DriverName,
			Host:       conf.DB.Host,
			Port:       conf.DB.Port,
			Username:   conf.DB.Username,
			Password:   conf.DB.Password,
			Name:       conf.DB.Name,
			SslMode:    conf.DB.SslMode,
		},
		ReplicaDB: &repositories.DbConfig{
			DriverName: conf.ReplicaDB.DriverName,
			Host:       conf.ReplicaDB.Host,
			Port:       conf.ReplicaDB.Port,
			Username:   conf.ReplicaDB.Username,
			Password:   conf.ReplicaDB.Password,
			Name:       conf.ReplicaDB.Name,
			SslMode:    conf.ReplicaDB.SslMode,
		},
		Cache: &repositories.CacheConfig{
			DriverName: conf.Cache.Driver,
			Host:       conf.Cache.Host,
			Port:       conf.Cache.Port,
			Username:   conf.Cache.Username,
			Password:   conf.Cache.Password,
			DB:         conf.Cache.DB,
		},
	})
	if err != nil {
		err = errors.Wrap(err, "[services][newBasicService] failed to init repositories")
		span.RecordError(err)
		return nil, err
	}

	// 6. init integrations

	// 7. init services
	resp := &basicService{
		r: repo,
	}

	return resp, nil
}
