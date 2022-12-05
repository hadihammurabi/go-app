package config

import (
	"github.com/gowok/gowok"
)

func Configure() (gowok.Config, error) {
	conf, err := gowok.Configure("config.yaml")
	if err != nil {
		return conf, err
	}

	return conf, err
}
