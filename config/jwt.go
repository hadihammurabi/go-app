package config

import "os"

// JWTConfig struct
type JWTConfig struct {
	Secret string
}

// ConfigureJWT func
func ConfigureJWT() (config *JWTConfig) {
	secret := os.Getenv("APP_KEY")
	if secret == "" {
		secret = "2Yu4i1lTSrmigPyb9RYxYJ35WcnxDOQsCBCOTfoo2Yu4i1lTSrmigPyb9RYx"
	}

	config = &JWTConfig{
		Secret: secret,
	}

	return
}
