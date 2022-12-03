package config

import (
	"github.com/gowok/gowok/driver"
)

type Config struct {
	APP       AppConfig
	JWT       JWTConfig
	Cache     driver.Cache
	Messaging driver.Messaging
}

func New() (Config, error) {
	appConf := ConfigureApp()
	jwtConf := ConfigureJWT()
	messagingConf := ConfigureMessaging()
	cacheConf := ConfigureCache()

	ConfigureDatabase()

	conf := Config{
		APP:       appConf,
		JWT:       jwtConf,
		Cache:     cacheConf,
		Messaging: messagingConf,
	}

	return conf, nil
}
