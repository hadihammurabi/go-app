package main

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web"
)

func main() {
	project := gowok.Get()
	project.Configures(
		web.Configure,
		grpc.Configure,
		messaging.Configure,
	)
	project.Run()
}
