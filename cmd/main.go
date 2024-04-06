package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/bagashiz/portfolio/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}
}

func run(ctx context.Context, addr string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	httpServer := app.NewServer(addr)

	go func() {
		fmt.Printf("listening on %s\n", addr)
		if err := httpServer.ListenAndServe(); err != nil {
			fmt.Printf("error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	addr := net.JoinHostPort(os.Args[1], os.Args[2])
	if err := run(ctx, addr); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
