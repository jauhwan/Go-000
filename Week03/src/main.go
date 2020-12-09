package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"jauhwan.com/Go-000/Week03/src/initial"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return initial.InitHttpServer(ctx, exit)
	})
	group.Wait()
}
