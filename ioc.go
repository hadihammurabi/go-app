package main

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/mq"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

// NewIOC func
func NewIOC(conf *config.Config) di.IOC {
	ioc := di.IOC{}
	ioc["config"] = conf
	ioc["repository"] = repository.NewRepository(ioc)
	ioc["service"] = service.NewService(ioc)
	ioc["delivery/http"] = rest.Init(ioc)
	ioc["delivery/mq"] = mq.Init(ioc)
	return ioc
}
