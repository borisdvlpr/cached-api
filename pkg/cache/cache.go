package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	"cachedapi/pkg/config"
)

type Client struct {
	client *redis.Client
	ttl    int
}

func NewClient(cfg *config.Config) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.Db,
		Protocol: 2,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	log.Printf("Connected to Redis.")
	return &Client{client: client, ttl: cfg.TTL}, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
