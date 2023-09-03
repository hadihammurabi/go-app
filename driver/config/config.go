package config

import (
	"os"

	"github.com/gowok/gowok"
)

func Configure() *gowok.Config {
	conf := gowok.Must(
		gowok.Configure(os.OpenFile("config.yaml", os.O_RDONLY, 600)),
	)

	return conf
}
