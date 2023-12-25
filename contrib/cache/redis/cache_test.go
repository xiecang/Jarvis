package redis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Person struct {
	Name string
}

var conf = Config{
	Addr:     "127.0.0.1:6379",
	UserName: "",
	Password: "",
	Database: 0,
}

func TestNewCache(t *testing.T) {
	cache := NewCache[int, string](conf)
	err := cache.Set(context.Background(), 1, "hello")
	assert.Nil(t, err)

	got, ok, err := cache.Get(context.Background(), 1)
	assert.Equal(t, ok, true)
	assert.Equal(t, got, "hello")
}

func TestNewPersonCache(t *testing.T) {
	cache := NewCache[string, Person](conf)
	err := cache.Set(context.Background(), "john", Person{Name: "John Smith"})
	assert.Nil(t, err)

	got, ok, err := cache.Get(context.Background(), "john")
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
	assert.NotNil(t, got)
	assert.Equal(t, got.Name, "John Smith")
}

func TestExpireCache(t *testing.T) {
	cache := NewCache[string, Person](conf)
	err := cache.SetWithExp(context.Background(), "john", Person{Name: "John Smith"}, time.Second)
	assert.Nil(t, err)

	got, ok, err := cache.Get(context.Background(), "john")
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
	assert.NotNil(t, got)
	assert.Equal(t, got.Name, "John Smith")

	time.Sleep(2 * time.Second)
	got, ok, err = cache.Get(context.Background(), "john")
	assert.Nil(t, err)
	assert.NotEqual(t, ok, true)
	assert.Empty(t, got)
}
