package main

import (
	"os"

	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/runner"
	"github.com/hadihammurabi/go-ioc/ioc"

	"net/http"
	_ "net/http/pprof"
)

func init() {
	runner.PrepareRuntime()
	runner.PrepareConfig()
}

func main() {
	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() config.Config {
		return conf
	})

	repo := repository.NewRepository()
	ioc.Set(func() repository.Repository {
		return repo
	})

	app := internal.NewApp()
	ioc.Set(func() internal.App {
		return *app
	})

	apiRest := rest.NewAPIRest()
	ioc.Set(func() rest.APIRest {
		return *apiRest
	})

	apiMessaging := messaging.NewAPIMessaging()
	ioc.Set(func() messaging.APIMessaging {
		return *apiMessaging
	})

	apiGrpc := grpc.NewAPIGrpc()
	ioc.Set(func() grpc.APIGrpc {
		return *apiGrpc
	})

	forever := make(chan bool)
	var gracefulStop = make(chan os.Signal)
	runner.GracefulStop(gracefulStop, func() {
		<-gracefulStop

		forever <- true
		apiRest.Stop()
		os.Exit(0)
	})

	go http.ListenAndServe("localhost:6060", nil)
	go apiRest.Run()
	go apiMessaging.Run()
	go apiGrpc.Run()
	<-forever
}
