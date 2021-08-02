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
	Redis *cache.Redis
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

	redis, err := ConfigureRedis()
	if err == nil {
		conf.Redis = redis
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
