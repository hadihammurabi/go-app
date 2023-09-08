package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
)

// Rest struct
type Rest struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
}

func NewAPIRest() *Rest {
	app := gowok.NewHTTP(&driver.Get().Config.App.Rest)

	api := &Rest{
		HTTP:        app,
		Middlewares: middleware.NewMiddleware(),
	}
	return api
}

func (d *Rest) Run() {
	restConf := driver.Get().Config.App.Rest
	if !restConf.Enabled {
		return
	}

	log.Println("API REST started at", restConf.Host)
	if err := d.HTTP.Listen(restConf.Host); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (d *Rest) Stop() {
	d.HTTP.Shutdown()
	log.Println("Server was stopped")
}
