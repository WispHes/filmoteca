package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func NewTestDB() *DB {
	db, err := sql.Open("postgres", "host=localhost user=test password=test dbname=test sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping test database: %v", err)
	}

	return &DB{db}
}
