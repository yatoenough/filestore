package pkg

import (
	"context"
	"log"
	"net/http"

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
	if err := a.srv.Start(a.cfg.Address); err != nil && err == http.ErrServerClosed {
		log.Println("Shutting down server...")
	}
}

func (a *App) Stop(ctx context.Context) error {
	log.Println("Stopping application...")
	err := a.db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = a.srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Application stopped gracefully")
	return nil
}
