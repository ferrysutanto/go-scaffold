package utils

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	vldtr     = validator.New()
	fieldsMap = make(map[string]string)
)

func init() {
	vldtr.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		fieldsMap[fld.Name] = name

		return name
	})
}

func StructCtx(ctx context.Context, s interface{}) error {
	if err := vldtr.StructCtx(ctx, s); err != nil {
		return errors.New(strings.Join(translateValidationErrors(err), "\n"))
	}

	return nil
}

func translateValidationErrors(err error) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = formatError(fe)
		}
		return out
	}

	return []string{err.Error()}
}

func formatError(fe validator.FieldError) string {
	jsonName := fe.Field()
	jsonParam := fe.Param()
	if name, ok := fieldsMap[fe.Param()]; ok {
		jsonParam = name
	}

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("Field '%s' is required", jsonName)
	case "email":
		return fmt.Sprintf("Field '%s' should be an email", jsonName)
	case "required_without":
		return fmt.Sprintf("Field '%s' is required if '%s' is not present", jsonName, jsonParam)
	case "required_with":
		return fmt.Sprintf("Field '%s' is required along with '%s'", jsonName, jsonParam)
	case "oneof":
		return fmt.Sprintf("Field '%s' should be one of '%s'", jsonName, jsonParam)
	}

	return fe.Error()
}
