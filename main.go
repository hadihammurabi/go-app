package main

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
	"github.com/hadihammurabi/go-ioc/ioc"

	"net/http"
	_ "net/http/pprof"
)

func init() {
	pkg.PrepareAll()
	internal.PrepareAll()
	api.PrepareAll()
}

func main() {
	go http.ListenAndServe("localhost:6060", nil)
	go (ioc.Get(rest.APIRest{})).Run()
	go (ioc.Get(grpc.APIGrpc{})).Run()
	go (ioc.Get(messaging.APIMessaging{})).Run()

	gowok.GracefulStop(func() {
		fmt.Println()
		fmt.Println("Stopping...")
	})
}
