package main

import (
	"campsite/internal/config"
	"campsite/internal/database"
	"campsite/internal/repository"
)

func main() {

	config := config.Get()

	dbConnection := database.NewMySqlConnection(config)

	_ = repository.NewUserRepository(dbConnection)
}
