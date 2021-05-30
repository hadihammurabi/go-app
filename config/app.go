package config

import (
	"fmt"
	"os"
)

// AppConfig struct
type AppConfig struct {
	Port string
}

// ConfigureApp func
func ConfigureApp() *AppConfig {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	return &AppConfig{
		Port: fmt.Sprintf(":%s", port),
	}
}
