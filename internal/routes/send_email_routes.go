package routes

import (
	"github.com/fabricio-cosati/portfolio/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func AppendSendEmailRoutes(app *fiber.App) {
	app.Post("/send-email", controllers.SendEmailHandler)
}
