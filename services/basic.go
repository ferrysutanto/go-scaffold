package services

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/repository/models"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type basicService struct {
	m models.Model
	v *validator.Validate
}

func newBasicService(ctx context.Context, conf Config) (*basicService, error) {
	resp := &basicService{
		v: validator.New(),
	}

	if conf.DB != nil {
		mpConf := models.Config{
			DriverName: conf.DB.DriverName,
			DB:         conf.DB.DB,

			DbHost:     conf.DB.DbHost,
			DbPort:     conf.DB.DbPort,
			DbUsername: conf.DB.DbUsername,
			DbPassword: conf.DB.DbPassword,
			DbName:     conf.DB.DbName,
			DbSSLMode:  conf.DB.DbSSLMode,

			ReplicationDB:         conf.DB.ReplicationDB,
			ReplicationDbHost:     conf.DB.ReplicationDbHost,
			ReplicationDbPort:     conf.DB.ReplicationDbPort,
			ReplicationDbUsername: conf.DB.ReplicationDbUsername,
			ReplicationDbPassword: conf.DB.ReplicationDbPassword,
			ReplicationDbName:     conf.DB.ReplicationDbName,
			ReplicationDbSSLMode:  conf.DB.ReplicationDbSSLMode,
		}

		modelsProvider, err := models.New(ctx, mpConf)
		if err != nil {
			return nil, errors.Wrap(err, "[services][newBasicService] failed to create models pg provider")
		}
		resp.m = modelsProvider
	}

	return resp, nil
}
