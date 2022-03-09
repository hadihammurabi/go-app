package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/cache"
)

// ConfigureRedis func
func ConfigureRedis() (*cache.Redis, error) {
	return cache.ConfigureRedis()
}
