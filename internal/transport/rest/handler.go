package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/silapdev/go_i18n_test/internal/service"
)

type Handler struct {
	services *service.Service
	i18n     *i18n.Bundle
}

func NewHandler(services *service.Service, i18n *i18n.Bundle) *Handler {
	return &Handler{services: services, i18n: i18n}
}

func (h *Handler) InitRoutes() *fiber.App {

	router := fiber.New()

	auth := router.Group("/auth")
	{
		auth.Post("/error", h.SignUp)
	}

	return router

}
