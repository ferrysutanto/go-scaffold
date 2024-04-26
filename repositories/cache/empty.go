package cache

import (
	"context"

	"github.com/ferrysutanto/go-errors"
)

var ErrNotImplemented = errors.NewWithCode("not implemented", 501)

type emptyCache struct{}

func (c *emptyCache) Get(ctx context.Context, key string) (string, error) {
	return "", ErrNotImplemented
}

func (c *emptyCache) Set(ctx context.Context, key, value string) error {
	return ErrNotImplemented
}

func newEmpty() ICache {
	return &emptyCache{}
}
