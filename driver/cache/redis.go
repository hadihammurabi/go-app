package cache

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func ConfigureRedis() (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})

	return &Redis{
		client: rdb,
	}, nil
}

func (c Redis) IsAvailable() error {
	if c.client == nil {
		return errors.New("cache is not available")
	}

	ctx := context.Background()
	if err := c.client.Ping(ctx).Err(); err != nil {
		return err
	}

	return nil
}

func (c Redis) Set(key string, val any, exp ...time.Duration) error {
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
		return c.client.Set(ctx, key, string(jsonMarshal), expireAt).Err()
	}

	return c.client.Set(ctx, key, val, expireAt).Err()
}

func (c Redis) Get(key string) (any, error) {
	if err := c.IsAvailable(); err != nil {
		return nil, err
	}
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var jsonUnmarshal any
	err = json.Unmarshal([]byte(val), &jsonUnmarshal)
	if err == nil {
		return jsonUnmarshal, nil
	}

	return val, nil
}
