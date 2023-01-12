package model

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var db *sqlx.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot collect .env: %s", err)
	}

	cfg := mysql.Config{
		User:   getEnvOrDefault("DB_USER", "root"),
		Passwd: getEnvOrDefault("DB_PASSWORD", ""),
		Net:    "tcp",
		Addr: fmt.Sprintf(
			"%s:%s",
			getEnvOrDefault("DB_HOST", "127.0.0.1"),
			getEnvOrDefault("DB_PORT", "3306"),
		),
		DBName: getEnvOrDefault("DB_NAME", "emoine"),
	}

	_db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	db = _db
}

func getEnvOrDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}
