package main

import (
	"log"
	"os"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	deliveryHttp "github.com/hadihammurabi/belajar-go-rest-api/internal/app/delivery/http"
	"github.com/hadihammurabi/belajar-go-rest-api/util"

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
	httpApp := deliveryHttp.Init(ioc)

	var gracefulStop = make(chan os.Signal)
	util.GracefulStop(gracefulStop, func() {
		sig := <-gracefulStop
		log.Printf("Caught SIG: %+v\n", sig)
		log.Println("Wait to finishing process")
		httpApp.HTTP.Shutdown()
		os.Exit(0)
	})

	forever := make(chan bool)
	go httpApp.Run()
	<-forever
}
