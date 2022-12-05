package internal

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/go-ioc/ioc"
)

type App struct {
	Service *service.Service
}

func NewApp() *App {
	return &App{
		Service: service.NewService(),
	}
}

func PrepareAll() {
	app := NewApp()
	ioc.Set(func() App {
		return *app
	})
}
