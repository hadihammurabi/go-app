package config

import (
	"os"

	"github.com/hadihammurabi/belajar-go-rest-api/pkg/cache"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/messaging"
	"gorm.io/gorm"
)

type Config struct {
	APP   *AppConfig
	JWT   *JWTConfig
	DB    *gorm.DB
	Cache cache.Cache
	MQ    messaging.Messaging
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

	mqDriver := os.Getenv("MQ_DRIVER")
	if mqDriver != "" {
		mqConf, err := ConfigureMQ()
		if err == nil {
			conf.MQ = mqConf
		}
	}

	return conf, nil
}
