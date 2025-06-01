package handler

import "github.com/labstack/echo/v4"

func InitV1Handler(e *echo.Echo) {
	e.GET("v1/healthcheck", Healthcheck)
	e.GET("v1/auth/login", BaseHandler(V1Login, &LoginRequest{}))
	e.GET("v1/auth/logout", BaseHandler(V1Logout, nil))
	e.POST("v1/problems", BaseHandler(V1GetProblems, &V1GetProblemsRequest{}))
	e.GET("v1/problems/:id", BaseHandler(V1GetProblem, nil))
}
