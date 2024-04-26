package services

import (
	"context"

	"github.com/ferrysutanto/go-errors"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/ferrysutanto/go-scaffold/integrations/tracers"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

type basicService struct {
	db db.IDB
}

type RDBMSConfig struct {
	// DB *sqlx.DB

	Host     string
	Port     string
	Username string
	Password string
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

	// 6. init integrations

	// 7. init services
	resp := &basicService{}

	return resp, nil
}
