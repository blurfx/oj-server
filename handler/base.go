package handler

import (
    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
    "net/http"
)

type BaseResponse struct {
    Code  int
    Data  interface{}
    Error int
}

type BaseHttpResponse struct {
    Data  interface{} `json:"data"`
    Error int         `json:"error"`
}

func BaseHandler(handler func(echo.Context) BaseResponse) func(c echo.Context) error {
    return func(c echo.Context) error {
        resp := handler(c)

        return c.JSON(resp.Code, BaseHttpResponse{
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
