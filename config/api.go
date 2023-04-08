package config

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type ApiConfig struct {
	Port string `env:"PORT,required"`
}

func LoadApiConfig() ApiConfig {
	apiConfig := ApiConfig{}

	if err := env.Parse(&apiConfig); err != nil {
		log.Fatal(err)
	}

	return apiConfig
}
