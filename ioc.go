package main

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/mq"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

// NewIOC func
func NewIOC(conf *config.Config) di.IOC {
	ioc := di.IOC{}
	ioc[di.DI_CONFIG] = conf
	ioc[di.DI_REPOSITORY] = repository.NewRepository(ioc)
	ioc[di.DI_SERVICE] = service.NewService(ioc)
	ioc[di.DI_DELIVERY_REST] = rest.Init(ioc)
	ioc[di.DI_DELIVERY_MQ] = mq.Init(ioc)
	ioc[di.DI_APP] = internal.NewApp(ioc)
	return ioc
}
