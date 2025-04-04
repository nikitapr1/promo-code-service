package postgres

import (
	"database/sql"
	"fmt"
	//"github.com/jackc/pgx/v5"
	//"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sql.DB
}

func New(storageURL string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("pgx", storageURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
