package main

import (
	"context"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Wait for Ctrl+C
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

}
