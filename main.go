package main

import (
	deliveryHttp "github.com/hadihammurabi/belajar-go-rest-api/delivery/http"

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

	ioc := NewIOC()
	httpApp := deliveryHttp.Init(ioc)

	forever := make(chan bool)
	go func() {
		httpApp.HTTP.Listen(":8080")
	}()
	<-forever
}
