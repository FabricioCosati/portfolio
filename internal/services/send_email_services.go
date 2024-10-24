package services

import (
	"fmt"

	"github.com/fabricio-cosati/portfolio/cmd/dto"
	gomail "gopkg.in/mail.v2"
)

func GetMailMessage(email dto.Email) *gomail.Message {
	message := gomail.NewMessage()

	message.SetHeader("From", email.Email)
	message.SetHeader("To", "fabricio18cosati@gmail.com")
	message.SetHeader("Subject", "Portfolio Contact")
	message.SetHeader("Reply-To", email.Email)
	message.SetBody("text/plain", fmt.Sprintf("Você recebeu um email de %v \n\n%v", email.Name, email.Message))

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
