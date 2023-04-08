package config

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type DatabaseConfig struct {
	Port     int    `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSW,required"`
	Schema   string `env:"DB_SCHEMA,required"`
}

func LoadDatabaseConfig() DatabaseConfig {
	dbConfig := DatabaseConfig{}

	if err := env.Parse(&dbConfig); err != nil {
		log.Fatal(err)
	}

	return dbConfig
}
