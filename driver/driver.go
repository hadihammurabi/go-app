package driver

import (
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/config"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/database"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/util/runner"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/validator"
)

func PrepareAll() {
	runner.PrepareRuntime()

	conf := config.Configure()
	ioc.Set(func() gowok.Config {
		return *conf
	})

	database.Configure(conf.Databases)

	ioc.Set(func() gowok.Validator {
		return validator.Configure()
	})

	repo := repository.NewRepository()
	ioc.Set(func() repository.Repository {
		return repo
	})

}
