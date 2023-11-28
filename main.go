package main

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web"
)

func main() {
	project := gowok.Get()
	web.ConfigureRoute()
	project.Runner.Run()
}
