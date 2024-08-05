package main

import (
	"campsite/internal/auth"
	"campsite/internal/config"
	"campsite/internal/database"
	"campsite/internal/middleware"
	"campsite/internal/repository"
	"campsite/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config := config.Get()

	// errCh := make(chan error)

	//buat dua buah tipe connection jika menggunakan orm maka return db Grom jika tidak maka return *sql.DB
	dbConnection := database.NewGormMySqlConnection(config)
	userRepository := repository.UseUserRepository(config.DB.IsUseOrm, nil, nil)
	reviewRepository := repository.NewReviewRepository(dbConnection)

	userService := service.NewUser(userRepository)
	reviewService := service.NewReview(reviewRepository)

	authMid := middleware.Authenticate()

	app := fiber.New()

	app.Use(logger.New())

	auth.NewUser(app, userService, authMid)
	auth.NewReview(app, reviewService, authMid)

	_ = app.Listen(config.SRV.Host + ":" + config.SRV.Port)

}
