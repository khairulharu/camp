package repository_test

import (
	"campsite/internal/config"
	"campsite/internal/database"
	"database/sql"
	"testing"

	"gorm.io/gorm"
)

var dbGorm *gorm.DB
var dbRareTest *sql.DB

func TestMain(m *testing.M) {
	configTest := config.GetConfigTest()

	// dbGorm = database.NewGormMySqlConnection(configTest)
	dbRareTest = database.NewMysqlConnection(configTest)

	m.Run()
}
