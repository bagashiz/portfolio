package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// The NewCache function creates a new Redis or other compatible in-memory store client.
func NewCache(ctx context.Context, url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
