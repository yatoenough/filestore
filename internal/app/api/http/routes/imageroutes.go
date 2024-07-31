package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yatoenough/filestore/internal/app/api/http/handler"
)

func RegisterImageRoutes(g *echo.Group, h handler.ImageHandler) {
	g.GET("/image", h.GetImage)
}
