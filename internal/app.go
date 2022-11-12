package internal

import "github.com/hadihammurabi/belajar-go-rest-api/internal/service"

type App struct {
	Service *service.Service
}

func NewApp() *App {
	return &App{
		Service: service.NewService(),
	}
}
