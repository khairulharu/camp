package database

import (
	"campsite/internal/config"
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	config := config.GetConfigTest()
	mysqlConnection := NewMysqlConnection(config)

	if err := mysqlConnection.Ping(); err != nil {
		t.Errorf("error when trying connect to database, %s", err.Error())
	}
}
