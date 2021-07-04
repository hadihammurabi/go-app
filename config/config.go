package config

import (
	"os"

	"github.com/hadihammurabi/belajar-go-rest-api/platform/cache"
	"gorm.io/gorm"
)

type Config struct {
	APP   *AppConfig
	JWT   *JWTConfig
	DB    *gorm.DB
	Cache cache.Cache
}

func New() (*Config, error) {
	jwtConf := ConfigureJWT()
	appConf := ConfigureApp()

	dbConf, err := ConfigureDatabase()
	if err != nil {
		return nil, err
	}

	conf := &Config{
		APP: appConf,
		JWT: jwtConf,
		DB:  dbConf,
	}

	cacheDriver := os.Getenv("CACHE_DRIVER")
	if cacheDriver != "" {
		cacheConf, err := ConfigureCache()
		if err == nil {
			conf.Cache = cacheConf
		}

	}

	return conf, nil
}
