package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/cache"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/database"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
)

type Config struct {
	APP       AppConfig
	JWT       JWTConfig
	DB        *database.Database
	Cache     cache.Cache
	Messaging messaging.Messaging
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

	// redis, err := ConfigureRedis()
	// if err != nil {
	// 	log.Printf("can not configure Redis. Caused by %s\n", err.Error())
	// } else {
	// 	conf.Redis = redis
	// }

	return conf, nil
}
