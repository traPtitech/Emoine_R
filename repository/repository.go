package repository

import "github.com/uptrace/bun"

type Repository struct {
	DB *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{DB: db}
}
