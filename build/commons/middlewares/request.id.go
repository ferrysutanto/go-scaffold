package middlewares

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func RequestIdMiddleware(c *gin.Context) {
	requestId := extractOrGenerateRequestID(c)

	c.Header("X-Request-ID", requestId)

	log.Printf("requestId: %s\n", requestId)

	c.Next()
}

func extractOrGenerateRequestID(c *gin.Context) string {
	requestID := c.GetHeader("X-Request-ID")
	if requestID == "" {
		requestID = generateRequestID()
	}

	return requestID
}

func generateRequestID() string {
	id := uuid.New()
	encoded := base64.RawURLEncoding.EncodeToString(id[:])
	return encoded
}
