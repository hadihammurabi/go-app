package internal

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

type App struct {
	Delivery   *delivery.Delivery
	Service    *service.Service
	Repository *repository.Repository
}

func NewApp(ioc di.IOC) *App {
	ioc[di.DI_REPOSITORY] = repository.NewRepository(ioc)
	ioc[di.DI_SERVICE] = service.NewService(ioc)

	return &App{
		Delivery: delivery.NewDelivery(ioc),
	}
}
