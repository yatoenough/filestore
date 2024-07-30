package pkg

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yatoenough/filestore/internal/app/api/http/server"
	"github.com/yatoenough/filestore/internal/app/config"
	"github.com/yatoenough/filestore/internal/app/database/pg"
)

type App struct {
	cfg *config.Config
	db  *pg.Storage
	srv *echo.Echo
}

func New() *App {
	a := &App{}
	a.cfg = config.MustLoad()

	db, err := pg.New(a.cfg.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	a.db = db

	a.srv = server.New()

	return a
}

func (a *App) Run() {
	err := a.srv.Start(a.cfg.Address)
	if err != nil {
		log.Fatal(err)
	}
}
