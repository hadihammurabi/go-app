package mq

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/service"
	"github.com/sarulabs/di"
)

// Delivery struct
type Delivery struct {
	Service *service.Service
	Config  *config.Config
}

// Init func
func Init(ioc di.Container) *Delivery {
	service := ioc.Get("service").(*service.Service)
	conf := ioc.Get("config").(*config.Config)

	delivery := &Delivery{
		Service: service,
		Config:  conf,
	}
	return delivery
}

func (d *Delivery) Run() {
	ConsumeHelloProcess(d)
}
