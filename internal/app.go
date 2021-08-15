package internal

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

type App struct {
	Delivery *delivery.Delivery
}

func NewApp(ioc di.IOC) *App {
	return &App{
		Delivery: delivery.NewDelivery(ioc),
	}
}
