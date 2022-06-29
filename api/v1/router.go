package api

import "github.com/labstack/echo/v4"

func Controller(
	e *echo.Echo,
) {

	// helper
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

}
