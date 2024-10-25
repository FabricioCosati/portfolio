package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env not found")
	}

	viper.AddConfigPath("./config")
	viper.SetConfigName(env)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return viper.ReadInConfig()
}
