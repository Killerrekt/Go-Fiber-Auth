package utils

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type cfg struct {
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
}

var Config cfg

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	if err := env.Parse(&Config); err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("Configuration successfully loaded")
}
