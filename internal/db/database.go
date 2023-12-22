package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"n1h41/auth-service/config"
)

func InitDB(config *config.Config) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v", config.DBHost, config.DBUser, config.DBPassword, config.DBName)
	db, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		return
	}

	fmt.Println("DATABASE CONNECTED")
	return
}
