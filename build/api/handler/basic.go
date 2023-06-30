package handler

import (
	"context"
	"net/http"

	"github.com/ferrysutanto/go-scaffold/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type basicHandler struct {
	serviceProvider service.Service
}

func newBasicHandler(ctx context.Context, conf Config) (*basicHandler, error) {
	svcPvdr, err := service.New(ctx, service.Config{
		Type: service.BASIC_SERVICE,
		DB: &service.DbObject{
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
	})
	if err != nil {
		return nil, errors.Wrap(err, "[handler] failed to create service")
	}

	resp := &basicHandler{
		serviceProvider: svcPvdr,
	}

	return resp, nil
}

func (provider *basicHandler) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthcheckResponse{Status: "OK"})
}
