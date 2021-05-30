package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config/database"
	"gorm.io/gorm"
)

type Config struct {
	JWT *JWTConfig
	DB  *gorm.DB
}

func New() (*Config, error) {
	jwtConf := ConfigureJWT()

	dbConf, err := database.ConfigureDatabase()
	if err != nil {
		return nil, err
	}

	conf := &Config{
		JWT: jwtConf,
		DB:  dbConf,
	}
	return conf, nil
}
