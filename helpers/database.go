package helpers

import (
	"fmt"
	"n1h41/auth-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var userSchema = `
  CREATE TABLE IF NOT EXISTS users (
    ID SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
  );
  `
var resetPasswordSchema = `
  CREATE TABLE IF NOT EXISTS reset_pass (
    ID SERIAL PRIMARY KEY,
    reset_code TEXT NOT NULL UNIQUE,
    user_id INTEGER REFERENCES users(ID),
    used BOOLEAN DEFAULT FALSE
  )
  `

func InitDB(config *config.Config) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v", config.DBHost, config.DBUser, config.DBPassword, config.DBName)
	db, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		return
	}

	createSchema(db)

	fmt.Println("DATABASE CONNECTED")
	return
}

func createSchema(db *sqlx.DB) {
	db.MustExec(userSchema)
  db.MustExec(resetPasswordSchema)
	fmt.Println("SCHEMA CREATED SUCCESSFULLY")
}
