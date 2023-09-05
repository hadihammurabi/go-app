package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
)

// Rest struct
type Rest struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
	Service     *service.Service
	Validator   *gowok.Validator
	Config      *gowok.Config
}

func NewAPIRest() *Rest {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${method} | ${path} | ${status} | ${ip} | ${latency}\n",
	}))
	// app.Use(recover.New())
	app.Use(cors.New())

	dr := driver.Get()
	conf := dr.Config
	validator := dr.Validator
	sv := service.Get()

	middlewares := middleware.NewMiddleware()
	api := &Rest{
		HTTP:        app,
		Middlewares: middlewares,
		Service:     sv,
		Config:      conf,
		Validator:   validator,
	}
	return api
}

func (d *Rest) Run() {
	if !d.Config.App.Rest.Enabled {
		return
	}

	log.Println("API REST started at", d.Config.App.Rest.Host)
	if err := d.HTTP.Listen(d.Config.App.Rest.Host); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (d *Rest) Stop() {
	d.HTTP.Shutdown()
	log.Println("Server was stopped")
}
