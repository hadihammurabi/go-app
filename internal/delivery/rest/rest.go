package rest

import (
	"fmt"
	"log"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Delivery struct
type Delivery struct {
	HTTP        *fiber.App
	Middlewares func(int) fiber.Handler
	Service     *service.Service
	Validator   *config.Validator
	Config      *config.Config
}

// Init func
func Init(ioc di.IOC) *Delivery {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${method} | ${path} | ${status} | ${ip} | ${latency}\n",
	}))
	app.Use(recover.New())
	app.Use(cors.New())

	service := ioc["service"].(*service.Service)
	conf := ioc["config"].(*config.Config)

	middleware.Middlewares = map[int]fiber.Handler{}
	middleware.Middlewares[middleware.AUTH] = middleware.Auth(ioc)

	delivery := &Delivery{
		HTTP:        app,
		Middlewares: middleware.Use,
		Service:     service,
		Config:      conf,
		Validator:   config.NewValidator(),
	}
	delivery.ConfigureRoute()
	return delivery
}

func (d *Delivery) Run() {
	connURL := fmt.Sprintf(":%s", d.Config.APP.Port)
	if err := d.HTTP.Listen(connURL); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}
