package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/bagashiz/portfolio/internal/app/cache"
	"github.com/bagashiz/portfolio/internal/app/server"
)

/**
 * The Run function sets up the server and starts it.
 * It also listens for an interrupt signal to shut down the server gracefully.
 */
func Run(ctx context.Context, getEnv func(string) string) error {
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

	fmt.Printf("using %s as cache with TTL of %s\n", config["CACHE_TYPE"], ttl)

	cache, err := cache.NewCache(ctx, cache.Config{
		Type: config["CACHE_TYPE"],
		URL:  config["CACHE_URL"],
		TTL:  ttl,
	})
	if err != nil {
		return err
	}

	httpServer := server.NewServer(config, cache)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("error listening and serving: %s\n", err)
		}
	}()

	fmt.Printf("listening on %s\n", httpServer.Addr)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("error shutting down server: %s\n", err)
			return
		}

		fmt.Println("server shut down gracefully")
	}()

	wg.Wait()

	return nil
}
