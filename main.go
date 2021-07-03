package main

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	deliveryHttp "github.com/hadihammurabi/belajar-go-rest-api/internal/app/delivery/http"

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

	forever := make(chan bool)
	go httpApp.Run()
	<-forever
}
