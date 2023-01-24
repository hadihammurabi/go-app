package cache

import (
	"github.com/gowok/gowok/cache"
	"github.com/gowok/gowok/driver"
)

func New() (driver.Cache, error) {
	return cache.NewMap()
}
