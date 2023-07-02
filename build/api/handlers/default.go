package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var isDefaultInit bool
var def Handler = &emptyHandler{}

func Init(ctx context.Context, conf Config) error {
	if isDefaultInit {
		return errors.New("[handler] default handler already initialized")
	}

	var err error
	def, err = New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "[handler] failed to init default handler")
	}

	return nil
}

func Healthcheck(c *gin.Context) {
	def.Healthcheck(c)
}
