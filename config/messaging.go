package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
	"github.com/spf13/viper"
)

// ConfigureMessaging func
func ConfigureMessaging() messaging.Messaging {
	messagingFromConfig := viper.Get("messaging").(map[string]any)

	driver, ok := messagingFromConfig["driver"].(string)
	if !ok {
		panic("messaging driver configuration failed")
	}

	url, ok := messagingFromConfig["url"].(string)
	if !ok {
		panic("messaging url configuration failed")
	}

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
