package api

import (
	"github.com/arifbugaresa/go-hexa/api/v1/user_info"
	"github.com/arifbugaresa/go-hexa/middleware/auth"
	"github.com/labstack/echo/v4"
)

func Controller(
	e *echo.Echo,
	userInfoController *user_info.Controller,
) {
	version := "/v1/api"

	// helper
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	// Public API
	// User Info
	userInfo := e.Group(version)
	userInfo.POST("/login", userInfoController.Login)

	// Authorized API
	// User Info
	userInfoJwt := e.Group(version)
	userInfoJwt.GET("/user-info", userInfoController.GetListUserInfo, auth.ValidateJwtMiddleware)

}