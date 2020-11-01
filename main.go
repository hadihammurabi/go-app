package main

import (
	"belajar-go-rest-api/config/database"
	deliveryHttp "belajar-go-rest-api/delivery/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db, err := database.ConfigureDatabase()

	if err != nil {
		panic(err)
	}

	database.MigrateDatabase(db)
	httpApp := deliveryHttp.Init(db)

	forever := make(chan bool)
	go func() {
		httpApp.HTTP.Listen(":8080")
	}()
	<-forever
}
