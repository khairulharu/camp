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
		JWT: SecretKey{
			Key: os.Getenv("JWT_SECRET_KEY"),
		},
	}
}

// Just get config for testing code see config file and loader.go
func GetConfigTest() *Config {
	//not use env so can run without env just
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error loading .env file get: %s", err.Error())
	// }

	return &Config{
		SRV: Server{
			Host: "",
			Port: "",
		},
		DB: Database{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "root",
			Name:     "campsite",
		},
		JWT: SecretKey{
			Key: os.Getenv("JWT_SECRET_KEY"),
		},
	}
}
