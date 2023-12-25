# Memory cache

> The memory cache implemented by [go-generics-cache](https://github.com/Code-Hex/go-generics-cache)

## Usage


```go

import (
    "github.com/xiecang/jarvis/cache"
    "github.com/xiecang/jarvis/contrib/cache/memory"
)
```

```go
func TestNewCache(t *testing.T) {
	cache := NewCache[int, string]()
	err := cache.Set(context.Background(), 1, "hello")
	assert.Nil(t, err)

	got, ok, err := cache.Get(context.Background(), 1)
	assert.Equal(t, ok, true)
	assert.Equal(t, got, "hello")
}
```