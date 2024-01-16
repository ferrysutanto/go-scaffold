package servers

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Healthcheck(c *gin.Context)
}

type HealthcheckResponse struct {
	Status string `json:"status" example:"OK"`
}

type GenericResponse struct {
	Status *string  `json:"status,omitempty" yaml:"status,omitempty" example:"OK"`
	Errors []string `json:"errors,omitempty" yaml:"errors,omitempty" example:"[\"error1\", \"error2\"]"`
}

type PageInfo struct {
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
	TotalRecord int  `json:"total_record"`
	TotalPage   *int `json:"total_page,omitempty"`
}
