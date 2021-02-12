package main

import (
	"log"

	"github.com/hadihammurabi/belajar-go-rest-api/config/database"
	deliveryHttp "github.com/hadihammurabi/belajar-go-rest-api/delivery/http"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/service"

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

	db, err := database.ConfigureDatabase()

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	httpApp := deliveryHttp.Init(service)

	forever := make(chan bool)
	go func() {
		httpApp.HTTP.Listen(":8080")
	}()
	<-forever
}
