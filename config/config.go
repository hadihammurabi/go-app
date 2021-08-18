package config

import (
	"gorm.io/gorm"
)

type Config struct {
	APP AppConfig
	JWT JWTConfig
	DB  *gorm.DB
	// Redis *cache.Redis
	// MQ    gorabbitmq.MQ
}

func New() (Config, error) {
	jwtConf := ConfigureJWT()
	appConf := ConfigureApp()

	dbConf, err := ConfigureDatabase()
	if err != nil {
		return Config{}, err
	}

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
