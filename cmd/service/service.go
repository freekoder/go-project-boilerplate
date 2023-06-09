package main

import (
	"context"
	"fmt"
	"github.com/freekoder/go-project-boilerplate/internal/config"
	"github.com/freekoder/go-project-boilerplate/internal/logger"
	"github.com/freekoder/go-project-boilerplate/internal/service"
	"github.com/freekoder/go-project-boilerplate/internal/web"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var cfg config.Config
	err := cleanenv.ReadConfig("configs/config.yml", &cfg)
	if err != nil {
		fmt.Printf("can not load configuration: %v", err)
		os.Exit(1)
	}

	log, err := logger.Configure(cfg.Logging)
	if err != nil {
		fmt.Printf("can not configure logger: %v", err)
		return
	}

	log.Info("starting service")

	// signal.NotifyContext added in go1.16
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// errgroup manage running of project subsystems
	g, gCtx := errgroup.WithContext(mainCtx)

	appService := service.New()

	webServer := web.New(gCtx, cfg.Web, log, appService)

	g.Go(webServer.Run)

	g.Go(func() error {
		<-gCtx.Done()
		// Reset os.Interrupt default behavior, similar to signal.Reset
		// use stop() here if you want to interrupt graceful shutdown on second SIGINT
		// stop()
		_ = webServer.Shutdown()
		_ = log.Sync()
		// use stop() here if you want to ignore the following SIGINT signal and fully complete shutdown
		stop()
		return gCtx.Err()
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
