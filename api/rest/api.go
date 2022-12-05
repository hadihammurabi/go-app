package rest

import (
	"log"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/go-ioc/ioc"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

// ConfigureRoute func
func (api *APIRest) ConfigureRoute() {
	Index(api)
	// Auth(api)
}

// APIRest struct
type APIRest struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
	Service     *service.Service
	Validator   *gowok.Validator
	Config      *gowok.Config
}

func NewAPIRest() *APIRest {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${method} | ${path} | ${status} | ${ip} | ${latency}\n",
	}))
	// app.Use(recover.New())
	app.Use(cors.New())

	internalApp := ioc.Get(internal.App{})
	service := internalApp.Service
	conf := ioc.Get(gowok.Config{})
	validator := ioc.Get(gowok.Validator{})

	middlewares := middleware.NewMiddleware()
	api := &APIRest{
		HTTP:        app,
		Middlewares: middlewares,
		Service:     service,
		Config:      conf,
		Validator:   validator,
	}
	api.ConfigureRoute()
	return api
}

func (d *APIRest) Run() {
	log.Println("API REST started")
	if err := d.HTTP.Listen(d.Config.App.Rest.Host); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (d *APIRest) Stop() {
	d.HTTP.Shutdown()
	log.Println("Server was stopped")
}
