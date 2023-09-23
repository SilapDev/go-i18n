package rest

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {

	return h.ErrorResponse(c, http.StatusInternalServerError, ServerError)
}
