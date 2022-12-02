package rest

import (
	"net/http"

	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/response"

	"github.com/gofiber/fiber/v2"
)

type Cache interface {
	Index(*fiber.Ctx) error
}

type cache struct {
	*APIRest
}

var _ Cache = &cache{}

// NewCacheHandler func
func NewCacheHandler(delivery *APIRest) {
	api := &cache{
		delivery,
	}
	router := api.HTTP.Group("/cache")
	router.Get("", api.Index)
}

// Index func
func (api cache) Index(c *fiber.Ctx) error {
	key := "users:1"

	err := api.Config.Cache.Set(c.Context(), key, fiber.Map{"username": "alexunder"})
	if err != nil {
		return response.Fail(c, err.Error(), http.StatusInternalServerError)
	}

	val, err := api.Config.Cache.Get(c.Context(), key)
	if err != nil {
		return response.Fail(c, err.Error(), http.StatusInternalServerError)
	}

	return response.Ok(c, fiber.Map{
		"key": key,
		"val": val,
	})
}
