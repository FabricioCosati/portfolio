package controllers

import (
	"log"

	"github.com/fabricio-cosati/portfolio/cmd/dto"
	"github.com/fabricio-cosati/portfolio/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SendEmailHandler(c *fiber.Ctx) error {
	email := new(dto.Email)
	if err := c.BodyParser(email); err != nil {
		return err
	}

	d := services.GetMailDialer(*email)
	message := services.GetMailMessage(*email)

	if err := d.DialAndSend(message); err != nil {
		log.Printf("Error sending email: %v", err)
		c.Status(500).SendString("Internal Server Error")
		panic(err)
	}

	return c.Redirect("/")
}
