package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file get: %s", err.Error())
	}

	return &Config{
		SRV: Server{
			Host: os.Getenv("SRV_HOST"),
			Port: os.Getenv("SRV_PORT"),
		},
		DB: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
