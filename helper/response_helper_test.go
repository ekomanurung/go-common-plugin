package helper

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestResponseHelper(t *testing.T) {
	var validate *validator.Validate
	t.Run("bad request", func(t *testing.T) {
		scenarios := []struct {
			Name string `validate:"required"`
			Age  int    `validate:"required,max=9"`
		}{
			{
				Name: "Eko",
				Age:  10,
			},
			{
				Name: "",
				Age:  10,
			},
		}

		validate = validator.New()
		for _, tc := range scenarios {
			err := validate.Struct(tc)
			if err != nil {
				response := BadRequest[any](err)
				assert.NotEqual(t, 0, len(response.Errors))
			}
		}
	})
	t.Run("success", func(t *testing.T) {
		scenarios := []struct {
			Name string `validate:"required"`
			Age  int    `validate:"required,max=9"`
		}{
			{
				Name: "Eko",
				Age:  6,
			},
		}

		validate = validator.New()
		for _, tc := range scenarios {
			err := validate.Struct(tc)
			assert.Equal(t, nil, err)
			response := Ok(tc)
			assert.Equal(t, nil, response.Errors)
		}
	})
	t.Run("not found", func(t *testing.T) {
		response := NotFound[any]()
		assert.Equal(t, nil, response.Errors)
		assert.Equal(t, 404, response.Code)
		assert.Equal(t, http.StatusText(http.StatusNotFound), response.Status)
	})
	t.Run("internal server error", func(t *testing.T) {
		response := InternalServerErrorp[any]()
		assert.Equal(t, nil, response.Errors)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), response.Status)
	})
}
