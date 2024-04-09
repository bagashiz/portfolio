package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

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
		"REDIS_URL":           getEnv("REDIS_URL"),
	}

	httpServer := server.NewServer(config)

	go func() {
		fmt.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()

	return nil
}
