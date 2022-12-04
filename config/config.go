package config

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
	"github.com/hadihammurabi/go-ioc/ioc"
)

func New() (gowok.Config, error) {
	conf, err := gowok.Configure("config.yaml")
	if err != nil {
		return conf, err
	}

	db, err := database.NewDatabase(database.Config{
		Driver: database.Driver(conf.Database.Driver),
		DSN:    conf.Database.DSN,
	})
	if err != nil {
		return conf, err
	}

	ioc.Set(func() gowok.Config { return conf })
	ioc.Set(func() database.Database { return db })

	return conf, err
}
