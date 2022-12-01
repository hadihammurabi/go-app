package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/database"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
)

type Config struct {
	APP AppConfig
	JWT JWTConfig
	DB  *database.Database
	// Redis *cache.Redis
	Messaging messaging.Messaging
}

func New() (Config, error) {
	appConf := ConfigureApp()
	jwtConf := ConfigureJWT()
	dbConf := ConfigureDatabase()
	messagingConf := ConfigureMessaging()

	conf := Config{
		APP:       appConf,
		JWT:       jwtConf,
		DB:        dbConf,
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
