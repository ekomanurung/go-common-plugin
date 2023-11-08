package helper

import (
	"errors"
	"fmt"
	"github.com/ekomanurung/go-common-plugin"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Ok[T interface{}](data T) common_plugin.Response[T] {
	return common_plugin.Response[T]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   data,
	}
}

func Status(status int) common_plugin.Response[interface{}] {
	return common_plugin.Response[interface{}]{
		Code:   status,
		Status: http.StatusText(status),
	}
}

func BusinessException(ex common_plugin.Error) common_plugin.Response[interface{}] {
	return common_plugin.Response[interface{}]{
		Code:   ex.Status,
		Status: ex.Err.Error(),
	}
}

func BadRequest(err error) common_plugin.Response[interface{}] {
	return common_plugin.Response[interface{}]{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: toMapError(err),
	}
}

func NotFound() common_plugin.Response[interface{}] {
	return common_plugin.Response[interface{}]{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
	}
}

func InternalServerError() common_plugin.Response[interface{}] {
	return common_plugin.Response[interface{}]{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
	}
}

func toMapError(err error) map[string][]string {
	groupOfErrors := make(map[string][]string, 0)

	v := err.(validator.ValidationErrors)

	if errors.As(err, &v) {
		for _, fieldError := range v {
			groupOfErrors[strings.ToLower(fieldError.Field())] = []string{toValidationErrorMessage(fieldError)}
		}
	}
	return groupOfErrors
}

func toValidationErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "must not be null or empty"
	case "lte":
		return fmt.Sprintf("%v should be less than %v", fe.Value(), fe.Param())
	case "gte":
		return fmt.Sprintf("%v should be greater than %v", fe.Value(), fe.Param())
	case "max":
		return fmt.Sprintf("should be max at %v", fe.Param())
	default:
		return fe.Error()
	}
}

func Max() {
}
