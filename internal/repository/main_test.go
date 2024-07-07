package repository_test

import (
	"campsite/internal/config"
	"campsite/internal/database"
	"testing"

	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func TestMain(m *testing.M) {
	configTest := config.GetConfigTest()

	dbGorm = database.NewGormMySqlConnection(configTest)

	m.Run()
}
