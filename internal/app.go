package internal

import (
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/db/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/util/di"
)

type App struct {
	APIRest *rest.APIRest
}

func NewIOC(conf config.Config) di.IOC {
	ioc := di.IOC{}
	ioc[di.DI_CONFIG] = conf
	ioc[di.DI_REPOSITORY] = repository.NewRepository(ioc)
	ioc[di.DI_SERVICE] = service.NewService(ioc)
	ioc[di.DI_APP] = NewApp(ioc)
	return ioc
}

func NewApp(ioc di.IOC) *App {
	return &App{
		APIRest: rest.NewAPIRest(ioc),
	}
}
