package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code  int
	Data  interface{}
	Error int
}

type response struct {
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
}

func BaseHandler[T any](handler func(echo.Context, *T) Response, req *T) func(c echo.Context) error {
	return func(c echo.Context) error {
		if req != nil {
			if err := c.Bind(req); err != nil {
				return c.JSON(http.StatusBadRequest, response{
					Data:  nil,
					Error: ErrBadRequest,
				})
			}
			if err := c.Validate(req); err != nil {
				return c.JSON(http.StatusBadRequest, response{
					Data:  nil,
					Error: ErrValidationFail,
				})
			}
		}

		resp := handler(c, req)
		return c.JSON(resp.Code, response{
			Data:  resp.Data,
			Error: resp.Error,
		})
	}
}

type RequestValidator struct {
	Validator *validator.Validate
}

func (v *RequestValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
