package handlers

import (
	"errors"

	"github.com/ferrysutanto/go-scaffold/utils"
	"github.com/gin-gonic/gin"
)

const (
	internalServerErrorMessage = "Internal server error. Please try again later."
	notFoundMessage            = "Resources not found"
	noParamMessage             = "No parameters supplied"
)

var (
	errInternal = errors.New(internalServerErrorMessage)
	errNotFound = errors.New(notFoundMessage)
	errNoParam  = errors.New(noParamMessage)
)

func injectAndValidateRequest(c *gin.Context, req interface{}) error {
	if c.Request.Method == "POST" || c.Request.Method == "PUT" {
		if err := c.BindJSON(req); err != nil {
			if err.Error() == "EOF" {
				err = errNoParam
			}

			return err
		}
	}

	if err := c.BindUri(req); err != nil {
		return err
	}

	if err := utils.StructCtx(c, req); err != nil {
		return err
	}

	return nil
}
