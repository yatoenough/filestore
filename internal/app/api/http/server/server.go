package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yatoenough/filestore/internal/app/config"
)

type Server struct {
	cfg  *config.Config
	echo *echo.Echo
}

func New(cfg *config.Config) *Server {
	e := echo.New()
	s := &Server{
		cfg:  cfg,
		echo: e,
	}

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.registerRoutes()

	return s
}

func (s *Server) registerRoutes() {
	s.echo.RouteNotFound("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]any{
			"code":      http.StatusNotFound,
			"message":   "Not found",
			"timestamp": time.Now(),
		})
	})
}

func (s *Server) Start() error {
	err := s.echo.Start(s.cfg.Address)
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Shutting down...")
			return nil
		}
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
