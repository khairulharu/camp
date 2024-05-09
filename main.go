package main

import (
	"campsite/internal/auth"
	"campsite/internal/config"
	"campsite/internal/database"
	"campsite/internal/middleware"
	"campsite/internal/repository"
	"campsite/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config := config.Get()

	errCh := make(chan error)

	err := <-errCh

	if err != nil {
		log.Fatal("error from errChan")
	}

	dbConnection := database.NewMySqlConnection(config)

	userRepository := repository.NewUserRepository(dbConnection)
	reviewRepository := repository.NewReviewRepository(dbConnection)

	userService := service.NewUser(userRepository)
	reviewService := service.NewReview(reviewRepository)

	authMid := middleware.Authenticate()

	app := fiber.New()

	app.Use(logger.New())

	auth.NewUser(app, userService, authMid)
	auth.NewReview(app, reviewService, authMid)

	errCh <- app.Listen(config.SRV.Host + ":" + config.SRV.Port)

}
