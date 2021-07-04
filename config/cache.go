package config

import (
	"errors"
	"os"

	"github.com/hadihammurabi/belajar-go-rest-api/platform/cache"
)

const (
	driverRedis = "redis"
)

// ConfigureCache func
func ConfigureCache() (cache.Cache, error) {
	driver := os.Getenv("CACHE_DRIVER")

	if driver == driverRedis {
		cache, err := cache.ConfigureRedis()
		return cache, err
	}

	return nil, errors.New("unknown cache driver")
}
