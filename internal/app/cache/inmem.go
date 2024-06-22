package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

/**
 * The Inmem type uses go-cache to cache data in memory.
 * It implements the Cache interface.
 */
type Inmem struct {
	cache *cache.Cache
}

// The newInmem function creates a new in-memory cache.
func newInmem(_ context.Context, ttl time.Duration) (Cache, error) {
	return &Inmem{
		cache: cache.New(ttl, ttl),
	}, nil
}

// The GetCache function retrieves value as bytes from the in-memory cache.
func (i *Inmem) GetCache(_ context.Context, key string) ([]byte, error) {
	if data, found := i.cache.Get(key); found {
		return data.([]byte), nil
	}

	return nil, fmt.Errorf("key not found")
}

// The SetCache function sets a value in the in-memory cache.
func (i *Inmem) SetCache(_ context.Context, key string, data []byte) error {
	i.cache.Set(key, data, cache.DefaultExpiration)
	return nil
}
