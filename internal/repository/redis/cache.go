package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache[T any] struct {
	client *redis.Client
	prefix string
	ttl    time.Duration
}

func NewCache[T any](client *redis.Client, prefix string, ttl time.Duration) *Cache[T] {
	return &Cache[T]{client: client, prefix: prefix, ttl: ttl}
}

func (c *Cache[T]) key(id string) string {
	return c.prefix + ":" + id
}

func (c *Cache[T]) Set(ctx context.Context, id string, data *T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("cant umarshal data adding to cache: %w", err)
	}

	err = c.client.Set(ctx, c.key(id), jsonData, c.ttl).Err()
	return err
}

func (c *Cache[T]) Get(ctx context.Context, id string) (*T, error) {
	var data T

	jsonData, err := c.client.Get(ctx, c.key(id)).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("no data in cache: %w", err)
	}

	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, fmt.Errorf("cant unmarshal cache: %w", err)
	}
	return &data, nil
}

func (c *Cache[T]) Del(ctx context.Context, id string) error {
	err := c.client.Del(ctx, c.key(id)).Err()
	return err
}
