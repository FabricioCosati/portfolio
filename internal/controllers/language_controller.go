package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func ChangeLanguageHandler(c *fiber.Ctx) error {
	lang := c.FormValue("language")

	c.Cookie(&fiber.Cookie{
		Name:  "lang",
		Value: lang,
		Path:  "/",
	})

	return c.Redirect("/")
}
