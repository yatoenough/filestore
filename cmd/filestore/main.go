package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yatoenough/filestore/internal/pkg"
)

func main() {
	app := pkg.New()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go app.Run()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.Stop(ctx)
}
