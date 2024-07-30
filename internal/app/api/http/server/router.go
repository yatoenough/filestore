package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	group.GET("/ping", func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
		return nil
	})
}
