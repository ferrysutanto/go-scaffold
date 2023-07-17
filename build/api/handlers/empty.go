package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type emptyHandler struct{}

func (p *emptyHandler) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, &GenericResponse{Errors: []string{"not implemented"}})
}
