package cache

import (
	"time"
)

type Driver string

const (
	DriverRedis Driver = "redis"
)

type Config struct {
	Driver Driver
	URL    string
}

type Cache interface {
	Set(key string, value any, ttl ...time.Duration) error
	Get(key string) (any, error)
	IsAvailable() bool
}

func New(config Config) (Cache, error) {
	if config.Driver == DriverRedis {
		return ConfigureRedis(config)
	}

	return ConfigureDefault(config)
}
