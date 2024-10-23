package middlewares

import (
	translator "github.com/fabricio-cosati/portfolio/internal/translations"
	"github.com/gofiber/fiber/v2"
)

func LocalizeMiddleware(initTranslator translator.Translator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		lang := c.Cookies("lang", "en")
		localizer := initTranslator.InitLocalize(lang)

		c.Locals("localizer", localizer)

		return c.Next()
	}
}
