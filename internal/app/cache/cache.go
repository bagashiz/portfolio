package cache

import (
	"context"
	"fmt"
	"time"
)

// The Cache interface provides a way to interact with an in-memory cache.
type Cache interface {
	// GetCache retrieves value as bytes from the cache.
	GetCache(ctx context.Context, key string) ([]byte, error)
	// SetCache sets a value in the cache.
	SetCache(ctx context.Context, key string, data []byte) error
}

// Config represents the configuration for a cache.
type Config struct {
	Type string
	URL  string
	TTL  time.Duration
}

// NewCache creates a new cache based on the provided configuration.
func NewCache(ctx context.Context, cfg Config) (Cache, error) {
	switch cfg.Type {
	case "redis":
		if cfg.URL == "" {
			return nil, fmt.Errorf("missing redis URL")
		}
		return newRedis(ctx, cfg.URL, cfg.TTL)
	case "inmem":
		return newInmem(ctx, cfg.TTL)
	default:
		return nil, fmt.Errorf("expected 'redis' or 'inmem', got %s", cfg.Type)
	}
}
