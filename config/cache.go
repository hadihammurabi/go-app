package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/cache"
	"github.com/spf13/viper"
)

// ConfigureCache func
func ConfigureCache() cache.Cache {
	messagingFromConfig := viper.Get("cache").(map[string]any)

	driver, ok := messagingFromConfig["driver"].(string)
	if !ok {
		panic("cache driver configuration failed")
	}

	url, ok := messagingFromConfig["url"].(string)
	if !ok {
		panic("cache url configuration failed")
	}

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
