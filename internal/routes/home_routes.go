package routes

import (
	"github.com/fabricio-cosati/portfolio/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func AppendHomeRoutes(app *fiber.App) {
	app.Get("/", controllers.RenderHomeHandler)
}
