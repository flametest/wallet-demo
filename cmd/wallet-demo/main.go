package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"
	"time"

	log "github.com/flametest/vita/vlog"
	"github.com/flametest/vita/vserver"
	"github.com/flametest/wallet-demo/internal/api"
	"github.com/flametest/wallet-demo/internal/config"
	"github.com/flametest/wallet-demo/internal/container"
)

var cfgFile = flag.String("config", "deploy/config.yaml", "config file")

func main() {
	var err error
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)
	defer stop()

	cfg, err := config.ParseConfig(*cfgFile)
	if err != nil {
		panic(err)
	}
	log.InitLogger(cfg.AppConfig.Name, cfg.LogLevel)
	log.Info().Msg("starting wallet-demo")
	srv, err := vserver.NewEchoServer(ctx, &cfg.AppConfig)
	if err != nil {
		panic(err)
	}
	c, err := container.NewContainer(cfg)
	if err != nil {
		panic(err)
	}
	app := api.NewApp(c)
	srv.Register(app.Router)
	go func() {
		_ = srv.Start(ctx)
	}()

	<-ctx.Done()

	log.Info().Msg("shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}
	log.Info().Msg("Server exiting")
}
