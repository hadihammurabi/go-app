package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
	"github.com/spf13/viper"
)

// ConfigureMessaging func
func ConfigureMessaging() messaging.Messaging {
	messagingFromConfig, _ := viper.Get("messaging").(map[string]any)
	driver, _ := messagingFromConfig["driver"].(string)
	url, _ := messagingFromConfig["url"].(string)

	conf := messaging.Config{
		Driver: driver,
		URL:    url,
	}

	m, err := messaging.New(conf)
	if err != nil {
		panic("messaging configuration failed")
	}

	return m
}
