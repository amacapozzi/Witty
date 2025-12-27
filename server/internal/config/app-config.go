package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppMode      string `mapstructure:"APP_MODE"`
	ServerPort   string `mapstructure:"APP_PORT"`
	DatabaseUrl  string `mapstructure:"DATABASE_URL"`
	ServerSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() *AppConfig {
	appConfig := AppConfig{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cant read .env file", err)
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &appConfig

}
