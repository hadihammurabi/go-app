package cache

import (
	"errors"
	"time"
)

type def struct {
}

func ConfigureDefault(config Config) (Cache, error) {
	return &def{}, nil
}

func (c def) IsAvailable() bool {
	return false
}

func (c def) Set(key string, val any, ttl ...time.Duration) error {
	return errors.New("cache is not available")
}

func (c def) Get(key string) (any, error) {
	return nil, errors.New("cache is not available")
}
