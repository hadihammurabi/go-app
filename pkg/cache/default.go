package cache

import (
	"context"
	"errors"
	"time"

	"github.com/gowok/gowok/driver"
)

type def struct {
}

func ConfigureDefault(config Config) (driver.Cache, error) {
	return &def{}, nil
}

func (c def) IsAvailable(ctx context.Context) bool {
	return false
}

func (c def) Set(ctx context.Context, key string, val any, ttl ...time.Duration) error {
	return errors.New("cache is not available")
}

func (c def) Get(ctx context.Context, key string) (any, error) {
	return nil, errors.New("cache is not available")
}
