package cache

import (
	"context"
	"time"
)

// Cache defines the generics cache interface for 1.8 or above version.
// Notice: SetWithExp function is not necessarily implemented
type Cache[K comparable, V any] interface {
	Get(ctx context.Context, key K) (value V, ok bool, err error)
	Set(ctx context.Context, key K, val V) (err error)
	SetWithExp(ctx context.Context, key K, val V, exp time.Duration) (err error)
	Delete(ctx context.Context, key K) (err error)
}
