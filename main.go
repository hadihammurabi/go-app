package main

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
)

func init() {
	driver.PrepareAll()
}

func main() {
	api.Run()

	go gowok.StartPProf()
	gowok.GracefulStop(func() {
		fmt.Println()
		fmt.Println("Stopping...")
	})
}
