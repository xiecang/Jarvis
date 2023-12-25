package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/xiecang/jarvis/cache"
	"time"
)

type Config struct {
	Addr     string `yaml:"addr"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

// Cache is a redis cache wrapper which implemented the cache interface
type Cache[K comparable, V any] struct {
	client *redis.Client
}

func (r Cache[K, V]) key(key K) (keyStr string) {
	keyStr = fmt.Sprintf("%v", key)
	return
}
func (r Cache[K, V]) Get(ctx context.Context, key K) (value V, ok bool, err error) {
	k := r.key(key)

	var data []byte
	data, err = r.client.Get(ctx, k).Bytes()
	if errors.Is(err, redis.Nil) {
		return value, false, nil
	}
	if err != nil {
		return value, false, err
	}

	if err := json.Unmarshal(data, &value); err != nil {
		return value, false, err
	}
	return value, true, nil
}

func (r Cache[K, V]) Set(ctx context.Context, key K, val V) (err error) {
	return r.SetWithExp(ctx, key, val, 0)
}

func (r Cache[K, V]) SetWithExp(ctx context.Context, key K, val V, exp time.Duration) (err error) {
	k := r.key(key)
	var value []byte
	value, err = json.Marshal(val)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, k, value, exp).Err()
	return
}

func (r Cache[K, V]) Delete(ctx context.Context, key K) (err error) {
	k := r.key(key)
	return r.client.Del(ctx, k).Err()
}

// NewCache returns a new redis cache interface with redis instance configured
func NewCache[K comparable, V any](config Config) cache.Cache[K, V] {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Username: config.UserName,
		Password: config.Password,
		DB:       config.Database,
	})

	return &Cache[K, V]{
		client: client,
	}
}
