package cache

import "context"

// ICache is an interface that defines cache operations
type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

var g ICache = newEmpty()

func SetGlobal(cache ICache) {
	g = cache
}

func Global() ICache {
	return g
}

func Get(ctx context.Context, key string) (string, error) {
	return g.Get(ctx, key)
}

func Set(ctx context.Context, key string, value string) error {
	return g.Set(ctx, key, value)
}
