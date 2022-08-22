package helper

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go-common-plugin/model"
)

func Ok[T any](data T) model.Response[T] {
	return model.Response[T]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   data,
	}
}

func BadRequest[T any](err error) model.Response[T] {
	return model.Response[T]{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: toMapError(err),
	}
}

func NotFound[T any]() model.Response[T] {
	return model.Response[T]{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
	}
}

func InternalServerErrorp[T any]() model.Response[T] {
	return model.Response[T]{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
	}
}

func toMapError(err error) map[string][]string {
	groupOfErrors := make(map[string][]string, 0)

	v := err.(validator.ValidationErrors)

	if errors.As(err, &v) {
		for _, fieldError := range v {
			groupOfErrors[fieldError.Field()] = []string{toValidationErrorMessage(fieldError)}
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
