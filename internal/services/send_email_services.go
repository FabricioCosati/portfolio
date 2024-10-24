package services

import (
	"bytes"
	"io"

	"github.com/fabricio-cosati/portfolio/cmd/dto"
	"github.com/fabricio-cosati/portfolio/internal/templates/pages"
	"github.com/gofiber/fiber/v2"
	gomail "gopkg.in/mail.v2"
)

func GetMailMessage(email dto.Email, body string) *gomail.Message {
	message := gomail.NewMessage()

	message.SetHeader("From", email.Email)
	message.SetHeader("To", "fabricio18cosati@gmail.com")
	message.SetHeader("Subject", "Portfolio Contact")
	message.SetHeader("Reply-To", email.Email)
	message.SetBody("text/html", body)

	return message
}

func GetMailDialer(email dto.Email) *gomail.Dialer {
	smtpCredentials := dto.GetSmtpCredentials()

	return gomail.NewDialer(
		smtpCredentials.Host,
		smtpCredentials.Port,
		smtpCredentials.Username,
		smtpCredentials.Password,
	)
}

func LoadEmailBody(c *fiber.Ctx, email dto.Email) (string, error) {
	var buf bytes.Buffer

	bodyTemplate := pages.Email("Email", email)
	writer := io.Writer(&buf)
	err := bodyTemplate.Render(c.Context(), writer)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
