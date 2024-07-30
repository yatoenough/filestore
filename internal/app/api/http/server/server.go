package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	registerRoutes(e)

	return e
}

func registerRoutes(e *echo.Echo) {
	e.RouteNotFound("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]any{
			"code":      http.StatusNotFound,
			"message":   "Not found",
			"timestamp": time.Now(),
		})
	})
}
