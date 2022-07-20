package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"t11/internal/cache"
	"t11/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cache := cache.NewCache()

	server.NewHTTPServer(ctx, cache)

	<-ctx.Done()

}
