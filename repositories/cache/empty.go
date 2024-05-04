package cache

import (
	"context"

	"github.com/ferrysutanto/go-errors"
)

var errNotImplemented = errors.NewWithCode("function not implemented", 501)

type emptyCache struct{}

func (c *emptyCache) Get(ctx context.Context, key string) (string, error) {
	return "", errNotImplemented
}

func (c *emptyCache) Set(ctx context.Context, key, value string) error {
	return errNotImplemented
}

func newEmpty() ICache {
	return &emptyCache{}
}
