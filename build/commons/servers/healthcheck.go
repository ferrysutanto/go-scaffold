package servers

import (
	"net/http"

	"github.com/ferrysutanto/go-errors"
	log "github.com/sirupsen/logrus"

	"github.com/ferrysutanto/go-scaffold/services"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

func Healthcheck(c *gin.Context) {
	// 1. fetch the context from gin
	ctx := c.Request.Context()

	// 2. create a span and defer its closure
	ctx, span := otel.Tracer("").Start(ctx, "[servers][Healthcheck]")
	defer span.End()

	// 3. execute the service
	if err := services.Healthcheck(ctx); err != nil {
		// 3.a. wrap the error with additional context
		err = errors.WrapWithCode(err, "failed to execute healthcheck", 500)
		// 3.b. add the error to the span
		span.RecordError(err)
		// 3.c. log the error
		log.Errorln(errors.RootCause(err))
		// 3.d. return the error to the client
		c.Status(http.StatusInternalServerError)
		return
	}

	// 4. return the success response to the client
	c.JSON(http.StatusOK, HealthcheckResponse{Status: "OK"})
}
