package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	viper.AddConfigPath("./config")
	viper.SetConfigName(env)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return viper.ReadInConfig()
}
