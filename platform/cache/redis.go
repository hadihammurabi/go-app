package cache

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func ConfigureRedis() (Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})

	return &cache{
		client: rdb,
	}, nil
}

func (c cache) IsAvailable() error {
	if c.client == nil {
		return errors.New("cache is not available")
	}

	return nil
}

func (c cache) Set(key string, val interface{}, exp ...time.Duration) error {
	if err := c.IsAvailable(); err != nil {
		return err
	}

	ctx := context.Background()
	expireAt := time.Duration(0)
	if len(exp) > 0 {
		expireAt = exp[0]
	}

	jsonMarshal, err := json.Marshal(val)
	if err == nil {
		return c.client.Set(ctx, key, val, expireAt).Err()
	}

	return c.client.Set(ctx, key, jsonMarshal, expireAt).Err()
}

func (c cache) Get(key string) (interface{}, error) {
	if err := c.IsAvailable(); err != nil {
		return nil, err
	}
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var jsonUnmarshal interface{}
	err = json.Unmarshal([]byte(val), &jsonUnmarshal)
	if err == nil {
		return jsonUnmarshal, nil
	}

	return val, nil
}
