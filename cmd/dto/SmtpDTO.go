package dto

import (
	"github.com/spf13/viper"
)

type Credentials struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Email struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Message string `json:"message" form:"message"`
}

func GetSmtpCredentials() *Credentials {

	return &Credentials{
		Host:     viper.GetString("smtp.host"),
		Port:     viper.GetInt("smtp.port"),
		Username: viper.GetString("smtp.username"),
		Password: viper.GetString("smtp.password"),
	}
}
