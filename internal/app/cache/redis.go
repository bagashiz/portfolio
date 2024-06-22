package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

/**
 * The Redis type represents a Redis or other compatible in-memory cache.
 * It implements the Cache interface.
 */
type Redis struct {
	client *redis.Client
	ttl    time.Duration
}

// The newRedis function creates a new Redis or other compatible in-memory cache client.
func newRedis(ctx context.Context, url string, ttl time.Duration) (Cache, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client, ttl}, nil
}

// The Get function retrieves value as bytes from the cache.
func (r *Redis) GetCache(ctx context.Context, key string) ([]byte, error) {
	return r.client.Get(ctx, key).Bytes()
}

// The Set function sets a value in the cache.
func (r *Redis) SetCache(ctx context.Context, key string, data []byte) error {
	status := r.client.Set(ctx, key, data, r.ttl)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}
