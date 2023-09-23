package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
)

const (
	ServerError = "ServerError"
)

func (h *Handler) ErrorResponse(c *fiber.Ctx, statusCode int, msg string) error {

	lang := c.Get("Accept-Language")
	log.Printf("language = %s\n", lang)
	msg = h.GetLocalizedMessage(msg, lang)

	return c.Status(statusCode).JSON(map[string]string{"Status": msg})
}

func (h *Handler) GetLocalizedMessage(msg string, lang string) string {

	//tag := language.Make(lang)

	localizer := i18n.NewLocalizer(h.i18n, lang)
	message := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msg})

	return message

}
