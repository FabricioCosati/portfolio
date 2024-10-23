package main

import (
	"log"

	fiber_middleware "github.com/fabricio-cosati/portfolio/cmd/middlewares"
	"github.com/fabricio-cosati/portfolio/internal/routes"
	"github.com/fabricio-cosati/portfolio/internal/services"
	translator "github.com/fabricio-cosati/portfolio/internal/translations"
	"github.com/gofiber/fiber/v2"
)

//go:generate goi18n extract -sourceLanguage en -outdir ../../internal/translations
//go:generate goi18n merge -outdir ../../internal/translations ../../internal/translations/active.en.toml ../../internal/translations/translate.pt.toml
//go:generate goi18n merge -sourceLanguage pt -outdir ../../internal/translations ../../internal/translations/active.pt.toml ../../internal/translations/translate.pt.toml
func main() {
	app := fiber.New()
	app.Static("/public", "./internal/public")
	app.Static("/img", "./internal/img")

	init_translator, err := translator.InitTranslator()
	if err != nil {
		log.Fatalf("Failed to init translator: %v", err)
	}

	services.InitCache()
	app.Use(fiber_middleware.LocalizeMiddleware(*init_translator))

	routes.AppendHomeRoutes(app)
	routes.AppendLanguageRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
