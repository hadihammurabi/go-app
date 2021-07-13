package mq

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/service"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/sarulabs/di"
)

// Delivery struct
type Delivery struct {
	Service *service.Service
	Config  *config.Config
	MQ      gorabbitmq.MQ
}

// Init func
func Init(ioc di.Container) (*Delivery, error) {
	service := ioc.Get("service").(*service.Service)
	conf := ioc.Get("config").(*config.Config)

	mq, err := gorabbitmq.NewMQ(conf.MQ.GetURL())
	if err != nil {
		return nil, err
	}

	delivery := &Delivery{
		Service: service,
		Config:  conf,
		MQ:      mq,
	}
	return delivery, nil
}

func (d *Delivery) Run() {
	ConsumeHelloProcess(d)
}
