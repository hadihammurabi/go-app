package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResponseStatus string

const (
	ResponseStatusFail ResponseStatus = "fail"
	ResponseStatusOk   ResponseStatus = "ok"
)

type FailResponse struct {
	Status ResponseStatus `json:"status"`
	Errors any            `json:"errors"`
}

type OkResponse struct {
	Status ResponseStatus `json:"status"`
	Data   any            `json:"data"`
}

func Fail(c *fiber.Ctx, errs any, status ...int) error {
	s := http.StatusBadRequest
	if len(status) > 0 {
		s = status[0]
	}

	return c.Status(s).JSON(FailResponse{
		Status: ResponseStatusFail,
		Errors: errs,
	})
}

func Ok(c *fiber.Ctx, data any, status ...int) error {
	s := http.StatusOK
	if len(status) > 0 {
		s = status[0]
	}

	return c.Status(s).JSON(OkResponse{
		Status: ResponseStatusOk,
		Data:   data,
	})
}
