package main

import (
	"log"
	"os"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/mq"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/runner"

	_ "github.com/hadihammurabi/belajar-go-rest-api/docs"

	"github.com/joho/godotenv"
)

// @title Belajar Go REST API
// @version 0.0.1
// @description Ini adalah projek untuk latihan REST API dengan Go
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	_ = godotenv.Load()

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ioc := NewIOC(conf)
	restApp := ioc.Get("delivery/http").(*rest.Delivery)
	mqAppFromDI, err := ioc.SafeGet("delivery/mq")
	var mqApp *mq.Delivery
	if err == nil {
		mqApp = mqAppFromDI.(*mq.Delivery)
	}

	var gracefulStop = make(chan os.Signal)
	runner.GracefulStop(gracefulStop, func() {
		sig := <-gracefulStop
		log.Printf("Caught SIG: %+v\n", sig)
		log.Println("Wait to finishing process")
		restApp.HTTP.Shutdown()
		os.Exit(0)
	})

	forever := make(chan bool)
	go restApp.Run()
	if mqAppFromDI != nil {
		go mqApp.Run()
	}
	<-forever
}
