package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// signal.NotifyContext added in go1.16
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// errgroup manage running of project subsystems
	g, gCtx := errgroup.WithContext(mainCtx)

	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("done")
		// Reset os.Interrupt default behavior, similar to signal.Reset
		// use stop() here if you want to interrupt graceful shutdown on second SIGINT
		// stop()
		time.Sleep(20 * time.Second)
		// use stop() here if you want to ignore the following SIGINT signal and fully complete shutdown
		stop()
		return gCtx.Err()
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
