package pkg

import (
	"context"
	"log"

	"github.com/yatoenough/filestore/internal/app/api/http/server"
	"github.com/yatoenough/filestore/internal/app/config"
	"github.com/yatoenough/filestore/internal/app/database/pg"
)

type App struct {
	cfg *config.Config
	db  *pg.Storage
	srv *server.Server
}

func New() *App {
	a := &App{}

	a.cfg = config.MustLoad()
	a.db = pg.New(a.cfg.ConnStr)
	a.srv = server.New(a.cfg)

	return a
}

func (a *App) Run() {
	a.srv.Start()
}

func (a *App) Stop(ctx context.Context) {
	err := a.db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = a.srv.Stop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
