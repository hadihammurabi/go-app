package config

import (
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

	dbConf, err := ConfigureDatabase()
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
