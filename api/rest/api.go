package rest

import (
	"log"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/service"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

// ConfigureRoute func
func (api *APIRest) ConfigureRoute() {
	api.HTTP.Mount("", Index(api).router)
	// api.HTTP.Mount("/auth", Auth(api).router)
}

// APIRest struct
type APIRest struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
	Service     *service.Service
	Validator   *gowok.Validator
	Config      *gowok.Config
}

var a *APIRest

func NewAPIRest() *APIRest {
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
	api := &APIRest{
		HTTP:        app,
		Middlewares: middlewares,
		Service:     sv,
		Config:      conf,
		Validator:   validator,
	}
	api.ConfigureRoute()
	return api
}

func Get() *APIRest {
	if a != nil {
		return a
	}

	a = NewAPIRest()
	return a
}
func (d *APIRest) Run() {
	if !d.Config.App.Rest.Enabled {
		return
	}

	log.Println("API REST started at", d.Config.App.Rest.Host)
	if err := d.HTTP.Listen(d.Config.App.Rest.Host); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (d *APIRest) Stop() {
	d.HTTP.Shutdown()
	log.Println("Server was stopped")
}
