package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiConfig
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		ApiConfig: LoadApiConfig(),
	}
}
