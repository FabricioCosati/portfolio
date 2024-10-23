package routes

import (
	"github.com/fabricio-cosati/portfolio/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func AppendLanguageRoutes(app *fiber.App) {
	app.Post("/change-lang", controllers.ChangeLanguageHandler)
}
