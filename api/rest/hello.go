package rest

import (
	"encoding/json"
	"net/http"

	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/response"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"

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
	message, _ := json.Marshal(map[string]string{
		"message": "hello bro!",
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
