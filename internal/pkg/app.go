package pkg

import (
	"log"

	"github.com/yatoenough/filestore/internal/app/config"
)

type App struct {
	cfg *config.Config
}

func New() (*App, error) {
	a := &App{}

	a.cfg = config.MustLoad()
	log.Print(a.cfg)

	return a, nil
}
