package pkg

import (
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/config"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/runner"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/validator"
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

	database.Configure(conf.Databases)

	ioc.Set(func() gowok.Validator {
		return validator.Configure()
	})

	repo := repository.NewRepository()
	ioc.Set(func() repository.Repository {
		return repo
	})

}
