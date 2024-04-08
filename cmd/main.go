package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bagashiz/portfolio/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}
}

func main() {
	ctx := context.Background()

	if err := app.Run(ctx, os.Getenv); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
