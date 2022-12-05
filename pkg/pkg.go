package pkg

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/config"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/runner"
	"github.com/hadihammurabi/go-ioc/ioc"
)

func PrepareAll() {
	runner.PrepareRuntime()

	conf, err := config.Configure()
	if err != nil {
		panic(err)
	}
	ioc.Set(func() gowok.Config {
		return conf
	})

	db, err := database.NewDatabase(database.Config{
		Driver: database.Driver(conf.Database.Driver),
		DSN:    conf.Database.DSN,
	})
	if err != nil {
		panic(err)
	}
	ioc.Set(func() database.Database { return db })

	repo := repository.NewRepository()
	ioc.Set(func() repository.Repository {
		return repo
	})

}
