package main

import (
	"belajar-go-rest-api/config/database"
	deliveryHttp "belajar-go-rest-api/delivery/http"
	"belajar-go-rest-api/repository"
	"belajar-go-rest-api/service"
	"log"

	_ "belajar-go-rest-api/docs"

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
