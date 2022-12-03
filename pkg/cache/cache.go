package cache

import (
	"github.com/gowok/gowok/driver"
)

type Driver string

const (
	DriverRedis Driver = "redis"
)

type Config struct {
	Driver Driver
	URL    string
}

func New(config Config) (driver.Cache, error) {
	if config.Driver == DriverRedis {
		return ConfigureRedis(config)
	}

	return ConfigureDefault(config)
}
