package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gowok/gowok/driver/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/response"

	"github.com/gofiber/fiber/v2"
)

type Hello interface {
	Index(*fiber.Ctx) error
}

type hello struct {
	*APIRest
}

var _ Hello = &hello{}

// NewHelloHandler func
func NewHelloHandler(delivery *APIRest) {

	api := &hello{
		delivery,
	}
	router := api.HTTP.Group("/hello")
	router.Post("", api.Index)
}

// Index func
func (api hello) Index(c *fiber.Ctx) error {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.QueryParser(&input); err != nil {
		return response.Fail(c, err.Error(), http.StatusBadRequest)
	}

	message, _ := json.Marshal(map[string]any{
		"message": "hello bro!",
		"data":    input,
	})

	err := api.Config.Messaging.Publish("hello", "", messaging.Message{
		Headers: messaging.Table{
			"content-type": "application/json",
		},
		Message: message,
	})

	if err != nil {
		return response.Fail(c, err, http.StatusInternalServerError)
	}

	return response.Ok(c, "message sent")
}