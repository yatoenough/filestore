package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group) {
	e.GET("/ping", func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
		return nil
	})
}
