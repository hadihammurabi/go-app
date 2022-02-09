package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
)

type Config struct {
	APP AppConfig
	JWT JWTConfig
	DB  *database.Database
	// Redis *cache.Redis
	// MQ    gorabbitmq.MQ
}

func New() (Config, error) {
	appConf := ConfigureApp()
	jwtConf := ConfigureJWT()
	dbConf := ConfigureDatabase()

	conf := Config{
		APP: appConf,
		JWT: jwtConf,
		DB:  dbConf,
	}

	// redis, err := ConfigureRedis()
	// if err != nil {
	// 	log.Printf("can not configure Redis. Caused by %s\n", err.Error())
	// } else {
	// 	conf.Redis = redis
	// }

	// mqConf, err := ConfigureRabbitMQ()
	// if err != nil {
	// 	log.Printf("can not configure MQ. Caused by %s\n", err.Error())
	// } else {
	// 	conf.MQ = mqConf
	// }

	return conf, nil
}
