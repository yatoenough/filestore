package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
}

func (ih *ImageHandler) GetImage(c echo.Context) error {
	return c.String(http.StatusOK, "Image...")
}
