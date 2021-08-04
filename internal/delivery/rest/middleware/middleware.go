package middleware

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

// Middlewares type
type Middlewares struct {
	config  *config.Config
	service *service.Service
}

func NewMiddleware(ioc di.IOC) *Middlewares {
	config := ioc[di.DI_CONFIG].(*config.Config)
	service := ioc[di.DI_SERVICE].(*service.Service)

	return &Middlewares{
		config:  config,
		service: service,
	}
}
