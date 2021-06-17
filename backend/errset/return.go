package errset

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ReturnError(err error, c *fiber.Ctx) error {
	switch err {
	case ErrBadRequest:
		return c.SendStatus(http.StatusBadRequest)
	case ErrNotFound:
		return c.SendStatus(http.StatusNotFound)
	case ErrUnAuthorized:
		return c.SendStatus(http.StatusUnauthorized)
	case ErrInternalServer:
		return c.SendStatus(http.StatusInternalServerError)
	}
	return nil
}
