package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *basicHandler) Healthcheck(c *gin.Context) {
	if err := h.s.Healthcheck(c); err != nil {
		log.Errorf("[build][api][handlers][basicHandler][Healthcheck] failed to execute service healthcheck: %v", err)
		// always return error in array to make it easier to parse for the client
		c.JSON(http.StatusInternalServerError, GenericResponse{Errors: []string{errInternal.Error()}})
		return
	}

	c.JSON(http.StatusOK, HealthcheckResponse{Status: "OK"})
}
