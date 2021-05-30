package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config/database"
	"gorm.io/gorm"
)

type Config struct {
	APP *AppConfig
	JWT *JWTConfig
	DB  *gorm.DB
}

func New() (*Config, error) {
	jwtConf := ConfigureJWT()
	appConf := ConfigureApp()

	dbConf, err := database.ConfigureDatabase()
	if err != nil {
		return nil, err
	}

	conf := &Config{
		APP: appConf,
		JWT: jwtConf,
		DB:  dbConf,
	}
	return conf, nil
}
