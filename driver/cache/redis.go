package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

type redis struct {
	client *goredis.Client
}

func ConfigureRedis(config Config) (Cache, error) {
	rdb := goredis.NewClient(&goredis.Options{
		Addr: config.URL,
	})

	return &redis{
		client: rdb,
	}, nil
}

func (c redis) IsAvailable() bool {
	if c.client == nil {
		return true
	}

	ctx := context.Background()
	err := c.client.Ping(ctx).Err()
	return err == nil
}

func (c redis) Set(key string, val any, ttl ...time.Duration) error {
	if !c.IsAvailable() {
		return errors.New("cache is not available")
	}

	expireAt := time.Duration(0)
	if len(ttl) > 0 {
		expireAt = ttl[0]
	}

	ctx := context.Background()
	jsonMarshal, err := json.Marshal(val)
	if err == nil {
		return c.client.Set(ctx, key, string(jsonMarshal), expireAt).Err()
	}

	return c.client.Set(ctx, key, val, expireAt).Err()
}

func (c redis) Get(key string) (any, error) {
	if !c.IsAvailable() {
		return nil, errors.New("cache is not available")
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
