package handlers

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/services"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var isDefaultInit bool
var def Handler = &emptyHandler{}

func Init(ctx context.Context, conf *Config) error {
	if err := services.Init(ctx, &services.Config{
		Type: services.BASIC_SERVICE,
		DB: &services.DbObject{
			DriverName:            conf.DbDriverName,
			DbHost:                conf.DbHost,
			DbPort:                conf.DbPort,
			DbName:                conf.DbName,
			DbUsername:            conf.DbUsername,
			DbPassword:            conf.DbPassword,
			DbSSLMode:             conf.DbSSLMode,
			ReplicationDbHost:     conf.ReplicationDbHost,
			ReplicationDbPort:     conf.ReplicationDbPort,
			ReplicationDbName:     conf.ReplicationDbName,
			ReplicationDbUsername: conf.ReplicationDbUsername,
			ReplicationDbPassword: conf.ReplicationDbPassword,
			ReplicationDbSSLMode:  conf.ReplicationDbSSLMode,
		},
	}); err != nil {
		return errors.Wrap(err, "[handlers][Init] failed to init services")
	}

	if isDefaultInit {
		return errors.New("[handlers][Init] default handler already initialized")
	}

	var err error
	def, err = New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "[handlers][Init] failed to init default handler")
	}

	isDefaultInit = true

	return nil
}

func Healthcheck(c *gin.Context) {
	def.Healthcheck(c)
}
