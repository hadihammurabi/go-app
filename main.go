package main

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
)

func init() {
	pkg.PrepareAll()
	internal.PrepareAll()
	api.PrepareAll()
}

func main() {
	go gowok.StartPProf()

	api.Run()

	gowok.GracefulStop(func() {
		fmt.Println()
		fmt.Println("Stopping...")
	})
}
