package controllers

import (
	"log"

	"github.com/fabricio-cosati/portfolio/cmd/dto"
	fiber_reder "github.com/fabricio-cosati/portfolio/cmd/fiber"
	"github.com/fabricio-cosati/portfolio/internal/services"
	"github.com/fabricio-cosati/portfolio/internal/templates/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func RenderHomeHandler(c *fiber.Ctx) error {
	lang := c.Cookies("lang", "en")

	if err := services.ValidateLanguage(lang); err != nil {
		log.Printf("Invalid language selected: %v", lang)
		c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	localizer := c.Locals("localizer").(*i18n.Localizer)
	initLocalizer := services.InitLocalize(*localizer, lang)

	infoData := dto.GetInfoData()
	LanguagesCollection := dto.GetSupportedLanguages()

	texts, textsErr := initLocalizer.LoadHomeTexts()
	if textsErr != nil {
		log.Printf("Error loading home texts: %v", textsErr)
		c.Status(500).SendString("Internal Server Error")
	}

	testimonials, testimonialsErr := initLocalizer.LoadTestimonialsTexts()
	if testimonialsErr != nil {
		log.Printf("Error loading testimonials texts: %v", testimonialsErr)
		c.Status(500).SendString("Internal Server Error")
	}

	return fiber_reder.Render(c, pages.Home("Fabricio Cosati", *infoData, texts, testimonials, *LanguagesCollection, lang))
}
