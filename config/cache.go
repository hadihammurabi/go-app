package config

import (
	"github.com/gowok/gowok/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/cache"
	"github.com/spf13/viper"
)

// ConfigureCache func
func ConfigureCache() driver.Cache {
	messagingFromConfig, _ := viper.Get("cache").(map[string]any)
	driver, _ := messagingFromConfig["driver"].(string)
	url, _ := messagingFromConfig["url"].(string)

	conf := cache.Config{
		Driver: cache.Driver(driver),
		URL:    url,
	}

	c, err := cache.New(conf)
	if err != nil {
		panic("cache configuration failed")
	}

	return c
}
