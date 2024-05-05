package database

import (
	"campsite/internal/config"
	"database/sql"
	"fmt"
	"log"
)

func NewMySqlConnection(config *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s/%s/%s:%s", config.DB.Host, config.SRV.Port, config.DB.User, config.DB.Password, config.DB.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("ERROR: cannot create mysql connection, Get : %s\n", err.Error())
	}

	return db
}
