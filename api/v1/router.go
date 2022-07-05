package api

import (
	"github.com/arifbugaresa/go-hexa/api/v1/health"
	"github.com/arifbugaresa/go-hexa/api/v1/user_info"
	"github.com/arifbugaresa/go-hexa/middleware/auth"
	"github.com/labstack/echo/v4"
)

func Controller(
	e *echo.Echo,
	healthController *health.Controller,
	userInfoController *user_info.Controller,
) {
	version := "/v1/api"

	// Health
	e.GET("/health", healthController.Health)

	// User Info
	userInfo := e.Group(version)
	userInfo.POST("/login", userInfoController.Login)
	userInfo.GET("/user-info", userInfoController.GetListUserInfo, auth.ValidateJwtMiddleware)

}
