package config

import (
	"github.com/spf13/viper"
)

// AppConfig struct
type AppConfig struct {
	Name string
	Port string
}

// ConfigureApp func
func ConfigureApp() AppConfig {
	name := viper.GetString("app.name")
	if name == "" {
		name = "SuperApp"
	}

	port := viper.GetString("app.port")
	if port == "" {
		port = "8080"
	}

	return AppConfig{
		Name: name,
		Port: port,
	}
}
