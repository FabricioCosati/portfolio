package controllers

import (
	"fmt"
	"log"

	"github.com/fabricio-cosati/portfolio/cmd/dto"
	"github.com/gofiber/fiber/v2"
	gomail "gopkg.in/mail.v2"
)

type Email struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Message string `json:"message" form:"message"`
}

func SendEmailHandler(c *fiber.Ctx) error {
	email := new(Email)
	if err := c.BodyParser(email); err != nil {
		return err
	}

	message := gomail.NewMessage()

	smtpCredentials := dto.GetSmtpCredentials()

	message.SetHeader("From", email.Email)
	message.SetHeader("To", "fabricio18cosati@gmail.com")
	message.SetHeader("Subject", "Portfolio Contact")
	message.SetHeader("Reply-To", email.Email)
	message.SetBody("text/plain", fmt.Sprintf("Você recebeu um email de %v \n\n%v", email.Name, email.Message))

	d := gomail.NewDialer(smtpCredentials.Host, smtpCredentials.Port, smtpCredentials.Username, smtpCredentials.Password)

	if err := d.DialAndSend(message); err != nil {
		log.Printf("Error sending email: %v", err)
		c.Status(500).SendString("Internal Server Error")
		panic(err)
	}

	return c.Redirect("/")
}
