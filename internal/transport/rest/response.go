package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	ServerError = "ServerError"
)

func (h *Handler) ErrorResponse(c *fiber.Ctx, statusCode int, msg string) error {

	lang := c.Get("Accept-Language")

	msg = h.GetMessage(msg, lang)

	return c.Status(statusCode).JSON(map[string]string{"Status": msg})
}

func (h *Handler) GetMessage(msg string, lang string) string {

	tag := language.Make(lang)

	localizer := i18n.NewLocalizer(h.i18n, tag.String())

	message := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msg})

	return message

}
