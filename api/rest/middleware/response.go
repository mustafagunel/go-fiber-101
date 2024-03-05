package middleware

import (
	"fiber-tutorial/pkg/i18n"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status      int
	Data        interface{}
	Error       string
	Description string
}

func (r *Response) Send(c *fiber.Ctx) error {
	r.Error = i18n.Translator.Translate(r.Error, nil)
	return c.JSON(r)
}
