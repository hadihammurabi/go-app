package main

import (
	"os"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/runner"
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
	conf := ioc.Get(gowok.Config{})
	apiRest := ioc.Get(rest.APIRest{})
	apiGrpc := ioc.Get(grpc.APIGrpc{})
	apiMessaging := ioc.Get(messaging.APIMessaging{})

	forever := make(chan bool)
	var gracefulStop = make(chan os.Signal)
	runner.GracefulStop(gracefulStop, func() {
		<-gracefulStop

		forever <- true
		apiRest.Stop()
		os.Exit(0)
	})

	go http.ListenAndServe("localhost:6060", nil)

	if conf.App.Rest.Enabled {
		go apiRest.Run()
	}

	if conf.App.Grpc.Enabled {
		go apiGrpc.Run()
	}

	go apiMessaging.Run()
	<-forever
}
