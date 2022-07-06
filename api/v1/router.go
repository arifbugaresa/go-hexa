package api

import (
	"github.com/arifbugaresa/go-hexa/api/v1/absensi"
	"github.com/arifbugaresa/go-hexa/api/v1/aktivitas"
	"github.com/arifbugaresa/go-hexa/api/v1/health"
	"github.com/arifbugaresa/go-hexa/api/v1/user_info"
	"github.com/arifbugaresa/go-hexa/middleware/auth"
	"github.com/labstack/echo/v4"
)

func Controller(
	e *echo.Echo,
	healthController *health.Controller,
	userInfoController *user_info.Controller,
	absensiController *absensi.Controller,
	aktivitasController *aktivitas.Controller,
) {
	version := "/v1/api"

	// Health
	e.GET("/health", healthController.Health)

	// User Info
	userInfo := e.Group(version)
	userInfo.POST("/login", userInfoController.Login)
	userInfo.POST("/logout", userInfoController.Logout)
	userInfo.GET("/user-info", userInfoController.GetListUserInfo, auth.ValidateJwtMiddleware)

	// Absensi
	absensi := e.Group(version)
	absensi.POST("/user/checkin", absensiController.CheckIn, auth.ValidateJwtMiddleware)
	absensi.POST("/user/checkout", absensiController.CheckOut, auth.ValidateJwtMiddleware)
	absensi.GET("/user/absensi", absensiController.GetListAbsensiByIDUser, auth.ValidateJwtMiddleware)

	// Aktivitas
	aktivitas := e.Group(version)
	aktivitas.POST("/user/aktivitas", aktivitasController.CreateAktivitas, auth.ValidateJwtMiddleware)
	aktivitas.PUT("/user/aktivitas/:id", aktivitasController.UpdateAktivitas, auth.ValidateJwtMiddleware)
	aktivitas.DELETE("/user/aktivitas/:id", aktivitasController.DeleteAktivitas, auth.ValidateJwtMiddleware)

}
