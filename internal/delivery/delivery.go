package delivery

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/mq"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

type Delivery struct {
	Rest *rest.Delivery
	MQ   *mq.Delivery
}

func NewDelivery(ioc di.IOC) *Delivery {
	return &Delivery{
		Rest: ioc[di.DI_DELIVERY_REST].(*rest.Delivery),
		MQ:   ioc[di.DI_DELIVERY_MQ].(*mq.Delivery),
	}
}

func (d Delivery) Start() {
	go d.Rest.Run()
	go d.MQ.Run()
}
