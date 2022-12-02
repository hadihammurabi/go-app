package config

import (
	"github.com/gowok/gowok/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/database"
)

type Config struct {
	APP       AppConfig
	JWT       JWTConfig
	DB        *database.Database
	Cache     driver.Cache
	Messaging driver.Messaging
}

func New() (Config, error) {
	appConf := ConfigureApp()
	jwtConf := ConfigureJWT()
	dbConf := ConfigureDatabase()
	messagingConf := ConfigureMessaging()
	cacheConf := ConfigureCache()

	conf := Config{
		APP:       appConf,
		JWT:       jwtConf,
		DB:        dbConf,
		Cache:     cacheConf,
		Messaging: messagingConf,
	}

	return conf, nil
}
