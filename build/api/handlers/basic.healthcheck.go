package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *basicHandler) Healthcheck(c *gin.Context) {
	if err := h.s.Healthcheck(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err500})
		return
	}

	c.JSON(http.StatusOK, HealthcheckResponse{Status: "OK"})
}
