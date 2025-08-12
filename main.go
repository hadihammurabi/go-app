package main

import (
	"github.com/gowok/gowok"
	"github.com/gowok/plugins/amqp"
	"github.com/gowok/plugins/gorm"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web"
	"gorm.io/driver/postgres"
)

func main() {
	project := gowok.Get()
	project.Configures(
		gorm.Configure(map[string]gorm.Opener{
			"postgres": postgres.Open,
		}),
		amqp.Configure,
		web.Configure,
		grpc.Configure,
		messaging.Configure,
	)
	project.Run()
}
