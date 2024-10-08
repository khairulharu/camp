package database

import (
	"campsite/internal/config"
	"campsite/internal/domain"
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func NewGormMySqlConnection(config *config.Config) *gorm.DB {

	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Loc
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)

	log.Println("Starting to connect database")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to database get: %s", err.Error())
	}
	log.Println("Database Connecting")

	log.Println("Starting migrations")
	// Run migrations
	err = db.AutoMigrate(
		&domain.Activity{},
		&domain.Admin{},
		&domain.Availability{},
		&domain.Booking{},
		&domain.Campsite{},
		&domain.Facility{},
		&domain.Review{},
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("errror migrate get: %s", err.Error())
	}
	log.Println("Migrate Succesfuly")

	return db
}

func NewMysqlConnection(config *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)

	log.Println("Starting to connect database")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connect to database get: %s", err.Error())
	}

	log.Println("Database Connecting")

	return db
}
