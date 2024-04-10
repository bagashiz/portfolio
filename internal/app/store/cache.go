package store

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// The Cache interface provides a way to interact with an in-memory store.
type Cache interface {
	// GetCache retrieves value as bytes from the cache.
	GetCache(ctx context.Context, key string) ([]byte, error)
	// SetCache sets a value in the cache.
	SetCache(ctx context.Context, key string, data any, ttl time.Duration) error
}

/**
 * The Redis type represents a Redis or other compatible in-memory store.
 * It implements the Cache interface.
 */
type Redis struct {
	*redis.Client
}

// The NewRedis function creates a new Redis or other compatible in-memory store client.
func NewRedis(ctx context.Context, url string) (Cache, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client}, nil
}

// The Get function retrieves value as bytes from the cache.
func (r *Redis) GetCache(ctx context.Context, key string) ([]byte, error) {
	return r.Get(ctx, key).Bytes()
}

// The Set function sets a value in the cache.
func (r *Redis) SetCache(ctx context.Context, key string, data any, ttl time.Duration) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	status := r.Set(ctx, key, dataJSON, 1*time.Hour)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}
