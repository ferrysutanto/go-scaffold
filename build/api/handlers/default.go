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
