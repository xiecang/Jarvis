package memory

import (
	"context"
	"time"

	genericsCache "github.com/Code-Hex/go-generics-cache"
	"github.com/xiecang/jarvis/cache"
)

// Cache is implemented by third-party packages which provide from github.com/Code-Hex/go-generics-cache
type Cache[K comparable, V any] struct {
	cache *genericsCache.Cache[K, V]
}

func (c Cache[K, V]) Get(ctx context.Context, key K) (value V, ok bool, err error) {
	value, ok = c.cache.Get(key)
	return
}

func (c Cache[K, V]) Set(ctx context.Context, key K, val V) (err error) {
	c.cache.Set(key, val)
	return nil
}

func (c Cache[K, V]) SetWithExp(ctx context.Context, key K, val V, exp time.Duration) (err error) {
	c.cache.Set(key, val, genericsCache.WithExpiration(exp))
	return nil
}

func (c Cache[K, V]) Delete(ctx context.Context, key K) (err error) {
	c.cache.Delete(key)
	return nil
}

// NewCache returns a new cache with the given types
func NewCache[K comparable, V any]() cache.Cache[K, V] {
	c := genericsCache.New[K, V]()
	return &Cache[K, V]{
		c,
	}
}
