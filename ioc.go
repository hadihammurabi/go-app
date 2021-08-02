package main

import (
	"log"

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
	deliveryMQ, err := mq.Init(ioc)
	if err != nil {
		log.Printf("can not configure MQ. Caused by %s\n", err.Error())
	} else {
		ioc["delivery/mq"] = deliveryMQ
	}
	return ioc
}
