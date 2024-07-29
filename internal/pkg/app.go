package pkg

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yatoenough/filestore/internal/app/api/http/server"
	"github.com/yatoenough/filestore/internal/app/config"
)

type App struct {
	c *config.Config
	s *echo.Echo
}

func New() *App {
	a := &App{}

	a.c = config.MustLoad()
	a.s = server.New()

	return a
}

func (a *App) Run() {
	err := a.s.Start(a.c.Address)
	if err != nil {
		log.Fatal(err)
	}
}
