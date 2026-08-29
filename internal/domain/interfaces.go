package domain

import "context"

type CacheRepo[T any] interface {
	Set(ctx context.Context, id string, data *T) error
	Get(ctx context.Context, id string) (*T, error)
	Del(ctx context.Context, id string) error
}
