package helpers

import (
	"fmt"
	"n1h41/auth-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
  CREATE TABLE IF NOT EXISTS USERS (
    ID SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    name TEXT not null,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
  );
  `

func InitDB(config *config.Config) (db *sqlx.DB, err error) {
	// dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=%v", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)
	db, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		return
	}

	createSchema(db)

	fmt.Println("Database connected")
	return
}

func createSchema(db *sqlx.DB) {
	db.MustExec(schema)
	fmt.Println("Schema created successfully")
}
