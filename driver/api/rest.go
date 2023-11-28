package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web/middleware"
)

// Rest struct
type Rest struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
}

func NewAPIRest() *Rest {
	api := &Rest{
		HTTP:        gowok.Get().Web,
		Middlewares: middleware.NewMiddleware(),
	}
	return api
}
