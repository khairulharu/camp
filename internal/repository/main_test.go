package repository

import (
	"campsite/internal/config"
	"campsite/internal/database"
	"testing"

	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func Main(m *testing.M) {
	configTest := config.GetConfigTest()

	dbGorm = database.NewMySqlConnection(configTest)
}
