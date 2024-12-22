package services

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/ferrysutanto/go-errors"
	"github.com/ferrysutanto/go-scaffold/integrations/tracers"
	"github.com/ferrysutanto/go-scaffold/repositories/cache"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

type srvImpl struct {
	db    db.IGenericDB
	cache cache.ICache
}

func validateConfig(ctx context.Context, cfg *Config) error {
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	if cfg == nil {
		return errors.NewWithCode("config is required", 400)
	}

	return nil
}

func New(ctx context.Context, cfg *Config) (IService, error) {
	if err := validateConfig(ctx, cfg); err != nil {
		return nil, err
	}

	// 3. init tracer provider
	if cfg.Tracer != nil {
		if err := tracers.Init(ctx, &tracers.Config{
			Host:     cfg.Tracer.Host,
			Port:     cfg.Tracer.Port,
			IsSecure: cfg.Tracer.IsSecure,
			AppName:  cfg.Tracer.AppName,
		}); err != nil {
			log.Println(ctx, "[services][New] failed to init tracers", err)
		} else {
			var span trace.Span
			// 4. init tracer
			ctx, span = otel.Tracer("").Start(ctx, "[services][New]")
			defer span.End()
		}
	}

	resp := &srvImpl{}

	// 5. init db
	if cfg.DB != nil {
		if cfg.DB.DB != nil {
			resp.db = cfg.DB.DB
		} else {
			dbDriver := cfg.DB.DriverName

			switch dbDriver {
			case "postgres", "pg":
			default:
				break
			}
		}

		return resp, nil
	}

	if cfg.Cache != nil {
		resp.cache = cfg.CacheClient
	}

	// 6. init integrations

	return resp, nil
}
