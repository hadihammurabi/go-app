package main

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/api"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
)

func init() {
	pkg.PrepareAll()
	internal.PrepareAll()
	api.PrepareAll()
}

func main() {
	go (new(rest.PProf)).Run()
	go (ioc.Get(rest.APIRest{})).Run()
	go (ioc.Get(grpc.APIGrpc{})).Run()
	go (ioc.Get(messaging.APIMessaging{})).Run()

	gowok.GracefulStop(func() {
		fmt.Println()
		fmt.Println("Stopping...")
	})
}
