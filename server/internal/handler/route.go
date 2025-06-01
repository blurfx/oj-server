package handler

import "github.com/labstack/echo/v4"

func InitV1Handler(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("/healthcheck", Healthcheck)
	v1.POST("/auth/login", BaseHandler(V1Login, &LoginRequest{}))
	v1.GET("/auth/logout", BaseHandler(V1Logout, nil))
	v1.POST("/auth/register", BaseHandler(V1Register, &RegisterRequest{}))
	v1.POST("/problems", BaseHandler(V1GetProblems, &V1GetProblemsRequest{}))
	v1.GET("/problems/:id", BaseHandler(V1GetProblem, nil))
}
