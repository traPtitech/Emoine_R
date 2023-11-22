package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/traPtitech/Emoine_R/repository/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type Repository struct {
	DB *bun.DB
}

func SetupRepository() (*Repository, error) {
	cfg := mysql.Config{
		User:   getEnvOrDefault("DB_USER", "root"),
		Passwd: getEnvOrDefault("DB_PASSWORD", ""),
		Net:    "tcp",
		Addr: fmt.Sprintf(
			"%s:%s",
			getEnvOrDefault("DB_HOST", "127.0.0.1"),
			getEnvOrDefault("DB_PORT", "3306"),
		),
		DBName:               getEnvOrDefault("DB_NAME", "emoine"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqldb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	if err := migrations.Migrate(context.Background(), db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Repository{DB: db}, nil
}

func getEnvOrDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}
