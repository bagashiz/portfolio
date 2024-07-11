package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/bagashiz/portfolio/internal/cache"
	"github.com/bagashiz/portfolio/internal/server"
	"github.com/joho/godotenv"
)

// init loads the environment variables from the .env file for local development.
func init() {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}
}

// entry point of the application.
func main() {
	ctx := context.Background()

	if err := run(ctx, os.Getenv); err != nil {
		slog.Error("error running application", "error", err)
		os.Exit(1)
	}
}

// run starts the server and listens for incoming requests.
func run(ctx context.Context, getEnv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config := map[string]string{
		"APP_HOST":            getEnv("APP_HOST"),
		"APP_PORT":            getEnv("APP_PORT"),
		"DEV_USERNAME":        getEnv("DEV_USERNAME"),
		"GITHUB_USERNAME":     getEnv("GITHUB_USERNAME"),
		"GITHUB_ACCESS_TOKEN": getEnv("GITHUB_ACCESS_TOKEN"),
		"CACHE_TYPE":          getEnv("CACHE_TYPE"),
		"CACHE_URL":           getEnv("CACHE_URL"),
		"CACHE_TTL":           getEnv("CACHE_TTL"),
	}

	ttl, err := time.ParseDuration(config["CACHE_TTL"])
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("using %s as cache with TTL of %s\n", config["CACHE_TYPE"], ttl))

	cache, err := cache.NewCache(ctx, cache.Config{
		Type: config["CACHE_TYPE"],
		URL:  config["CACHE_URL"],
		TTL:  ttl,
	})
	if err != nil {
		return err
	}

	httpServer := server.New(ctx, config, cache)

	if err := httpServer.Start(ctx); err != nil {
		return err
	}

	return nil
}
