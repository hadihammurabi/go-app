package http

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/delivery/http/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
	utils "github.com/hadihammurabi/belajar-go-rest-api/util"
	"github.com/sarulabs/di"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

// Delivery struct
type Delivery struct {
	HTTP        *fiber.App
	Middlewares func(int) fiber.Handler
	Service     *service.Service
	Validator   *utils.Validator
}

// Init func
func Init(ioc di.Container) *Delivery {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[\"${time}\", \"${method}\", \"${path}\", \"${status}\", \"${ip}\", \"${latency}\"]\n",
	}))
	app.Use(cors.New())

	middleware.Middlewares = map[int]fiber.Handler{}
	middleware.Middlewares[middleware.AUTH] = middleware.NewAuthMiddleware(config.ConfigureJWT())

	service := ioc.Get("service").(*service.Service)

	delivery := &Delivery{
		HTTP:        app,
		Middlewares: middleware.Use,
		Service:     service,
		Validator:   utils.NewValidator(),
	}
	delivery.ConfigureRoute()
	return delivery
}
