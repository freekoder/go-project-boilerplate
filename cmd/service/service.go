package main

import (
	"context"
	"fmt"
	"github.com/freekoder/go-project-boilerplate/internal/web"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// signal.NotifyContext added in go1.16
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// errgroup manage running of project subsystems
	g, gCtx := errgroup.WithContext(mainCtx)

	server := web.New(gCtx)

	g.Go(func() error {
		return server.Run()
	})

	g.Go(func() error {
		<-gCtx.Done()
		// Reset os.Interrupt default behavior, similar to signal.Reset
		// use stop() here if you want to interrupt graceful shutdown on second SIGINT
		// stop()
		_ = server.Shutdown()
		// use stop() here if you want to ignore the following SIGINT signal and fully complete shutdown
		stop()
		return gCtx.Err()
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
