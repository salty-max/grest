package config

import (
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	DB_HOST       string `mapstructure:"DB_HOST"`
	DB_PORT       string `mapstructure:"DB_PORT"`
	DB_NAME       string `mapstructure:"DB_NAME"`
	DB_USER       string `mapstructure:"DB_USER"`
	DB_PASSWORD   string `mapstructure:"DB_PASSWORD"`
	PORT          string `mapstructure:"PORT"`
	TIMEOUT_DELAY string `mapstructure:"TIMEOUT_DELAY"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			DB_HOST:       os.Getenv("DB_HOST"),
			DB_PORT:       os.Getenv("DB_PORT"),
			DB_NAME:       os.Getenv("DB_NAME"),
			DB_USER:       os.Getenv("DB_USER"),
			DB_PASSWORD:   os.Getenv("DB_PASSWORD"),
			PORT:          os.Getenv("PORT"),
			TIMEOUT_DELAY: os.Getenv("TIMEOUT_DELAY"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// validate config here

	return
}
