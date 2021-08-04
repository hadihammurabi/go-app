package mq

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
)

// Delivery struct
type Delivery struct {
	Service *service.Service
	Config  *config.Config
	MQ      gorabbitmq.MQ
}

// Init func
func Init(ioc di.IOC) *Delivery {
	service := ioc[di.DI_SERVICE].(*service.Service)
	conf := ioc[di.DI_CONFIG].(*config.Config)

	delivery := &Delivery{
		Service: service,
		Config:  conf,
		// MQ: conf.MQ,
	}
	return delivery
}

func (d *Delivery) Run() {
	// TODO: Call all MQ consumer here
}
