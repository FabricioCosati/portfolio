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

func GetSmtpCredentials() *Credentials {

	return &Credentials{
		Host:     viper.GetString("smtp.host"),
		Port:     viper.GetInt("smtp.port"),
		Username: viper.GetString("smtp.username"),
		Password: viper.GetString("smtp.password"),
	}
}
