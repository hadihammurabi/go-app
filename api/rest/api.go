package rest

import (
	"fmt"
	"log"

	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/util/di"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// APIRest struct
type APIRest struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
	Service     service.Service
	Validator   *config.Validator
	Config      config.Config
}

func NewAPIRest(ioc di.IOC) *APIRest {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${method} | ${path} | ${status} | ${ip} | ${latency}\n",
	}))
	app.Use(recover.New())
	app.Use(cors.New())

	service := ioc[di.DI_SERVICE].(service.Service)
	conf := ioc[di.DI_CONFIG].(config.Config)

	middlewares := middleware.NewMiddleware(ioc)
	api := &APIRest{
		HTTP:        app,
		Middlewares: middlewares,
		Service:     service,
		Config:      conf,
		Validator:   config.NewValidator(),
	}
	api.ConfigureRoute()
	return api
}

func (d *APIRest) Run() {
	connURL := fmt.Sprintf(":%s", d.Config.APP.Port)
	if err := d.HTTP.Listen(connURL); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (d *APIRest) Stop() {
	d.HTTP.Shutdown()
	log.Println("Server was stopped")
}

// ConfigureRoute func
func (api *APIRest) ConfigureRoute() {
	api.HTTP.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
		})
	})

	NewAuthHandler(api)
	NewUserHandler(api)
}
